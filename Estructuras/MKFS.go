package Estructuras

type MKFS struct {
	MKFS_Success bool
	MKFS_Id      [4]byte
	MKFS_type    [4]byte
	Filesystem   int64
}
