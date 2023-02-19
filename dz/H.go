package main

import (
	"fmt"
)

func main() {
	var dro int
	var z []int
	var up []int
	fmt.Scan(&dro)
	for i := 0; i < dro; i++ {
		var x int
		var y int

		fmt.Scan(&x)
		fmt.Scan(&y)
		z = append(z, x)
		up = append(up, y)
	}
	//	fmt.Println(z)
	//	fmt.Println(up)

	for k := range [100000]int{} {
		k++
		if dro == 1 {
			if k%up[dro-1] == 0 {
				fmt.Println(k)
				break
			} else {
				break
			}
		} else if dro == 2 {
			if k%up[dro-1] == 0 && k%up[dro-2] == 0 {
				fmt.Println(k)
				break
			} else {
				break
			}
		} else if dro == 3 {
			if k%up[dro-1] == 0 && k%up[dro-2] == 0 && k%up[dro-3] == 0 {

				result := z[0]*(k/up[0]) + z[1]*(k/up[1]) + z[2]*(k/up[2])
				fmt.Printf("%d/%d", result, k)
				break
			}
		}

	}
}
