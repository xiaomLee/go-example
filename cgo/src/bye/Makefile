CC = gcc

SRCS = $(shell find . -type f -name '*.c')
OBJS = $(SRCS:.c=.o)
TARGET = libbye.so.1.0.0

CFLAGS = -Wall -g -O -fPIC
INCLUDE = -I$(PWD)

build: pre
	$(CC) -o bye bye.c libprint.so

lib: clean pre $(OBJS)
	$(CC) -shared -fPIC -o $(TARGET) $(OBJS)
	@ln -s $(TARGET) libbye.so

$(OBJS): $(SRCS)
	$(CC) $(CFLAGS)  $(INCLUDE) -c $^

clean:
	rm -rf $(TARGET) $(OBJS) libbye.so print.so print.h

pre: export-go

export-go:
#	# static library
#	go build -buildmode=c-archive -o print.a ./go-print/print.go
	# dynamic library
	go build -buildmode=c-shared -o libprint.so ./go-print/print.go