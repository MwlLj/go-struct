package copy

import (
	"fmt"
	"testing"
)

var _ = fmt.Println

type CPig struct {
	Name string
	Age  int
}

type CDog struct {
	Name string
	Age  int
	Pig  CPig
}

type CCat struct {
	Name string
	Age  int
	Pig  CPig
}

func TestOrderCopy(t *testing.T) {
	dog := CDog{
		Name: "wangcai",
		Age:  1,
		Pig: CPig{
			Name: "hanghang1",
			Age:  1,
		},
	}
	cat := CCat{
		Name: "tomon",
		Age:  2,
		Pig: CPig{
			Name: "hanghang2",
			Age:  3,
		},
	}
	fmt.Println(cat)
	OrderCopy(&dog, &cat)
	fmt.Println(cat)
}
