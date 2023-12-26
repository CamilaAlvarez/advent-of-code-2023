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
    int valid_arrangements = 0;
    Record *records = parse_records(argv[1], &number_records);
    printf("Number records: %zu\n", number_records);
    for (int i = 0; i < number_records; i++)
    {
        int record_length = records[i].record_length;
        records[i].record_length++;
        // We try two different options:
        // Add a ? at the end:
        records[i]
            .condition_record[record_length] = UNKNOWN;
        int un_at_the_end = number_different_arrangements(&records[i]);
        printf("Arrangements with ? at the end: %d\n", un_at_the_end);
        // Add a ? at the start:
        for (int j = record_length - 1; j >= 0; j--)
        {
            records[i].condition_record[j + 1] = records[i].condition_record[j];
        }
        records[i].condition_record[0] = UNKNOWN;
        int un_at_the_start = number_different_arrangements(&records[i]);
        printf("Arrangements with ? at the start: %d\n", un_at_the_start);
        int arrangements;
        if (un_at_the_end > un_at_the_start)
        {
            arrangements = un_at_the_end * un_at_the_end * un_at_the_end * un_at_the_end * un_at_the_start;
        }
        else
        {
            arrangements = un_at_the_start * un_at_the_start * un_at_the_start * un_at_the_start * un_at_the_end;
        }
        printf("Valid arrangements: %d\n", arrangements);
        valid_arrangements += arrangements;
    }
    printf("Sum of valid arrangements: %d\n", valid_arrangements);
    return 0;
}
