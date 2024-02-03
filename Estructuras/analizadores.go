package Estructuras

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Analyze(command string) {
	//variables
	var flagExit bool

	//aqui al comando, le quitamos los espacios y lo devolvemos como un token

	token_ := strings.Split(command, " ")

	//CONVERTIMOS EL TOKEN EN MINUSCULAS
	if len(token_) > 0 {
		token_[0] = strings.ToLower(token_[0])
	}
	fmt.Println(token_[0])

	switch token_[0] {
	case "mkdisk":
		//estamos aca
		Analyze_Mkdisk(token_[1:])

	case "exit":
		flagExit = true

	default:
		fmt.Println("error al leer el comando")
	}

	if !flagExit {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Ingresa un comando: ")

		comando, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			return
		}
		Analyze(comando)
	}

}

func Analyze_Mkdisk(list_tokens []string) {

	fmt.Println(list_tokens)

}
