package Estructuras

import (
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

			ReporteSuperBlock(id, ruta, name)
		} else if name == "ebr" {
			ReporteEBR(id, path, name)
		} else {
			fmt.Println("Error: El nombre del reporte es incorrecto")
		}
	}

	//ReporteDisk(name, path, id)

}

func ReporteMBR(id string, path string, name string) {

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
	dot += "}"

	//Crear el archivo .dot
	dotName := path + letter + ".dot"
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
	cmd := exec.Command("dot", "-T", "jpg", dotName, "-o", path)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al generar la imagen: ", err)
		return
	}

	fmt.Println("Reporte generado con exito")
}

func ReporteEBR(id string, path string, name string) {

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

			archivo.Seek(int64(Desplazamiento), 0)
			binary.Read(archivo, binary.LittleEndian, &ebr)

			dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion 1</td></tr>"

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

		}

	}
	if disk.MBR_PART2.PART_TYPE == [1]byte{'E'} && disk.MBR_PART2.PART_SIZE != 0 {
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

			archivo.Seek(int64(Desplazamiento), 0)
			binary.Read(archivo, binary.LittleEndian, &ebr)

			dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion 2</td></tr>"
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

		}
	}
	if disk.MBR_PART3.PART_TYPE == [1]byte{'E'} && disk.MBR_PART3.PART_SIZE != 0 {
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

			archivo.Seek(int64(Desplazamiento), 0)
			binary.Read(archivo, binary.LittleEndian, &ebr)

			dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion 3</td></tr>"
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
		}

	}
	if disk.MBR_PART4.PART_TYPE == [1]byte{'E'} && disk.MBR_PART4.PART_SIZE != 0 {
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

			archivo.Seek(int64(Desplazamiento), 0)
			binary.Read(archivo, binary.LittleEndian, &ebr)
			dot += "<tr><td colspan=\"2\" bgcolor=\"darksalmon\">Particion 4</td></tr>"
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
		}

	}
	dot += "</table>>];"
	dot += "}"

	//Crear el archivo .dot
	dotName := path + letter + ".dot"
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
	cmd := exec.Command("dot", "-T", "jpg", dotName, "-o", path)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al generar la imagen: ", err)
		return
	}

	fmt.Println("Reporte generado con exito")

}

func ReporteDisk(name string, path string, id string) {

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

	//Crear el archivo .dot
	DotName := "Reportes/ReporteDisk.dot"
	archivoDot, err := os.Create(DotName)
	if err != nil {
		fmt.Println("Error al crear el archivo .dot: ", err)
		return
	}
	defer archivoDot.Close()
	_, err = archivoDot.WriteString(Dot)
	if err != nil {
		fmt.Println("Error al escribir el archivo .dot: ", err)
		return
	}
	//Generar la imagen
	cmd := exec.Command("dot", "-T", "pdf", DotName, "-o", "Reportes/ReporteDisk.pdf")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al generar la imagen: ", err)
		return
	}

	fmt.Println("Reporte generado con exito")

}

func ReporteSuperBlock(id string, ruta string, name string) {

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
	err = binary.Read(archivo, binary.BigEndian, &sb)
	if err != nil {
		fmt.Println("Error al abrir el superblock: ", err)
		return
	}
	//leemos el superbloque y mostramos en consola para ver si estan bien los datos
	fmt.Println("Sistema de archivos utilizado", sb.SFilesystemType)
	fmt.Println("fecha que se monto el sistema", sb.SMtime)
	fmt.Println("Tamano Superbloque: ", sb.SBlockS)
	fmt.Println("Numero de inodos: ", sb.SInodesCount)
	fmt.Println("Inicio de tabla de inodos: ", sb.SInodeStart)
	fmt.Println("Numero de bloques: ", sb.SBlocksCount)
	fmt.Println("Bloques libres: ", sb.SFreeBlocksCount)
	fmt.Println("Inodos libres: ", sb.SFreeInodesCount)
	fmt.Println("Primer bloque de datos: ", sb.SBmBlockStart)
	fmt.Println("Tamano del inodo: ", sb.SInodeS)
	fmt.Println("Numero magico: ", sb.SMagic)
	fmt.Println("inicio de superbloque : ", sb.SBlockStart)
	fmt.Println("inicio de bitmap de bloques : ", sb.SBmBlockStart)
	fmt.Println("inicio de bitmap de inodos : ", sb.SBmInodeStart)
	fmt.Println("posicion del primer inodo libre : ", sb.SFirstIno)
	fmt.Println("posicion del primer bloque libre : ", sb.SFirstBlo)
	formattedDate := time.Unix(0, int64(binary.LittleEndian.Uint64(sb.SMtime[:]))*1000000).String()

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
	dot += formattedDate
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

	//Crear el archivo .dot
	dotName := "Reportes/ReporteSuperBloque.dot"
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
	cmd := exec.Command("dot", "-T", "jpg", dotName, "-o", "Reportes/ReporteSuperBloque.jpg")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al generar la imagen: ", err)
		return
	}
}
