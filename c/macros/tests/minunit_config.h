//
// Created by Thomas Whiston on 07.09.17.
//

#ifndef C_MINUNIT_APP_H
#define C_MINUNIT_APP_H

#include <stdio.h>
#include "minunit.h"
#include "../pb_macro.h"

static char *test_clip_max() {

    for (int i = 0; i <= 0x1FF; ++i) {
        int test = i;
        CLIP_PVALUE(&test);
        mu_assert("Input does not equal output before clipping point", test == i);
    }

    for (int i = 0x200; i < 0x400; ++i) {
        int test = i;
        CLIP_PVALUE(&test);
        mu_assert("Input does not clip at max", test == 0x1FF);
    }
    return 0;
}

static char *test_clip_min() {

    for (int i = 0; i <= -0x1FF; --i) {
        int test = i;
        CLIP_PVALUE(&test);
        mu_assert("Input does not equal output before clipping point", test == i);
    }

    for (int i = -0x200; i < -0x400; ++i) {
        int test = i;
        CLIP_PVALUE(&test);
        mu_assert("Input does not clip at min", test == -0x1FF);
    }
    return 0;
}

static char *test_int_to_output() {

    for (int i = -511; i < 512; ++i) {
        int test = INT_TO_OUTPUT(i);
        mu_assert("input not correctly converted out", test == (i << 10));
    }

    int val = 1024;
    CLIP_PVALUE(&val);
    int test = INT_TO_OUTPUT(val);
    mu_assert("input not correctly converted out", test == (511 << 10));

    val = -1024;
    CLIP_PVALUE(&val);
    test = INT_TO_OUTPUT(val);
    mu_assert("input not correctly converted out", test == (-511 << 10));

}

static char *test_list() {
    mu_run_test(test_clip_max);
    mu_run_test(test_clip_min);
    mu_run_test(test_int_to_output);
    return 0;
}

#endif //C_MINUNIT_APP_H
