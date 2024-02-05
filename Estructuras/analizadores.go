package Estructuras

import (
	"fmt"
	"strconv"
	"strings"
)

func Analyze(command string) {
	//variables
	//var flagExit = false

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
	case "execute":

		Analyze_execute(token_[1:])

	case "rmdisk":
		Analyze_Rmdisk(token_[1:])

	case "exit":
		//flagExit = true
		fmt.Println("gracias por usar el programa")

	default:
		fmt.Println("error al leer el comando")

	}

	/*
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
	*/
}

func Analyze_Mkdisk(list_tokens []string) {

	fmt.Println(list_tokens)

	//variables del mkdisk
	var size_int int
	var size, fit, unit string
	//vamos a separar el valor igual

	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-size":
			size = tokens[1]
			size_int, _ = strconv.Atoi(tokens[1])

			fmt.Println("El size es: " + size)
		case "-unit":
			unit = tokens[1]
			fmt.Println("El unit es: " + unit)

		case "-fit":
			fit = tokens[1]
			fmt.Println("el fit es; " + fit)
		}
		if size_int == 0 {

			fmt.Println("Faltan parametros obligatorios")

		} else {
			//creamos archivo vacio

		}
	}

}

func Analyze_execute(list_tokens []string) {
	fmt.Println(list_tokens)

}
func Analyze_Rmdisk(list_tokens []string) {
	fmt.Println(list_tokens)
}

func Confirmacion(msg string) bool {
	fmt.Println(msg + "(y/n)")
	//var respuesta string
	return true

}

func WriteInBytes() {
	fmt.Println("jeje")
}
