package Estructuras

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Analyze(command string) {
	fmt.Println(command)
	//aqui al comando, le quitamos los espacios y lo devolvemos como un token

	token_ := strings.Split(command, " ")

	//CONVERTIMOS EL TOKEN EN MINUSCULAS
	if len(token_) > 0 {
		token_[0] = strings.ToLower(token_[0])
	}

	//verificamos si el token tiene un #
	if strings.Contains(token_[0], "#") {
		fmt.Println("comentario: ")
		token_[0] = strings.Replace(token_[0], "#", "", -1)
		for _, element := range token_ {
			fmt.Println(element + " ")
		}

	}

	//fmt.Println(token_[0])
	//fmt.Println("token 0: ", token_[0])
	switch token_[0] {
	case "mkdisk":
		//estamos aca

		Analyze_Mkdisk(token_[1:])
	case "execute":

		Analyze_execute(token_[1:])

	case "rmdisk":
		Analyze_Rmdisk(token_[1:])
	case "fdisk":
		Analyze_Fdisk(token_[1:])

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
	FlagObligatorio := true
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
			//pasar aminuscula
			unit = strings.ToLower(unit)
			fmt.Println("El unit es: " + unit)

		case "-fit":
			fit = tokens[1]
			//pasar a minuscula
			fit = strings.ToLower(fit)
			fmt.Println("el fit es; " + fit)
		}

	}
	if size_int == 0 {

		fmt.Println("Faltan parametros obligatorios")
		FlagObligatorio = false

	}

	if FlagObligatorio == true {
		//creamos archivo
		fmt.Println("creando disco...")

		CrearDisco(size_int, unit, fit)

	}

}

func Analyze_execute(list_tokens []string) {
	fmt.Println(list_tokens)

}
func Analyze_Rmdisk(list_tokens []string) {
	fmt.Println(list_tokens)
	//con este comando borramos el disco que se le indique

	//separa el nombre del disco
	tokens := strings.Split(list_tokens[0], "=")

	if tokens[0] == "-driveletter" {
		//preguntamos si quiere borra el disco
		if Confirmacion("Desea borrar el disco") {
			fmt.Println("borrando disco")
			nombreArchivo := "Discos/" + tokens[1] + ".dsk" // El nombre o ruta absoluta del archivo
			err := os.Remove(nombreArchivo)
			if err != nil {
				fmt.Printf("Error eliminando el disco: %v\n", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}
		} else {
			fmt.Println("No se elimino el disco")
		}

	} else {
		fmt.Println("Error en el comando rmdisk")

	}
}
func Analyze_Fdisk(list_tokens []string) {
	var FlagObligatorio bool = true
	//variables del mkdisk
	var size_int int
	var size, fit, unit, drive, name, type_, delete, add string
	//vamos a separar el valor igual

	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-size":
			size = tokens[1]
			size_int, _ = strconv.Atoi(tokens[1])

			fmt.Println("El size es: " + size)
		case "-driveletter":
			drive = tokens[1]
			fmt.Println("El drive es: " + drive)

		case "-name":
			name = tokens[1]
			fmt.Println("El name es: " + name)

		case "-unit":
			unit = tokens[1]
			//pasar aminuscula
			unit = strings.ToLower(unit)
			fmt.Println("El unit es: " + unit)

		case "-type":
			type_ = tokens[1]
			//pasar aminuscula
			type_ = strings.ToLower(type_)
			fmt.Println("El type es: " + type_)

		case "-fit":
			fit = tokens[1]
			//pasar a minuscula
			fit = strings.ToLower(fit)
			fmt.Println("el fit es; " + fit)
		case "-delete":
			delete = tokens[1]
			//pasar a minuscula
			delete = strings.ToLower(delete)
			fmt.Println("el delete es; " + delete)

		case "-add":
			add = tokens[1]
			//pasar a minuscula
			add = strings.ToLower(add)
			fmt.Println("el add es; " + add)
		}

	}
	if size_int == 0 || size_int < 0 {

		fmt.Println("no se encontro el parametro -size o el valor es negativo")
		FlagObligatorio = false

	}
	if drive == "" {
		fmt.Println("no se encontro el parametro -driveletter")
		FlagObligatorio = false
	}
	if name == "" {
		fmt.Println("no se encontro el parametro -name")
		FlagObligatorio = false
	}

	if FlagObligatorio == true {
		// creamos la particion
		fmt.Println("creando particion...")
		CrearParticion(size_int, unit, fit, drive, name, type_, delete, add)

	}

}
func Confirmacion(msg string) bool {
	fmt.Println(msg + "(y/n)")
	var respuesta string
	fmt.Scanln(&respuesta)
	if respuesta == "y" {
		return true
	} else {
		return false

	}

}

func WriteInBytes() {
	fmt.Println("jeje")
}

