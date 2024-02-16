package main

import (
	"MIA_P1_202004822/Estructuras"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("********************")
	fmt.Println("**** BIENVENIDO PRY1 Sebastian Solares *****")
	fmt.Println("********************")

	//send_console()
	Open_File()

}

func send_console() {
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Ingresa un comando: ")

		comando, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			return
		}

		// Eliminar espacios en blanco y nueva línea de la entrada
		comando = strings.TrimSpace(comando)

		if comando == "exit" {
			break
		}
		Estructuras.Analyze(comando)
	}

}
func Open_File() {

	// pedimos la ruta del archivo
	fmt.Println("Ingrese la ruta del archivo")
	var ruta string
	fmt.Scanln(&ruta)
	fmt.Println("La ruta es:", ruta)
	//Abrimos el archivo

	readFile, err := os.Open(ruta)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for _, line := range lines {
		//fmt.Println("The name is:", line)
		Estructuras.Analyze(line)

	}
}
