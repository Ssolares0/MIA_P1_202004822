package Estructuras

type PARTITIONS struct {
	PART_SIZE int64

	PART_NAME   [16]byte
	PART_STATUS [1]byte
	PART_TYPE   [1]byte
	PART_FIT    [1]byte
	PART_START  int64
}

func NewPartition() PARTITIONS {
	return PARTITIONS{
		PART_SIZE:   -1,
		PART_NAME:   [16]byte{'~'},
		PART_STATUS: [1]byte{'0'},
		PART_TYPE:   [1]byte{'P'},
		PART_FIT:    [1]byte{'F'},
		PART_START:  -1,
	}
}
