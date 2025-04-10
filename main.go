package main

import "fmt"

func main() {
	var fullGolang = "Watch Nana's Golang Full Course"
	var shortGolang = "Watch Go crash course"
	rewardDessert := "Reward myself with a donut"

	// var taskItems = []string{"Watch Nana's Golang Full Course", "Watch Go crash course", "Reward myself with a donut"}
	taskItems := []string{fullGolang, shortGolang, rewardDessert}

	fmt.Println("###### Welcome to our TodoList App! ######")

	fmt.Println()
	fmt.Println("- List of my Todos -")
	//fmt.Println(shortGolang)
	//fmt.Println(fullGolang)
	//fmt.Println(rewardDessert)
	// fmt.Println("Tasks:", taskItems)

	for index, task := range taskItems {
		//fmt.Println(index+1, ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}

	//fmt.Println()
	//fmt.Println("Tutorials:")
	//fmt.Println(shortGolang)
	//fmt.Println(fullGolang)

	//fmt.Println()
	//fmt.Println("Rewards:")
	//fmt.Println(rewardDessert)

	//fmt.Println()
	//fmt.Println("My Project:")
	//fmt.Println(fullGolang)
}
