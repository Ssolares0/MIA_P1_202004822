package Estructuras

type Content struct {
	BName  string // nombre de carpeta o archivo
	BInodo int    // apuntador a un inodo carpeta o archivo, -1 por defecto
}

type FolderBlock struct {
	BContent [4]Content
}
