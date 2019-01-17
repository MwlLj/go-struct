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
	Name    string
	Age     int
	Pig     CPig
	Collect string
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
	// fmt.Println(cat)
	OrderCopy(&dog, &cat)
	// fmt.Println(cat)
}

type Src struct {
	Name string
	ID   string
	Pig  *CPig
}

type Dst struct {
	Name string
	ID   string
	Pig  CDog
}

// 测试复制结构体
func TestStructCopy(t *testing.T) {
	src := &Src{"katy", "000", &CPig{"angang", 1}}
	// fmt.Println("src:" + util.StringifyJson(src))
	var dst Dst
	StructCopy(&dst, src)
	fmt.Println(dst)
	// fmt.Println("dst:" + util.StringifyJson(dst))
}
