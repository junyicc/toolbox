package typeconv

import (
	"reflect"
	"runtime"
	"testing"
)

func TestStringToBytesWithGC(t *testing.T) {
	res := getHelloSlice()
	expectedBytes := []byte("hello")
	if !reflect.DeepEqual(res, expectedBytes) {
		t.Errorf("StringToBytes() = %v, want %v", res, expectedBytes)
	}
}

func getHelloSlice() []byte {
	// force gc
	defer runtime.GC()
	x := make([]byte, 5)
	x[0] = 'h'
	x[1] = 'e'
	x[2] = 'l'
	x[3] = 'l'
	x[4] = 'o'
	return StringToBytes(string(x))
}
