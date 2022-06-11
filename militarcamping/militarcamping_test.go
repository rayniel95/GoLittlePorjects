package militarcamping

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	jerarquia := buildJerarquia("Resoplez")
	jerarquia.asignaJefe("Resoplez", "Cuco")

	fmt.Println("Probando asignacion")
	want := true
	got := jerarquia.esSuperior("Resoplez", "Cuco")

	fmt.Println()

	jerarquia.ordena("Resoplez", "Cuco", "Hacer el campamento")
	jerarquia.asignaJefe("Resoplez", "Paco")
	jerarquia.ordena("Cuco", "Paco", "Ayudar")
	jerarquia.asignaJefe("Paco", "Pepe")
	jerarquia.ordena("Paco", "Pepe", "Ensillar las yeguas")
	jerarquia.ordena("Cuco", "Pepe", "Hacer café")
	jerarquia.ordena("Resoplez", "Pepe", "Ayudar a Cuco")
	jerarquia.ordena("Paco", "Pepe", "Volver")

}
