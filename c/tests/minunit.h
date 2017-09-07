//
// Created by Thomas Whiston on 07.09.17.
//
/*
* file: minunit.h
* http://www.jera.com/techinfo/jtns/jtn002.html#Source_Code
*/
int tests_run;
int assertions_made;
#define mu_assert(message, test) do { assertions_made++; if (!(test)) return message; } while (0)
#define mu_run_test(test) do { char *message = test(); tests_run++; \
                                if (message) return message; } while (0)

