# go-smpp

## Client Usage

```
go get github.com/Jason-Hendry/go-smpp
```

Example app <main.go>
```
package main

import (
	"github.com/Jason-Hendry/go-smpp"
)

main() {
  go_smpp.Client("smsglobal.com:1775","<username>","<password>","<source>","<destination>","Test SMS from Go")
}
```