# go-mystem [![Build Status](https://travis-ci.org/dveselov/mystem.svg?branch=master)](https://travis-ci.org/dveselov/mystem) [![GoDoc](https://godoc.org/github.com/dveselov/mystem?status.svg)](https://godoc.org/github.com/dveselov/mystem) [![Go Report Card](https://goreportcard.com/badge/github.com/dveselov/go-mystem)](https://goreportcard.com/report/github.com/dveselov/mystem) [![Coverage Status](https://coveralls.io/repos/github/dveselov/go-mystem/badge.svg)](https://coveralls.io/github/dveselov/go-mystem)
CGo bindings to Yandex.Mystem - russian morphology analyzer.

# Install
```bash
$ wget https://github.com/yandex/tomita-parser/releases/download/v1.0/libmystem_c_binding.so.linux_x64.zip
$ unzip libmystem_c_binding.so.linux_x64.zip
$ sudo cp libmystem_c_binding.so /usr/lib/
$ sudo ln -s /usr/lib/libmystem_c_binding.so /usr/lib/libmystem_c_binding.so.1
$ go get -u github.com/dveselov/mystem
```

# Usage
```go
package main

import (
    "fmt"
)

import "github.com/dveselov/mystem"

func main() {
    analyses := mystem.NewAnalyses("маша")
    defer analyses.Close()
    fmt.Println(fmt.Sprintf("Analyze of '%s':", "маша"))
    for i := 0; i < analyses.Count(); i++ {
        lemma := analyses.GetLemma(i)
        grammemes := lemma.StemGram()
        fmt.Println(fmt.Sprintf("%d. %s - %v", i+1, lemma.Text(), grammemes))
    }
}
```
Output'll looks like this:
```
Analyze of 'маша':
1. маша - [136 155 191 201]
2. махать - [137 196 206]
```

# License

Source code of `go-mystem` is licensed under MIT license, but Yandex.Mystem have their own [EULA](https://yandex.ru/legal/mystem/) (allows commercial use), that you must accept.
