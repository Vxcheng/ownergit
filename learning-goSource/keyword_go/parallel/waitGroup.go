package parallel

import (
	"log"
	"sync"
)

type User struct {
	wg  *sync.WaitGroup
	num int
}

func NewUser() *User {
	return &User{
		wg: new(sync.WaitGroup),
	}
}

func Stu_waitGroup(u *User, nums int) error {
	log.Println("学习waitGroup")

	for i := 0; i < nums; i++ {
		u.wg.Add(1)
		go u.Count()
	}
	u.wg.Wait()
	log.Println("num:", u.num)
	return nil
}
func (u *User) Count() {
	defer u.wg.Done()
	u.num++
}
