package ex00

import (
	"errors"
	"unsafe"
)

func GetElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty slice")
	}
	if idx < 0 {
		return 0, errors.New("negative idx")
	}
	root := unsafe.Pointer(&arr[0])

	for i := 0; i < len(arr); i++ {
		if i == idx {
			res := (*int)(unsafe.Pointer(uintptr(root) + uintptr(uintptr(i)*unsafe.Sizeof(root))))
			return *res, nil
		}
	}
	return 0, errors.New("out of range")
}