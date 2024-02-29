package Estructuras

var MountList []*MOUNT

type MOUNT struct {
	Name_part  string
	path_part  string
	ID_part    [4]byte
	type_part  [1]byte
	Size_Part  int64
	Start_part int64
}

func NewMount() *MOUNT {
	return &MOUNT{
		Name_part:  "",
		path_part:  "",
		ID_part:    [4]byte{'0', '0', '0', '0'},
		type_part:  [1]byte{'P'},
		Size_Part:  -1,
		Start_part: -1,
	}
}
