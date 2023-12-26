#ifndef RECORDS_H
#define RECORDS_H
typedef enum RecordType
{
    OPERATIONAL,
    DAMAGED,
    UNKNOWN
} RecordType;
typedef struct Record
{
    int number_groups;
    int record_length;
    RecordType condition_record[256];
    int continuous_damaged_items[256];
} Record;
int number_different_arrangements(Record *record);
#endif