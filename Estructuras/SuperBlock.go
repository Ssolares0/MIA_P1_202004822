package Estructuras

import (
	"unsafe"
)

type Superblock struct {
	SFilesystemType  int64 // 2 | 3 (identifica el sistema de archivos utilizado)
	SInodesCount     int64 // total de inodos
	SBlocksCount     int64 // total de bloques
	SFreeBlocksCount int64
	SFreeInodesCount int64
	SUMtime          [19]byte // última fecha en el que el sistema fue montado
	SMtime           [19]byte // última fecha en el que el sistema fue montado
	SMntCount        int64    // cantidad de veces que el sistema se ha montado
	SMagic           int64    // identifica al sistema de archivos, 0xEF53
	SInodeS          int64    // tamaño del inodo
	SBlockS          int64    // tamaño del bloque
	SFirstIno        int64    // posición del primer inodo libre
	SFirstBlo        int64    // posición del primer bloque libre
	SBmInodeStart    int64    // posición de inicio del bitmap de inodos
	SBmBlockStart    int64    // posición de inicio del bitmap de bloques
	SInodeStart      int64    // posición de inicio de la tabla de inodos
	SBlockStart      int64    // posición de inicio de la tabla de bloques
}

func NewSuperblock() Superblock {
	return Superblock{
		SFilesystemType:  0,
		SInodesCount:     0,
		SBlocksCount:     0,
		SFreeBlocksCount: 0,
		SFreeInodesCount: 0,
		SUMtime:          [19]byte{},
		SMtime:           [19]byte{},
		SMntCount:        0,
		SMagic:           0xEF53,
		SInodeS:          int64(unsafe.Sizeof(Inode{})),
		SBlockS:          int64(unsafe.Sizeof(FolderBlock{})),
		SFirstIno:        0,
		SFirstBlo:        0,
		SBmInodeStart:    0,
		SBmBlockStart:    0,
		SInodeStart:      0,
		SBlockStart:      0,
	}
}
