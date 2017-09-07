
#include <stdio.h>
#include <time.h>

#include "minunit_config.h"

int main(int argc, char* argv[]) {

    clock_t begin = clock();
    char *result = test_list();
    clock_t end = clock();
    double time_spent = (double)(end - begin) / CLOCKS_PER_SEC;

    if (result != 0) {
        printf("!!! ERROR !!!\n");
        printf("%s\n", result);
    }
    else {
        printf("ALL TESTS PASSED\n");
    }
    printf("Tests run: %d\n", tests_run);
    printf("Assertions made: %d\n", assertions_made);
    printf("Tests took: %f seconds\n", time_spent);

    return result != 0;
}