func CrearDisco(size_int int, unit string, fit string) {

	var size_bytes int64
	var fit_mod string

	if unit == "m" && size_int != 0 {
		size_bytes = int64(size_int * 1024 * 1024)

	} else if unit == "k" && size_int != 0 {
		size_bytes = int64(size_int * 1024)

	} else if unit == "" && size_int != 0 {
		size_bytes = int64(size_int * 1024 * 1024)

	} else {
		fmt.Println("Error en el tamaño del disco")
		return

	}
	if fit != "bf" && fit != "ff" && fit != "wf" {
		fmt.Println("No SE ENCONTRO EL PARAMETRO -fit")
		fit_mod = "F"

	} else {
		if fit == "bf" {
			fit_mod = "B"
		} else if fit == "ff" {
			fit_mod = "F"
		} else if fit == "wf" {
			fit_mod = "W"
		}
	}

	//Si no existe el directorio Discos, entonces crearlo
	if _, err := os.Stat("Discos"); os.IsNotExist(err) {
		err = os.Mkdir("Discos", 0777)
		if err != nil {
			fmt.Println("Error al crear el directorio Discos: ", err)
			return
		}
	}
	//Contar la cantidad de discos para asignar el nombre
	archivos, err := ioutil.ReadDir("Discos")
	if err != nil {
		fmt.Println("Error al leer el directorio: ", err)
		return
	}

	letter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//crear nombre del disco a partir de la cantidad de discos que hay en la carpeta
	nameDisk := string(letter[len(archivos)])

	//crear el archivo binario
	file, err := os.Create("Discos/" + nameDisk + ".dsk")

	if err != nil {
		fmt.Println("error al crear el disco", err)
		return
	}
	defer file.Close()

	//en este apartado emepzamos en la creacion del MBR en el disk

	randomNum := rand.Intn(99) + 1
	var disk MBR

	dateNow := time.Now()
	date := dateNow.Format("2006-01-02 15:04:05")
	disk.MBR_SIZE = (size_bytes)
	disk.MBR_ID = (int64(randomNum))
	copy(disk.MBR_DATE[:], date)

	fitaux := []byte(fit_mod)

	disk.DSK_FIT = [1]byte{fitaux[0]}

	//ESCRIBIMOS EN LAS PARTICIONES
	disk.MBR_PART1.PART_STATUS = [1]byte{'0'}
	disk.MBR_PART2.PART_STATUS = [1]byte{'0'}
	disk.MBR_PART3.PART_STATUS = [1]byte{'0'}
	disk.MBR_PART4.PART_STATUS = [1]byte{'0'}

	disk.MBR_PART1.PART_TYPE = [1]byte{'0'}
	disk.MBR_PART2.PART_TYPE = [1]byte{'0'}
	disk.MBR_PART3.PART_TYPE = [1]byte{'0'}
	disk.MBR_PART4.PART_TYPE = [1]byte{'0'}

	disk.MBR_PART1.PART_FIT = [1]byte{'0'}
	disk.MBR_PART2.PART_FIT = [1]byte{'0'}
	disk.MBR_PART3.PART_FIT = [1]byte{'0'}
	disk.MBR_PART4.PART_FIT = [1]byte{'0'}

	disk.MBR_PART1.PART_START = 0
	disk.MBR_PART2.PART_START = 0
	disk.MBR_PART3.PART_START = 0
	disk.MBR_PART4.PART_START = 0

	disk.MBR_PART1.PART_SIZE = 0
	disk.MBR_PART2.PART_SIZE = 0
	disk.MBR_PART3.PART_SIZE = 0
	disk.MBR_PART4.PART_SIZE = 0

	disk.MBR_PART1.PART_NAME = [16]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
	disk.MBR_PART2.PART_NAME = [16]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
	disk.MBR_PART3.PART_NAME = [16]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
	disk.MBR_PART4.PART_NAME = [16]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}

	//llenamos el archivo en bytes
	bufer := new(bytes.Buffer)
	for i := 0; i < 1024; i++ {
		bufer.WriteByte(0)
	}

	var totalBytes int = 0
	for totalBytes < int(size_bytes) {
		c, err := file.Write(bufer.Bytes())
		if err != nil {
			fmt.Println("Error al escribir en el archivo: ", err)
			return
		}
		totalBytes += c
	}
	fmt.Println("Archivo llenado con 0s")
	//Escribir el MBR en el disco
	file.Seek(0, 0)
	err = binary.Write(file, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al escribir el MBR en el disco: ", err)
		return
	}
	fmt.Println("Disco", nameDisk, "creado con exito")
}

func CrearParticion(size_int int, unit string, fit string, drive string, name string, type_ string, delete string, add string) {
	fmt.Println("Creando pARTICION...")

}
