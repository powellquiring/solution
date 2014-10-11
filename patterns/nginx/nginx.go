package nginx

import "github.com/powellquiring/solution"
import "gopkg.in/yaml.v2"
import "fmt"

// Produce a single computer fig file for the solution.
// Each container is fronted by an nginx proxy to allow A/B upgrade
//func fig(solution solution.Solution) string {
//	return ""
//}

func Doit(solution solution.Solution) {
	type T struct {
		F int `yaml:"a,omitempty"`
		B int
	}
	var t T
	yaml.Unmarshal([]byte("a: 1\nb: 2"), &t)
	fmt.Println(t)
	fmt.Println("1")
	out, err := yaml.Marshal(solution)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
	fmt.Println("2")
}
