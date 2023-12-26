#include "include/parser.h"
#include "include/records.h"
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
    int valid_arrangements = 0;
    Record *records = parse_records(argv[1], &number_records);
    printf("Number records: %zu\n", number_records);
    for (int i = 0; i < number_records; i++)
    {
        valid_arrangements += number_different_arrangements(&records[i]);
    }
    printf("Sum of valid arrangements: %d\n", valid_arrangements);
    return 0;
}
