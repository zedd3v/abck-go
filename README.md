# Abck

Abck is a library that makes generating akamai cookies easier.

Currently does not support all of the functions/methods. Expect much more in an upcoming update.


## Installation

```bash
go get github.com/zedd3v/abck-go
```

## Usage

```go
package main

import (
  "fmt"
  abck "github.com/zedd3v/abck-go"
)

func main() {
  fmt.Println(abck.GetCfDate()) // Equivalent of Date.now()
}
```

```go
func D3() int64

func Getmr() string // Broken, make a pr if you find a fix

func Ab(t string) int

func O9(d3 int64) int

func GetCfDate() int64

func GetType(t string) int

func Sed(browser string) string

func Uar(userAgent string) string

func Jrs(t int) ([]float64, error)

func Od(t string, a string) string

func Z1(startTimestamp int64) int64

func Fas(browser string) (int, error)

func Np(browser string) (string, error)

func Gd(browser string, browserData BrowserData, startTimestamp int64, d3 int64) (string, error)
```

## Maintainer

[![ZedDev](https://github.com/zedd3v.png?size=100)](https://abck.dev/)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)