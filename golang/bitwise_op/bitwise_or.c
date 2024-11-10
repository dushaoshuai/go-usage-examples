#include <stdio.h>
#include <stdint.h>
#include "integer_bits.h"

int main(int argc, char const *argv[])
{
    uint8_t u81 = 12;
    uint8_t u82 = 35;
    uint8_t u8Or = u81 | u82;

    int8_t positiveI8 = 12;
    int8_t negativeI8 = -35;
    int8_t i8Or = positiveI8 | negativeI8;

    printf("%s | %s = %s\n", integerBits(sizeof(u81), &u81), integerBits(sizeof(u82), &u82), integerBits(sizeof(u8Or), &u8Or));
    printf("%s | %s = %s\n", integerBits(sizeof(positiveI8), &positiveI8), integerBits(sizeof(negativeI8), &negativeI8), integerBits(sizeof(i8Or), &i8Or));

    // $ gcc ./bitwise_or.c ./integer_bits.c && ./a.out
	// 1100 | 100011 = 101111
	// 1100 | 11011101 = 11011101

    return 0;
}
