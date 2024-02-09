package Estructuras

type PARTITIONS struct {
	PART_STATUS [1]byte
	PART_TYPE   [1]byte
	PART_FIT    [1]byte
	PART_START  int64
	PART_SIZE   int64
	PART_NAME   [16]byte
	PART_CORR   int
	PART_ID     [1]byte
}

func NewPartition() PARTITIONS {
	return PARTITIONS{
		PART_STATUS: [1]byte{},
		PART_TYPE:   [1]byte{},
		PART_FIT:    [1]byte{},
		PART_START:  0,
		PART_SIZE:   0,
		PART_NAME:   [16]byte{},
		PART_CORR:   0,
		PART_ID:     [1]byte{},
	}
}
