1. web
http://localhost:6060/debug/pprof


2. cmd
# 只采集特定时间段的CPU数据（需Go 1.20+）
curl --output cpu.pprof "http://localhost:6060/debug/pprof/profile?seconds=30&delta=10"
# 生成两个版本的profile
go tool pprof -base old.pprof new.pprof

# 网页对比模式
go tool pprof -http=:8080 -diff_base old.pprof new.pprof

# list web top

go tool pprof "http://localhost:6060/debug/pprof/goroutine"

GODEBUG=gctrace=1 ./go-pprof-practice | grep gc


资料：
https://github.com/wolfogre/go-pprof-practice


