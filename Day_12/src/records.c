#include "include/records.h"
#include <strings.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

#define MAX_INVALID_SEQUENCE 4096
#define SEQ_LEN 256

static bool is_valid(char *sequence, int sequence_length, Record *record, char invalid_sequences[MAX_INVALID_SEQUENCE][SEQ_LEN], int number_invalid)
{
    for (int i = 0; i < number_invalid; i++)
    {
        // We already check that this case is invalid
        if (strcmp(sequence, invalid_sequences[i]) == 0)
        {
            return false;
        }
    }

    int damaged_seq_len = 0;
    int damaged_group_index = 0;
    bool seen_damaged = false;
    for (int i = 0; i < sequence_length; i++)
    {
        char curr_char = sequence[i];
        if (!seen_damaged && curr_char == '#')
        {
            seen_damaged = true;
            damaged_seq_len++;
        }
        else if (seen_damaged && curr_char == '#')
        {
            damaged_seq_len++;
        }
        else if (seen_damaged && curr_char == '.')
        {
            if (damaged_group_index >= record->number_groups)
            {
                // There are more groups than there should be
                return false;
            }
            if (damaged_seq_len != record->continuous_damaged_items[damaged_group_index])
            {
                return false;
            }
            damaged_group_index++;
            seen_damaged = false;
            damaged_seq_len = 0;
        }
    }
    // If we're at the end of the string and we couldn't check the last damaged section
    if (sequence_length == record->record_length && seen_damaged)
    {
        if (damaged_group_index >= record->number_groups)
        {
            // There are more groups than there should be
            return false;
        }
        if (damaged_seq_len != record->continuous_damaged_items[damaged_group_index])
        {
            return false;
        }
        damaged_group_index++;
    }
    if (sequence_length == record->record_length && damaged_group_index < record->number_groups)
    {
        // We didn't see all groups even though we look over the whole string
        // If we had seen all groups our indexed would have reached the number of groups
        return false;
    }
    return true;
}

static void recurse_number_different_arrangements(Record record, char invalid_sequences[MAX_INVALID_SEQUENCE][SEQ_LEN], int *invalid_index, int *valid_arrangements)
{
    int record_length = record.record_length;
    char temp_sequence[SEQ_LEN], current_sequence[SEQ_LEN] = "\0";
    char *ptemp_sequence = temp_sequence;
    char *pcurrent_sequence = current_sequence;
    for (int i = 0; i < record_length; i++)
    {
        switch (record.condition_record[i])
        {
        case OPERATIONAL:
            strcat(pcurrent_sequence, ".");
            break;
        case DAMAGED:
            strcat(pcurrent_sequence, "#");
            break;
        case UNKNOWN:
            // When we find a ? we try to options, but stop the iteration.
            // The inner calls to the function will check the innermost ? symbols
            if (!is_valid(pcurrent_sequence, i, &record, invalid_sequences, *invalid_index))
            {
                return;
            }
            // We need to try with . and #
            // First .:
            strcpy(ptemp_sequence, pcurrent_sequence);
            strcat(ptemp_sequence, ".");
            // If it is valid we can continue down this path
            if (is_valid(ptemp_sequence, i + 1, &record, invalid_sequences, *invalid_index))
            {
                record.condition_record[i] = OPERATIONAL;
                recurse_number_different_arrangements(record, invalid_sequences, invalid_index, valid_arrangements);
            }
            else
            {
                strcpy(invalid_sequences[*invalid_index], ptemp_sequence);
                (*invalid_index)++;
            }
            // Second #:
            strcpy(ptemp_sequence, pcurrent_sequence);
            strcat(ptemp_sequence, "#");
            if (is_valid(ptemp_sequence, i + 1, &record, invalid_sequences, *invalid_index))
            {
                record.condition_record[i] = DAMAGED;
                recurse_number_different_arrangements(record, invalid_sequences, invalid_index, valid_arrangements);
            }
            else
            {
                strcpy(invalid_sequences[*invalid_index], ptemp_sequence);
                (*invalid_index)++;
            }
            // We take a look at every ? one by one
            return;
        default:
            printf("Invalid symbol\n");
            exit(3);
        }
    }
    if (strlen(pcurrent_sequence) == record_length && is_valid(pcurrent_sequence, record_length, &record, invalid_sequences, *invalid_index))
    {
        (*valid_arrangements)++;
        printf("Valid arrangement: %s\n", pcurrent_sequence);
    }
    else if (!is_valid(pcurrent_sequence, record_length, &record, invalid_sequences, *invalid_index))
    {
        strcpy(invalid_sequences[*invalid_index], pcurrent_sequence);
        (*invalid_index)++;
    }
}
int number_different_arrangements(Record *record)
{
    int valid_arrangements = 0;
    char invalid_sequences[MAX_INVALID_SEQUENCE][SEQ_LEN];
    int invalid_index = 0;
    recurse_number_different_arrangements(*record, invalid_sequences, &invalid_index, &valid_arrangements);
    return valid_arrangements;
}