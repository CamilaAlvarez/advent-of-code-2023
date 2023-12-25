#include "include/parser.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[])
{
    if (argc < 2)
    {
        printf("No input file given\n");
        exit(1);
    }
    size_t number_records;
    Record *records = parse_records(argv[1], &number_records);
    printf("Number records: %zu\n", number_records);
    return 0;
}
