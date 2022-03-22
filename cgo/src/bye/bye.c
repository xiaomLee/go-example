#include <stdio.h>
#include <string.h>
#include "bye.h"

void SayBye(char* from, char* to) {
    //printf("call go print.(print by bye.c)\n");
    GoPrint(from, to);
}

int main() {
    char from[] = "Tom";
    char to[] = "Jack";
    SayBye(from, to);
    return 0;
}