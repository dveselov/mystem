# go-mystem [![Build Status](https://travis-ci.org/dveselov/go-mystem.svg?branch=master)](https://travis-ci.org/dveselov/go-mystem) [![GoDoc](https://godoc.org/github.com/dveselov/go-mystem?status.svg)](https://godoc.org/github.com/dveselov/go-mystem) [![Go Report Card](https://goreportcard.com/badge/github.com/dveselov/go-mystem)](https://goreportcard.com/report/github.com/dveselov/go-mystem)
CGo bindings to Yandex.Mystem - russian morphology analyzer.

# Usage
```go
package main

import (
    "fmt"
)

import "github.com/dveselov/go-mystem"

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
Output'll look like this:
```
Analyze of 'маша':
1. маша - [136 155 191 201]
2. махать - [137 196 206]
```

# License

Source code of `go-mystem` is licensed under MIT license, but Yandex.Mystem have their own [EULA](https://yandex.ru/legal/mystem/) (allows commercial usage), that you must accept.
