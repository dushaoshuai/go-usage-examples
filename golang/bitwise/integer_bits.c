// Assumes little endian
#include <stdio.h>
#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "integer_bits.h"

// https://stackoverflow.com/a/3974138
char *integerBits(size_t const size, void const *const ptr)
{
    unsigned char *b = (unsigned char *)ptr;
    unsigned char byte;
    int i, j;

    size_t str_size = size * 8 + 1;
    char *bit_string = (char *)malloc(str_size);
    if (bit_string == NULL)
    {
        return NULL;
    }

    memset(bit_string, 0, str_size);
    size_t index = 0;

    for (i = size - 1; i >= 0; i--)
    {
        for (j = 7; j >= 0; j--)
        {
            byte = (b[i] >> j) & 1;
            if (index != 0 || byte != 0 || (i == 0 && j == 0 && index == 0))
            {
                bit_string[index++] = byte + '0';
            }
        }
    }
    bit_string[index] = '\0';

    return bit_string;
}
