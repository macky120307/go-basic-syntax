package main

import "math/rand"
import "time"
import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Print(i)
			fmt.Print(" - 偶数\n")
			} else {
				fmt.Print(i)
				fmt.Print(" - 奇数\n")
		}
	}

	rand.Seed(time.Now().Unix())
	n := rand.Intn(6) + 1

	switch n {
		case 1:
			fmt.Print("凶")
		case 2, 3:
			fmt.Print("吉")
		case 4, 5:
			fmt.Print("中吉")
		case 6:
			fmt.Print("大吉")
	}
}