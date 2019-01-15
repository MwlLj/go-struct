package copy

import (
	"fmt"
	"testing"
)

type CDog struct {
	Name string
	Age  int
}

type CCat struct {
	Name string
	Age  int
}

func TestOrderCopy(t *testing.T) {
	dog := CDog{
		Name: "wangcai",
		Age:  1,
	}
	cat := CCat{
		Name: "tomon",
		Age:  2,
	}
	fmt.Println(dog)
	OrderCopy(&dog, &cat)
	fmt.Println(dog)
}
