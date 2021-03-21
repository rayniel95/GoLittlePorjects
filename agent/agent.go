package agent

/*
tratando de recrear IEnumerable y IEnumerator en go, un dolor de cabeza,
todo un reto
*/
type list struct {
	index int
	items []*interface{}
}

func buildList() *list {
	return &list{index: -1, items: make([]*interface{}, 0)}
}

func (list *list) reset() {
	(*list).index = -1
}

func (list *list) moveNext() bool {
	if (*list).index+1 >= len((*list).items) {
		return false
	}
	(*list).index++
	return true
}

func (list *list) current() *interface{} {
	if (*list).index < len((*list).items) {
		return (*list).items[(*list).index]
	}
	return nil
}

type suspectedList struct {
	suspected *IEnumerator
}

func (suspectedList *suspectedList) getEnumerator() *IEnumerator {
	return (*suspectedList).suspected
}

type IEnumerator interface {
	moveNext() bool
	current() *interface{}
	reset()
}

type IEnumerable interface {
	getEnumerator() *IEnumerator
}

type IAgent interface {
	loadData(data []*Person)
	search()
	addData(data string, value string)
	suspect() *IEnumerable
}

type Person struct {
	name     string
	features map[string]string
}

// odio tener que ponerlo en espanol pero no tengo diccionario a mano
type agent struct {
	delincuentes  []*Person
	loadedData    bool
	startedSearch bool
	pistas        map[string]string
}

func (agent *agent) loadData(data []*Person) {
	agent = buildAgent()
	(*agent).delincuentes = data
	(*agent).loadedData = true
}

func (agent *agent) search() {
	if !(*agent).loadedData {
		return
	}
	(*agent).startedSearch = true
	(*agent).pistas = make(map[string]string)
}

func (agent *agent) addData(data string, value string) {
	if !(*agent).startedSearch {
		return
	}
	(*agent).pistas[data] = value
}

func (agent *agent) suspect() *suspectedList {
	var suspect []*Person
	max := 0
	for _, person := range (*agent).delincuentes {
		numberPista := 0
		for _, pista := range (*agent).pistas {
			value, ok := (*person).features[pista]
			if ok && value == (*agent).pistas[pista] {
				numberPista++
			}
		}
		if max != 0 && numberPista == max {
			suspect = append(suspect, person)
		} else if numberPista > max {
			max = numberPista
			suspect = make([]*Person, 0)
			suspect = append(suspect, person)
		}
	}
	return &suspectedList{suspected: buildList()}
}

func buildPerson() *Person {
	return &Person{features: make(map[string]string)}
}

func buildAgent() *agent {
	return &agent{
		delincuentes: make([]*Person, 0),
		pistas:       make(map[string]string),
	}
}
