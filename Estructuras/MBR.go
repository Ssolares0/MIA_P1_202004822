package Estructuras

type MBR struct {
	MBR_SIZE  int64    // 8
	MBR_DATE  [16]byte // 16
	MBR_ID    int64    // 4
	DSK_FIT   [1]byte  // 1
	MBR_PART1 PARTITIONS
	MBR_PART2 PARTITIONS
	MBR_PART3 PARTITIONS
	MBR_PART4 PARTITIONS
}

func NewMBR() MBR {
	return MBR{
		MBR_SIZE:  0,
		MBR_DATE:  [16]byte{},
		MBR_ID:    0,
		DSK_FIT:   [1]byte{'W'},
		MBR_PART1: NewPartition(),
		MBR_PART2: NewPartition(),
		MBR_PART3: NewPartition(),
		MBR_PART4: NewPartition(),
	}

}
