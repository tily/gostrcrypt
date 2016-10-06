# gostrcrypt

[![Build Status](https://travis-ci.org/tily/gostrcrypt.svg?branch=master)](https://travis-ci.org/tily/gostrcrypt)
[![Code Climate](https://codeclimate.com/github/tily/gostrcrypt/badges/gpa.svg)](https://codeclimate.com/github/tily/gostrcrypt)
[![Issue Count](https://codeclimate.com/github/tily/gostrcrypt/badges/issue_count.svg)](https://codeclimate.com/github/tily/gostrcrypt)
[![Coverage Status](https://coveralls.io/repos/github/tily/gostrcrypt/badge.svg?branch=master)](https://coveralls.io/github/tily/gostrcrypt?branch=master)
[![GoDoc](https://godoc.org/github.com/tily/gostrcrypt?status.svg)](http://godoc.org/github.com/tily/gostrcrypt)

Golang library to encrypt/decrypt short string.

## Usage

```go
package main

import (
        "fmt"
        "github.com/tily/gostrcrypt"
)

func main() {
        key := "CAFEBABECAFEBABE"
        sc := gostrcrypt.StrCrypt{key}

        plaintext := "hello, world"
        encrypted, err := sc.Encrypt(plaintext)
        if err != nil {
                panic(err)
        }
        decrypted, err := sc.Decrypt(encrypted)
        if err != nil {
                panic(err)
        }
        fmt.Println(decrypted) // => "hello, world"
}
```

## Note

* Only supports AES CFB encryption currently
* Large-size string encryption will possibly exhaust the memory on your system. Use lower stream API instead
