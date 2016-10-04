dst=dist

build:
	go build -ldflags "-s -X main.version=`git rev-parse --short HEAD`" -o $(dst)/magnolia demo/main.go
	-cp -rv demo/locales $(dst)/

clean:
	-rm -rv $(dst)
