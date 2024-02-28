package Estructuras

type FileBlock struct {
	BContent [64]byte
	/* capacidad máxima para un archivo -> 4380 FileBlock máximo en el inodo * 64 -> 280320 bytes */
}

type PointerBlock struct {
	BPointers [16]int
	// Array con pointers a bloques (FolderBlock | FileBlock | PointerBlock (para dobles y triples))
}
