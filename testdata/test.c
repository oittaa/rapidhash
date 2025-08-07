/*
 * Test program for the original rapidhash C implementation.
 * Computes hashes for predefined test inputs and prints them.
 * 
 * wget https://raw.githubusercontent.com/Nicoshev/rapidhash/bc4b4baa48a15ff52ff4725e1ccdcda62815221c/rapidhash.h
 * gcc -O3 -o test test.c && ./test
 */

#include <inttypes.h> 
#include <stdio.h>
#include <string.h>
#include "rapidhash.h"

int main() {
    const char* tests[] = {
        "",
        "a",
        "abc",
        "hello",
        "message digest",
        "abcdefghijklmnopqrstuvwxyz",
        "Test_string_with_various_chars_and_length_of_48!",
        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
        "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
        "#This is a sample string for testing purposes. It has been created to contain one hundred and twelve characters.",
        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
    };
    size_t num = sizeof(tests) / sizeof(char*);
    for (size_t j = 0; j < num; j++) {
        if (j > 0) {
            printf("###\n");
        }
        const char* key = tests[j];
        size_t len = strlen(key);
        uint64_t h1 = rapidhash(key, len);
        uint64_t h2 = rapidhashMicro(key, len);
        uint64_t h3 = rapidhashNano(key, len);
        printf("{[]byte(\"%s\"), 0x%" PRIx64 "}, // Rapidhash\n", key, h1);
        printf("{[]byte(\"%s\"), 0x%" PRIx64 "}, // RapidhashMicro\n", key, h2);
        printf("{[]byte(\"%s\"), 0x%" PRIx64 "}, // RapidhashNano\n", key, h3);
        
    }
    const uint64_t seed = 42069;
    for (size_t j = 0; j < num; j++) {
        printf("###\n");
        const char* key = tests[j];
        size_t len = strlen(key);
        uint64_t h1 = rapidhash_withSeed(key, len, seed);
        uint64_t h2 = rapidhashMicro_withSeed(key, len, seed);
        uint64_t h3 = rapidhashNano_withSeed(key, len, seed);
        printf("{[]byte(\"%s\"), %" PRIu64 ", 0x%" PRIx64 "}, // RapidhashWithSeed\n", key, seed, h1);
        printf("{[]byte(\"%s\"), %" PRIu64 ", 0x%" PRIx64 "}, // RapidhashMicroWithSeed\n", key, seed, h2);
        printf("{[]byte(\"%s\"), %" PRIu64 ", 0x%" PRIx64 "}, // RapidhashNanoWithSeed\n", key, seed, h3);
    }
    return 0;
}
