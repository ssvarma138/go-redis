package core_test

import (
	"fmt"
	"testing"
	"go-redis/core"
)
func TestSimpleStringDecode(t *testing.T) {
	cases := map[string]string{
		"+OK\r\n": "OK",
	}
	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}
func TestError(t *testing.T) {
	cases := map[string]string{
		"-Error message\r\n": "Error message",
	}
	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}
func TestInt64(t *testing.T) {
	cases := map[string]int64{
		":0\r\n":    0,
		":1000\r\n": 1000,
	}
	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}
func TestBulkStringDecode(t *testing.T) {
	cases := map[string]string{
		"$5\r\nhello\r\n": "hello",
		"$0\r\n\r\n":      "",
	}
	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}
func TestArrayDecode(t *testing.T) {
	cases := map[string][]interface{}{
		"*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n":                     {"hello", "world"},
	}
	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		array := value.([]interface{})

		if len(array) != len(v) {
			t.Fail()
		}
		for i := range array {
			fmt.Println(array...)
			if fmt.Sprintf("%v", v[i]) != fmt.Sprintf("%v", array[i]) {
				t.Fail()
			}
		}
	}
}