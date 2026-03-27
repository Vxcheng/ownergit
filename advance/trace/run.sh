 go tool trace trace.out

curl http://127.0.0.1:8082/debug/pprof/trace?seconds=20 > trace.out

http://127.0.0.1:60767/sched
http://127.0.0.1:60767/trace


GODEBUG=schedtrace=1000 ./trace
SCHED 0ms: gomaxprocs=8 idleprocs=7 threads=6 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [ 0 0 0 0 0 0 0 0 ] schedticks=[ 1 0 0 0 0 0 1 3 ]
SCHED 1007ms: gomaxprocs=8 idleprocs=4 threads=11 spinningthreads=1 needspinning=0 idlethreads=4 runqueue=0 [ 0 0 0 0 0 0 0 0 ] schedticks=[ 1954 2035 54 371 995 2046 3859 3470 ]

资料
https://blog.csdn.net/u013911096/article/details/139156238
https://cloud.tencent.com/developer/article/2502912