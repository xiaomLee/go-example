#include "bye.h"
#include <stdio.h>
#include <string.h>

void SayBye(const char* s) {
    printf("call go print.(print by bye.c)\n");
    int length = strlen(s);
    GoString name = {s, length};
    GoPrint(name);
}

int main() {
    char name[] = "Tom";
    int length = strlen(name);
    SayBye(name);
    return 0;
}