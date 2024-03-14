package Estructuras

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

		token_[0] = strings.Replace(token_[0], "#", "", -1)

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
	case "mount":
		Analyze_Mount(token_[1:])
	case "unmount":
		Analyze_Unmount(token_[1:])

	case "mkfs":
		Analyze_mkfs(token_[1:])

	case "login":
		Analyze_Login(token_[1:])
	case "logout":
		Analyze_Logout(token_[1:])

	case "showmount":
		ShowMount()

	case "rep":
		Analyze_Reportes(token_[1:])

	case "pause":
		fmt.Println("Presione enter para continuar")
		fmt.Scanln()

	case "exit":
		//flagExit = true
		fmt.Println("gracias por usar el programa")

	default:
		fmt.Println("error: el comando no existe")

	}

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

		default:
			fmt.Println("no existe el parametro en mkdisk")
		}

	}
	if size_int == 0 {

		fmt.Println("Faltan parametros obligatorios")
		FlagObligatorio = false

	}

	if FlagObligatorio == true {
		//creamos archivo
		fmt.Println("creando disco...")

		CreateNewDisk(size_int, unit, fit)

	}

}

func Analyze_execute(list_tokens []string) {
	var path string

	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-path":
			path = tokens[1]
			fmt.Println("El path es: " + path)

		default:
			fmt.Println("no existe el parametro en execute")

		}

	}
	if path == "" {
		fmt.Println("Falta el parametro -path")
		return

	}

	//Abrimos el archivo

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line != "" {
			lines = append(lines, line)

		}
	}
	readFile.Close()
	for _, line := range lines {
		//fmt.Println("The name is:", line)
		Analyze(line)

	}

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
			nombreArchivo := "MIA/P1/" + tokens[1] + ".dsk" // El nombre o ruta absoluta del archivo
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
	var add_flag, delete_flag = false, false
	//variables del mkdisk
	var size_int int
	var size, fit, unit, drive, name, type_, delete, add string

	add = "0"

	//vamos a separar el valor igual

	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-size":
			size = tokens[1]
			size_int, _ = strconv.Atoi(size)

		case "-driveletter":
			drive = tokens[1]

		case "-name":
			name = tokens[1]

		case "-unit":
			unit = tokens[1]
			//pasar aminuscula
			unit = strings.ToLower(unit)

		case "-type":
			type_ = tokens[1]
			//pasar aminuscula
			type_ = strings.ToLower(type_)

		case "-fit":
			fit = tokens[1]
			//pasar a minuscula
			fit = strings.ToLower(fit)

		case "-delete":
			delete = tokens[1]
			//pasar a minuscula
			delete = strings.ToLower(delete)

			delete_flag = true

		case "-add":
			add = tokens[1]

			fmt.Println("el add es; " + add)
			add_flag = true
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

		CreateFdisk(size_int, unit, fit, drive, name, type_, delete, add, add_flag, delete_flag)

	}

}
func Analyze_Mount(list_tokens []string) {
	fmt.Println(list_tokens)
	var FlagObligatorio bool = true
	var drive, name string

	//vamos a separar el valor igual
	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-driveletter":
			drive = tokens[1]

			fmt.Println("El drive es: " + drive)

		case "-name":
			name = tokens[1]
			fmt.Println("El name es: " + name)
		}

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
		fmt.Println("Montando particion...")
		//montamos la particion
		Mount(drive, name)

	}

}
func Confirmacion(msg string) bool {
	fmt.Println(msg + "(y/n)")
	var respuesta string
	fmt.Scanln(&respuesta)
	if respuesta == "y" {
		return true
	} else if respuesta == "n" {
		return false

	} else {
		fmt.Println("respuesta no valida")
		return false

	}

}

func WriteinBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateNewDisk(size_int int, unit string, fit string) {

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
	if _, err := os.Stat("MIA/P1"); os.IsNotExist(err) {
		err = os.Mkdir("MIA/P1", 0777)
		if err != nil {
			fmt.Println("Error al crear el directorio Discos: ", err)
			return
		}
	}
	//Contar la cantidad de discos para asignar el nombre
	archivos, err := ioutil.ReadDir("MIA/P1")
	if err != nil {
		fmt.Println("Error al leer el directorio: ", err)
		return
	}

	letter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//crear nombre del disco a partir de la cantidad de discos que hay en la carpeta
	nameDisk := string(letter[len(archivos)])

	//crear el archivo binario
	file, err := os.Create("MIA/P1/" + nameDisk + ".dsk")

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

	disk.MBR_PART1.PART_ID = [4]byte{'0', '0', '0', '0'}
	disk.MBR_PART2.PART_ID = [4]byte{'0', '0', '0', '0'}
	disk.MBR_PART3.PART_ID = [4]byte{'0', '0', '0', '0'}
	disk.MBR_PART4.PART_ID = [4]byte{'0', '0', '0', '0'}

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

func CreateFdisk(size_int int, unit string, fit string, drive string, name string, type_ string, delete string, add string, add_flag bool, delete_flag bool) {
	fmt.Println("Creando pARTICION...")

	var size_bytes int64
	var fit_mod string

	if unit == "m" && size_int != 0 {
		size_bytes = int64(size_int * 1024 * 1024)
		unit = "M"

	} else if unit == "k" && size_int != 0 {
		size_bytes = int64(size_int * 1024)
		unit = "K"

	} else if unit == "" && size_int != 0 {
		size_bytes = int64(size_int * 1024)
		unit = "K"
	} else {
		fmt.Println("Error en el tamaño del disco")
		return
	}

	if fit != "bf" && fit != "ff" && fit != "wf" {
		fmt.Println("No SE ENCONTRO EL PARAMETRO -fit")
		fit_mod = "W"

	} else {
		if fit == "bf" {
			fit_mod = "B"
		} else if fit == "ff" {
			fit_mod = "F"
		} else if fit == "wf" {
			fit_mod = "W"
		}
	}

	if type_ != "p" && type_ != "e" && type_ != "l" {
		fmt.Println("No SE ENCONTRO EL PARAMETRO -type")
		type_ = "P"

	} else {
		if type_ == "p" {
			type_ = "P"
		} else if type_ == "e" {
			type_ = "E"
		} else if type_ == "l" {
			type_ = "L"
		}
	}

	if delete_flag {
		if delete != "full" {
			fmt.Println("No SE ENCONTRO EL PARAMETRO -delete")
			return
		}

	}

	//pasamos a entero el valor de add
	var add_int int
	if add_flag {
		var err error
		add_int, err = strconv.Atoi(add)
		if err != nil {
			fmt.Println("Error al convertir el valor de add, el parametro no es valido")
			return
		}
		if add_int <= 0 {
			fmt.Println("como el parametro -add es negativo restamos el tamano de la particion")

		}
	}

	//imprimimos para ver si estan correctos los valores
	fmt.Println("size: ", size_bytes)
	fmt.Println("fit: ", fit_mod)
	fmt.Println("unit: ", unit)
	fmt.Println("drive: ", drive)
	fmt.Println("name: ", name)
	fmt.Println("type: ", type_)
	fmt.Println("delete: ", delete)
	fmt.Println("add: ", add_int)

	//Abrimos el disco
	file, err := os.OpenFile("MIA/P1/"+drive+".dsk", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer file.Close()
	//leemos el mbr
	var disk MBR
	file.Seek(int64(0), 0)
	err = binary.Read(file, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al leer el MBR: ", err)
		return
	}
	if delete == "" && add_int == 0 {
		//no se esta haciendo ninguna operacion, entonces creamos una
		TempD := 1 + binary.Size(MBR{})
		var PartExt PARTITIONS
		indicePart := 0
		var nameRepeat, verificarEspacio bool
		if disk.MBR_PART1.PART_SIZE != 0 {
			if disk.MBR_PART1.PART_TYPE == [1]byte{'E'} {
				PartExt = disk.MBR_PART1

			}
			if strings.Contains(string(disk.MBR_PART1.PART_NAME[:]), name) {
				nameRepeat = true
			}
			TempD += int(disk.MBR_PART1.PART_SIZE) + 1

		} else {
			verificarEspacio = true
			indicePart = 1
		}
		if disk.MBR_PART2.PART_SIZE != 0 {
			if disk.MBR_PART2.PART_TYPE == [1]byte{'E'} {
				PartExt = disk.MBR_PART2

			}
			if strings.Contains(string(disk.MBR_PART2.PART_NAME[:]), name) {
				nameRepeat = true
			}
			TempD += int(disk.MBR_PART2.PART_SIZE) + 1

		} else if !verificarEspacio {
			verificarEspacio = true
			indicePart = 2
		}
		if disk.MBR_PART3.PART_SIZE != 0 {
			if disk.MBR_PART3.PART_TYPE == [1]byte{'E'} {
				PartExt = disk.MBR_PART3

			}
			if strings.Contains(string(disk.MBR_PART3.PART_NAME[:]), name) {
				nameRepeat = true
			}
			TempD += int(disk.MBR_PART3.PART_SIZE) + 1
		} else if !verificarEspacio {
			verificarEspacio = true
			indicePart = 3
		}
		if disk.MBR_PART4.PART_SIZE != 0 {
			if disk.MBR_PART4.PART_TYPE == [1]byte{'E'} {
				PartExt = disk.MBR_PART4

			}
			if strings.Contains(string(disk.MBR_PART4.PART_NAME[:]), name) {
				nameRepeat = true
			}
			TempD += int(disk.MBR_PART4.PART_SIZE) + 1

		} else if !verificarEspacio {
			verificarEspacio = true
			indicePart = 4
		}
		// si el indice sigue estnado en 0 es por que no hau espacio libre
		if indicePart == 0 && type_ != "L" {
			fmt.Println("No hay espacio para crear la particion")
			return

		}
		//verificamos si el nombre de la particion ya existe
		if nameRepeat {
			fmt.Println("Ya existe una particion con ese nombre")
			return
		}
		//validamos que no exista alguna particion extendida
		if type_ == "E" && PartExt.PART_TYPE == [1]byte{'E'} {
			fmt.Println("La particion extendida ya existe")
			return
		}

		if type_ != "L" {
			newPartition := NewPartition()
			newPartition.PART_STATUS = [1]byte{'1'}
			newPartition.PART_TYPE = [1]byte{type_[0]}
			newPartition.PART_FIT = [1]byte{fit_mod[0]}
			newPartition.PART_START = int64(TempD)

			newPartition.PART_SIZE = size_bytes
			copy(newPartition.PART_NAME[:], name)

			if int64(TempD)+newPartition.PART_SIZE+1 > disk.MBR_SIZE {
				fmt.Println("No hay espacio suficiente para crear la particion")
				return
			}
			if indicePart == 1 {
				disk.MBR_PART1 = newPartition
			}
			if indicePart == 2 {
				disk.MBR_PART2 = newPartition
			}
			if indicePart == 3 {
				disk.MBR_PART3 = newPartition
			}
			if indicePart == 4 {
				disk.MBR_PART4 = newPartition
			}
			file.Seek(0, 0)
			binary.Write(file, binary.LittleEndian, &disk)
			file.Close()

			if type_ == "P" {
				fmt.Println("Se creo la particion primaria con exito!!!")

			} else {
				fmt.Println("Se creo la particion extendida con exito!!!")
			}

		} else {
			fmt.Println("Creando particion logica....")

			if PartExt.PART_TYPE != [1]byte{'E'} {
				fmt.Println("No existe una particion extendida")
				return
			}
			//creamos la particion logica
			ebr := NewEBR()
			TempD = int(PartExt.PART_START)
			//leemos ebrs en un for
			for {
				file.Seek(int64(TempD), 0)
				binary.Read(file, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE != 0 {
					if strings.Contains(string(ebr.EBR_NAME[:]), name) {
						fmt.Println("Ya existe una particion logica con ese nombre")
						return
					}
					TempD += int(ebr.EBR_SIZE) + 1 + binary.Size(EBR{})
				}
				if ebr.EBR_NEXT == 0 {
					break
				}
			}
			//creamos un nuevo ebr
			var size_ebr int
			if unit == "K" {
				size_ebr = size_int * 1024
			} else if unit == "M" {
				size_ebr = size_int * 1024 * 1024
			} else {
				size_ebr = size_int
			}
			fmt.Println("el size de la particion logica es: ", size_ebr)
			fmt.Println("el tamano de la extendida es", (PartExt.PART_START + PartExt.PART_SIZE))
			if int64(TempD)+int64(size_ebr)+1 > PartExt.PART_START+PartExt.PART_SIZE {
				fmt.Println("No hay espacio suficiente para crear la particion")
				return
			}
			ebrNEw := NewEBR()
			ebrNEw.EBR_MOUNT = [1]byte{'1'}
			ebrNEw.EBR_FIT = [1]byte{fit_mod[0]}
			ebrNEw.EBR_START = int64(TempD) + 1 + int64(binary.Size(EBR{}))
			ebrNEw.EBR_SIZE = int64(size_ebr)
			ebrNEw.EBR_NEXT = int64(TempD) + 1 + int64(binary.Size(EBR{})) + ebrNEw.EBR_SIZE
			copy(ebrNEw.EBR_NAME[:], name)
			//escribimos el ebr
			file.Seek(int64(TempD), 0)
			binary.Write(file, binary.LittleEndian, &ebrNEw)
			file.Close()
			fmt.Println("Se creo la particion logica con exito!!!")
			return

		}
	}

}
func Mount(drive string, name string) {
	//creamos el id
	//estructura del ID: letra del disco+ correlativo de la particion + ultimos dos digitos del carnet 202004822
	var Id string
	var path string
	var ultimoCaracter byte

	if drive == "" {
		fmt.Println("No se encontro el parametro -driveletter")
		return

	}
	if name == "" {
		fmt.Println("No se encontro el parametro -name")
		return

	}

	//obtenemos el numero del nombre
	name_mod := name

	// Convertir el string a un slice de bytes
	bytes := []byte(name_mod)

	//guardamos toda la ruta
	path = "MIA/P1/" + drive + ".dsk"

	// Obtener el último carácter
	if len(bytes) > 0 {
		ultimoCaracter = bytes[len(bytes)-1]
		fmt.Printf("El último carácter es: %c\n", ultimoCaracter)
	} else {
		fmt.Println("El nombre esta vacio.")
	}
	ultimoCa := string(ultimoCaracter)
	// Asignar valores al ID
	Id = drive + ultimoCa + "22"
	fmt.Println("El id es: ", Id)

	//Abrimos el disco
	file, err := os.OpenFile("MIA/P1/"+drive+".dsk", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer file.Close()
	//leemos el mbr
	var disk MBR
	file.Seek(int64(0), 0)
	err = binary.Read(file, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al leer el MBR: ", err)
		return
	}

	if disk.MBR_SIZE == 0 {
		fmt.Println("El disco no es valido")

		return

	}

	//verificar si ya esta montada la partcion
	for _, element := range MountList {
		if strings.Contains(string(element.Name_part[:]), name) && strings.Contains(string(element.ID_part[:]), Id) {
			fmt.Println("Error: La particion ya esta montada!")

			return
		}

	}
	//primero vemos que partcion se va montar con el correlativo del ID
	if ultimoCaracter == '1' {
		if disk.MBR_PART1.PART_TYPE == [1]byte{'P'} && strings.Contains(string(disk.MBR_PART1.PART_NAME[:]), name) {
			//modificamos el estado de la particion
			disk.MBR_PART1.PART_STATUS = [1]byte{'1'}

			//y modificamos el id
			copy(disk.MBR_PART1.PART_ID[:], Id)

			file.Seek(0, 0)
			binary.Write(file, binary.LittleEndian, &disk)
			//estos valores los almacenamos en la estructura  mount
			mount1 := NewMount()
			mount1.Name_part = name
			mount1.path_part = path
			copy(mount1.ID_part[:], Id)
			mount1.type_part = disk.MBR_PART1.PART_TYPE
			mount1.Size_Part = disk.MBR_PART1.PART_SIZE
			mount1.Start_part = disk.MBR_PART1.PART_START
			MountList = append(MountList, mount1)

			fmt.Println("Se monto la particion1 con exito ID: ", Id)
		}
	} else if ultimoCaracter == '2' {
		if disk.MBR_PART2.PART_TYPE == [1]byte{'P'} && strings.Contains(string(disk.MBR_PART2.PART_NAME[:]), name) {
			//modificamos el estado de la particion
			disk.MBR_PART2.PART_STATUS = [1]byte{'1'}
			//y modificamos el id
			copy(disk.MBR_PART2.PART_ID[:], Id)
			file.Seek(0, 0)
			binary.Write(file, binary.LittleEndian, &disk)
			//estos valores los almacenamos en la estructura  mount
			mount2 := NewMount()
			mount2.Name_part = name
			mount2.path_part = path
			copy(mount2.ID_part[:], Id)
			mount2.type_part = disk.MBR_PART2.PART_TYPE
			mount2.Size_Part = disk.MBR_PART2.PART_SIZE
			mount2.Start_part = disk.MBR_PART2.PART_START
			MountList = append(MountList, mount2)
			//escribimos en una lista nativa de golang

			fmt.Println("Se monto la particion2 con exito ID: ", Id)

		} else {
			fmt.Println("No se encontro la particion")
			return
		}
	} else if ultimoCaracter == '3' {
		if disk.MBR_PART3.PART_TYPE == [1]byte{'P'} && strings.Contains(string(disk.MBR_PART3.PART_NAME[:]), name) {
			//modificamos el estado de la particion
			disk.MBR_PART3.PART_STATUS = [1]byte{'1'}
			//y modificamos el id
			copy(disk.MBR_PART3.PART_ID[:], Id)
			file.Seek(0, 0)
			binary.Write(file, binary.LittleEndian, &disk)

			//estos valores los almacenamos en la estructura  mount
			mount3 := NewMount()
			mount3.Name_part = name
			mount3.path_part = path
			copy(mount3.ID_part[:], Id)
			mount3.type_part = disk.MBR_PART3.PART_TYPE
			mount3.Size_Part = disk.MBR_PART3.PART_SIZE
			mount3.Start_part = disk.MBR_PART3.PART_START
			MountList = append(MountList, mount3)
			//escribimos en una lista nativa de golang

			fmt.Println("Se monto la particion3 con exito ID: ", Id)

		} else {
			fmt.Println("No se encontro la particion ")
			return
		}

	} else if ultimoCaracter == '4' {
		if disk.MBR_PART4.PART_TYPE == [1]byte{'P'} && strings.Contains(string(disk.MBR_PART4.PART_NAME[:]), name) {
			//modificamos el estado de la particion
			disk.MBR_PART4.PART_STATUS = [1]byte{'1'}
			//y modificamos el id
			copy(disk.MBR_PART4.PART_ID[:], Id)
			file.Seek(0, 0)
			binary.Write(file, binary.LittleEndian, &disk)

			//estos valores los almacenamos en la estructura  mount
			mount4 := NewMount()
			mount4.Name_part = name
			mount4.path_part = path
			copy(mount4.ID_part[:], Id)
			mount4.type_part = disk.MBR_PART4.PART_TYPE
			mount4.Size_Part = disk.MBR_PART4.PART_SIZE
			mount4.Start_part = disk.MBR_PART4.PART_START
			MountList = append(MountList, mount4)
			//escribimos en una lista nativa de golang

			fmt.Println("Se monto la particion4 con exito ID: ", Id)

		} else {
			fmt.Println("No se encontro la particion")
			return
		}

	}

	//

}

func Analyze_Unmount(list_tokens []string) {
	var FlagObligatorio bool = true
	var id string

	//vamos a separar el valor igual
	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-id":
			id = tokens[1]

			fmt.Println("El id es: " + id)

		}

	}
	if id == "" {
		fmt.Println("no se encontro el parametro -driveletter")
		FlagObligatorio = false

	}

	if FlagObligatorio == true {
		fmt.Println("Desmontando  particion...")
		//montamos la particion
		UnMount(id)

	}

}

func UnMount(id string) {

	//Obtener el segundo carácter
	segundoCaracter := id[1]
	primerCaracter := id[0]

	segundoNumero, err := strconv.Atoi(string(segundoCaracter))
	if err != nil {
		fmt.Println("Error al convertir el segundo carácter a entero:", err)
		return
	}

	//Abrimos el disco
	file, err := os.OpenFile("MIA/P1/"+string(primerCaracter)+".dsk", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer file.Close()
	//leemos el mbr
	var disk MBR
	file.Seek(int64(0), 0)
	err = binary.Read(file, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al leer el MBR: ", err)
		return
	}

	if disk.MBR_SIZE == 0 {
		fmt.Println("El disco no es valido")

		return

	}

	if segundoNumero == 1 {

		//modificamos el estado de la particion
		disk.MBR_PART1.PART_STATUS = [1]byte{'0'}

	} else if segundoNumero == 2 {
		//modificamos el estado de la particion
		disk.MBR_PART2.PART_STATUS = [1]byte{'0'}
	} else if segundoNumero == 3 {
		//modificamos el estado de la particion
		disk.MBR_PART3.PART_STATUS = [1]byte{'0'}
	} else if segundoNumero == 4 {
		//modificamos el estado de la particion
		disk.MBR_PART4.PART_STATUS = [1]byte{'0'}

	} else {
		fmt.Println("No se encontro la particion")
		return
	}

	//buscamos el id en la lista
	for i, element := range MountList {
		if string(element.ID_part[:]) == id {
			//modificamos el estado de la particion
			MountList[i].Name_part = ""
			MountList[i].path_part = ""
			MountList[i].ID_part = [4]byte{'0', '0', '0', '0'}
			MountList[i].type_part = [1]byte{'P'}
			MountList[i].Size_Part = -1
			MountList[i].Start_part = -1
			fmt.Println("Se desmonto la particion con exito")
			return
		}
	}
	fmt.Println("No se encontro la particion con el id: ", id)

}

func Analyze_mkfs(list_tokens []string) {
	var FlagObligatorio bool = true
	var id, type_, fs string

	//vamos a separar el valor igual
	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-id":
			id = tokens[1]

			fmt.Println("El id es: " + id)

		case "-type":
			type_ = tokens[1]
			type_ = strings.ToLower(type_)
			fmt.Println("El type es: " + type_)

		case "-fs":
			fs = tokens[1]
			fs = strings.ToLower(fs)

		}

	}
	if id == "" {
		fmt.Println("no se encontro el parametro -id")
		FlagObligatorio = false

	}

	if FlagObligatorio == true {
		fmt.Println("Formateando particion...")
		//montamos la particion
		Mkfs(id, type_, fs)

	}

}

func Mkfs(id string, type_ string, fs string) {
	var n float64

	if type_ == "" {
		type_ = "full"
	}
	if fs == "" {
		fs = "2fs"

	}

	//empezamos con el formateo EXT2
	indice := VerificarPartMontada(id)
	part_size := ObtenerTamano(id)

	if indice == -1 {
		fmt.Println("la particion no esta montada: ", id)
		return

	}
	fmt.Println("el indice es; " + strconv.Itoa(indice))
	MountActual := MountList[indice]

	if fs == "2fs" {
		numerador := float64(part_size) - float64(binary.Size(Superblock{}))
		denominador := float64(4 + binary.Size(Inode{}) + 3*binary.Size(FileBlock{}))
		n = math.Floor(numerador / denominador)

	} else {
		numerador := float64(part_size) - float64(binary.Size(Superblock{}))
		denominador := float64(4 + binary.Size(Inode{}) + 3*binary.Size(FolderBlock{}) + binary.Size(Journal{}))
		n = math.Floor(numerador / denominador)
	}

	//parte para crear superblock
	sb := NewSuperblock()
	sb.SInodesCount = int64(n)
	sb.SBlocksCount = int64(n * 3)
	sb.SFreeBlocksCount = int64(n * 3)
	sb.SFreeInodesCount = int64(n)
	fechaActual := time.Now()
	fecha := fechaActual.Format("2006-01-02 15:04:05")
	copy(sb.SMtime[:], fecha)
	copy(sb.SUMtime[:], fecha)
	sb.SMntCount = 1

	if fs == "2fs" {
		Create2fs(sb, MountActual, int64(n))
	}
	if fs == "3fs" {
		Create3fs(sb, MountActual, int64(n))
	}

}

func Create2fs(superblock Superblock, MountActual *MOUNT, n int64) {
	//zeros := 0

	//creamos el superbloque
	superblock.SFilesystemType = 2
	superblock.SBmInodeStart = MountActual.Start_part + int64(binary.Size(Superblock{}))
	superblock.SBmBlockStart = superblock.SBmInodeStart + n
	superblock.SInodeStart = superblock.SBmBlockStart + (3 * n)
	superblock.SBlockStart = superblock.SBmInodeStart + (n * int64(binary.Size(Inode{})))

	superblock.SFreeBlocksCount--
	superblock.SFreeInodesCount--
	superblock.SFreeBlocksCount--
	superblock.SFreeInodesCount--

	file, err := os.OpenFile(MountActual.path_part, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//escribir  el superbloque
	file.Seek(int64(MountActual.Start_part), 0)

	binary.Write(file, binary.LittleEndian, &superblock)

	//Crear el bitmap de inodos
	var llenar byte = 0
	file.Seek(int64(superblock.SBmInodeStart), 0)
	for i := 0; i < int(n); i++ {
		binary.Write(file, binary.LittleEndian, &llenar)
	}

	//Crear el bitmap de bloques
	file.Seek(int64(superblock.SBmBlockStart), 0)
	for i := 0; i < int(n*3); i++ {
		binary.Write(file, binary.LittleEndian, &llenar)
	}

	//Crear el inodo 0
	inodo0 := NewInode()

	//Crear el bloque 0
	var bloque0 FileBlock

	//Formatear inodos
	file.Seek(int64(superblock.SInodeStart), 0)
	for i := 0; i < int(n); i++ {
		binary.Write(file, binary.LittleEndian, &inodo0)
	}

	//Formatear bloques
	file.Seek(int64(superblock.SBlockStart), 0)
	for i := 0; i < int(n*3); i++ {
		binary.Write(file, binary.LittleEndian, &bloque0)
	}

	//Crear el directorio raíz
	//Crear el inodo
	inodo0.IUid = 1
	inodo0.IGid = 1
	fechaActual := time.Now()
	fecha := fechaActual.Format("2006-01-02 15:04:05")
	copy(inodo0.IAtime[:], fecha)
	copy(inodo0.ICtime[:], fecha)
	copy(inodo0.IMtime[:], fecha)
	inodo0.IType = ([1]byte{'0'})
	inodo0.IPerm = 664
	inodo0.IBlock[0] = 0

	//Crear el bloque carpeta

	var bloqueCarpeta FolderBlock
	bloqueCarpeta.BContent[0].BInodo = 0
	copy(bloqueCarpeta.BContent[0].BName[:], ".")
	bloqueCarpeta.BContent[1].BInodo = 0
	copy(bloqueCarpeta.BContent[1].BName[:], "..")
	bloqueCarpeta.BContent[2].BInodo = 1
	copy(bloqueCarpeta.BContent[2].BName[:], "users.txt")
	bloqueCarpeta.BContent[3].BInodo = -1

	data := "1,G,root\n1,U,root,root,123\n"

	//Escribir el inodo y el bloque en el archivo

	inodo1 := NewInode()
	inodo1.IUid = 1
	inodo1.IGid = 1
	fechaActual = time.Now()
	fecha = fechaActual.Format("2006-01-02 15:04:05")
	copy(inodo1.IAtime[:], fecha)
	copy(inodo1.ICtime[:], fecha)
	copy(inodo1.IMtime[:], fecha)
	inodo1.IType = [1]byte{'1'}
	inodo1.IPerm = 664
	inodo1.IBlock[0] = 1
	inodo1.IS = int64(len(data)) + int64(binary.Size(FileBlock{}))

	inodo0.IS = int64(inodo1.IS) + int64(binary.Size(FolderBlock{})) + int64(binary.Size(FolderBlock{}))

	var bloqueArchivo FileBlock
	copy(bloqueArchivo.BContent[:], data)

	//Escribir el inodo en el archivo
	file.Seek(int64(superblock.SBmInodeStart), 0)
	var bit byte = 1
	binary.Write(file, binary.LittleEndian, &bit)
	binary.Write(file, binary.LittleEndian, &bit)

	file.Seek(int64(superblock.SBmBlockStart), 0)
	binary.Write(file, binary.LittleEndian, &bit)
	binary.Write(file, binary.LittleEndian, &bit)

	file.Seek(int64(superblock.SInodeStart), 0)
	binary.Write(file, binary.LittleEndian, &inodo0)
	binary.Write(file, binary.LittleEndian, &inodo1)

	file.Seek(int64(superblock.SBlockStart), 0)
	binary.Write(file, binary.LittleEndian, &bloqueCarpeta)
	binary.Write(file, binary.LittleEndian, &bloqueArchivo)

	fmt.Println("Sistema de archivos 2FS creado con éxito en el disco: ")

}

func Create3fs(superblock Superblock, MountActual *MOUNT, n int64) {
	superblock.SFilesystemType = 3
	superblock.SBmInodeStart = (MountActual.Start_part) + int64(binary.Size(Superblock{})) + int64(binary.Size(Journal{}))
	superblock.SBmBlockStart = (superblock.SBmInodeStart) + n
	superblock.SInodeStart = (superblock.SBmBlockStart) + (3 * n)
	superblock.SBlockStart = superblock.SInodeStart + int64(n*int64(binary.Size(Inode{})))
	//Crear el bloque 0, inodo 0 y el usuario root
	superblock.SFreeBlocksCount--
	superblock.SFreeInodesCount--
	superblock.SFreeBlocksCount--
	superblock.SFreeInodesCount--

	//Creacion Journaling
	var journal Journal

	startJournal := "mkdir"
	rutaJournal := "/"
	contenidoJournaling := "-"

	fechaActual := time.Now()
	fecha := fechaActual.Format("2006-01-02 15:04:05")

	copy(journal.Journal[0].Tipo_operacion[:], startJournal)
	copy(journal.Journal[0].Path[:], rutaJournal)
	copy(journal.Journal[0].Contenido[:], contenidoJournaling)
	copy(journal.Journal[0].Time[:], fecha)

	startJournal = "mkfile"
	rutaJournal = "/users.txt"
	contenidoJournaling = "1,G,root\n1,U,root,root,123\n"

	journal.Journal_size = 2
	journal.Journal_last = 1

	copy(journal.Journal[1].Tipo_operacion[:], startJournal)
	copy(journal.Journal[1].Path[:], rutaJournal)
	copy(journal.Journal[1].Contenido[:], contenidoJournaling)
	copy(journal.Journal[1].Time[:], fecha)

	//Creación del superbloque
	//Abrir el archivo

	file, err := os.OpenFile(MountActual.path_part, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Seek(int64(MountActual.Start_part), 0)
	file.Seek(int64(MountActual.Start_part), 0)
	binary.Write(file, binary.LittleEndian, &superblock)
	binary.Write(file, binary.LittleEndian, &journal)

	//Crear el bitmap de inodos
	var llenar byte = 0
	file.Seek(int64(superblock.SBmInodeStart), 0)
	for i := 0; i < int(n); i++ {
		binary.Write(file, binary.LittleEndian, &llenar)
	}

	//Crear el bitmap de bloques
	file.Seek(int64(superblock.SBmBlockStart), 0)
	for i := 0; i < int(n*3); i++ {
		binary.Write(file, binary.LittleEndian, &llenar)
	}

	inodo0 := NewInode()
	var bloque0 FileBlock

	//Formatear inodos
	file.Seek(int64(superblock.SInodeStart), 0)
	for i := 0; i < int(n); i++ {
		binary.Write(file, binary.LittleEndian, &inodo0)
	}

	//Formatear bloques
	file.Seek(int64(superblock.SBlockStart), 0)
	for i := 0; i < int(n*3); i++ {
		binary.Write(file, binary.LittleEndian, &bloque0)
	}

	//Crear el inodo 0
	inodo0.IUid = 1
	inodo0.IGid = 1
	fechaActual = time.Now()
	fecha = fechaActual.Format("2006-01-02 15:04:05")
	copy(inodo0.IAtime[:], fecha)
	copy(inodo0.ICtime[:], fecha)
	copy(inodo0.IMtime[:], fecha)
	inodo0.IType = [1]byte{'0'}
	inodo0.IPerm = 664
	inodo0.IBlock[0] = 0

	//Crear el bloque carpeta

	var bloqueCarpeta FolderBlock
	bloqueCarpeta.BContent[0].BInodo = 0
	copy(bloqueCarpeta.BContent[0].BName[:], ".")
	bloqueCarpeta.BContent[1].BInodo = 0
	copy(bloqueCarpeta.BContent[1].BName[:], "..")
	bloqueCarpeta.BContent[2].BInodo = 1
	copy(bloqueCarpeta.BContent[2].BName[:], "users.txt")
	bloqueCarpeta.BContent[3].BInodo = -1

	data := "1,G,root\n1,U,root,root,123\n"

	inodo1 := NewInode()
	inodo1.IUid = 1
	inodo1.IGid = 1
	fechaActual = time.Now()
	fecha = fechaActual.Format("2006-01-02 15:04:05")
	copy(inodo1.IAtime[:], fecha)
	copy(inodo1.ICtime[:], fecha)
	copy(inodo1.IMtime[:], fecha)
	inodo1.IType = [1]byte{'1'}
	inodo1.IPerm = 664
	inodo1.IBlock[0] = 1
	inodo1.IS = int64(len(data)) + int64(binary.Size(FileBlock{}))

	inodo0.IS = int64(inodo1.IS) + int64(binary.Size(FolderBlock{})) + int64(binary.Size(FolderBlock{}))

	var bloqueArchivo FileBlock
	copy(bloqueArchivo.BContent[:], data)

	//Escribir el inodo en el archivo
	file.Seek(int64(superblock.SBmInodeStart), 0)
	var bit byte = 1
	binary.Write(file, binary.LittleEndian, &bit)
	binary.Write(file, binary.LittleEndian, &bit)

	file.Seek(int64(superblock.SBmBlockStart), 0)
	binary.Write(file, binary.LittleEndian, &bit)
	binary.Write(file, binary.LittleEndian, &bit)

	file.Seek(int64(superblock.SInodeStart), 0)
	binary.Write(file, binary.LittleEndian, &inodo0)
	binary.Write(file, binary.LittleEndian, &inodo1)

	file.Seek(int64(superblock.SBlockStart), 0)
	binary.Write(file, binary.LittleEndian, &bloqueCarpeta)
	binary.Write(file, binary.LittleEndian, &bloqueArchivo)
	fmt.Println("Sistema de archivos 3FS creado con éxito en el disco: ")

}

func VerificarPartMontada(id string) int {
	//buscamos el id en la lista
	var indice int = -1
	for _, element := range MountList {
		if string(element.ID_part[:]) == id {
			indice = indice + 1
			//devolver el indice

			return indice

		}
	}
	return -1

}
func ObtenerTamano(id string) int {
	//buscamos el id en la lista

	for _, element := range MountList {
		if string(element.ID_part[:]) == id {
			tamano := element.Size_Part
			//devolver el indice

			return int(tamano)

		}
	}
	return int(0)

}

func Analyze_Login(list_tokens []string) {
	var FlagObligatorio bool = true
	var user, password, id string

	//vamos a separar el valor igual
	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-user":
			user = tokens[1]

			fmt.Println("El user es: " + user)

		case "-pass":
			password = tokens[1]
			fmt.Println("El password es: " + password)

		case "-id":
			id = tokens[1]
			fmt.Println("El id es: " + id)

		}

	}
	if user == "" {
		fmt.Println("no se encontro el parametro -user")
		FlagObligatorio = false

	}
	if password == "" {
		fmt.Println("no se encontro el parametro -password")
		FlagObligatorio = false
	}
	if id == "" {
		fmt.Println("no se encontro el parametro -id")
		FlagObligatorio = false

	}

	if FlagObligatorio == true {
		fmt.Println("Logeando...")
		//montamos la particion
		Logged(user, password, id)

	}

}

func Analyze_Logout(list_tokens []string) {
	fmt.Println("Deslogeando...")
	LogOut()

}

func Analyze_mkgrp(list_tokens []string) {
	fmt.Println("Creando grupo...")
	var FlagObligatorio bool = true
	var name string

	//vamos a separar el valor igual
	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-name":
			name = tokens[1]

			fmt.Println("El nombre del grupo es: " + name)

		default:
			fmt.Println("Error: Parametro no reconocido")

		}

	}
	if name == "" {
		fmt.Println("no se encontro el parametro -name")
		FlagObligatorio = false
	}

	if FlagObligatorio == true {

		//montamos la particion
		Mkgrp(name)
	}

}

