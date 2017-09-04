static inline void i2o1_delay8(i2o1 *data){

	int32_t smp;
	smp = ((((SMP_RATE << 10) / 1000) * (*data->in1 >> 10)) >> 10) + 1;
	if(smp > 2999){ // delay time clip
		smp = 2999;
	}
	if(data->index > 2999){ // read loop
		data->index = 0;
	}
	smp = data->index - smp; // delay index
	if(smp < 0){
		smp = 3000 + smp;
	}
	if(*data->in0 > 0x1FF){ // input clip
		*data->in0 = 0x1FF;
	}else if(*data->in0 < -0x1FF){
		*data->in0 = -0x1FF;
	}
	data->data[data->index] = *data->in0 >> 2; // bit reduction
	data->out0 = data->data[smp] << 2;
	data->index++;

}