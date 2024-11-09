#include <stdio.h>
#include <stdint.h>
#include "integer_bits.h"

void mask();
void turning_bits_on();

// $ gcc ./usage.c ./integer_bits.c && ./a.out
int main(int argc, char const *argv[])
{
    // 10101111 & 11010011 = 10000011
    // 1111 & 11010011 = 11
    // 11111000 & 11010011 = 11010000
    mask();

	// 10101111 | 11010011 = 11111111
    // 1111 | 11010011 = 11011111
    // 11111000 | 11010011 = 11111011
    turning_bits_on();

    return 0;
}

void mask()
{
    uint8_t mask = 0b11010011;

    uint8_t flags1 = 0b10101111;
    uint8_t v1 = flags1 & mask;

    uint8_t flags2 = 0b1111;
    uint8_t v2 = flags2 & mask;

    uint8_t flags3 = 0b11111000;
    uint8_t v3 = flags3 & mask;

    printf("%s & %s = %s\n", integerBits(sizeof(flags1), &flags1), integerBits(sizeof(mask), &mask), integerBits(sizeof(v1), &v1));
    printf("%s & %s = %s\n", integerBits(sizeof(flags2), &flags1), integerBits(sizeof(mask), &mask), integerBits(sizeof(v2), &v2));
    printf("%s & %s = %s\n", integerBits(sizeof(flags3), &flags1), integerBits(sizeof(mask), &mask), integerBits(sizeof(v3), &v3));
}

void turning_bits_on()
{
    uint8_t mask = 0b11010011;

    uint8_t flags1 = 0b10101111;
    uint8_t v1 = flags1 | mask;

    uint8_t flags2 = 0b1111;
    uint8_t v2 = flags2 | mask;

    uint8_t flags3 = 0b11111000;
    uint8_t v3 = flags3 | mask;

    printf("%s | %s = %s\n", integerBits(sizeof(flags1), &flags1), integerBits(sizeof(mask), &mask), integerBits(sizeof(v1), &v1));
    printf("%s | %s = %s\n", integerBits(sizeof(flags2), &flags1), integerBits(sizeof(mask), &mask), integerBits(sizeof(v2), &v2));
    printf("%s | %s = %s\n", integerBits(sizeof(flags3), &flags1), integerBits(sizeof(mask), &mask), integerBits(sizeof(v3), &v3));
}