func Mkgrp(name string) {
	//solo el usuario root puede acceder a este comando
	if Logeado.Uid == -1 {
		fmt.Println("No hay sesion iniciada")
		return

	}
	if Logeado.User != "root" {
		fmt.Println("Solo el usuario root puede crear grupos")
		return

	}
	//buscamos el id en la lista
	partition := VerificarPartMontada(Logeado.Id)

	if partition == -1 {
		fmt.Println("La particion no esta montada")
		return

	}
	MountActual := MountList[partition]

	//abrimos el archivo
	archivo, err := os.OpenFile(MountActual.path_part, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return

	}
	sb := NewSuperblock()
	archivo.Seek(int64(MountList[partition].Start_part), 0)
	err = binary.Read(archivo, binary.LittleEndian, &sb)
	if err != nil {
		fmt.Println("Error al leer el superbloque: ", err)
		return

	}
	inodo := NewInode()
	archivo.Seek(int64(sb.SInodeStart)+int64(binary.Size(Inode{})), 0)
	err = binary.Read(archivo, binary.LittleEndian, &inodo)

	if err != nil {
		fmt.Println("Error al leer el inodo: ", err)
		return

	}
	var barch FileBlock
	txt := ""
	for bloque := 1; bloque < 16; bloque++ {
		if inodo.IBlock[bloque-1] != -1 {
			break

		}
		archivo.Seek(int64(sb.SBlockStart)+int64(binary.Size(FolderBlock{}))+int64(binary.Size(FileBlock{}))*int64(bloque-1), 0)

		err = binary.Read(archivo, binary.LittleEndian, &barch)
		if err != nil {
			fmt.Println("Error al leer el bloque de carpeta: ", err)
			return

		}
		for i := 0; i < len(barch.BContent); i++ {
			if barch.BContent[i] != 0 {
				txt += string(barch.BContent[i])
			}
		}

	}

}

