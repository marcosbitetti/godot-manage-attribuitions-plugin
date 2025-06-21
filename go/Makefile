default: cross_compile

cross_compile:
	GOOS=linux GOARCH=amd64 go build -o ../bin/attribuitions-amd64-linux ./cmd/commandline/main.go
	upx -9 ../bin/attribuitions-amd64-linux
	#GOOS=linux GOARCH=386 go build -o ../bin/attribuitions-386-linux ./cmd/commandline/main.go
	#upx -9 ../bin/attribuitions-386-linux
	GOOS=windows GOARCH=amd64 go build -o ../bin/attribuitions-amd64.exe ./cmd/commandline/main.go
	upx -9 ../bin/attribuitions-amd64.exe
	#GOOS=windows GOARCH=386 go build -o ../bin/attribuitions-386.exe ./cmd/commandline/main.go
	#upx -9 ../bin/attribuitions-386.exe
	GOOS=darwin GOARCH=amd64 go build -o ../bin/attribuitions-amd64-darwin ./cmd/commandline/main.go
	upx -9 ../bin/attribuitions-amd64-darwin

webserver_compile:
	GOOS=linux GOARCH=amd64 go build -o ../bin/attribuitions-local-server-amd64-linux ./cmd/webserver/main.go

test:
	go test -v -count=1 ./cmd/...
