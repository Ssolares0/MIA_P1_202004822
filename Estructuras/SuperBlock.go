package Estructuras

import "time"

type Superblock struct {
	SFilesystemType  int // 2 | 3 (identifica el sistema de archivos utilizado)
	SInodesCount     int // total de inodos
	SBlocksCount     int // total de bloques
	SFreeBlocksCount int
	SFreeInodesCount int
	SMtime           time.Time // última fecha en el que el sistema fue montado
	SUmtime          time.Time // última fecha en que el sistema fue desmontado
	SMntCount        int       // cantidad de veces que el sistema se ha montado
	SMagic           int       // identifica al sistema de archivos, 0xEF53
	SInodeS          int       // tamaño del inodo
	SBlockS          int       // tamaño del bloque
	SFirstIno        int       // posición del primer inodo libre
	SFirstBlo        int       // posición del primer bloque libre
	SBmInodeStart    int       // posición de inicio del bitmap de inodos
	SBmBlockStart    int       // posición de inicio del bitmap de bloques
	SInodeStart      int       // posición de inicio de la tabla de inodos
	SBlockStart      int       // posición de inicio de la tabla de bloques
}

func NewSuperblock() Superblock {
	return Superblock{
		SFilesystemType:  0,
		SInodesCount:     0,
		SBlocksCount:     0,
		SFreeBlocksCount: 0,
		SFreeInodesCount: 0,
		SMtime:           time.Now(),
		SUmtime:          time.Now(),
		SMntCount:        0,
		SMagic:           0,
		SInodeS:          0,
		SBlockS:          0,
		SFirstIno:        0,
		SFirstBlo:        0,
		SBmInodeStart:    0,
		SBmBlockStart:    0,
		SInodeStart:      0,
		SBlockStart:      0,
	}
}
