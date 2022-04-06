package typeconv

import "unsafe"

func StringToBytes(s string) []byte {
	strPtr := (*[2]uintptr)(unsafe.Pointer(&s))
	byteSlicePtr := [3]uintptr{strPtr[0], strPtr[1], strPtr[1]}
	return *(*[]byte)(unsafe.Pointer(&byteSlicePtr))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
