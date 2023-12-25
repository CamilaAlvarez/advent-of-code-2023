#include "include/parser.h"
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

static int MAX_RECORDS = 1000;

Record *parse_records(const char *filename, size_t *number_records)
{
    FILE *input = fopen(filename, "r");
    if (input == NULL)
    {
        printf("Invalid input file: %s\n", filename);
        exit(1);
    }
    Record *records = (Record *)malloc(sizeof(Record) * MAX_RECORDS);
    int record_index = 0;
    char *line = NULL;
    size_t n = 0;
    size_t read;
    *number_records = 0;
    while ((read = getline(&line, &n, input)) != -1)
    {
        records[record_index].number_groups = 0;
        records[record_index].record_length = 0;
        bool reading_groups = false;
        for (int i = 0; i < read; i++)
        {
            char read_char = line[i];
            if (read_char == '\n')
            {
                break;
            }
            else if (read_char == ' ')
            {
                reading_groups = true;
            }
            else if (read_char == ',')
            {
                continue;
            }
            else if (reading_groups)
            {
                records[record_index].continuous_damaged_items[records[record_index].number_groups] = read_char - '0';
                records[record_index].number_groups++;
            }
            else
            {
                switch (read_char)
                {
                case '.':
                    records[record_index].condition_record[i] = OPERATIONAL;
                    break;
                case '#':
                    records[record_index].condition_record[i] = DAMAGED;
                    break;
                case '?':
                    records[record_index].condition_record[i] = UNKNOWN;
                    break;
                default:
                    printf("Invalid symbol: %c\n", read_char);
                    exit(2);
                    break;
                }
                records[record_index].record_length++;
            }
        }
        record_index++;
    }
    *number_records = record_index;
    return records;
}
