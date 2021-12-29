TAG=$(shell git rev-parse --short HEAD)
build:
	rm -rf bins || true
	mkdir bins
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ./bins/tunaps_linux cmd/tunaps/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ./bins/tunapc_linux cmd/tunapc/*.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o ./bins/tunaps_mac cmd/tunaps/*.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o ./bins/tunapc_mac cmd/tunapc/*.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o ./bins/tunaps_win.exe cmd/tunaps/*.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o ./bins/tunapc_win.exe cmd/tunapc/*.go

testing:
	go test -v -count 1 ./...
	golangci-lint run -v
