<h1 align="center">Is Even</h1>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/AvicennaJr/isEven?style=plastic">
<img src="https://img.shields.io/github/actions/workflow/status/avicennajr/iseven/tests.yml?style=plastic">
<img src="https://img.shields.io/github/last-commit/avicennajr/iseven?style=plastic">
<img src="https://img.shields.io/github/stars/avicennajr/iseven?style=plastic">
</p>
A golang package to check is a number is even.

## Installation
```bash
go get github.com/AvicennaJr/IsEven
```
## How To Use
```go
package main

import (
	"fmt"

	iseven "github.com/AvicennaJr/isEven"
)

func main() {
	n := 4
	if iseven.IsEven(n) {
		fmt.Println("It is even")
	} else {
		fmt.Println("It is not even")
	}
}
```
## Issues
Do open issues for any complaints
## Contributions
If you want to contribute, simply open a pull request.
## Credits
Thanks to the hundreds of developers who made this project possible.
