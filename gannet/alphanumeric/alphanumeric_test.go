package alphanumeric

import (
	"fmt"
	"testing"
)

func TestRemoveNonAlphanumeric(t *testing.T) {
	fmt.Println(RemoveNonAlphanumeric("ab$cd1!#2*3?4%]"))
	fmt.Println(RemoveNonAlphanumeric2("ab$cd1!#2*3?4%]"))
}
