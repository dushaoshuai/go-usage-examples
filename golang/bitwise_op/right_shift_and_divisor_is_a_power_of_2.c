#include <stdio.h>
#include <stdint.h>

int main(int argc, char const *argv[])
{
    uint8_t u8 = 77;
    int8_t i81 = 77;
	int8_t i82 = -77;

	printf("%d >> 1 = %d\n", u8, u8>>1);
	printf("%d / 2 = %d\n", u8, u8/2);

	printf("%d >> 1 = %d\n", i81, i81>>1);
	printf("%d / 2 = %d\n", i81, i81/2);

	printf("%d >> 1 = %d\n", i82, i82>>1);
	printf("%d / 2 = %d\n", i82, i82/2);

	// gcc ./right_shift_and_divisor_is_a_power_of_2.c && ./a.out
	// 77 >> 1 = 38
	// 77 / 2 = 38
	// 77 >> 1 = 38
	// 77 / 2 = 38
	// -77 >> 1 = -39
	// -77 / 2 = -38

    return 0;
}
