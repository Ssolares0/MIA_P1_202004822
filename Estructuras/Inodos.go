package Estructuras

import "time"

type Inode struct {
	IUid   int       // UID del propietario del archivo o carpeta
	IGid   int       // GID del grupo al que pertenece el archivo o carpeta
	IS     int       // tamaño del archivo en bytes, -1 si es una carpeta
	IAtime time.Time // última fecha en que se leyó el inodo sin modificarlo
	ICtime time.Time // fecha en la que se creó el inodo
	IMtime time.Time // última fecha en la que se modificó el inodo
	IBlock [15]int   // -1 si no son utilizados
	IType  byte      // 0: carpeta, 1: archivo
	IPerm  int       // permisos del archivo o carpeta (conjuntos de 3 bits: RWX)
}

func NewInode() Inode {
	return Inode{
		IUid:   -1,
		IGid:   -1,
		IS:     -1,
		IAtime: time.Now(),
		ICtime: time.Now(),
		IMtime: time.Now(),
		IBlock: [15]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		IType:  0,
		IPerm:  0,
	}
}
