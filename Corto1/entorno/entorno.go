package entorno

import "fmt"

type Entorno struct {
	Ant   *Entorno
	Tabla map[string]*Simbolo
}

type IEntorno interface {
	Put(string, *Simbolo)
	Get(string)
}

func NewEntorno(ant *Entorno) *Entorno {
	return &Entorno{Ant: ant, Tabla: make(map[string]*Simbolo)}
}

func (ent *Entorno) Put(key string, sim *Simbolo) {
	ent.Tabla[key] = sim
}

func (ent *Entorno) Get(key string) *Simbolo {
	for e := ent; e != nil; e = e.Ant {
		sim, ok := e.Tabla[key]

		if ok {
			return sim
		}
	}

	fmt.Println("variable no ha sido encontrada")
	return nil
}

func Mostrar(ent *Entorno) {
	for _, v := range ent.Tabla {
		fmt.Print("Identificador: ", v.Id)
		fmt.Print("\tTipo: ", v.Tipo)
		fmt.Println("\tDireccion: ", v.Dir)
	}
}
