package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chrisviana/golang-champions/athlete"
	"github.com/chrisviana/golang-champions/championship"
)

func main() {
	menu()
}

func menu() {
	fmt.Println("****************************************************************")
	fmt.Println("1. Create a new championship")
	fmt.Println("2. Add a athlete")
	fmt.Println("3. Create game")
	fmt.Println("4. List athletes")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		createNewChampionship()
	case 2:
		addAthlete()
	case 3:
		createGame()
	case 4:
		listAthletes()
	case 5:
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
		menu()
	}
}

var currentChampionship *championship.Championship

func createNewChampionship() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Creating a new championship...")
	fmt.Println("Enter the championship's name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter the start date (YYYY-MM-DD):")
	start, _ := reader.ReadString('\n')
	start = strings.TrimSpace(start)

	fmt.Printf("Enter the end date (YYYY-MM-DD):")
	end, _ := reader.ReadString('\n')
	end = strings.TrimSpace(end)

	fmt.Print("Enter the number of athletes:")
	numberStr, _ := reader.ReadString('\n')
	numberAthlete, _ := strconv.Atoi(strings.TrimSpace(numberStr))

	fmt.Print("Enter a unique hash code for the championship:")
	hashCode, _ := reader.ReadString('\n')
	hashCode = strings.TrimSpace(hashCode)

	currentChampionship = championship.NewChampionship(name, start, end, hashCode, numberAthlete)
	fmt.Printf("Campeonato criado com sucesso: %+v\n", currentChampionship)

	menu()
}

func addAthlete() {
	if currentChampionship == nil {
		fmt.Println("É preciso criar um campeonato primeiro")
		menu()
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Adicionado novo atleta...")
	fmt.Print("Digite o nome do atleta: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	newAthlete := athlete.NewAthlete(name)
	err := currentChampionship.AddAthlete(newAthlete)
	if err != nil {
		fmt.Printf("Erro adicionar atleta: %v\n", err)
	} else {
		fmt.Printf("Atleta %s adicionado com sucesso!\n", name)
	}

	menu()
}

func createGame() {
	if currentChampionship == nil {
		fmt.Println("É preciso criar um campeonato primeiro")
		menu()
		return
	}

	if len(currentChampionship.Athletes) < 2 {
		fmt.Println("Erro: Adicione pelo menos 2 atletas antes de cirar um jogo")
		menu()
		return
	}

	fmt.Println("Atletas disponíveis:")
	for i, athlete := range currentChampionship.Athletes {
		fmt.Printf("%d. %s\n", i+1, athlete.Name)
	}

	var athlete1Index, athlete2Index int
	fmt.Print("Selecione o número do primeiro atleta: ")
	fmt.Scan(&athlete1Index)
	fmt.Print("Selecione o número do segundo atleta: ")
	fmt.Scan(&athlete2Index)

	athlete1Index--
	athlete2Index--

	err := currentChampionship.CreateGame(athlete1Index, athlete2Index)
	if err != nil {
		fmt.Printf("Erro ao criar jogo: %v\n", err)
	} else {
		fmt.Printf("Jogo criado com sucesso entre %s e %s!\n",
			currentChampionship.Athletes[athlete1Index].Name,
			currentChampionship.Athletes[athlete2Index].Name)
	}

	menu()
}

func listAthletes() {
	if currentChampionship == nil {
		fmt.Println("É preciso criar um campeonato primeiro")
		menu()
		return
	}

	fmt.Println("\nAtletas no campeonato:  ")
	if len(currentChampionship.Athletes) == 0 {
		fmt.Println("Nenhum atleta cadastrado ainda")
	} else {
		for i, atathlete := range currentChampionship.Athletes {
			fmt.Printf("%d. %s\n", i+1, atathlete.Name)
		}
	}

	menu()
}
