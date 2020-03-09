package gaw

import (
	"net/http"
)

//GetHeaderSize returns the size in bytes of the given header
func GetHeaderSize(headers http.Header) uint32 {
	var size uint32
	for k, v := range headers {
		size += uint32(len(k))
		for _, val := range v {
			size += uint32(len(val))
		}
	}
	return size
}