func Logged(user string, password string, id string) bool {

	//verificar si ya hay una sesion iniciada
	if Logeado.Uid != -1 {
		fmt.Println("Ya hay una sesion iniciada,cierre sesion primero")
		return false
	}
	partition := VerificarPartMontada(id)

	MountActual := MountList[partition]
	sb := NewSuperblock()

	if partition == -1 {
		fmt.Println("La particion no esta montada")
		return false

	}
	//abrimos el archivo
	archivo, err := os.OpenFile(MountList[partition].path_part, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return false

	}
	defer archivo.Close()

	archivo.Seek(int64(MountActual.Start_part), 0)

	err = binary.Read(archivo, binary.LittleEndian, &sb)
	if err != nil {
		fmt.Println("Error al leer el superbloque: ", err)
		return false

	}

	if !(sb.SFilesystemType == int64(2) || sb.SFilesystemType == int64(3)) {
		fmt.Println("El sistema de archivos no es 2fs o 3fs, debe formatear primero")
		return false

	}

	// buscamos primero el archivo /users.txt

	numInode := FindFile("/users.txt", *MountActual, sb, archivo)

	if numInode == -1 {
		fmt.Println("No se encontro el archivo /users.txt")
		return false

	}
	//leemos el archivo
	var contenido string

	inodo := NewInode()
	archivo.Seek(int64(sb.SInodeStart+int64(numInode)*int64(binary.Size(Inode{}))), 0)

	err = binary.Read(archivo, binary.LittleEndian, &inodo)
	if err != nil {
		fmt.Println("Error al leer el inodo first: ", err)
		return false

	}

	if inodo.IS == 0 {
		fmt.Println("la particion no tiene nada")
		return false

	}
	for _, i := range inodo.IBlock {

		if i != -1 {
			daspl2 := int64(sb.SBlockStart) + int64(i)*int64(binary.Size(FileBlock{}))
			var block FileBlock
			archivo.Seek(daspl2, 0)
			err = binary.Read(archivo, binary.LittleEndian, &block)
			if err != nil {
				log.Fatal("error al leer el bloque")
				return false
			}
			read := strings.TrimRight(string(block.BContent[:]), string(rune(0)))

			var contFinal string
			cantidadcaracteres := len(read)
			if cantidadcaracteres < 64 {
				contFinal = read
			} else {
				for i := 0; i < 64; i++ {
					contFinal += string(read[i])
					read = read[1:]
				}
			}
			contenido += contFinal
		}

	}

	if contenido == "" {
		fmt.Println("El archivo /users.txt esta vacio")
		return false

	}

	//dividimos el contenido del archivo
	lines := strings.Split(contenido, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		if line[2] == 'U' || line[2] == 'u' {
			//dividimos la linea
			words := strings.Split(line, ",")
			if words[3] == user && words[4] == password {
				uid, _ := strconv.Atoi(words[0])
				Logeado.Uid = int(uid)
				Logeado.User = user
				Logeado.Id = id
				Logeado.Password = password
				Logeado.Grp = (words[2])
				break
			}
		}
		if len(line) == 0 {
			break

		}
		if line[2] == 'G' || line[2] == 'g' {
			//dividimos la linea
			words := strings.Split(line, ",")
			if words[1] == Logeado.Grp {
				gid, _ := strconv.Atoi(words[0])
				Logeado.Gid = int(gid)

			}

		}
	}
	fmt.Println("Sesion iniciada con exito en la particion: " + id + "Usuario: " + user)
	return true
}

