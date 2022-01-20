# Go Password hasher

Password hasher for GO.

Installation and usage
----------------------

To install it, run:

    go get github.com/vetal4ik10/hasher

Example
-------

```Go
package main

import "fmt"
import "github.com/vetal4ik10/hasher"

func main() {
	password := "secret"
	hash, _ := hasher.HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := hasher.CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
```

This example will generate the following output:

```
Password: secret
Hash:     $2a$14$IQ1JrAW9Saq5JGHN/Rv6pe8/Lond8SwE02.bIZpeZap/FJAwLy3P.
Match:    true

```

