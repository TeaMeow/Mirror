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
	err := Cast(any, &i)
	assert.Equal(123456, i)
	assert.NoError(err)
}

func TestInt64ToInt(t *testing.T) {
	assert := assert.New(t)
	var i int
	any = int64(123456)
	err := Cast(any, &i)
	assert.Equal(123456, i)
	assert.NoError(err)
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	var s string
	any = "myString"
	err := Cast(any, &s)
	assert.Equal("myString", s)
	assert.NoError(err)
}

func TestFloat64(t *testing.T) {
	assert := assert.New(t)
	var f float64
	any = 3.141592653
	err := Cast(any, &f)
	assert.Equal(3.141592653, f)
	assert.NoError(err)
}

func TestByteSlice(t *testing.T) {
	assert := assert.New(t)
	var b []byte
	any = []byte{32, 64, 128}
	err := Cast(any, &b)
	assert.Equal([]byte{32, 64, 128}, b)
	assert.NoError(err)
}

func TestIntSlice(t *testing.T) {
	assert := assert.New(t)
	var is []int
	any = []int{1, 2, 3}
	err := Cast(any, &is)
	assert.Equal([]int{1, 2, 3}, is)
	assert.NoError(err)
}

func TestMapStringInterface(t *testing.T) {
	assert := assert.New(t)
	var m map[string]interface{}
	any = map[string]interface{}{
		"string": "myString",
		"int":    123456,
	}
	err := Cast(any, &m)
	assert.Equal(map[string]interface{}{
		"string": "myString",
		"int":    123456,
	}, m)
	assert.NoError(err)
}

func TestStruct(t *testing.T) {
	assert := assert.New(t)
	var st structType
	any = structType{"myString", 123456}
	err := Cast(any, &st)
	assert.Equal(structType{"myString", 123456}, st)
	assert.NoError(err)
}

func TestError(t *testing.T) {
	assert := assert.New(t)
	err := Cast(any, nil)
	assert.Error(err)
}
