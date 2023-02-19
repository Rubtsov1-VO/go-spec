package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x3, y3, x4, y4 int
	var n int
	fmt.Scan(&x1,
		&y1,
		&x2,
		&y2,
		&x3,
		&y3,
		&x4,
		&y4,
		&n)
	for i := 0; i < n; i++ {
		var x int
		var y int
		fmt.Scan(&x)
		fmt.Scan(&y)
		switch {
		case x1 <= x && x <= x4 && y1 <= y && y <= y3:
			fmt.Printf("Точка с координатами %d,%d принадлежит исследуемому квадрату\n", x, y)
			continue
		case x > x4 || x1 > x || y > y3:
			fmt.Printf("Точка с координатами %d,%d не принадлежит исследуемому квадрату\n", x, y)
			continue
		}

	}
}
