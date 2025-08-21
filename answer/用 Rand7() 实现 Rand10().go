package answer

import "math/rand/v2"

func rand7() int {
	return rand.IntN(7) + 1
}

func rand10() int {
	for {
		num := (rand7()-1)*7 + rand7()
		if num <= 40 {
			return (num-1)%10 + 1
		}
	}
}
