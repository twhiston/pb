#define CLIP_PVALUE(data_input)	\
    if(*data_input > 0x1FF){ \
        *data_input = 0x1FF; \
	} else if(*data_input < -0x1FF){ \
		*data_input = -0x1FF; \
    }

#define UINPUT_TO_INT(input) \
    input >> 10;

#define INPUT_TO_INT(input) \
    input / 1024;

#define INT_TO_OUTPUT(input) \
    input << 10;