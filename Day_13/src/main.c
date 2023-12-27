#include "include/parser.h"
#include <stdlib.h>
#include <stdio.h>

int main(int argc, char const *argv[])
{
    if (argc < 2)
    {
        printf("Missing input file\n");
        exit(1);
    }
    int total_patterns;
    Pattern **patterns = parse_patterns(argv[1], &total_patterns);
    for (int i = 0; i < total_patterns; i++)
    {
        printf("Pattern %d\n", i);
        Pattern *p = patterns[i];
        for (int j = 0; j < p->rows; j++)
        {
            printf("%s", p->data[j]);
        }
        printf("\n");
    }
    // TODO: free all memory
    return 0;
}
