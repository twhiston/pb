cmake_minimum_required(VERSION 3.8)
project(pbmacro C)

set(SOURCE_FILES
        tests/minunit_runner.c
        tests/minunit_config.h
        tests/minunit.h)

add_executable(pbmacro_test ${SOURCE_FILES})