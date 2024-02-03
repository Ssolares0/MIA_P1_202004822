package main

import (
	"MIA_P1_202004822/Estructuras"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("********************")
	fmt.Println("**** BIENVENIDO PRY1 Sebastian Solares *****")
	fmt.Println("********************")

	Open_File()

}

func send_console() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingresa un comando: ")

	comando, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	}

	Estructuras.Analyze(comando)

}
func Open_File() {
	readFile, err := os.Open("entrada")
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
