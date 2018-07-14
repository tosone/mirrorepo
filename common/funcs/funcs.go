package funcs

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"strings"
)

// Uint2byte uint to byte
func Uint2byte(num uint64) (ret []byte, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, num); err != nil {
		return
	}
	return buf.Bytes(), nil
}

// BucketName BucketName
func BucketName(val interface{}) string {
	valueOf := reflect.ValueOf(val)
	if valueOf.Type().Kind() == reflect.Ptr {
		return strings.ToLower(reflect.Indirect(valueOf).Type().Name())
	}
	return strings.ToLower(valueOf.Type().Name())
}
