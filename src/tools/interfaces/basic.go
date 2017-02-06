package gg

import (
	"fmt"
)

// Point 基本的な Struct
type Point struct {
	X int
	Y int
}

// Coordinate Struct のメソッド
func (p Point) Coordinate() string {
	// p がレシーバ
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	var a Point = Point{2, 3}
	fmt.Println(a.Coordinate())
}
