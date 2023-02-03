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
	query := fmt.Sprintf(`query Person{
		persons(func: eq(name, "%s")) {
			uid
			name
		}
	}`, name)

	j := QueryDgraph(query)

	var r dgraphResponse
	err := json.Unmarshal(j, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r.Persons
}

func reloadData() {
	DropAll()
	SetSchema(`name: string @index(term) .

	type Region {
		regionName: string
		partOf: Region
		coorX: int
		coorY: int
		coorZ: int
	}

	regionName: string .
	partOf: uid .
	coorX: int .
	coorY: int .
	coorZ: int .	
	`)
	loadSeed()
}

func loadSeed() {
	for _, p := range fellowship {
		j, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		MutateDgraph(j)
	}
}

var fellowship = []person{
	{
		Name: "Aragorn",
	},
	{
		Name: "Boromir",
	},
	{
		Name: "Frodo",
	},
	{
		Name: "Gandalf",
	},
	{
		Name: "Gimli",
	},
	{
		Name: "Legolas",
	},
	{
		Name: "Merry",
	},
	{
		Name: "Pippin",
	},
	{
		Name: "Samwise",
	},
}
