package core

import (
	"errors"
)

func readSimpleString(data []byte) (string, int, error) {
	pos := 1
	for ; data[pos] != '\r'; pos++ {
	}
	return string(data[1:pos]), pos + 2, nil
}

func readErrorString(data []byte) (string, int, error) {
	return readSimpleString(data)
}

func readInt64(data []byte) (int64, int, error) {
	pos := 1
	var value int64 = 0
	for ; data[pos] != '\r'; pos++ {
		value = value * 10 + int64(data[pos]-'0')
	}
	return value, pos + 2, nil
}

func readLength(data []byte) (int, int) {
     pos := 0
	 length := 0
	 
	 for pos = range data {
		if data[pos] >= '0' && data[pos] <= '9' {
			length = length * 10 + int(data[pos]-'0')
		} else {
			return length, pos + 2
		}
	 }
	 return 0,0
}

func readBulkString(data []byte) (string, int, error) {
	pos := 1
	length, delta := readLength(data[pos:])
    pos += delta
	return string(data[pos : (pos + length)]), pos + length + 2, nil
}
// "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n":  
func readArray(data []byte) ([]interface{}, int, error) {
	pos := 1
	count, delta := readLength(data[pos:])
    pos += delta
	var elems []interface{} = make([]interface{}, count)
	for i := 0; i < count; i++ {
	   elem, delta, err := decodeOne(data[pos:])
	   println(pos)
	   if err != nil {
		return nil, 0, err
	   }
	   elems[i] = elem
	   pos += delta
	}

	return elems, pos, nil

}

func decodeOne(data []byte) (interface{}, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}
	switch data[0] {
		case '+':
			return readSimpleString(data)
		case '-':
			return readErrorString(data)
		case ':':
			return readInt64(data)
		case '$': 
			return readBulkString(data)
		case '*':
			return readArray(data)
	}
	return nil, 0, nil
}

func Decode(data []byte) (interface{}, error) {	
   if len(data) == 0 {
	return nil, errors.New("no data")
   }
   value, _, err := decodeOne(data)

   return value, err
} 