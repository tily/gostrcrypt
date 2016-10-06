# gostrcrypt

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
