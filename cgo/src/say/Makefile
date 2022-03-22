CC = g++

SRCS = $(shell find . -type f -name '*.cpp')
OBJS = $(SRCS:.cpp=.o)
TARGET = libsay.so.1.0.0

CFLAGS = -Wall -g -O -fPIC
INCLUDE = -I/usr/local/include

all: $(OBJS)
	$(CC) -shared -fPIC -o $(TARGET) $(OBJS)
	@ln -s $(TARGET) libsay.so

$(OBJS): $(SRCS)
	$(CC) $(CFLAGS) $(INCLUDE) -c $^

clean:
	rm -rf $(TARGET) $(OBJS)