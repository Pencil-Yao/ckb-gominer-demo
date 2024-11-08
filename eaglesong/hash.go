package eaglesong

/*
void EaglesongHash(unsigned char *output, const unsigned char *input, unsigned int input_length);
#include "src/eaglesong.c"
*/
import "C"
import "unsafe"

func EaglesongHash(data []byte) []byte {
	output := make([]byte, 32)
	length := len(data)
	C.EaglesongHash((*C.uchar)(unsafe.Pointer(&output[0])), (*C.uchar)(unsafe.Pointer(&data[0])), C.uint(length))
	return output
}
