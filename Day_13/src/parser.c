#include "include/parser.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <strings.h>

#define MAX_PATTERNS 1000
Pattern **parse_patterns(const char *filename, int *total_patterns)
{
    Pattern *patterns[MAX_PATTERNS];
    int number_patterns = 0;
    FILE *file = fopen(filename, "r");
    if (file == NULL)
    {
        printf("Couldn't open file with filename: %s\n", filename);
        exit(2);
    }
    size_t n;
    size_t read = 0;
    char *line = NULL;
    bool first_line_in_pattern = true;
    int rows = 0;
    int cols;
    char *pattern_data[MAX_PATTERNS];
    while ((read = getline(&line, &n, file)) != -1)
    {
        // space between patterns
        if (strcmp(line, "\n") == 0)
        {
            Pattern *p = (Pattern *)malloc(sizeof(Pattern));
            p->cols = cols;
            p->rows = rows;
            p->data = (char **)malloc(sizeof(char **) * rows);
            for (int i = 0; i < rows; i++)
            {
                p->data[i] = pattern_data[i];
            }
            patterns[number_patterns] = p;
            first_line_in_pattern = true;
            number_patterns++;
            rows = 0;
            continue;
        }
        if (first_line_in_pattern)
        {
            cols = read;
            if (line[read - 1] == '\n')
            {
                cols--;
            }
        }
        char *str = (char *)malloc(sizeof(char) * cols);
        // I'm including the \n
        strcpy(str, line);
        pattern_data[rows] = str;
        rows++;
    }
    if (rows != 0)
    {
        Pattern *p = (Pattern *)malloc(sizeof(Pattern));
        p->cols = cols;
        p->rows = rows;
        p->data = (char **)malloc(sizeof(char *) * rows);
        for (int i = 0; i < rows; i++)
        {
            p->data[i] = pattern_data[i];
        }
        patterns[number_patterns] = p;
        number_patterns++;
    }
    Pattern **ppatterns = (Pattern **)malloc(sizeof(Pattern *) * number_patterns);
    for (int i = 0; i < number_patterns; i++)
    {
        ppatterns[i] = patterns[i];
    }
    *total_patterns = number_patterns;
    return ppatterns;
}