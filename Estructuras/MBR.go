package Estructuras

type MBR struct {
	MBR_SIZE int64    // 8
	MBR_DATE [16]byte // 16
	MBR_ID   int64    // 4
	DSK_FIT  [1]byte  // 1

}

func NewMBR() MBR {
	return MBR{
		MBR_SIZE: 0,
		MBR_DATE: [16]byte{},
		MBR_ID:   0,
		DSK_FIT:  [1]byte{},
	}

}
