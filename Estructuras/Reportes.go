package Estructuras

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	//"time"
)

func Analyze_Reportes(list_tokens []string) {
	flagObligatorio := false

	//variables que usaremos para analizar
	var id string
	var path string
	var name string
	var ruta string

	//aqui empezamos el analizaor de reportes

	for x := 0; x < len(list_tokens); x++ {
		tokens := strings.Split(list_tokens[x], "=")
		switch tokens[0] {
		case "-name":
			name = tokens[1]

		case "-path":
			path = tokens[1]

		case "-id":
			id = tokens[1]

		case "ruta":
			ruta = tokens[1]
			fmt.Println("La ruta especial es: " + ruta)

		default:
			fmt.Println("error al leer el comando")

		}

	}
	if id == "" {
		fmt.Println("Error: El id es obligatorio")
		flagObligatorio = true
		return
	}
	if path == "" {
		fmt.Println("Error: El path es obligatorio")
		flagObligatorio = true
		return
	}
	if name == "" {
		fmt.Println("Error: El name es obligatorio")
		flagObligatorio = true
		return

	}

	if flagObligatorio == false {
		//analizamos el nombre
		if name == "disk" {
			ReporteDisk(name, path, id)

		} else if name == "mbr" {

			ReporteMBR(id, path, name)

		} else if name == "sb" {

			ReporteSuperBlock(id, path, name)
		} else if name == "ebr" {
			ReporteEBR(id, path, name)

		} else if name == "inode" {
			ReporteInode(id, path, name)

		} else if name == "block" {
			ReporteBlock(id, path, name)

		} else if name == "bm_inode" {

			ReporteBitmapInode(id, path, name)
		} else if name == "bm_bloc" {
			ReporteBitmapBlock(id, path, name)

		} else if name == "tree" {
			ReporteTree(id, path, name)

		} else {
			fmt.Println("Error: El nombre del reporte es incorrecto")
		}
	}

	//ReporteDisk(name, path, id)

}

