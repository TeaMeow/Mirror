package mirror

import (
	"errors"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

var (
	// ErrUnknownPanic 表示發生了未知的 Panic。
	ErrUnknownPanic = errors.New("mirror: unknown panic")
)

// Cast 能將一個 `interface{}` 宣告到另一個 `interface{}` 指針上。
func Cast(from interface{}, to interface{}) (err error) {
	// 擷取任何的 `panic` 錯誤。
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = ErrUnknownPanic
			}
		}
	}()
	// 將 `interface{}` 宣告到 `interface{}` 指針。
	switch from.(type) {
	case map[string]interface{}:
		err = mapstructure.Decode(from, to)
	default:
		reflect.ValueOf(to).Elem().Set(reflect.ValueOf(from))
	}
	return
}
