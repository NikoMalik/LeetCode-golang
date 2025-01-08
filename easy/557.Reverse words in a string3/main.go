package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

func reverseWords(s string) string {
	_b := make([]byte, len(s))
	copy(_b, *(*[]byte)(unsafe.Pointer(&s)))
	_blen := len(_b)
	_bQ := bytes.LastIndexByte(_b, ' ')

	for _bQ >= 0 {
		reverseRange(_b, _bQ+1, _blen-1)
		_blen = _bQ
		_bQ = bytes.LastIndexByte(_b[:_blen], ' ')
	}

	reverseRange(_b, 0, _blen-1)

	return *(*string)(unsafe.Pointer(&_b))
}

func reverseRange(b []byte, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
func main() {
	word := "Let's take LeetCode contest"

	fmt.Println(reverseWords(word))
}
