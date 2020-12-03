package structs

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

type Point struct {
	X, Y int
}

type Circle struct {
	 Center Point
	 Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func Test_Struct_inherit(t *testing.T){
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Spokes = 20
	fmt.Printf("weel %s", gconv.String(w))

	w = Wheel{
		Circle: Circle{
			Center: Point{
				X: 0,
				Y: 0,
			},
			Radius: 10,
		},
	}
}
