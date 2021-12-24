// 表示该函数的链接符号遵循C语言的规则。
extern "C" {
    #include "any.h"
}
#include <iostream>

void SaySomething(const char* s) {
    std::cout << s;
    std::cout << "(print by any.cpp)";
    std::cout << "\n";
}

int main(void) {
    SaySomething("my name is Jack.");
    return 0;
}