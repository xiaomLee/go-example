.PHONY: build

build:
	CGO_ENABLE=1 GOOS=linux GOARCH=amd64 go build -o main main.go

lib-any:
	g++ lib-src/any/any.cpp -fPIC -shared -o ./libs/libany.so

lib-bye:
	g++ lib-src/bye/bye.c lib-src/bye/print.a -fPIC -shared -o ./libs/libbye.so

export-go-print:
	go build -buildmode=c-archive -o lib-src/bye/print.a lib-src/bye/gen-go-print/print.go

build-bye:
	g++ lib-src/bye/bye.c lib-src/bye/print.a -o bye.out

build-any:
	g++ lib-src/any/any.cpp -o any.out