package main

import (
	"fmt"
	"github.com/powellquiring/patterns/noop"
	"github.com/powellquiring/solution"
	"gopkg.in/yaml.v2"
	"log"
)

func main() {
	container := solution.Container{
		Image:       "powellquiring/pfq",
		Ports:       []solution.Port{solution.Port{HostPort: 80, ContainerPort: 5000}},
		Environment: map[string]string{"RACK_ENV": "development"},
	}
	solution := solution.Solution{"pfq": container}

	fmt.Println(solution)
	fmt.Println("hello")

	out, err := noop.Yaml(solution)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", out)

	fig := `web:
  build: .
  command: python app.py
  links:
   - db
  ports:
   - "8000:8000"
db:
  image: postgres
`

	fmt.Println(fig)
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(fig), m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%#v\n\n", m)
}
