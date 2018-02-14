## UUID package for Go language

A Go library that generates concise, unambiguous, URL-safe UUIDs. Based on and compatible with the Python library shortuuid.

## Usage

```go
import (
	"fmt"
	"strings"

	"github.com/houndgo/suuid"
)

func main() {
  uid := suuid.New()
  fmt.Println("\n uuid is: ", uid)
}

```


## License

MIT