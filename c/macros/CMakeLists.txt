cmake_minimum_required(VERSION 3.8)
project(pbmacro C)

##set(CMAKE_CXX_STANDARD 11)

set(SOURCE_FILES
        tests/minunit_runner.c
       # pb_macro.h
       # pb_macro_test.h
        tests/minunit_config.h tests/minunit.h)

add_executable(pbmacro_test ${SOURCE_FILES})