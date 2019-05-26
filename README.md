# helper
[![Build Status](https://travis-ci.org/jeanmorais/helper.svg?branch=master)](https://travis-ci.org/jeanmorais/helper) [![codecov](https://codecov.io/gh/jeanmorais/helper/branch/master/graph/badge.svg?token=r3cEOnuN7T)](https://codecov.io/gh/jeanmorais/helper)
[![Go Report Card](https://goreportcard.com/badge/github.com/jeanmorais/helper)](https://goreportcard.com/report/github.com/jeanmorais/helper)
[![Documentation](https://godoc.org/github.com/jeanmorais/helper?status.svg)](https://godoc.org/github.com/jeanmorais/helper)

helper is a collection of utilitarian functions, to facilitate working with strings and slices in Golang.

## Install

```
go get github.com/jeanmorais/helper
```

## Usage

Here are some examples of usage.

```go
package main

import (
	"fmt"
	"github.com/jeanmorais/helper"
)

func main() {
	fmt.Println(helper.OnlyNumbers("abc123"))
	fmt.Println(helper.OnlyLetters("abc123"))
	fmt.Println(helper.RightPad("golang", 8, "!"))

}

```

```bash
123
abc
golang!!
```
