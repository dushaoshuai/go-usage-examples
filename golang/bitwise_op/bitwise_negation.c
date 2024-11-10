#include <stdio.h>
#include <stdint.h>
#include "integer_bits.h"

int main(int argc, char const *argv[])
{
    uint8_t u8 = 12;
    uint8_t negation_u8 = ~u8;

    int8_t positiveI8 = 12;
    int8_t negation_positiveI8 = ~positiveI8;

    int8_t negativeI8 = -12;
    int8_t negation_negativeI8 = ~negativeI8;

    printf("uint8 = %s, ~uint8 = %s\n", integerBits(sizeof(u8), &u8), integerBits(sizeof(negation_u8), &negation_u8));
    printf("positiveI8 = %s, ~positiveI8 = %s\n", integerBits(sizeof(positiveI8), &positiveI8), integerBits(sizeof(negation_positiveI8), &negation_positiveI8));
    printf("negativeI8 = %s, ~negativeI8 = %s\n", integerBits(sizeof(negativeI8), &negativeI8), integerBits(sizeof(negation_negativeI8), &negation_negativeI8));

    // $ gcc ./bitwise_negation.c ./integer_bits.c && ./a.out
    // uint8 = 1100, ~uint8 = 11110011
    // positiveI8 = 1100, ~positiveI8 = 11110011
    // negativeI8 = 11110100, ~negativeI8 = 1011

    return 0;
}
