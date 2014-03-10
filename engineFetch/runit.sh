go build -o bin
go run *.go
go tool pprof bin *.prof