CC = g++


OUTPUT = output
SRCS = $(shell find ./say-cpp-src -type f -name '*.cpp')
OBJS = $(SRCS:.cpp=.o)
TARGET = $(OUTPUT)/libsay.so

CFLAGS = -Wall -g -O -fPIC
INCLUDE = -I/usr/local/include

all: clean output c-dynamic go-plugin build

build:
	go build main.go

go-plugin:
	GOOS=linux go build -buildmode=plugin -o ./output/libgo_say_plugin.so ./say/say_plugin.go

c-dynamic: $(OBJS)
	$(CC) -shared -fPIC -o $(TARGET) $(OBJS)
	@#ln -s $(TARGET) libsay.so

$(OBJS): $(SRCS)
	$(CC) $(CFLAGS) $(INCLUDE) -c $^

output:
	mkdir -p output

clean:
	rm -rf output