# Quick start
## setup env
- wget -c last-ver.tar.gz
- unzip to /usr/local/go
- add env 
```
export GOPATH=$HOME/projects/go
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```
## run program
- go run main.go
## compile for linux
- GOOS=linux GOARCH=amd64 go build -o linuxapp 
## format code
- gofmt