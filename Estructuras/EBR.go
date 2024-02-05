package Estructuras

type EBR struct {
	EBR_STATUS byte
	EBR_FIT    byte
	EBR_START  byte
	EBR_SIZE   int64
	EBR_NEXT   int64
	EBR_NAME   [16]byte
}

func NewEBR() {
	var ebr EBR
	ebr.EBR_STATUS = '0'
	ebr.EBR_SIZE = 0
	ebr.EBR_NEXT = -1

}
