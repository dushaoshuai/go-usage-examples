#include <stdio.h>
#include <stdint.h>
#include "integer_bits.h"

int main(int argc, char const *argv[])
{
    uint8_t u81 = 12;
    uint8_t u82 = 35;
    uint8_t u8And = u81 & u82;

    int8_t positiveI8 = 12;
    int8_t negativeI8 = -35;
    int8_t i8And = positiveI8 & negativeI8;

    printf("%s & %s = %s\n", integerBits(sizeof(u81), &u81), integerBits(sizeof(u82), &u82), integerBits(sizeof(u8And), &u8And));
    printf("%s & %s = %s\n", integerBits(sizeof(positiveI8), &positiveI8), integerBits(sizeof(negativeI8), &negativeI8), integerBits(sizeof(i8And), &i8And));

    // $ gcc ./bitwise_and.c ./integer_bits.c && ./a.out
    // 1100 & 100011 = 0
    // 1100 & 11011101 = 1100

    return 0;
}
