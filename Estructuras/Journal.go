package Estructuras

type ContenidoJournal struct {
	Tipo_operacion [10]byte
	Path           [100]byte
	Contenido      [100]byte
	Time           [19]byte
}

type Journal struct {
	Journal_size int32
	Journal_last int32
	Journal      [50]ContenidoJournal
}

func NewJournal() Journal {
	return Journal{
		Journal_size: 0,
		Journal_last: 1,
		Journal:      [50]ContenidoJournal{},
	}
}
