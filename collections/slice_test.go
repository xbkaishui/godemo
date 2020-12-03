package collections

import (
	"fmt"
	"testing"
)

func Test_slice(t *testing.T){
	values := [3]int{1,2,3}
	fmt.Println(values[:0])
	t.Log("try testing slice")
	var x, y []int
	for i:=0 ; i< 10 ; i++ {
		y = AppendInt(x,i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func AppendInt(x []int, i int) []int {
	var z [] int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	z[len(x)] = i
	return z
}
