#ifndef PARSER_H
#define PARSER_H
#include "include/patterns.h"

Pattern **parse_patterns(const char *filename, int *total_patterns);
#endif