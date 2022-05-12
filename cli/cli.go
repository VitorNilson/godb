package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/vitornilson1998/godb/persistence"
)
const logo = 
`
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░     ░░░░░░░░░     ░░░░░░      ░░░░░   ░░░░░░
▒▒  ▒▒▒▒   ▒▒▒▒▒   ▒▒▒▒   ▒▒▒   ▒▒▒   ▒▒   ▒▒▒▒▒▒
▒  ▒▒▒▒▒▒▒▒▒▒▒   ▒▒▒▒▒▒▒▒   ▒   ▒▒▒▒   ▒   ▒▒▒▒▒▒
▓   ▓▓▓▓▓▓▓▓▓▓   ▓▓▓▓▓▓▓▓   ▓   ▓▓▓▓   ▓   ▓   ▓▓
▓   ▓▓▓      ▓   ▓▓▓▓▓▓▓▓   ▓   ▓▓▓▓   ▓   ▓▓▓   
▓▓   ▓▓▓▓  ▓▓▓▓▓   ▓▓▓▓▓   ▓▓   ▓▓▓   ▓▓   ▓▓▓   
███      █████████     ██████      █████   █   ██
█████████████████████████████████████████████████
`

// Enum to define SQL clause
const (
	Select  = "select"
	Create  = "create"
	Update  = "update"
	Drop  = "drop"


	Exit = "exit"
)

const (
	colorPurple = "\033[35m"
)

// Enum to define SQL preposition
const (
	Table = "table"
	Index = "index"
	Sequence = "sequence"
	View = "view"
)

const (
	defaultMessage = "Not implemented yet."
	doesNotExists = "Clause or preposition does not exist."
)

var tableMethods = persistence.TableOperations{}

var logoPrinted bool

func main() {
	printLogo()
	processSql(entryPoint())
}

func printLogo(){
	if(!logoPrinted){
		logoPrinted = true
		fmt.Print(string(colorPurple), logo)
	}
	
}

func processSql(sql string){
	sqlSplited := strings.Split(sql, " ")

	switch  strings.ToLower(sqlSplited[0]) {
		case Select, Update:
			notImplementedYet()

		case Drop:
			handleDeleteClause(sqlSplited[1], sqlSplited)

		case Create:
			handleCreateClause(sqlSplited[1], sql, sqlSplited)
			
		case Exit:
			exit()

		default:
			doesNotExistsFunc()
	}

}

func exit(){
	fmt.Println("bye")
	os.Exit(0)
}

func notImplementedYet(){
	fmt.Println(defaultMessage)
			main()

}

func doesNotExistsFunc(){
	fmt.Println(doesNotExists)	
	main()
}
func handleDeleteClause(delete string, sqlDropTable []string){
	switch strings.ToLower(delete){
		case Table:
			tableMethods.DropTable(sqlDropTable[2])
			main()
	}
}

func handleCreateClause(create string, sql string, sqlCreateTable []string){
	
	switch strings.ToLower(create){
			
		case Index, Sequence, View:
			notImplementedYet()
		
		case Table:
			tableMethods.CreateTable(sql, sqlCreateTable)
			main()

		default:
			doesNotExistsFunc()

	}
}




func entryPoint() string{
	fmt.Print(string(colorPurple),"godb: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n")

	return input
}