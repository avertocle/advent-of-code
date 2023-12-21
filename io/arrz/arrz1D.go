package arrz

import (
	"fmt"
	"github.com/avertocle/contests/io/cmz"
)

func Key1D[T cmz.Primitive](keys []T) string {
	return fmt.Sprintf("%v", keys)
}
