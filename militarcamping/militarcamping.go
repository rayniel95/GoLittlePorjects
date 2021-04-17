package militarcamping

import "errors"

type militar struct {
	grado        int
	subordinados []*militar
	superior     *militar
	name         string
	ordenDada    interface{}
	dadaPor      *militar
}

type IJerarquia interface {
	asignaJefe(jefe string, subordinado string)
	esSuperior(superior string, subordinado string) bool
	ordena(superior string, subordinado string, orden interface{})
	tieneOrden(militar string) bool
	ordenesPorCumplir() []interface{}
	existe(militar string) bool
	orden(militar string, superior *string) interface{}
}

func (oficial *militar) asignaJefe(jefe string, subordinado string) error {
	superior := oficial.busca(jefe)
	inferior := oficial.busca(subordinado)

	if inferior == nil && superior != nil {
		nuevoSubordinado := &militar{
			grado:        (*superior).grado + 1,
			subordinados: make([]*militar, 0),
			superior:     superior,
			name:         subordinado,
		}
		(*superior).subordinados = append((*superior).subordinados, nuevoSubordinado)
		return nil
	}
	return errors.New("InvalidOperationException")
}

func (oficial *militar) esSuperior(superior string, subordinado string) (bool, error) {
	jefe := oficial.busca(superior)
	inferior := oficial.busca(subordinado)

	if jefe == nil || inferior == nil {
		return false, errors.New("InvalidOperationException")
	}
	node := oficial.busca(subordinado)
	for ; (*node).superior != nil; node = (*node).superior {
		if (*(*node).superior).name == superior {
			return true, nil
		}
	}
	return false, nil
}

func (oficial *militar) orden(subordinado string, superior *string) (interface{}, error) {
	sub := oficial.busca(subordinado)
	if sub == nil || (*sub).dadaPor == nil {
		errors.New("InvalidOperationException")
	}
	*superior = (*((*sub).dadaPor)).name
	return (*sub).orden, nil
}

func (oficial *militar) tieneOrden(subordinado string) (bool, error) {
	sub := oficial.busca(subordinado)
	if sub == nil {
		errors.New("InvalidOperationException")
	}
	return (*sub).dadaPor != nil, nil
}

func (oficial *militar) existe(subordinado string) bool {
	sub := oficial.busca(subordinado)
	return sub != nil
}

func (oficial *militar) ordena(
	superior string, integrante string, orden interface{},
) error {
	jefe := oficial.busca(superior)
	subordinado := oficial.busca(integrante)

	if jefe == nil || subordinado == nil {
		return errors.New("InvalidOperationException")
	}
	value, err := oficial.esSuperior((*jefe).name, (*subordinado).name)
	if err != nil {
		return errors.New("InvalidOperationException")
	}

	if value {
		(*subordinado).ordenDada = orden
		(*subordinado).dadaPor = jefe
	}
	return nil
}

func (militar *militar) busca(name string) *militar {
	node := militar
	for ; (*node).superior != nil; node = (*node).superior {
	}
	return node.searchSon(name)
}

func (node *militar) searchSon(name string) *militar {
	if (*node).name == name {
		return node
	}
	for _, value := range (*node).subordinados {
		son := value.searchSon((*value).name)
		if son != nil {
			return son
		}
	}
	return nil
}

func (oficial *militar) ordenesPorCumplir() []interface{} {
	if len((*oficial).subordinados) == 0 {
		if (*oficial).ordenDada != nil {
			return []interface{}{(*oficial).ordenDada}
		}
		return []interface{}{}
	}
	ordenes := make([]interface{}, 0)
	for _, value := range (*oficial).subordinados {
		ordenesSub := value.ordenesPorCumplir()
		ordenes = append(ordenes, ordenesSub...)
	}
	return ordenes
}

func buildJerarquia(supremo string) *militar {
	return &militar{
		grado:        0,
		subordinados: make([]*militar, 0),
		name:         supremo,
	}
}
