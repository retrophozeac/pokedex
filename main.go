package main

func main() {
	startRepl()
	// scanner := bufio.NewScanner(os.Stdin)
	// commands := map[string]cliCommand{
	// 	"exit": {
	// 		name:        "exit",
	// 		description: "Exit the Pokedex",
	// 		callback:    commandExit,
	// 	},
	// 	"help": {
	// 		name:        "help",
	// 		description: "Show the commands",
	// 		callback:    commandHelp,
	// 	},
	// }
	// for {
	// 	fmt.Printf("Pokedex > ")
	// 	ok := scanner.Scan()
	// 	if ok {
	// 		text := scanner.Text()
	// 		text_array := cleanInput((text))
	// 		if s, ok := commands[text_array[0]]; ok {
	// 			err := s.callback()
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}
	// 			continue
	// 		}
	// 		fmt.Printf("Unknown command\n")

	// 	}
	// }
}

// func cleanInput(text string) []string {
// 	output := strings.ToLower(text)
// 	words := strings.Fields(output)
// 	// fmt.Println(words)
// 	return words
// }
// func commandExit() error {
// 	fmt.Printf("Closing the Pokedex... Goodbye!")
// 	os.Exit(0)
// 	return nil
// }

// type cliCommand struct {
// 	name        string
// 	description string
// 	callback    func() error
// }

// func commandHelp() error {
// 	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
// 	return nil
// }
