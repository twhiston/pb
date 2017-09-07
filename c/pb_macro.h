/*
 * Clip a value to 10bits.
 * See http://wiki.patchblocks.com/index.php/Programming_Custom_Blocks#Signals
 * Works on a pointer and does not return a value
 */
#define CLIP_PVALUE(data_input)	\
    if(*data_input > 0x1FF){ *data_input = 0x1FF; } else if(*data_input < -0x1FF){ *data_input = -0x1FF; }

/*
 * Macros for Signal to Integer Conversion.
 * See http://wiki.patchblocks.com/index.php/Programming_Custom_Blocks#Converting_to_and_from_Integers
 */

/* Does not work on signed integers but is computationally cheaper */
#define UINPUT_TO_INT(input) input >> 10;
/* Works on signed and unsigned integers */
#define INPUT_TO_INT(input) input / 1024;
/* Shift left 10 places ready for output */
#define INT_TO_OUTPUT(input) input << 10;