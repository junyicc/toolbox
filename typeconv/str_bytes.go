package typeconv

import (
	"reflect"
	"unsafe"
)

func UnsafeStringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func UnsafeBytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// Deprecated: cannot return determined []byte when gc happens
// func StringToBytes(s string) []byte {
// 	strPtr := (*[2]uintptr)(unsafe.Pointer(&s))
// 	byteSlicePtr := [3]uintptr{strPtr[0], strPtr[1], strPtr[1]}
// 	return *(*[]byte)(unsafe.Pointer(&byteSlicePtr))
// }

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts string to byte slice without a memory allocation, from Gin
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// ForceStrigToBytes converts string to byte by type conversion
func ForceStringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// ReflectStringToBytes converts string to byte slice by reflect.SliceHeader & reflect.StringHeader
func ReflectStringToBytes(s string) []byte {
	var b []byte
	l := len(s)
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	p.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	p.Len = l
	p.Cap = l
	return b
}
