# go-smpp

## Client Usage


```
import (
	"github.com/Jason-Hendry/go-smpp"
)

main() {
  go_smpp.Client("smsglobal.com:1775","<username>","<password>","<source>","<destination>","Test SMS from Go")
}
```