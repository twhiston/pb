# PB
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/189f2cf2145c4629bc9cf55c79cf0c28)](https://www.codacy.com?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=twhiston/pb&amp;utm_campaign=Badge_Grade)

Patchblocks CLI helper tool for developing new blocks

## Features

- Generated readable .yml and .c files to describe and implement blocks with
- Generate .xml from .yml and .c files
- Generate .yml and .c files from .xml file

## Macros

If you want to encapsulate common functionality as macros to re-use across multiple blocks pb supports that via a simple inlining function.
Any files that you pass to render via the --macro flag will be added to your function before the main code.
This is technically valid as #defines are preprocessor code, but its extremely important to understand that any non preprocessor code will be
merged in before the body of your function, and therefore may affect behaviour. So don't do it, preprocessor defines only!

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

## Roadmap

- Render time expansion of common patchblock operation macros
- Add more core macros (accepting submissions)
- Docker container for testing project golang and C files in CI
- Work out how to increase reload iteration time
- Test framework
- Improved simulator debugging options