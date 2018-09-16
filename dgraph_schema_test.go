package agora

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	UID  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
}

type persons struct {
	Persons []person `json:"persons,omitempty"`
}

type dgraphResponse struct {
	Persons []person `json:"persons,omitempty"`
}

func TestResolvePerson(t *testing.T) {
	reloadData()
	persons := resolvePerson("Aragorn")
	assert.Equal(t, persons[0].Name, "Aragorn")
}

func resolvePerson(name string) []person {
	conn := Dial()
	defer conn.Close()

	query := fmt.Sprintf(`query Person{
		persons(func: eq(name, "%s")) {
			uid
			name
		}
	}`, name)

	j := QueryDgraph(conn, query)

	var r dgraphResponse
	err := json.Unmarshal(j, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r.Persons
}

func reloadData() {
	DropAll()
	SetSchema(`name: string @index(term) .`)
	loadSeed()
}

func loadSeed() {
	conn := Dial()
	defer conn.Close()

	for _, p := range fellowship {
		j, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		_ = MutateDgraph(conn, j)
	}
}

var fellowship = []person{
	person{
		Name: "Aragorn",
	},
	person{
		Name: "Boromir",
	},
	person{
		Name: "Frodo",
	},
	person{
		Name: "Gandalf",
	},
	person{
		Name: "Gimli",
	},
	person{
		Name: "Legolas",
	},
	person{
		Name: "Merry",
	},
	person{
		Name: "Pippin",
	},
	person{
		Name: "Samwise",
	},
}
