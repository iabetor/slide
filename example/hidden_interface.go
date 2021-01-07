package example

import "fmt"

//START OMIT
type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t *T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = &T{"hello"}
	i.M()
}

//END OMIT
