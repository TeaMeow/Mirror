package mirror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// structType 是一個測試用的建構體。
type structType struct {
	stringField string
	intField    int
}

// any 是存放任意值用的暫存變數。
var any interface{}

func TestInt(t *testing.T) {
	assert := assert.New(t)
	var i int
	any = 123456
	Cast(any, &i)
	assert.Equal(123456, i)
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	var s string
	any = "myString"
	Cast(any, &s)
	assert.Equal("myString", s)
}

func TestFloat64(t *testing.T) {
	assert := assert.New(t)
	var f float64
	any = 3.141592653
	Cast(any, &f)
	assert.Equal(3.141592653, f)
}

func TestByteSlice(t *testing.T) {
	assert := assert.New(t)
	var b []byte
	any = []byte{32, 64, 128}
	Cast(any, &b)
	assert.Equal([]byte{32, 64, 128}, b)
}

func TestIntSlice(t *testing.T) {
	assert := assert.New(t)
	var is []int
	any = []int{1, 2, 3}
	Cast(any, &is)
	assert.Equal([]int{1, 2, 3}, is)
}

func TestMapStringInterface(t *testing.T) {
	assert := assert.New(t)
	var m map[string]interface{}
	any = map[string]interface{}{
		"string": "myString",
		"int":    123456,
	}
	Cast(any, &m)
	assert.Equal(map[string]interface{}{
		"string": "myString",
		"int":    123456,
	}, m)
}

func TestStruct(t *testing.T) {
	assert := assert.New(t)
	var st structType
	any = structType{"myString", 123456}
	Cast(any, &st)
	assert.Equal(structType{"myString", 123456}, st)
}

func TestError(t *testing.T) {
	assert := assert.New(t)
	err := Cast(any, nil)
	assert.Error(err)
}
