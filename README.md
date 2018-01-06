# Mirror

米洛能以最簡單的方式協助你將一個未知的 `interface{}` 宣告到另一個已知的 `interface{}` 指針所指向的建構體或資料。這很適合設計套件時將一個未知型態的資料綁定到未知型態的指針上。

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/TeaMeow/Mirror
```

# 使用方式

本套件以最少的函式提供最大的實用性。

## 宣告

透過 `Cast` 能將一個 `interface{}` 宣告到一個已知的 `interface{}` 指針所指向的建構體或資料上。

```go
package main

import (
	"fmt"

	mirror "github.com/TeaMeow/Mirror"
)

func main() {
	var any interface{}

	// `int` 型態。
	var i int
	any = 123456
	mirror.Cast(any, &i)
	fmt.Println(i) // 輸出：123456
}
```