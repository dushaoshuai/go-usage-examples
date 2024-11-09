#include <stdio.h>
#include <stdint.h>
#include "integer_bits.h"

void mask();
void turning_bits_on();
void turning_bits_off();
void toggling_bits();
void checking_the_value_of_a_bit();

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

    // 10101111 & ~11010011 = 101100
    // 1111 & ~11010011 = 1100
    // 11111000 & ~11010011 = 101000
    turning_bits_off();

    // 10101111 ^ 11111111 = 1010000
    // 1111 ^ 11111111 = 11110000
    // 11111000 ^ 11111111 = 111
    toggling_bits();

    // 10101111
    // 1 is at position 0
    // 1 is at position 1
    // 1 is at position 2
    // 1 is at position 3
    // 0 is at position 4
    // 1 is at position 5
    // 0 is at position 6
    // 1 is at position 7
    checking_the_value_of_a_bit();

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
    printf("%s & %s = %s\n", integerBits(sizeof(flags2), &flags2), integerBits(sizeof(mask), &mask), integerBits(sizeof(v2), &v2));
    printf("%s & %s = %s\n", integerBits(sizeof(flags3), &flags3), integerBits(sizeof(mask), &mask), integerBits(sizeof(v3), &v3));
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
    printf("%s | %s = %s\n", integerBits(sizeof(flags2), &flags2), integerBits(sizeof(mask), &mask), integerBits(sizeof(v2), &v2));
    printf("%s | %s = %s\n", integerBits(sizeof(flags3), &flags3), integerBits(sizeof(mask), &mask), integerBits(sizeof(v3), &v3));
}

// clearing bits
// bit clear
void turning_bits_off()
{
    uint8_t mask = 0b11010011;

    uint8_t flags1 = 0b10101111;
    uint8_t v1 = flags1 & ~mask;

    uint8_t flags2 = 0b1111;
    uint8_t v2 = flags2 & ~mask;

    uint8_t flags3 = 0b11111000;
    uint8_t v3 = flags3 & ~mask;

    printf("%s & ~%s = %s\n", integerBits(sizeof(flags1), &flags1), integerBits(sizeof(mask), &mask), integerBits(sizeof(v1), &v1));
    printf("%s & ~%s = %s\n", integerBits(sizeof(flags2), &flags2), integerBits(sizeof(mask), &mask), integerBits(sizeof(v2), &v2));
    printf("%s & ~%s = %s\n", integerBits(sizeof(flags3), &flags3), integerBits(sizeof(mask), &mask), integerBits(sizeof(v3), &v3));
}

void toggling_bits()
{
    uint8_t mask = 0b11111111;

    uint8_t flags1 = 0b10101111;
    uint8_t v1 = flags1 ^ mask;

    uint8_t flags2 = 0b1111;
    uint8_t v2 = flags2 ^ mask;

    uint8_t flags3 = 0b11111000;
    uint8_t v3 = flags3 ^ mask;

    printf("%s ^ %s = %s\n", integerBits(sizeof(flags1), &flags1), integerBits(sizeof(mask), &mask), integerBits(sizeof(v1), &v1));
    printf("%s ^ %s = %s\n", integerBits(sizeof(flags2), &flags2), integerBits(sizeof(mask), &mask), integerBits(sizeof(v2), &v2));
    printf("%s ^ %s = %s\n", integerBits(sizeof(flags3), &flags3), integerBits(sizeof(mask), &mask), integerBits(sizeof(v3), &v3));
}

void checking_the_value_of_a_bit()
{
    uint8_t flags = 0b10101111;

    printf("%s\n", integerBits(sizeof(flags), &flags));

    for (int i = 0; i < 8; i++)
    {
        if ((flags & (1 << i)) != 0)
        {
            printf("%d is at position %d\n", 1, i);
        }
        else
        {
            printf("%d is at position %d\n", 0, i);
        }
    }
}