func LogOut() bool {
	if Logeado.Uid == -1 {
		fmt.Println("No hay sesion iniciada")
		return false
	}
	Logeado = NewUserActual()
	fmt.Println("-Sesion cerrada con exito-")
	return true
}

func ShowMount() {

	fmt.Println("-----------------------------------Mostrando particiones montadas--------------------------")
	//mostramos la lista de particiones montadas
	for _, element := range MountList {
		fmt.Println("Nombre de la particion: ", string(element.Name_part[:]))
		fmt.Println("Ruta de la particion: ", string(element.path_part[:]))
		fmt.Println("ID de la particion: ", string(element.ID_part[:]))
		fmt.Println("Tipo de la particion: ", element.type_part)
		fmt.Println("Start de la particion: ", element.Start_part)
		fmt.Println("Tamaño de la particion: ", element.Size_Part)
		fmt.Println("----------------------------------------------------------------------------------------------")
	}
	fmt.Println("----------------------------------------------------------------------------------------------")

}
func FindFile(name string, mount MOUNT, sb Superblock, file *os.File) int {
	splitruta := strings.Split(name, "/")
	//buscamos el archivo
	var newruta []string
	for _, ru := range splitruta {
		if ru != "" {
			newruta = append(newruta, ru)
		}

	}
	splitruta = newruta
	inodo := NewInode()
	file.Seek(int64(sb.SInodeStart), 0)
	err := binary.Read(file, binary.LittleEndian, &inodo)
	if err != nil {
		log.Fatal("error al leer el inodo")
		return -1
	}
	numInodo := FindIndiceInodo(inodo, splitruta, sb, file)
	return numInodo
}

