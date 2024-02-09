# go-nuban

Generate NUBAN (Nigeria Uniform Bank Account Number) using Go based on the offical/revised CBN algorithm.

## Installation

```bash
go get -u github.com/bigmeech/go-nuban
```

## Usage

```go
package main

import (
	"fmt"
	nuban "github.com/bigmeech/go-nuban"
)

const (
	bankCode = "011"
	serialNumber = "1234567890"
)

func main() {
	nuban := nuban.Generate(bankCode, serialNumber)
	fmt.Println(nuban)
}

```

## License

MIT


Made with love from the Yunobank Team

