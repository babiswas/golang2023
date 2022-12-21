package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var test string
	count := 0
	daemon := map[string]int{"dracula": 40, "shapeshifter": 100, "angel": 150}
	weapon := map[string]int{"sword": 10, "knife": 5, "stick": 8, "colt": 12}
	monsters := []string{"dracula", "shapeshifter", "angel"}
	weapons := []string{"sword", "knife", "stick", "colt"}
	for {
		num := rand.Intn(3)
		monster := monsters[num]
		fmt.Println("Monster appeared:", monster)
		if daemon[monster] <= 0 {
			daemon[monster] = rand.Intn(1000)
		}
		for daemon[monster] > 0 {
			point := rand.Intn(4)
			fmt.Println("Obtained weapon:", weapons[point])
			daemon[monster] -= weapon[weapons[point]]
			count += 1
		}
		fmt.Println("No of strikes:", count)
		count = 0
		fmt.Scanln(&test)
	}
}
