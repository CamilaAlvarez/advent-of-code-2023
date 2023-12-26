#ifndef PARSER_H
#define PARSER_H
#include "include/records.h"
#include <sys/types.h>

Record *parse_records(const char *filename, size_t *number_records);
#endif