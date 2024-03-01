package Estructuras

type Content struct {
	BName  [12]byte // nombre de carpeta o archivo
	BInodo int64    // apuntador a un inodo carpeta o archivo, -1 por defecto
}

type FolderBlock struct {
	BContent [4]Content
}

func NewFolderBlock() FolderBlock {

	var block FolderBlock
	for i := 0; i < len(block.BContent); i++ {
		block.BContent[i] = NewContent()
	}
	return block
}

func NewContent() Content {
	return Content{
		BInodo: -1,
	}
}
