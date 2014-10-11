package solution

import "testing"
import "fmt"

func testSolution1(t *testing.T) {
	container := Container{name: "nginx"}
	ports := [...]Port{}
	t.Log("hello")
	fmt.Println(container)
	fmt.Println(ports)
	fmt.Println("hello")
	//	container := Container{name: "powellquiring/nginx", ports: [...]Port{Port{80, 80}}}
	//solution := Solution{name: "apacheab", containers: [...]Container{container}}
}
