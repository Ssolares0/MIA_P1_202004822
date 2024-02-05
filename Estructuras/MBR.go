package Estructuras

type MBR struct {
	MBR_SIZE int64
	MBR_DATE [16]byte
	MBR_ID   int64
	//MBR_FIT  char[2]

}

func NewMBR() {
	var nmb MBR
	nmb.MBR_SIZE = 1000
	//nmb.MBR_DATE = '2020'
	nmb.MBR_ID = 10
}
