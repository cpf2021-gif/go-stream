# go-stream
> A simple stream api for Go, powered by **Rangefunc Experiment**.

## Usage
> [!NOTE]
> Using Go 1.22, build your program using **GOEXPERIMENT=rangefunc**, as in
>```bash
>GOEXPERIMENT=rangefunc go build
>GOEXPERIMENT=rangefunc go run
> 
>```

##### Quick Start
```go
package main

import (
	"fmt"
	"github.com/cpf2021-gif/go-stream"
)

func main() {
	nums := []int{
		1, 1, 9, 4, 4, 6, 3, 5, 15,
	}

	fmt.Println(stream.StreamOf(nums).
		Filter(func(i int, v int) bool {
			return v%2 == 1
		}). // -> 1, 1, 9, 3, 5, 15
		Sorted(func(a, b any) bool {
			return a.(int) < b.(int)
		}). // -> 1, 1, 3, 5, 9, 15
		Distinct(func(v any) any {
			return v.(int) % 5
		}).        // -> 1, 3, 5, 9
		Tail(3).   // -> 3, 5, 9
		ToSlice()) // -> [3, 5, 9]
}

```