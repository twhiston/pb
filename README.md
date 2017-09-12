# PB
[ ![Codeship Status for twhiston/pb](https://app.codeship.com/projects/0e3cf160-79cd-0135-bdc7-52a114685316/status?branch=master)](https://app.codeship.com/projects/245072)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/189f2cf2145c4629bc9cf55c79cf0c28)](https://www.codacy.com?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=twhiston/pb&amp;utm_campaign=Badge_Grade)

Patchblocks CLI helper tool for developing new blocks.

Moves block definitions into a more friendly .yml and .c file configuration and provides easy rendering into .xml patchblock files.

More to come...

## Features

- Generated readable .yml and .c files to describe and implement blocks
- Generate .xml from .yml and .c files
- Generate .yml and .c files from .xml file
- Render time expansion of common patchblock operation macros

## Installation

Download a binary for your platform from the release pages.
It has been build for MacOsX, Linux x86, Linux //TODO - FIX THIS
and place it in your path, usually something like `/usr/local/bin`

Windows users - I don't have a windows machine to do a build on or a way to test a cross compilation but it should work ok if you compile it yourself.
If you are a windows user who knows their golang please reach out.

Alternatively build the code from source
```
go get github.com/twhiston/pb
```

## Usage

`pb --help` gets info about the available commands and top level flags.
You can always get more information about a command by running `pb <subcommand> --help`

## Macros

If you want to encapsulate common functionality as macros to re-use across multiple blocks pb supports that via a simple inlining function.
Any files that you pass to render via the --macro flag will be added to your function before the main code.
This is technically valid as #defines are preprocessor code, but its extremely important to understand that any non preprocessor code will be
merged in before the body of your function, and therefore may affect behaviour. So don't do it, preprocessor defines only!
It would be nicer to do this in another way, but it's not possible within the block xml structure currently

Some example macros are supplied in the c folder in the source code and can be include in builds with the special --macro-inbuilt flag. These cover

#### Clip

CLIP_PVALUE(data_input)
clip an input value (input taken by pointer) see http://wiki.patchblocks.com/index.php/Programming_Custom_Blocks#Signals

#### Integer Conversion
http://wiki.patchblocks.com/index.php/Programming_Custom_Blocks#Converting_to_and_from_Integers

UINPUT_TO_INT(input)
input >> 10;

INPUT_TO_INT(input)
input / 1024;

INT_TO_OUTPUT(input)
input << 10;

### Example

A do nothing example that just shows the macros in use. Garbage in, Garbage out ;)
```
int32_t input = *data->in0;
CLIP_PVALUE(&input)
input = INPUT_TO_INT(input);
data->out0 = INT_TO_OUTPUT(input);
```

### Developing Macros

The macro file is kept in the c/macros folder of the source code, and as well as the macros contains a miniature test framework and test cases
to ensure the validity of the macros. These macros were developed using [CLion IDE](https://www.jetbrains.com/clion/) and the makefiles should work as is within their build system.

If you wish to contribute macros to the inbuilt library please contribute them to this copy of pb_macro.h and not the one in the assets folder.
All contributed macros should have an accompanying test in
c/macros/tests/minunit_config.h
and the test should be added in the test_list() function at the bottom of the file.
You should not need to alter any of the other test files.


## Roadmap

- Docker container for testing project golang and C files in CI
- Add loads more tests of functionality
- Docs
- Add more core macros (accepting submissions)
- Add versioning to release script (inject version info in file)
- Work out how to decrease reload iteration time
- Test framework
- Improved simulator debugging options