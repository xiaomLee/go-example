// 表示该函数的链接符号遵循C语言的规则。
extern "C" {
    #include "say.h"
}
#include <iostream>

void SaySomething(const char* sb, const char* sth) {
    std::cout << sb;
    std::cout << ": ";
    std::cout << sth;
    std::cout << "\n";
}

int main(void) {
    SaySomething("Jack", "my name is Jack.");
    return 0;
}