func ReporteMBR(id string, path string, name string) {

	var extension string

	// Obtener el primer carácter
	primerCaracter := id[0]
	//pasar astring
	letter := string(primerCaracter)
	//Abrir el disco A
	archivo, err := os.Open("MIA/P1/" + letter + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer archivo.Close()
	disk := NewMBR()
	archivo.Seek(int64(0), 0)
	err = binary.Read(archivo, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al leer el MBR del disco: ", err)
		return
	}
	tamano := strconv.Itoa(int(disk.MBR_SIZE))
	//LEEMOS LOS DATOS DEL MBR Y LOS PONEMS EN GRAPHVIZ

	dot := "digraph { graph [pad=\"0.5\", nodesep=\"0.5\", ranksep=\"2\", splines=\"ortho\"];"
	dot += "node [shape=plain]"
	dot += "rankdir=LR;"

	dot += "Foo [label=<"
	dot += "<table border=\"0\" cellborder=\"1\" cellspacing=\"0\">"
	dot += "<tr><td colspan=\"2\" bgcolor=\"lightblue\">Reporte del MBR</td></tr>"
	dot += "<tr><td>Tamano_MBR</td><td>"
	dot += tamano
	dot += "</td></tr>"
	dot += "<tr><td>Fecha_CreacionMBR</td><td>"
	dot += string(disk.MBR_DATE[:])
	dot += "</td></tr>"
	dot += "<tr><td>Signature MBR</td><td>"
	dot += strconv.Itoa(int(disk.MBR_ID))
	dot += "</td></tr>"
	dot += "<tr><td>Fit MBR</td><td>"
	dot += strconv.Itoa(int(disk.DSK_FIT[0]))
	dot += "</td></tr>"

	if disk.MBR_PART1.PART_SIZE != 0 && disk.MBR_PART1.PART_TYPE == [1]byte{'P'} {
		dot += "<tr><td colspan=\"2\" bgcolor=\"lightblue\">Particion 1 PRIMARIA</td></tr>"
		dot += "<tr><td>Nombre</td><td>"
		dot += string(disk.MBR_PART1.PART_NAME[:])
		dot += "</td></tr>"
		dot += "<tr><td>Tipo</td><td>"
		dot += string(disk.MBR_PART1.PART_TYPE[:])
		dot += "</td></tr>"
		dot += "<tr><td>Inicio</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART1.PART_START))
		dot += "</td></tr>"
		dot += "<tr><td>Tamano</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART1.PART_SIZE))
		dot += "</td></tr>"

	} else if disk.MBR_PART1.PART_SIZE != 0 && disk.MBR_PART1.PART_TYPE == [1]byte{'E'} {
		dot += "<tr><td colspan=\"2\" bgcolor=\"aquamarine1\">Particion 1 EXTENDIDA</td></tr>"
		dot += "<tr><td>Nombre</td><td>"
		dot += string(disk.MBR_PART1.PART_NAME[:])
		dot += "</td></tr>"
		dot += "<tr><td>Tipo</td><td>"
		dot += string(disk.MBR_PART1.PART_TYPE[:])
		dot += "</td></tr>"
		dot += "<tr><td>Inicio</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART1.PART_START))
		dot += "</td></tr>"
		dot += "<tr><td>Tamano</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART1.PART_SIZE))
		dot += "</td></tr>"

	}
	if disk.MBR_PART2.PART_SIZE != 0 && disk.MBR_PART2.PART_TYPE == [1]byte{'P'} {
		dot += "<tr><td colspan=\"2\" bgcolor=\"lightblue\">Particion 2 PRIMARIA</td></tr>"
		dot += "<tr><td>Nombre</td><td>"
		dot += string(disk.MBR_PART2.PART_NAME[:])
		dot += "</td></tr>"
		dot += "<tr><td>Tipo</td><td>"
		dot += string(disk.MBR_PART2.PART_TYPE[:])
		dot += "</td></tr>"
		dot += "<tr><td>Inicio</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART2.PART_START))
		dot += "</td></tr>"
		dot += "<tr><td>Tamano</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART2.PART_SIZE))
		dot += "</td></tr>"

	} else if disk.MBR_PART2.PART_SIZE != 0 && disk.MBR_PART2.PART_TYPE == [1]byte{'E'} {
		dot += "<tr><td colspan=\"2\" bgcolor=\"aquamarine1\">Particion 2 EXTENDIDA</td></tr>"
		dot += "<tr><td>Nombre</td><td>"
		dot += string(disk.MBR_PART2.PART_NAME[:])
		dot += "</td></tr>"
		dot += "<tr><td>Tipo</td><td>"
		dot += string(disk.MBR_PART2.PART_TYPE[:])
		dot += "</td></tr>"
		dot += "<tr><td>Inicio</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART2.PART_START))
		dot += "</td></tr>"
		dot += "<tr><td>Tamano</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART2.PART_SIZE))
		dot += "</td></tr>"
	}

	if disk.MBR_PART3.PART_SIZE != 0 && disk.MBR_PART3.PART_TYPE == [1]byte{'P'} {
		dot += "<tr><td colspan=\"2\" bgcolor=\"lightblue\">Particion 3 PRIMARIA</td></tr>"
		dot += "<tr><td>Nombre</td><td>"
		dot += string(disk.MBR_PART3.PART_NAME[:])
		dot += "</td></tr>"
		dot += "<tr><td>Tipo</td><td>"
		dot += string(disk.MBR_PART3.PART_TYPE[:])
		dot += "</td></tr>"
		dot += "<tr><td>Inicio</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART3.PART_START))
		dot += "</td></tr>"
		dot += "<tr><td>Tamano</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART3.PART_SIZE))
		dot += "</td></tr>"

	} else if disk.MBR_PART3.PART_SIZE != 0 && disk.MBR_PART3.PART_TYPE == [1]byte{'E'} {
		dot += "<tr><td colspan=\"2\" bgcolor=\"aquamarine1\">Particion 3 EXTENDIDA</td></tr>"
		dot += "<tr><td>Nombre</td><td>"
		dot += string(disk.MBR_PART3.PART_NAME[:])
		dot += "</td></tr>"
		dot += "<tr><td>Tipo</td><td>"
		dot += string(disk.MBR_PART3.PART_TYPE[:])
		dot += "</td></tr>"
		dot += "<tr><td>Inicio</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART3.PART_START))
		dot += "</td></tr>"
		dot += "<tr><td>Tamano</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART3.PART_SIZE))
		dot += "</td></tr>"
	}
	if disk.MBR_PART4.PART_SIZE != 0 && disk.MBR_PART4.PART_TYPE == [1]byte{'P'} {
		dot += "<tr><td colspan=\"2\" bgcolor=\"lightblue\">Particion 4 PRIMARIA </td></tr>"
		dot += "<tr><td>Nombre</td><td>"
		dot += string(disk.MBR_PART4.PART_NAME[:])
		dot += "</td></tr>"
		dot += "<tr><td>Tipo</td><td>"
		dot += string(disk.MBR_PART4.PART_TYPE[:])
		dot += "</td></tr>"
		dot += "<tr><td>Inicio</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART4.PART_START))
		dot += "</td></tr>"
		dot += "<tr><td>Tamano</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART4.PART_SIZE))
		dot += "</td></tr>"

	} else if disk.MBR_PART4.PART_SIZE != 0 && disk.MBR_PART4.PART_TYPE == [1]byte{'E'} {
		dot += "<tr><td colspan=\"2\" bgcolor=\"aquamarine1\">Particion 4 EXTENDIDA</td></tr>"
		dot += "<tr><td>Nombre</td><td>"
		dot += string(disk.MBR_PART4.PART_NAME[:])
		dot += "</td></tr>"
		dot += "<tr><td>Tipo</td><td>"
		dot += string(disk.MBR_PART4.PART_TYPE[:])
		dot += "</td></tr>"
		dot += "<tr><td>Inicio</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART4.PART_START))
		dot += "</td></tr>"
		dot += "<tr><td>Tamano</td><td>"
		dot += strconv.Itoa(int(disk.MBR_PART4.PART_SIZE))
		dot += "</td></tr>"
	}
	dot += "</table>>];"

	//LEEMOS LOS DATOS DEL eBR Y LOS PONEMS EN GRAPHVIZ

	dot += "Foo2 [label=<"
	dot += "<table border=\"0\" cellborder=\"1\" cellspacing=\"0\">"
	dot += "<tr><td colspan=\"2\" bgcolor=\"lightblue\">Reporte del EBR</td></tr>"

	if disk.MBR_PART1.PART_TYPE == [1]byte{'E'} && disk.MBR_PART1.PART_SIZE != 0 {
		var ebr EBR
		var count int = 1

		TempD := int64(disk.MBR_PART1.PART_START)
		archivo.Seek(TempD, 0)
		err := binary.Read(archivo, binary.LittleEndian, &ebr)
		if err != nil {
			fmt.Println("Error al leer el EBR:", err)
			return
		}
		if ebr.EBR_SIZE != 0 {
			for {
				dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion logica</td></tr>"

				dot += "<tr><td>Part_name</td><td>"
				dot += string(ebr.EBR_NAME[:])
				dot += "</td></tr>"

				dot += "<tr><td>FIT</td><td>"
				dot += string(ebr.EBR_FIT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_status</td><td>"
				dot += string(ebr.EBR_MOUNT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_type</td><td>"
				dot += "L"
				dot += "</td></tr>"

				dot += "<tr><td>Inicio</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_START))
				dot += "</td></tr>"

				dot += "<tr><td>Tamano</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_SIZE))
				dot += "</td></tr>"
				dot += "<tr><td>Next</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_NEXT))
				dot += "</td></tr>"

				//TempD += int64(ebr.EBR_SIZE) + 1 + int64(binary.Size(EBR{}))
				archivo.Seek(ebr.EBR_NEXT, 0)
				binary.Read(archivo, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE == 0 {
					fmt.Println("Se termino de leer el EBR")
					break
				}
				count++

			}
			fmt.Println("El contador es: ", count)
		}

	}

	if disk.MBR_PART2.PART_TYPE == [1]byte{'E'} && disk.MBR_PART2.PART_SIZE != 0 {

		var ebr EBR
		var count int = 1

		TempD := int64(disk.MBR_PART2.PART_START)
		archivo.Seek(TempD, 0)
		err := binary.Read(archivo, binary.LittleEndian, &ebr)
		if err != nil {
			fmt.Println("Error al leer el EBR:", err)
			return
		}
		if ebr.EBR_SIZE != 0 {
			for {
				dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion logica</td></tr>"

				dot += "<tr><td>Part_name</td><td>"
				dot += string(ebr.EBR_NAME[:])
				dot += "</td></tr>"

				dot += "<tr><td>FIT</td><td>"
				dot += string(ebr.EBR_FIT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_status</td><td>"
				dot += string(ebr.EBR_MOUNT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_type</td><td>"
				dot += "L"
				dot += "</td></tr>"

				dot += "<tr><td>Inicio</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_START))
				dot += "</td></tr>"

				dot += "<tr><td>Tamano</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_SIZE))
				dot += "</td></tr>"
				dot += "<tr><td>Next</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_NEXT))
				dot += "</td></tr>"

				//TempD += int64(ebr.EBR_SIZE) + 1 + int64(binary.Size(EBR{}))
				archivo.Seek(ebr.EBR_NEXT, 0)
				binary.Read(archivo, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE == 0 {
					fmt.Println("Se termino de leer el EBR")
					break
				}
				count++

			}
			fmt.Println("El contador es: ", count)
		}
	}
	if disk.MBR_PART3.PART_TYPE == [1]byte{'E'} && disk.MBR_PART3.PART_SIZE != 0 {
		//Leer el EBR

		var ebr EBR
		var count int = 1

		TempD := int64(disk.MBR_PART3.PART_START)
		archivo.Seek(TempD, 0)
		err := binary.Read(archivo, binary.LittleEndian, &ebr)
		if err != nil {
			fmt.Println("Error al leer el EBR:", err)
			return
		}
		if ebr.EBR_SIZE != 0 {
			for {
				dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion logica</td></tr>"

				dot += "<tr><td>Part_name</td><td>"
				dot += string(ebr.EBR_NAME[:])
				dot += "</td></tr>"

				dot += "<tr><td>FIT</td><td>"
				dot += string(ebr.EBR_FIT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_status</td><td>"
				dot += string(ebr.EBR_MOUNT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_type</td><td>"
				dot += "L"
				dot += "</td></tr>"

				dot += "<tr><td>Inicio</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_START))
				dot += "</td></tr>"

				dot += "<tr><td>Tamano</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_SIZE))
				dot += "</td></tr>"
				dot += "<tr><td>Next</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_NEXT))
				dot += "</td></tr>"

				//TempD += int64(ebr.EBR_SIZE) + 1 + int64(binary.Size(EBR{}))
				archivo.Seek(ebr.EBR_NEXT, 0)
				binary.Read(archivo, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE == 0 {
					fmt.Println("Se termino de leer el EBR")
					break
				}
				count++

			}
			fmt.Println("El contador es: ", count)
		}

	}
	if disk.MBR_PART4.PART_TYPE == [1]byte{'E'} && disk.MBR_PART4.PART_SIZE != 0 {
		var ebr EBR
		var count int = 1

		TempD := int64(disk.MBR_PART4.PART_START)
		archivo.Seek(TempD, 0)
		err := binary.Read(archivo, binary.LittleEndian, &ebr)
		if err != nil {
			fmt.Println("Error al leer el EBR:", err)
			return
		}
		if ebr.EBR_SIZE != 0 {
			for {
				dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion logica</td></tr>"

				dot += "<tr><td>Part_name</td><td>"
				dot += string(ebr.EBR_NAME[:])
				dot += "</td></tr>"

				dot += "<tr><td>FIT</td><td>"
				dot += string(ebr.EBR_FIT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_status</td><td>"
				dot += string(ebr.EBR_MOUNT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_type</td><td>"
				dot += "L"
				dot += "</td></tr>"

				dot += "<tr><td>Inicio</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_START))
				dot += "</td></tr>"

				dot += "<tr><td>Tamano</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_SIZE))
				dot += "</td></tr>"
				dot += "<tr><td>Next</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_NEXT))
				dot += "</td></tr>"

				//TempD += int64(ebr.EBR_SIZE) + 1 + int64(binary.Size(EBR{}))
				archivo.Seek(ebr.EBR_NEXT, 0)
				binary.Read(archivo, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE == 0 {
					fmt.Println("Se termino de leer el EBR")
					break
				}
				count++

			}
			fmt.Println("El contador es: ", count)
		}

	}
	dot += "</table>>];"
	dot += "}"

	//obtenemos los valores de la ruta path

	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)

	//separar la extension de la ruta

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "jpg" {
			extension = "jpg"

		} else if exten[1] == "png" {
			extension = "png"
		} else if exten[1] == "pdf" {
			extension = "pdf"
		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}
	// Verifica si la carpeta existe, y si no, la crea
	if _, err := os.Stat(path + "." + extension); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}

	//Crear el archivo .dot
	CrearGraphviz(path, dot, extension)
	ReporteEBR(id, path, "ebr")
}

func ReporteEBR(id string, path string, name string) {

	var extension string

	//ABRIMOS EL EBR SI ES QUE SE UTILIZO
	// Obtener el primer carácter
	primerCaracter := id[0]
	//pasar astring
	letter := string(primerCaracter)
	//Abrir el disco A
	archivo, err := os.Open("MIA/P1/" + letter + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer archivo.Close()
	var disk MBR
	archivo.Seek(int64(0), 0)
	err = binary.Read(archivo, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al leer el MBR del disco: ", err)
		return
	}
	//LEEMOS LOS DATOS DEL eBR Y LOS PONEMS EN GRAPHVIZ
	dot := "digraph { graph [pad=\"0.5\", nodesep=\"0.5\", ranksep=\"2\", splines=\"ortho\"];"
	dot += "node [shape=plain]"
	dot += "rankdir=LR;"

	dot += "Foo [label=<"
	dot += "<table border=\"0\" cellborder=\"1\" cellspacing=\"0\">"
	dot += "<tr><td colspan=\"2\" bgcolor=\"lightblue\">Reporte del EBR</td></tr>"

	if disk.MBR_PART1.PART_TYPE == [1]byte{'E'} && disk.MBR_PART1.PART_SIZE != 0 {
		var ebr EBR
		var count int = 1

		TempD := int64(disk.MBR_PART1.PART_START)
		archivo.Seek(TempD, 0)
		err := binary.Read(archivo, binary.LittleEndian, &ebr)
		if err != nil {
			fmt.Println("Error al leer el EBR:", err)
			return
		}
		if ebr.EBR_SIZE != 0 {
			for {
				dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion logica</td></tr>"

				dot += "<tr><td>Part_name</td><td>"
				dot += string(ebr.EBR_NAME[:])
				dot += "</td></tr>"

				dot += "<tr><td>FIT</td><td>"
				dot += string(ebr.EBR_FIT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_status</td><td>"
				dot += string(ebr.EBR_MOUNT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_type</td><td>"
				dot += "L"
				dot += "</td></tr>"

				dot += "<tr><td>Inicio</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_START))
				dot += "</td></tr>"

				dot += "<tr><td>Tamano</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_SIZE))
				dot += "</td></tr>"
				dot += "<tr><td>Next</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_NEXT))
				dot += "</td></tr>"

				//TempD += int64(ebr.EBR_SIZE) + 1 + int64(binary.Size(EBR{}))
				archivo.Seek(ebr.EBR_NEXT, 0)
				binary.Read(archivo, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE == 0 {
					fmt.Println("Se termino de leer el EBR")
					break
				}
				count++

			}
			fmt.Println("El contador es: ", count)
		}

	}

	if disk.MBR_PART2.PART_TYPE == [1]byte{'E'} && disk.MBR_PART2.PART_SIZE != 0 {

		var ebr EBR
		var count int = 1

		TempD := int64(disk.MBR_PART2.PART_START)
		archivo.Seek(TempD, 0)
		err := binary.Read(archivo, binary.LittleEndian, &ebr)
		if err != nil {
			fmt.Println("Error al leer el EBR:", err)
			return
		}
		if ebr.EBR_SIZE != 0 {
			for {
				dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion logica</td></tr>"

				dot += "<tr><td>Part_name</td><td>"
				dot += string(ebr.EBR_NAME[:])
				dot += "</td></tr>"

				dot += "<tr><td>FIT</td><td>"
				dot += string(ebr.EBR_FIT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_status</td><td>"
				dot += string(ebr.EBR_MOUNT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_type</td><td>"
				dot += "L"
				dot += "</td></tr>"

				dot += "<tr><td>Inicio</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_START))
				dot += "</td></tr>"

				dot += "<tr><td>Tamano</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_SIZE))
				dot += "</td></tr>"
				dot += "<tr><td>Next</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_NEXT))
				dot += "</td></tr>"

				//TempD += int64(ebr.EBR_SIZE) + 1 + int64(binary.Size(EBR{}))
				archivo.Seek(ebr.EBR_NEXT, 0)
				binary.Read(archivo, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE == 0 {
					fmt.Println("Se termino de leer el EBR")
					break
				}
				count++

			}
			fmt.Println("El contador es: ", count)
		}
	}
	if disk.MBR_PART3.PART_TYPE == [1]byte{'E'} && disk.MBR_PART3.PART_SIZE != 0 {
		//Leer el EBR

		var ebr EBR
		var count int = 1

		TempD := int64(disk.MBR_PART3.PART_START)
		archivo.Seek(TempD, 0)
		err := binary.Read(archivo, binary.LittleEndian, &ebr)
		if err != nil {
			fmt.Println("Error al leer el EBR:", err)
			return
		}
		if ebr.EBR_SIZE != 0 {
			for {
				dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion logica</td></tr>"

				dot += "<tr><td>Part_name</td><td>"
				dot += string(ebr.EBR_NAME[:])
				dot += "</td></tr>"

				dot += "<tr><td>FIT</td><td>"
				dot += string(ebr.EBR_FIT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_status</td><td>"
				dot += string(ebr.EBR_MOUNT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_type</td><td>"
				dot += "L"
				dot += "</td></tr>"

				dot += "<tr><td>Inicio</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_START))
				dot += "</td></tr>"

				dot += "<tr><td>Tamano</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_SIZE))
				dot += "</td></tr>"
				dot += "<tr><td>Next</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_NEXT))
				dot += "</td></tr>"

				//TempD += int64(ebr.EBR_SIZE) + 1 + int64(binary.Size(EBR{}))
				archivo.Seek(ebr.EBR_NEXT, 0)
				binary.Read(archivo, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE == 0 {
					fmt.Println("Se termino de leer el EBR")
					break
				}
				count++

			}
			fmt.Println("El contador es: ", count)
		}

	}
	if disk.MBR_PART4.PART_TYPE == [1]byte{'E'} && disk.MBR_PART4.PART_SIZE != 0 {
		var ebr EBR
		var count int = 1

		TempD := int64(disk.MBR_PART4.PART_START)
		archivo.Seek(TempD, 0)
		err := binary.Read(archivo, binary.LittleEndian, &ebr)
		if err != nil {
			fmt.Println("Error al leer el EBR:", err)
			return
		}
		if ebr.EBR_SIZE != 0 {
			for {
				dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion logica</td></tr>"

				dot += "<tr><td>Part_name</td><td>"
				dot += string(ebr.EBR_NAME[:])
				dot += "</td></tr>"

				dot += "<tr><td>FIT</td><td>"
				dot += string(ebr.EBR_FIT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_status</td><td>"
				dot += string(ebr.EBR_MOUNT[:])
				dot += "</td></tr>"

				dot += "<tr><td>Part_type</td><td>"
				dot += "L"
				dot += "</td></tr>"

				dot += "<tr><td>Inicio</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_START))
				dot += "</td></tr>"

				dot += "<tr><td>Tamano</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_SIZE))
				dot += "</td></tr>"
				dot += "<tr><td>Next</td><td>"
				dot += strconv.Itoa(int(ebr.EBR_NEXT))
				dot += "</td></tr>"

				//TempD += int64(ebr.EBR_SIZE) + 1 + int64(binary.Size(EBR{}))
				archivo.Seek(ebr.EBR_NEXT, 0)
				binary.Read(archivo, binary.LittleEndian, &ebr)
				if ebr.EBR_SIZE == 0 {
					fmt.Println("Se termino de leer el EBR")
					break
				}
				count++

			}
			fmt.Println("El contador es: ", count)
		}

	}
	dot += "</table>>];"
	dot += "}"
	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)

	//separar la extension de la ruta

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "jpg" {
			extension = "jpg"

		} else if exten[1] == "png" {
			extension = "png"
		} else if exten[1] == "pdf" {
			extension = "pdf"
		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}

	// Verifica si la carpeta existe, y si no, la crea
	if _, err := os.Stat(path + "." + extension); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}

	//Crear el archivo .dot
	CrearGraphviz(path, dot, extension)

}

func ReporteDisk(name string, path string, id string) {
	var extension string

	bytesTexto := []byte(id)

	// Accede al primer elemento del slice (primer carácter)
	primerCaracter := bytesTexto[0]
	segundoCaracter := bytesTexto[1]

	// Imprime el resultado
	fmt.Println("El primer carácter es:", string(primerCaracter))
	fmt.Println("El segundo carácter es:", string(segundoCaracter))

	letter := string(primerCaracter)

	//Abrir el disco A
	archivo, err := os.Open("MIA/P1/" + letter + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer archivo.Close()
	disk := NewMBR()
	archivo.Seek(int64(0), 0)
	err = binary.Read(archivo, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al leer el MBR del disco: ", err)
		return
	}
	sizeMBR := int(disk.MBR_SIZE)
	libre := int(disk.MBR_SIZE)

	Dot := "digraph grid {bgcolor=\"antiquewhite\" fontname=\"Comic Sans MS \" label=\" Reporte Disco\"" + letter + "layout=dot "
	Dot += "labelloc = \"t\"edge [weigth=1000 style=dashed color=red4 dir = \"both\" arrowtail=\"open\" arrowhead=\"open\"]"
	Dot += "node[shape=record, color=black]a0[label=\"MBR"

	//creamos el grafico de la particion 1 <---------------------

	if disk.MBR_PART1.PART_SIZE != 0 {
		libre -= int(disk.MBR_PART1.PART_SIZE)
		Dot += "|"
		if disk.MBR_PART1.PART_TYPE == [1]byte{'P'} {
			Dot += "Primaria"
			porcentaje := (float64(disk.MBR_PART1.PART_SIZE) * float64(100)) / float64(sizeMBR)
			Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
		} else {
			libreExtendida := int(disk.MBR_PART1.PART_SIZE)
			Dot += "{Extendida"
			//Leer el EBR
			ebr := NewEBR()
			Desplazamiento := int(disk.MBR_PART1.PART_START)
			archivo.Seek(int64(Desplazamiento), 0)
			err = binary.Read(archivo, binary.LittleEndian, &ebr)
			if err != nil {
				fmt.Println("Error al leer el EBR del disco: ", err)
				return
			}
			if ebr.EBR_SIZE != 0 {
				Dot += "|{"
				PrimerEbr := true
				for {
					if !PrimerEbr {
						Dot += "|EBR"
					} else {
						PrimerEbr = false
						Dot += "EBR"
					}
					Dot += "|Logica"
					porcentaje := (float64(ebr.EBR_SIZE) * float64(100)) / float64(sizeMBR)
					Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
					libre -= int(ebr.EBR_SIZE)

					Desplazamiento += int(ebr.EBR_SIZE) + 1 + binary.Size(EBR{})
					archivo.Seek(int64(Desplazamiento), 0)
					binary.Read(archivo, binary.LittleEndian, &ebr)
					if ebr.EBR_SIZE == 0 {
						break
					}
				}
				if libreExtendida > 0 {
					Dot += "|Libre"
					porcentaje := (float64(libreExtendida) * float64(100)) / float64(sizeMBR)
					Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
				}
				Dot += "}}"
			} else {
				Dot += "|Libre"
				porcentaje := (float64(libreExtendida) * float64(100)) / float64(sizeMBR)
				Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
				Dot += "}"
			}
		}
	}
	//creamos el grafico de la particion 2 <---------------------
	if disk.MBR_PART2.PART_SIZE != 0 {
		libre -= int(disk.MBR_PART2.PART_SIZE)
		Dot += "|"
		if disk.MBR_PART2.PART_TYPE == [1]byte{'P'} {
			Dot += "Primaria"
			porcentaje := (float64(disk.MBR_PART2.PART_SIZE) * float64(100)) / float64(sizeMBR)
			Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
		} else {
			libreExtendida := int(disk.MBR_PART2.PART_SIZE)
			Dot += "{Extendida"
			//Leer el EBR
			ebr := NewEBR()
			Desplazamiento := int(disk.MBR_PART2.PART_START)
			archivo.Seek(int64(Desplazamiento), 0)
			err = binary.Read(archivo, binary.LittleEndian, &ebr)
			if err != nil {
				fmt.Println("Error al leer el EBR del disco: ", err)
				return
			}
			if ebr.EBR_SIZE != 0 {
				Dot += "|{"
				PrimerEbr := true
				for {
					if !PrimerEbr {
						Dot += "|EBR"
					} else {
						PrimerEbr = false
						Dot += "EBR"
					}
					Dot += "|Logica"
					porcentaje := (float64(ebr.EBR_SIZE) * float64(100)) / float64(sizeMBR)
					Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
					libre -= int(ebr.EBR_SIZE)

					Desplazamiento += int(ebr.EBR_SIZE) + 1 + binary.Size(EBR{})
					archivo.Seek(int64(Desplazamiento), 0)
					binary.Read(archivo, binary.LittleEndian, &ebr)
					if ebr.EBR_SIZE == 0 {
						fmt.Println("Se termino de leer el EBR EN REPDISK")
						break
					}
				}
				if libreExtendida > 0 {
					Dot += "|Libre"
					porcentaje := (float64(libreExtendida) * float64(100)) / float64(sizeMBR)
					Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
				}
				Dot += "}}"
			} else {
				Dot += "|Libre"
				porcentaje := (float64(libreExtendida) * float64(100)) / float64(sizeMBR)
				Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
				Dot += "}"
			}
		}
	}
	//creamos el grafico de la particion 3 <---------------------
	if disk.MBR_PART3.PART_SIZE != 0 {
		libre -= int(disk.MBR_PART3.PART_SIZE)
		Dot += "|"
		if disk.MBR_PART3.PART_TYPE == [1]byte{'P'} {
			Dot += "Primaria"
			porcentaje := (float64(disk.MBR_PART3.PART_SIZE) * float64(100)) / float64(sizeMBR)
			Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
		} else {
			libreExtendida := int(disk.MBR_PART3.PART_SIZE)
			Dot += "{Extendida"
			//Leer el EBR
			ebr := NewEBR()
			Desplazamiento := int(disk.MBR_PART3.PART_START)
			archivo.Seek(int64(Desplazamiento), 0)
			err = binary.Read(archivo, binary.LittleEndian, &ebr)
			if err != nil {
				fmt.Println("Error al leer el EBR del disco: ", err)
				return
			}
			if ebr.EBR_SIZE != 0 {
				Dot += "|{"
				PrimerEbr := true
				for {
					if !PrimerEbr {
						Dot += "|EBR"
					} else {
						PrimerEbr = false
						Dot += "EBR"
					}
					Dot += "|Logica"
					porcentaje := (float64(ebr.EBR_SIZE) * float64(100)) / float64(sizeMBR)
					Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
					libre -= int(ebr.EBR_SIZE)

					Desplazamiento += int(ebr.EBR_SIZE) + 1 + binary.Size(EBR{})
					archivo.Seek(int64(Desplazamiento), 0)
					binary.Read(archivo, binary.LittleEndian, &ebr)
					if ebr.EBR_SIZE == 0 {
						break
					}
				}
				if libreExtendida > 0 {
					Dot += "|Libre"
					porcentaje := (float64(libreExtendida) * float64(100)) / float64(sizeMBR)
					Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
				}
				Dot += "}}"
			} else {
				Dot += "|Libre"
				porcentaje := (float64(libreExtendida) * float64(100)) / float64(sizeMBR)
				Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
				Dot += "}"
			}
		}
	}
	//creamos el grafico de la particion 4 <---------------------
	if disk.MBR_PART4.PART_SIZE != 0 {
		libre -= int(disk.MBR_PART4.PART_SIZE)
		Dot += "|"
		if disk.MBR_PART4.PART_TYPE == [1]byte{'P'} {
			Dot += "Primaria"
			porcentaje := (float64(disk.MBR_PART4.PART_SIZE) * float64(100)) / float64(sizeMBR)
			Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
		} else {
			libreExtendida := int(disk.MBR_PART4.PART_SIZE)
			Dot += "{Extendida"
			//Leer el EBR
			ebr := NewEBR()
			Desplazamiento := int(disk.MBR_PART4.PART_START)
			archivo.Seek(int64(Desplazamiento), 0)
			err = binary.Read(archivo, binary.LittleEndian, &ebr)
			if err != nil {
				fmt.Println("Error al leer el EBR del disco: ", err)
				return
			}
			if ebr.EBR_SIZE != 0 {
				Dot += "|{"
				PrimerEbr := true
				for {
					if !PrimerEbr {
						Dot += "|EBR"
					} else {
						PrimerEbr = false
						Dot += "EBR"
					}
					Dot += "|Logica"
					porcentaje := (float64(ebr.EBR_SIZE) * float64(100)) / float64(sizeMBR)
					Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
					libre -= int(ebr.EBR_SIZE)

					Desplazamiento += int(ebr.EBR_SIZE) + 1 + binary.Size(EBR{})
					archivo.Seek(int64(Desplazamiento), 0)
					binary.Read(archivo, binary.LittleEndian, &ebr)
					if ebr.EBR_SIZE == 0 {
						break
					}
				}
				if libreExtendida > 0 {
					Dot += "|Libre"
					porcentaje := (float64(libreExtendida) * float64(100)) / float64(sizeMBR)
					Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
				}
				Dot += "}}"
			} else {
				Dot += "|Libre"
				porcentaje := (float64(libreExtendida) * float64(100)) / float64(sizeMBR)
				Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
				Dot += "}"
			}
		}
	}
	if libre > 0 {
		Dot += "|Libre"
		porcentaje := (float64(libre) * float64(100)) / float64(sizeMBR)
		Dot += "\\n" + fmt.Sprintf("%.2f", porcentaje) + "%\\n"
	}
	Dot += "\"];\n}"

	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)
	//separar la extension de la ruta

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "jpg" {
			extension = "jpg"

		} else if exten[1] == "png" {
			extension = "png"
		} else if exten[1] == "pdf" {
			extension = "pdf"
		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}

	// Verifica si la carpeta existe, y si no, la crea
	if _, err := os.Stat(path + "." + extension); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}

	//Crear el archivo .dot
	CrearGraphviz(path, Dot, extension)

}

func ReporteInode(id string, path string, name string) {
	// Obtener el primer carácter
	primerCaracter := id[0]

	// Obtener el segundo carácter
	segundoCaracter := id[1]

	segundoNumero, err := strconv.Atoi(string(segundoCaracter))
	if err != nil {
		fmt.Println("Error al convertir el segundo carácter a entero:", err)
		return
	}

	//vemos si la particion esta montada
	partition := VerificarPartMontada(id)

	if partition == -1 {
		fmt.Println("La particion no esta montada")
		return

	}
	MountActual := MountList[segundoNumero-1]

	//Abrir el disco A

	archivo, err := os.Open("MIA/P1/" + string(primerCaracter) + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}

	var sb Superblock

	_, err = archivo.Seek(MountActual.Start_part, 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero en el inicio de la partición: ", err)
		return
	}
	err = binary.Read(archivo, binary.LittleEndian, &sb)
	if err != nil {
		fmt.Println("Error al leer el SuperBlock: ", err)
		return
	}
	InodoAnterior := 0
	Dot := "digraph G {\n"
	DireccionInodo := sb.SInodeStart

	_, err = archivo.Seek(DireccionInodo, 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero en la dirección del Inodo: ", err)
		return
	}

	total_inodes := sb.SInodesCount - sb.SFreeInodesCount

	//var inodoAnterior int = 0
	for i := 0; i < int(total_inodes); i++ {
		inode := NewInode()
		err = binary.Read(archivo, binary.LittleEndian, &inode)
		if err != nil {
			fmt.Println("Error al leer el Inodo: ", err)
			return
		}

		// Imprimir la información del inodo
		Dot += "a" + strconv.Itoa(i) + "[shape=none, color=lightgrey, label=<\n"
		Dot += "<TABLE cellspacing=\"3\" cellpadding=\"2\" style=\"rounded\">\n"
		Dot += "<TR><TD>Inodo_" + strconv.Itoa(i) + "</TD><TD></TD></TR>\n"
		Dot += "<TR><TD> i_uid: </TD><TD> " + strconv.Itoa(int(inode.IUid)) + "</TD></TR>\n"
		Dot += "<TR><TD> i_gid: </TD><TD> " + strconv.Itoa(int(inode.IGid)) + "</TD></TR>\n"
		Dot += "<TR><TD> i_size: </TD><TD> " + strconv.Itoa(int(inode.IS)) + "</TD></TR>\n"
		Dot += "<TR><TD> i_atime: </TD><TD> " + string(sb.SMtime[:]) + "</TD></TR>\n"
		Dot += "<TR><TD> i_ctime: </TD><TD> " + string(sb.SMtime[:]) + "</TD></TR>\n"
		Dot += "<TR><TD> i_mtime: </TD><TD> " + string(sb.SMtime[:]) + "</TD></TR>\n"
		Dot += "<TR><TD> i_block: </TD><TD> " + strconv.Itoa(int(inode.IBlock[0])) + "</TD></TR>\n"
		Dot += "<TR><TD> i_type: </TD><TD> " + string(inode.IType[:]) + "</TD></TR>\n"
		Dot += "<TR><TD> i_perm: </TD><TD> " + strconv.Itoa(int(inode.IPerm)) + "</TD></TR>\n"
		Dot += "</TABLE>>]; \n\n"

		if i-1 >= 0 {
			Dot += "a" + strconv.Itoa(InodoAnterior) + "-> a" + strconv.Itoa(i) + "\n\n"

		}

		InodoAnterior = i

	}
	Dot += "}"
	defer archivo.Close()
	var extension string

	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "jpg" {
			extension = "jpg"

		} else if exten[1] == "png" {
			extension = "png"
		} else if exten[1] == "pdf" {
			extension = "pdf"
		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}
	// Verifica si la carpeta existe, y si no, la crea
	if _, err := os.Stat(path + "." + extension); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}

	CrearGraphviz(path, Dot, extension)

}

func ReporteBlock(id string, path string, name string) {
	// Obtener el primer carácter
	primerCaracter := id[0]

	// Obtener el segundo carácter
	segundoCaracter := id[1]

	segundoNumero, err := strconv.Atoi(string(segundoCaracter))
	if err != nil {
		fmt.Println("Error al convertir el segundo carácter a entero:", err)
		return
	}

	//vemos si la particion esta montada
	partition := VerificarPartMontada(id)

	if partition == -1 {
		fmt.Println("La particion no esta montada")
		return

	}
	MountActual := MountList[segundoNumero-1]

	//Abrir el disco A

	archivo, err := os.Open("MIA/P1/" + string(primerCaracter) + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}

	var sb Superblock

	_, err = archivo.Seek(MountActual.Start_part, 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero en el inicio de la partición: ", err)
		return
	}
	err = binary.Read(archivo, binary.LittleEndian, &sb)
	if err != nil {
		fmt.Println("Error al leer el SuperBlock: ", err)
		return
	}
	var inodo Inode
	_, err = archivo.Seek(sb.SInodeStart, 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero en la dirección del Inodo: ", err)
		return
	}
	err = binary.Read(archivo, binary.LittleEndian, &inodo)
	if err != nil {
		fmt.Println("Error al leer el Inodo: ", err)
		return

	}
	//bloqueanterior := 0
	count := 0
	Dot := "digraph G {\n"

	count = 0

	for _, i := range inodo.IBlock {

		if i != -1 {
			Dot += "Bloque" + strconv.Itoa(int(i)) + "[label = <\n"
			Dot += "<table border=\"0\" cellborder=\"1\" cellspacing=\"0\">\n"

			DesplazamientoBloque := sb.SBlockStart + int64(int(i)*binary.Size(FolderBlock{}))
			FolderBlock := NewFolderBlock()
			_, err = archivo.Seek(DesplazamientoBloque, 0)
			binary.Read(archivo, binary.LittleEndian, &FolderBlock)

			if inodo.IType == [1]byte{'0'} {

				Dot += "<tr><td colspan=\"2\" port='0' bgcolor=\"cadetblue1\">Bloque" + strconv.Itoa(int(i)) + "</td></tr>\n"
				Contador2 := 0
				for _, j := range FolderBlock.BContent {
					nam := strings.TrimRight(string(j.BName[:]), string(rune(0)))
					fmt.Println("El nombre del bloque es: ", nam)
					if Contador2 == 0 {
						nam = "."
					}
					if Contador2 == 1 {
						nam = ".."
					}
					if j.BInodo == -1 {
						nam = ""
					}

					//fmt.Println("Nombre: ", nam)
					Dot += "<tr><td>" + nam + "</td><td port='" + strconv.Itoa(Contador2+1) + "'>" + strconv.Itoa(int(j.BInodo)) + "</td></tr>\n"

					Contador2++

				}
				Dot += "</table>>];\n"
				Contador2 = 0

				for _, j := range FolderBlock.BContent {
					if j.BInodo != -1 {
						if j.BName[0] != '.' {
							//Leer el inodo
							Dot += "Bloque" + strconv.Itoa(int(i)) + ":" + strconv.Itoa(Contador2+1) + " -> Inodo" + strconv.Itoa(int(j.BInodo)) + ":0;\n"
							//Buscar el inodo siguiente
							DesplazamientoInodo := int(sb.SInodeStart) + (int(j.BInodo) * binary.Size(Inode{}))
							inodoNext := NewInode()
							archivo.Seek(int64(DesplazamientoInodo), 0)
							binary.Read(archivo, binary.LittleEndian, &inodoNext)

						}
					}
					Contador2++
				}

			} else {
				file := FileBlock{}
				archivo.Seek(int64(DesplazamientoBloque), 0)
				binary.Read(archivo, binary.LittleEndian, &file)
				Dot += "<tr><td colspan=\"2\" port='0' bgcolor=\"cadetblue1\">Bloque" + strconv.Itoa(int(i)) + "</td></tr>\n"
				Dot += "<tr><td>" + strings.TrimRight(string(file.BContent[:]), string(rune(0))) + "</td></tr>\n"
				Dot += "</table>>];\n"

			}

		}
		count++

	}

	Dot += "\n}"

	defer archivo.Close()
	var extension string

	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "jpg" {
			extension = "jpg"

		} else if exten[1] == "png" {
			extension = "png"
		} else if exten[1] == "pdf" {
			extension = "pdf"
		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}
	// Verifica si la carpeta existe, y si no, la crea
	if _, err := os.Stat(path + "." + extension); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}

	CrearGraphviz(path, Dot, extension)

}

func ReporteBitmapInode(id string, path string, name string) {
	// Obtener el primer carácter
	primerCaracter := id[0]

	// Obtener el segundo carácter
	segundoCaracter := id[1]

	segundoNumero, err := strconv.Atoi(string(segundoCaracter))
	if err != nil {
		fmt.Println("Error al convertir el segundo carácter a entero:", err)
		return
	}

	//Abrir el disco A
	MountActual := MountList[segundoNumero-1]

	//fmt.Println("El indice del mount es: ", MountActual)

	archivo, err := os.Open("MIA/P1/" + string(primerCaracter) + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}

	var sb Superblock

	_, err = archivo.Seek(MountActual.Start_part, 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero en el inicio de la partición: ", err)
		return
	}
	err = binary.Read(archivo, binary.LittleEndian, &sb)
	if err != nil {
		fmt.Println("Error al leer el SuperBlock: ", err)
		return
	}
	FileText := ""

	Direccion := sb.SBmInodeStart

	for i := 0; i <= int(sb.SInodesCount); i++ {
		var bit byte = '0'

		_, err := archivo.Seek(Direccion+int64(i), 0)
		if err != nil {
			fmt.Println("Error al mover el puntero :", err)
			return
		}
		_, err = archivo.Read([]byte{bit})
		if err != nil {
			fmt.Println("Error al leer el bit del bitmap de inodos: ", err)
			return
		}

		FileText += "\t"
		FileText += string([]byte{bit})
		//fmt.Println("El bit es: ", bit)

		if (i+1)%20 == 0 {

			FileText += "\n"

		}

	}

	defer archivo.Close()
	//fmt.Println(FileText)
	var extension string

	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "txt" {
			extension = "txt"

		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}
	//crear el txt
	Creartxt(path, FileText, extension)

}

func ReporteBitmapBlock(id string, path string, name string) {
	// Obtener el primer carácter
	primerCaracter := id[0]

	// Obtener el segundo carácter
	segundoCaracter := id[1]

	segundoNumero, err := strconv.Atoi(string(segundoCaracter))
	if err != nil {
		fmt.Println("Error al convertir el segundo carácter a entero:", err)
		return
	}

	//Abrir el disco A
	MountActual := MountList[segundoNumero-1]

	//fmt.Println("El indice del mount es: ", MountActual)

	archivo, err := os.Open("MIA/P1/" + string(primerCaracter) + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}

	var sb Superblock

	_, err = archivo.Seek(MountActual.Start_part, 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero en el inicio de la partición: ", err)
		return
	}
	err = binary.Read(archivo, binary.LittleEndian, &sb)
	if err != nil {
		fmt.Println("Error al leer el SuperBlock: ", err)
		return
	}
	FileText := ""

	Direccion := sb.SBmBlockStart

	for i := 0; i <= int(sb.SBlocksCount); i++ {
		var bit byte = '0'

		_, err := archivo.Seek(Direccion+int64(i), 0)
		if err != nil {
			fmt.Println("Error al mover el puntero :", err)
			return
		}
		_, err = archivo.Read([]byte{bit})
		if err != nil {
			fmt.Println("Error al leer el bit del bitmap de bloques: ", err)
			return
		}

		FileText += "\t"
		FileText += string(bit)
		//fmt.Println("El bit es: ", bit)

		if (i+1)%20 == 0 {

			FileText += "\n"

		}

	}

	defer archivo.Close()
	//fmt.Println(FileText)
	var extension string

	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "txt" {
			extension = "txt"

		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}
	//crear el txt
	Creartxt(path, FileText, extension)
}

func ReporteTree(id string, path string, name string) {
	// Obtener el primer carácter
	primerCaracter := id[0]

	// Obtener el segundo carácter
	segundoCaracter := id[1]

	segundoNumero, err := strconv.Atoi(string(segundoCaracter))
	if err != nil {
		fmt.Println("Error al convertir el segundo carácter a entero:", err)
		return
	}

	//vemos si la particion esta montada
	partition := VerificarPartMontada(id)

	if partition == -1 {
		fmt.Println("La particion no esta montada")
		return

	}
	MountActual := MountList[segundoNumero-1]

	archivo, err := os.Open("MIA/P1/" + string(primerCaracter) + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer archivo.Close()
	sb := NewSuperblock()
	fmt.Println("la posicion de inicio es: ", MountActual.Start_part)
	archivo.Seek(int64(MountActual.Start_part), 0)
	err = binary.Read(archivo, binary.LittleEndian, &sb)
	if err != nil {
		fmt.Println("Error al abrir el superblock: ", err)
		return
	}

	//buscamos inodo
	inodo := NewInode()
	archivo.Seek(int64(sb.SInodeStart), 0)
	err = binary.Read(archivo, binary.LittleEndian, &inodo)
	if err != nil {
		fmt.Println("Error al leer el inodo raiz: ", err)
		return
	}

	//creamos el archivo dot

	Dot := "digraph H {\n"
	Dot += "node [pad=\"0.5\", nodesep=\"0.5\", ranksep=\"1\"];\n"
	Dot += "node [shape=plaintext];\n"
	Dot += "graph [bb=\"0,0,352,154\"];\n"
	Dot += "rankdir=LR;\n"
	Dot += RecursividadTree(inodo, sb, archivo, 0)
	Dot += "}"

	var extension string

	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "jpg" {
			extension = "jpg"

		} else if exten[1] == "png" {
			extension = "png"
		} else if exten[1] == "pdf" {
			extension = "pdf"
		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}

	CrearGraphviz(path, Dot, extension)

}

func RecursividadTree(inodo Inode, sb Superblock, archivo *os.File, numeroInodo int) string {
	Dot := "Inodo" + strconv.Itoa(numeroInodo) + "[label = <\n"
	Dot += "<table border=\"0\" cellborder=\"1\" cellspacing=\"0\">\n"
	Dot += "<tr><td bgcolor=\"antiquewhite\">Inodo" + strconv.Itoa(numeroInodo) + "</td></tr>\n"
	Dot += "<tr><td>i_uid</td><td>" + strconv.Itoa(int(inodo.IUid)) + "</td></tr>\n"
	Dot += "<tr><td>i_gid</td><td>" + strconv.Itoa(int(inodo.IGid)) + "</td></tr>\n"
	Dot += "<tr><td>i_size</td><td>" + strconv.Itoa(int(inodo.IS)) + "</td></tr>\n"
	Dot += "<tr><td>i_atime</td><td>" + string(inodo.IAtime[:]) + "</td></tr>\n"
	Dot += "<tr><td>i_ctime</td><td>" + string(inodo.ICtime[:]) + "</td></tr>\n"
	Dot += "<tr><td>i_mtime</td><td>" + string(inodo.IMtime[:]) + "</td></tr>\n"
	Dot += "<tr><td>i_type</td><td>" + string(inodo.IType[:]) + "</td></tr>\n"
	Dot += "<tr><td>i_perm</td><td>" + strconv.Itoa(int(inodo.IPerm)) + "</td></tr>\n"

	count := 0
	for _, i := range inodo.IBlock {
		Dot += "<tr><td>i_block" + strconv.Itoa(count+1) + "</td><td port='" + strconv.Itoa(count+1) + "'>" + strconv.Itoa(int(i)) + "</td></tr>\n"
		count++
	}
	Dot += "</table>>];\n"

	//reseteamos el count
	count = 0
	for _, i := range inodo.IBlock {
		if i != -1 {
			//Leer el bloque
			Dot += "Inodo" + strconv.Itoa(numeroInodo) + ":" + strconv.Itoa(count+1) + " -> Bloque" + strconv.Itoa(int(i)) + ":0;\n"
			Dot += "Bloque" + strconv.Itoa(int(i)) + "[label = <\n"
			Dot += "<table border=\"0\" cellborder=\"1\" cellspacing=\"0\">\n"
			DesplazamientoBloque := int(sb.SBlockStart) + (int(i) * binary.Size(FolderBlock{}))
			fldblock := FolderBlock{}
			archivo.Seek(int64(DesplazamientoBloque), 0)
			binary.Read(archivo, binary.LittleEndian, &fldblock)
			if inodo.IType == [1]byte{'0'} {
				Dot += "<tr><td colspan=\"2\" port='0' bgcolor=\"cadetblue1\">Bloque" + strconv.Itoa(int(i)) + "</td></tr>\n"
				Contador2 := 0
				for _, j := range fldblock.BContent {
					//fmt.Println("Nombre: ", string(j.BName[:]))
					nam := strings.TrimRight(string(j.BName[:]), string(rune(0)))

					if Contador2 == 0 {
						nam = "."
					}
					if Contador2 == 1 {
						nam = ".."
					}
					if j.BInodo == -1 {
						nam = ""
					}
					//fmt.Println("Nombre: ", nam)
					Dot += "<tr><td>" + nam + "</td><td port='" + strconv.Itoa(Contador2+1) + "'>" + strconv.Itoa(int(j.BInodo)) + "</td></tr>\n"
					Contador2++
				}
				Dot += "</table>>];\n"
				Contador2 = 0
				for _, j := range fldblock.BContent {
					if j.BInodo != -1 {
						if j.BName[0] != '.' {
							//Leer el inodo
							Dot += "Bloque" + strconv.Itoa(int(i)) + ":" + strconv.Itoa(Contador2+1) + " -> Inodo" + strconv.Itoa(int(j.BInodo)) + ":0;\n"
							//Buscar el inodo siguiente
							DesplazamientoInodo := int(sb.SInodeStart) + (int(j.BInodo) * binary.Size(Inode{}))
							inodoNext := NewInode()
							archivo.Seek(int64(DesplazamientoInodo), 0)
							binary.Read(archivo, binary.LittleEndian, &inodoNext)
							Dot += RecursividadTree(inodoNext, sb, archivo, int(j.BInodo))
						}
					}
					Contador2++
				}
			} else {
				file := FileBlock{}
				archivo.Seek(int64(DesplazamientoBloque), 0)
				binary.Read(archivo, binary.LittleEndian, &file)
				Dot += "<tr><td colspan=\"2\" port='0' bgcolor=\"cadetblue1\">Bloque" + strconv.Itoa(int(i)) + "</td></tr>\n"
				Dot += "<tr><td>" + strings.TrimRight(string(file.BContent[:]), string(rune(0))) + "</td></tr>\n"
				Dot += "</table>>];\n"
			}

		}
		count++
	}
	return Dot

}
func ReporteSuperBlock(id string, path string, name string) {

	// Obtener el primer carácter
	primerCaracter := id[0]

	// Obtener el segundo carácter
	segundoCaracter := id[1]

	segundoNumero, err := strconv.Atoi(string(segundoCaracter))
	if err != nil {
		fmt.Println("Error al convertir el segundo carácter a entero:", err)
		return
	}

	//Abrir el disco A
	MountActual := MountList[segundoNumero-1]

	archivo, err := os.Open("MIA/P1/" + string(primerCaracter) + ".dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer archivo.Close()
	sb := NewSuperblock()
	fmt.Println("la posicion de inicio es: ", MountActual.Start_part)
	archivo.Seek(int64(MountActual.Start_part), 0)
	err = binary.Read(archivo, binary.LittleEndian, &sb)
	if err != nil {
		fmt.Println("Error al abrir el superblock: ", err)
		return
	}

	//LEEMOS LOS DATOS DEL SUPERBLOQUE Y LOS PONEMS EN GRAPHVIZ
	dot := "digraph { graph [pad=\"0.5\", nodesep=\"0.5\", ranksep=\"2\", splines=\"ortho\"];"
	dot += "node [shape=plain]"
	dot += "rankdir=LR;"

	dot += "Foo [label=<"
	dot += "<table border=\"0\" cellborder=\"1\" cellspacing=\"0\">"
	dot += "<tr><td colspan=\"2\" bgcolor=\"lightblue\">Reporte del SuperBloque</td></tr>"

	dot += "<tr><td>Sistema de archivos</td><td>"
	dot += strconv.Itoa(int(sb.SFilesystemType))
	dot += "</td></tr>"
	dot += "<tr><td>Fecha de montura</td><td>"
	dot += string(sb.SMtime[:])
	dot += "</td></tr>"
	dot += "<tr><td>Tamano Superbloque</td><td>"
	dot += strconv.Itoa(int(sb.SBlockS))
	dot += "</td></tr>"
	dot += "<tr><td>Numero de inodos</td><td>"
	dot += strconv.Itoa(int(sb.SInodesCount))
	dot += "</td></tr>"
	dot += "<tr><td>Inicio de tabla de inodos</td><td>"
	dot += strconv.Itoa(int(sb.SInodeStart))
	dot += "</td></tr>"
	dot += "<tr><td>Numero de bloques</td><td>"
	dot += strconv.Itoa(int(sb.SBlocksCount))
	dot += "</td></tr>"
	dot += "<tr><td>Bloques libres</td><td>"
	dot += strconv.Itoa(int(sb.SFreeBlocksCount))
	dot += "</td></tr>"
	dot += "<tr><td>Inodos libres</td><td>"
	dot += strconv.Itoa(int(sb.SFreeInodesCount))
	dot += "</td></tr>"
	dot += "<tr><td>Primer bloque de datos</td><td>"
	dot += strconv.Itoa(int(sb.SBmBlockStart))
	dot += "</td></tr>"
	dot += "<tr><td>Tamano del inodo</td><td>"
	dot += strconv.Itoa(int(sb.SInodeS))
	dot += "</td></tr>"
	dot += "<tr><td>Numero magico</td><td>"
	dot += strconv.Itoa(int(sb.SMagic))
	dot += "</td></tr>"
	dot += "<tr><td>Inicio de superbloque</td><td>"
	dot += strconv.Itoa(int(sb.SBlockStart))
	dot += "</td></tr>"
	dot += "<tr><td>Inicio de bitmap de bloques</td><td>"
	dot += strconv.Itoa(int(sb.SBmBlockStart))
	dot += "</td></tr>"
	dot += "<tr><td>Inicio de bitmap de inodos</td><td>"
	dot += strconv.Itoa(int(sb.SBmInodeStart))
	dot += "</td></tr>"
	dot += "<tr><td>Posicion del primer inodo libre</td><td>"
	dot += strconv.Itoa(int(sb.SFirstIno))
	dot += "</td></tr>"
	dot += "<tr><td>Posicion del primer bloque libre</td><td>"
	dot += strconv.Itoa(int(sb.SFirstBlo))
	dot += "</td></tr>"
	dot += "</table>>];"
	dot += "}"

	var extension string

	//quitamos comillas por si trae en path
	path = strings.Replace(path, "\"", "", -1)

	if strings.Contains(path, ".") {
		exten := strings.Split(path, ".")
		//guardamos la ruta sin la extension
		path = exten[0]
		fmt.Println("La extension es: " + exten[1])
		if exten[1] == "jpg" {
			extension = "jpg"

		} else if exten[1] == "png" {
			extension = "png"
		} else if exten[1] == "pdf" {
			extension = "pdf"
		} else {
			fmt.Println("Error: La extension del archivo no es valida")
			return
		}

	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return
		}
	}

	CrearGraphviz(path, dot, extension)

}
func CrearGraphviz(path string, dot string, extension string) {
	//Crear el archivo .dot
	dotName := path + ".dot"
	archivoDot, err := os.Create(dotName)
	if err != nil {
		fmt.Println("Error al crear el archivo .dot: ", err)
		return
	}
	defer archivoDot.Close()
	_, err = archivoDot.WriteString(dot)
	if err != nil {
		fmt.Println("Error al escribir el archivo .dot: ", err)
		return
	}
	//Generar la imagen
	cmd := exec.Command("dot", "-T", extension, dotName, "-o", path+"."+extension)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al generar la imagen: ", err)
		return
	}

	fmt.Println("Reporte Super Bloque generado con exito")
}

func Creartxt(path string, FileText string, extension string) {
	txtName := path + "." + extension
	archivoTxt, err := os.Create(txtName)
	if err != nil {
		fmt.Println("Error al crear el archivo .txt: ", err)
		return
	}
	defer archivoTxt.Close()

	//escribimos en el txt
	escribir := bufio.NewWriter(archivoTxt)

	_, err = escribir.WriteString(FileText)
	if err != nil {
		fmt.Println("Error al escribir el archivo .txt: ", err)
		return
	}
	escribir.Flush()
	if err != nil {
		fmt.Println("Error al escribir en el txt: ", err)
		return

	}
	fmt.Println("Reporte Bitmap Inode generado con exito")

}
