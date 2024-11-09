#include <stdio.h>
#include <stdint.h>
#include "integer_bits.h"

void bitwise_shift_unsigned();

int main(int argc, char const *argv[])
{
    bitwise_shift_unsigned();

    return 0;
}

void bitwise_shift_unsigned()
{
    uint8_t u81 = 12;
    uint8_t u82 = 35;

    for (int i = 0; i < 8; i++)
    {
        uint8_t shift_result = u81 >> i;
        printf("%s >> %d = %s\n", integerBits(sizeof(u81), &u81), i, integerBits(sizeof(shift_result), &shift_result));
    }
    for (int i = 0; i < 8; i++)
    {
        uint8_t shift_result = u82 << i;
        printf("%s << %d = %s\n", integerBits(sizeof(u82), &u82), i, integerBits(sizeof(shift_result), &shift_result));
    }

    // $ gcc ./bitwise_shift.c ./integer_bits.c && ./a.out
    // 1100 >> 0 = 1100
    // 1100 >> 1 = 110
    // 1100 >> 2 = 11
    // 1100 >> 3 = 1
    // 1100 >> 4 = 0
    // 1100 >> 5 = 0
    // 1100 >> 6 = 0
    // 1100 >> 7 = 0
    // 100011 << 0 = 100011
    // 100011 << 1 = 1000110
    // 100011 << 2 = 10001100
    // 100011 << 3 = 11000
    // 100011 << 4 = 110000
    // 100011 << 5 = 1100000
    // 100011 << 6 = 11000000
    // 100011 << 7 = 10000000
}