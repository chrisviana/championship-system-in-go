package main

import "fmt"

func main() {
	menu()
}

func menu() {
	fmt.Println("****************************************************************")
	fmt.Println("1. Create a new championship")
	fmt.Println("2. Add a athlete")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		createNewChampionship()
	case 2:
		// listChampionships()
	case 3:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
		menu()
	}
}

func createNewChampionship() {

	var name string
	var start string
	var end string
	var numberAthlete int
	var hashCode string

	fmt.Println("Creating a new championship...")
	fmt.Println("Enter the championship's name:")
	fmt.Scan(&name)
	fmt.Printf("Nmae the championship is: %s \n", name)
	fmt.Println("Enter the start date (YYYY-MM-DD):")
	fmt.Scan(&start)
	fmt.Printf("Enter the end ate (YYYY-MM-DD):")
	fmt.Scan(&end)
	fmt.Println("Enter the number of athletes:")
	fmt.Scan(&numberAthlete)
	fmt.Println("Enter a unique hash code for the championship:")
	fmt.Scan(&hashCode)
}
