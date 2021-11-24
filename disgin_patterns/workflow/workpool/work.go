package workpool

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

type Job interface {
	Do(workerID int) error
}

type Payload struct {
	id int
}

func (p *Payload) Do(workerID int) error {
	fmt.Printf("job-%d doing by worker-%d\n", p.id, workerID)
	return nil
}

type Worker struct {
	id          int64
	name        string
	state       State
	pendingJobC chan Job
	workerC     chan chan Job
	finished    chan struct{}
	quit        bool
}

type State int64

const (
	Ready State = iota + 1
	Running
	Destroy
)

func NewWorker(workerC chan chan Job, id int64) *Worker {
	return &Worker{
		id:          id,
		name:        fmt.Sprintf("worker-%d", id),
		pendingJobC: make(chan Job),
		workerC:     workerC,
		finished:    make(chan struct{}),
	}
}

func (w *Worker) Name() string {
	return w.name
}

func (w *Worker) Start(workerStateC chan<- *Worker) (err error) {
	w.workerC <- w.pendingJobC // attach

	select {
	case job, opening := <-w.pendingJobC:
		if !opening {
			return
		}
		w.state = Running
		workerStateC <- w

		defer func() {
			w.state = Ready
			workerStateC <- w
		}()
		if err = job.Do(int(w.id)); err != nil { // load
			return fmt.Errorf("%s do job failed, err: %v", w.name, err)
		}
	case <-w.finished:
		w.state = Destroy
		fmt.Printf("%s stop\n", w.name)

		workerStateC <- w
		close(w.pendingJobC)
		w.quit = true
		return
	}

	return
}

func (w *Worker) Stop() {
	w.finished <- struct{}{}
}

func (w *Worker) GetState() State {
	return w.state
}

func (w *Worker) Quit() bool { return w.quit }

type Scheduler struct {
	maxWorkerNum, minWorkerNum uint64
	workerNum, idleWorkerNum   int64
	allJobC                    chan Job
	workbenchC                 chan chan Job
	workers                    map[string]*Worker
	workerStateC               chan *Worker
	quit                       chan struct{}

	rwmu sync.RWMutex
}

const (
	maxWorkerNum = 5
	maxJobNum    = 100
	minWorkerNum = 3
)

func NewScheduler() *Scheduler {
	return &Scheduler{
		maxWorkerNum: maxWorkerNum,
		minWorkerNum: minWorkerNum,
		allJobC:      make(chan Job, maxJobNum),
		workbenchC:   make(chan chan Job, maxWorkerNum),
		workers:      make(map[string]*Worker),
		workerStateC: make(chan *Worker),
		quit:         make(chan struct{}),
	}
}

func (s *Scheduler) Run() {
	go s.schedule()
}

func (s *Scheduler) schedule() {
	fmt.Println("start schedule")
	// go s.monitor()

	for {
		select {
		case job, opening := <-s.allJobC:
			if !opening {
				return
			}
			if atomic.LoadInt64(&s.workerNum) < int64(s.minWorkerNum) {
				s.applyWorker(int64(s.minWorkerNum) - s.workerNum)
			} else {
				s.applyWorker(1)
			}

			pendingJobC, ok := <-s.workbenchC
			if ok {
				pendingJobC <- job // store
			}
		case <-s.quit:
			fmt.Println("Scheduler exit")

			close(s.allJobC)
			// close(s.workbenchC)
			s.closeWorkers()
			return
		}
	}
}

func (s *Scheduler) monitor() {
	for {
		select {
		case worker := <-s.workerStateC:
			switch worker.GetState() {
			case Running:
				atomic.AddInt64(&s.idleWorkerNum, -1)
			case Ready:
				atomic.AddInt64(&s.idleWorkerNum, 1)
			case Destroy:
				atomic.AddInt64(&s.workerNum, -1)
				atomic.AddInt64(&s.idleWorkerNum, -1)
				s.rwmu.RLock()
				delete(s.workers, worker.Name())
				s.rwmu.RUnlock()
			}
		}

		if s.workerNum == 0 {
			return
		}
	}
}

func (s *Scheduler) applyWorker(num int64) {
	min := int(math.Min(float64(int64(s.maxWorkerNum)-s.workerNum), float64(num)))
	for i := 0; i < min; i++ {
		atomic.AddInt64(&s.workerNum, 1)
		atomic.AddInt64(&s.idleWorkerNum, 1)

		go func(workerID int64) {
			w := NewWorker(s.workbenchC, workerID)
			s.rwmu.Lock()
			_, exist := s.workers[w.Name()]
			if !exist {
				s.workers[w.Name()] = w
			}
			s.rwmu.Unlock()

			var err error
			for {
				if err = w.Start(s.workerStateC); err != nil {
					fmt.Println(err)
				}
				if w.Quit() {
					return
				}
			}
		}(s.workerNum)
	}
}

func (s *Scheduler) applyWorker2(num int64) {
	min := int(math.Min(float64(int64(s.maxWorkerNum)-s.workerNum), float64(num)))
	for i := 0; i < min; i++ {
		atomic.AddInt64(&s.workerNum, 1)
		atomic.AddInt64(&s.idleWorkerNum, 1)

		go func(workerID int64) {
			w := NewWorker(s.workbenchC, workerID)
			s.rwmu.Lock()
			_, exist := s.workers[w.Name()]
			if !exist {
				s.workers[w.Name()] = w
			}
			s.rwmu.Unlock()

			for {
				w.workerC <- w.pendingJobC // attach

				select {
				case job, opening := <-w.pendingJobC:
					if !opening {
						return
					}
					w.state = Running
					atomic.AddInt64(&s.idleWorkerNum, -1)

					if err := job.Do(int(w.id)); err != nil { // load
						fmt.Printf("%s do job failed, err: %v\n", w.name, err)
					}

					w.state = Ready
					atomic.AddInt64(&s.idleWorkerNum, 1)
				case <-w.finished:
					w.state = Destroy
					fmt.Printf("%s stop\n", w.name)
					atomic.AddInt64(&s.workerNum, -1)
					atomic.AddInt64(&s.idleWorkerNum, -1)
					s.rwmu.RLock()
					delete(s.workers, w.Name())
					s.rwmu.RUnlock()

					close(w.pendingJobC)
					return
				}
			}
		}(s.workerNum)
	}
}

func (s *Scheduler) closeWorkers() {
	for _, w := range s.workers {
		go w.Stop()
	}
}

func (s *Scheduler) Close() {
	s.quit <- struct{}{}
}

func (s *Scheduler) AddJob(job Job) {
	s.allJobC <- job
}
