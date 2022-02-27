package entorno

type Simbolo struct {
	Id   string
	Tipo string
	Dir  int
}

func NewSimbolo(id string, tipo string, dir int) *Simbolo {
	return &Simbolo{Id: id, Tipo: tipo, Dir: dir}
}
