#include "include/parser.h"
#include "include/records.h"
#include <stdio.h>
#include <stdlib.h>

static void print_record(Record *record)
{
    for (int i = 0; i < record->record_length; i++)
    {
        switch (record->condition_record[i])
        {
        case OPERATIONAL:
            printf(".");
            break;
        case DAMAGED:
            printf("#");
            break;
        case UNKNOWN:
            printf("?");
            break;
        default:
            break;
        }
    }
    printf("\n");
}

int main(int argc, char const *argv[])
{
    if (argc < 2)
    {
        printf("No input file given\n");
        exit(1);
    }
    size_t number_records;
    unsigned long valid_arrangements = 0;
    Record *records = parse_records(argv[1], &number_records);
    printf("Number records: %zu\n", number_records);
    for (int i = 0; i < number_records; i++)
    {
        int record_length = records[i].record_length;
        int number_groups = records[i].number_groups;
        print_record(&records[i]);
        // We try two different options:
        // Normal input:
        int initial_arrangements = number_different_arrangements(&records[i]);
        printf("Arrangements initial input: %d\n", initial_arrangements);
        // Two instances of the initial input:
        records[i].record_length = 2 * records[i].record_length + 1;
        records[i].condition_record[record_length] = UNKNOWN;
        for (int j = 0; j < record_length; j++)
        {
            records[i].condition_record[record_length + 1 + j] = records[i].condition_record[j];
        }
        records[i].number_groups = 2 * records[i].number_groups;
        for (int j = 0; j < number_groups; j++)
        {
            records[i].continuous_damaged_items[number_groups + j] = records[i].continuous_damaged_items[j];
        }
        print_record(&records[i]);
        for (int j = 0; j < records[i].number_groups; j++)
        {
            printf("%d\n", records[i].continuous_damaged_items[j]);
        }
        int double_input_arrangements = number_different_arrangements(&records[i]);
        printf("Arrangements with double input: %d\n", double_input_arrangements);
        // we divide the results to find the pattern (p = r2/r1)
        // then we do: r1 * p * p * p * p
        unsigned long increase_ratio = (unsigned long)double_input_arrangements / initial_arrangements;
        printf("ratio: %lu\n", increase_ratio);
        unsigned long arrangements = initial_arrangements * increase_ratio * increase_ratio * increase_ratio * increase_ratio;
        printf("Valid arrangements: %lu\n", arrangements);
        valid_arrangements += arrangements;
    }
    printf("Sum of valid arrangements: %lu\n", valid_arrangements);
    return 0;
}
