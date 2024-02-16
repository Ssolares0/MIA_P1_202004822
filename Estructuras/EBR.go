package Estructuras

type EBR struct {
	EBR_MOUNT [1]byte
	EBR_FIT   [1]byte
	EBR_START int64
	EBR_SIZE  int64
	EBR_NEXT  int64
	EBR_NAME  [16]byte
}

func NewEBR() EBR {
	return EBR{
		EBR_MOUNT: [1]byte{0},
		EBR_FIT:   [1]byte{'W'},
		EBR_START: -1,
		EBR_SIZE:  0,
		EBR_NEXT:  -1,
		EBR_NAME:  [16]byte{'~', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
	}
}
