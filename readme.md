## UUID package for Go language

A Go library that generates concise, unambiguous, URL-safe UUIDs. Based on and compatible with the Python library shortuuid.


## Install

```
go get github.com/houndgo/suuid
```

## Usage

```go
import (
	"fmt"
	"strings"

	"github.com/houndgo/suuid"
)

func main() {
  uid := suuid.New()
  fmt.Println("\n uuid is: ", uid) //fnB6fRCHrPqStSXYEs7W73 by uuid is 6e8e463e-d39c-460e-9e62-35192ff11f89
}

```


## License

MIT