package Estructuras

type Inode struct {
	IUid   int64     // UID del propietario del archivo o carpeta
	IGid   int64     // GID del grupo al que pertenece el archivo o carpeta
	IS     int64     // tamaño del archivo en bytes, -1 si es una carpeta
	IAtime [16]byte  // última fecha en que se leyó el inodo sin modificarlo
	ICtime [16]byte  // fecha en la que se creó el inodo
	IMtime [16]byte  // última fecha en la que se modificó el inodo
	IBlock [16]int64 // -1 si no son utilizados
	IType  int64     // 0: carpeta, 1: archivo
	IPerm  int64     // permisos del archivo o carpeta (conjuntos de 3 bits: RWX)
}

func NewInode() Inode {

	var inode Inode
	inode.IUid = -1
	inode.IGid = -1
	inode.IS = -1
	for i := 0; i < 16; i++ {
		inode.IBlock[i] = -1
	}
	inode.IType = -1
	inode.IPerm = -1
	return inode

}
