#include <stdio.h>
#include <stdint.h>
#include "integer_bits.h"

void bitwise_shift_unsigned();
void bitwise_shift_signed();

int main(int argc, char const *argv[])
{
    bitwise_shift_unsigned();
    bitwise_shift_signed();

    return 0;
}

void bitwise_shift_unsigned()
{
    uint8_t u81 = 12;
    uint8_t u82 = 35;

    for (int i = 0; i < 9; i++)
    {
        uint8_t shift_result = u81 >> i;
        printf("%s >> %d = %s(%d)\n", integerBits(sizeof(u81), &u81), i, integerBits(sizeof(shift_result), &shift_result), shift_result);
    }
    printf("\n");

    for (int i = 0; i < 9; i++)
    {
        uint8_t shift_result = u82 << i;
        printf("%s << %d = %s(%d)\n", integerBits(sizeof(u82), &u82), i, integerBits(sizeof(shift_result), &shift_result), shift_result);
    }
    printf("\n");

    // $ gcc ./bitwise_shift.c ./integer_bits.c && ./a.out
    // 1100 >> 0 = 1100(12)
    // 1100 >> 1 = 110(6)
    // 1100 >> 2 = 11(3)
    // 1100 >> 3 = 1(1)
    // 1100 >> 4 = 0(0)
    // 1100 >> 5 = 0(0)
    // 1100 >> 6 = 0(0)
    // 1100 >> 7 = 0(0)
    // 1100 >> 8 = 0(0)
    //
    // 100011 << 0 = 100011(35)
    // 100011 << 1 = 1000110(70)
    // 100011 << 2 = 10001100(140)
    // 100011 << 3 = 11000(24)
    // 100011 << 4 = 110000(48)
    // 100011 << 5 = 1100000(96)
    // 100011 << 6 = 11000000(192)
    // 100011 << 7 = 10000000(128)
    // 100011 << 8 = 0(0)
}

void bitwise_shift_signed()
{
    int8_t i81 = -35;
    int8_t i82 = 35;

    for (int i = 0; i < 9; i++)
    {
        int8_t shift_result = i81 >> i;
        printf("%s >> %d = %s(%d)\n", integerBits(sizeof(i81), &i81), i, integerBits(sizeof(shift_result), &shift_result), shift_result);
    }
    printf("\n");

    for (int i = 0; i < 9; i++)
    {
        int8_t shift_result = i81 << i;
        printf("%s << %d = %s(%d)\n", integerBits(sizeof(i81), &i81), i, integerBits(sizeof(shift_result), &shift_result), shift_result);
    }
    printf("\n");

    for (int i = 0; i < 9; i++)
    {
        int8_t shift_result = i82 >> i;
        printf("%s >> %d = %s(%d)\n", integerBits(sizeof(i82), &i82), i, integerBits(sizeof(shift_result), &shift_result), shift_result);
    }
    printf("\n");

    for (int i = 0; i < 9; i++)
    {
        int8_t shift_result = i82 << i;
        printf("%s << %d = %s(%d)\n", integerBits(sizeof(i82), &i82), i, integerBits(sizeof(shift_result), &shift_result), shift_result);
    }

    // $ gcc ./bitwise_shift.c ./integer_bits.c && ./a.out
    // 11011101 >> 0 = 11011101(-35)
    // 11011101 >> 1 = 11101110(-18)
    // 11011101 >> 2 = 11110111(-9)
    // 11011101 >> 3 = 11111011(-5)
    // 11011101 >> 4 = 11111101(-3)
    // 11011101 >> 5 = 11111110(-2)
    // 11011101 >> 6 = 11111111(-1)
    // 11011101 >> 7 = 11111111(-1)
    // 11011101 >> 8 = 11111111(-1)
    //
    // 11011101 << 0 = 11011101(-35)
    // 11011101 << 1 = 10111010(-70)
    // 11011101 << 2 = 1110100(116)
    // 11011101 << 3 = 11101000(-24)
    // 11011101 << 4 = 11010000(-48)
    // 11011101 << 5 = 10100000(-96)
    // 11011101 << 6 = 1000000(64)
    // 11011101 << 7 = 10000000(-128)
    // 11011101 << 8 = 0(0)
    //
    // 100011 >> 0 = 100011(35)
    // 100011 >> 1 = 10001(17)
    // 100011 >> 2 = 1000(8)
    // 100011 >> 3 = 100(4)
    // 100011 >> 4 = 10(2)
    // 100011 >> 5 = 1(1)
    // 100011 >> 6 = 0(0)
    // 100011 >> 7 = 0(0)
    // 100011 >> 8 = 0(0)
    //
    // 100011 << 0 = 100011(35)
    // 100011 << 1 = 1000110(70)
    // 100011 << 2 = 10001100(-116)
    // 100011 << 3 = 11000(24)
    // 100011 << 4 = 110000(48)
    // 100011 << 5 = 1100000(96)
    // 100011 << 6 = 11000000(-64)
    // 100011 << 7 = 10000000(-128)
    // 100011 << 8 = 0(0)
}