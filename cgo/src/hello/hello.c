#include "hello.h"
#include <stdio.h>

void SayHello(const char* from, const char* to) {
    printf("%s: hello %s.\n", from, to);
}