func FindIndiceInodo(inodo Inode, splitruta []string, sb Superblock, file *os.File) int {
	count := 0
	if len(splitruta) == 0 {
		return count

	}
	actual := splitruta[0]
	ruta := splitruta[1:]

	for _, i := range inodo.IBlock {
		if i != -1 {
			despl := int64(sb.SBlockStart) + (int64(i) * int64(binary.Size(FileBlock{})))
			file.Seek(despl, 0)
			var fldblock FolderBlock
			err := binary.Read(file, binary.LittleEndian, &fldblock)
			if err != nil {
				log.Fatal("error al leer el bloque")
				return -1
			}
			for _, content := range fldblock.BContent {
				if content.BInodo != -1 && strings.Contains(string(content.BName[:]), actual) {
					if len(ruta) == 0 {
						return int(content.BInodo)
					}
					nextInode := NewInode()
					file.Seek(int64(sb.SInodeStart)+int64(content.BInodo*int64(binary.Size(Inode{}))), 0)
					err := binary.Read(file, binary.LittleEndian, &nextInode)
					if err != nil {
						log.Fatal("error al leer el inodo 222")
						return -1
					}
					return FindIndiceInodo(nextInode, ruta, sb, file)

				}
			}
		}
	}
	return -1
}
func leerBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number) //array de bytes

	_, err := file.Read(bytes) // Leido -> bytes
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func Comparacion(name string, name2 string) bool {
	if strings.ToUpper(name) == strings.ToUpper(name2) {
		return true
	}

	return false
}
