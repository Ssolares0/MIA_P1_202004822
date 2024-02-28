package Estructuras

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var MountList []*MOUNT

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
	case "mount":
		Analyze_Mount(token_[1:])
	case "unmount":
		Analyze_Unmount(token_[1:])

	case "mkfs":
		Analyze_mkfs(token_[1:])

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
		fmt.Println("error al leer el comando")

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
	} else {
		return false

	}

}

func WriteInBytes() {
	fmt.Println("jeje")
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
					TempD = int(ebr.EBR_SIZE) + 1 + binary.Size(EBR{})
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

	var n int

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
	MountActual := MountList[1]

	if fs == "2fs" {
		char := byte('1')

		numerator := part_size - binary.Size(Superblock{})
		denominator := 4*binary.Size(char) + binary.Size(Inode{}) + 3*binary.Size(FileBlock{})
		result := numerator / denominator
		n = int(math.Floor(float64(result)))

	} else {
		n = 0
	}

	//parte para crear superblock
	sp := NewSuperblock()
	sp.SInodesCount = n
	sp.SBlocksCount = (n * 3)
	sp.SFreeBlocksCount = (n * 3)
	sp.SFreeInodesCount = (n)

	if fs == "2fs" {
		Create2fs(sp, MountActual, n)
	}

}

func Create2fs(superblock Superblock, MountActual *MOUNT, n int) {
	llenar := byte('0')
	//creamos el superbloque
	superblock.SFilesystemType = 2
	superblock.SBmInodeStart = binary.Size(Superblock{}) + int(MountActual.Start_part)
	superblock.SBmBlockStart = superblock.SBmInodeStart + n
	superblock.SInodeStart = superblock.SBmBlockStart + (3 * n)
	superblock.SBlockStart = superblock.SInodeStart + (n * binary.Size(Inode{}))
	superblock.SFreeBlocksCount--
	superblock.SFreeInodesCount--
	superblock.SFreeInodesCount--
	superblock.SFreeBlocksCount--

	archivo, err := os.OpenFile(MountActual.path_part, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	//escribir  el superbloque
	archivo.Seek(int64(MountActual.Start_part), 0)

	binary.Write(archivo, binary.LittleEndian, &superblock)

	fmt.Println("Se creo el superbloque con exito")

	//crear el bitmap de inodos

	archivo.Seek(int64(superblock.SBmInodeStart), 0)
	//llenamos el bitmap de inodos
	for i := 0; i < n; i++ {
		_, err := archivo.Write([]byte{llenar})
		if err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}

	}
	fmt.Println("Se creo el bitmap de inodos con exito")

	//creamos el bitmap de bloques
	archivo.Seek(int64(superblock.SBmBlockStart), 0)
	//llenamos el bitmap de bloques
	for i := 0; i < (3 * n); i++ {
		_, err := archivo.Write([]byte{llenar})
		if err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}

	}
	fmt.Println("Se creo el bitmap de bloques con exito")
	//crear los inodos
	var inodo Inode
	archivo.Seek(int64(superblock.SInodeStart), 0)
	for i := 0; i < n; i++ {
		binary.Write(archivo, binary.LittleEndian, &inodo)

	}
	defer archivo.Close()

	//creamos los bloques
	var flblock FileBlock
	archivo.Seek(int64(superblock.SBlockStart), 0)
	for i := 0; i < (3 * n); i++ {
		binary.Write(archivo, binary.LittleEndian, &flblock)

	}

	//creamos el user root
	var superblock2 Superblock
	archivo2, err := os.OpenFile(MountActual.path_part, os.O_RDWR, 0644)
	if err != nil {
		panic(err)

	}
	archivo2.Seek(int64(MountActual.Start_part), 0)
	binary.Read(archivo2, binary.LittleEndian, &superblock2)
	defer archivo2.Close()

	inodo.IUid = 1
	inodo.IGid = 1
	inodo.IS = 0
	inodo.IAtime = time.Now()
	inodo.ICtime = time.Now()
	inodo.IMtime = time.Now()
	inodo.IType = 0
	inodo.IPerm = 664
	inodo.IBlock[0] = 0

	//crear el bloque carpeta
	var fldblock FolderBlock
	// Asignar los valores a los elementos del bloque de la carpeta
	fldblock.BContent[0].BName = "."
	fldblock.BContent[0].BInodo = 0
	fldblock.BContent[1].BName = ".."
	fldblock.BContent[1].BInodo = 0
	fldblock.BContent[2].BName = "users.txt"
	fldblock.BContent[2].BInodo = 1

	// Crear un string para usuarioRoot
	usuarioRoot := "1,G,root\n1,U,root,root,123\n"

	//crear inodo tipo 1 archivo
	var inodo2 Inode
	inodo2.IUid = 1
	inodo2.IGid = 1
	inodo2.IS = len(usuarioRoot) + binary.Size(FileBlock{})
	inodo2.IAtime = time.Now()
	inodo2.ICtime = time.Now()
	inodo2.IMtime = time.Now()
	inodo2.IType = 1
	inodo2.IPerm = 664
	inodo2.IBlock[0] = 1

	//cambiamos a inodo valor
	FolderBlockSize := binary.Size(FolderBlock{})
	InodesSize := binary.Size(Inode{})
	inodo.IS = inodo2.IS + FolderBlockSize + InodesSize

	//creamos el bloque del archivo
	var flblock2 FileBlock
	copy(flblock2.BContent[:], usuarioRoot)

	//abrimos el archivo
	archivo3, err := os.OpenFile(MountActual.path_part, os.O_RDWR, 0644)

	if err != nil {
		panic(err)

	}
	archivo3.Seek(int64(superblock.SBmInodeStart), 0)
	bit := byte('1')
	binary.Write(archivo3, binary.LittleEndian, &bit)
	binary.Write(archivo3, binary.LittleEndian, &bit)

	archivo3.Seek(int64(superblock.SBlockStart), 0)
	binary.Write(archivo3, binary.LittleEndian, &bit)
	binary.Write(archivo3, binary.LittleEndian, &bit)

	archivo3.Seek(int64(superblock.SInodeStart), 0)
	binary.Write(archivo3, binary.LittleEndian, &inodo)
	binary.Write(archivo3, binary.LittleEndian, &inodo2)

	archivo3.Seek(int64(superblock.SInodeStart), 0)
	binary.Write(archivo3, binary.LittleEndian, &fldblock)
	binary.Write(archivo3, binary.LittleEndian, &flblock)
	defer archivo3.Close()

	fmt.Println("se creo el sistema de archivos ext2")

}

func VerificarPartMontada(id string) int {
	//buscamos el id en la lista
	var indice int = 1
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
