package noop

import "github.com/powellquiring/solution"
import "gopkg.in/yaml.v2"
import "strconv"
import "strings"

// Produce a single computer fig file for the solution.
// Each container is fronted by an nginx proxy to allow A/B upgrade
//func fig(solution solution.Solution) string {
//	return ""
//}

type localPort solution.Port

func Yaml(sol solution.Solution) (out []byte, err error) {
	//figMap := make(map[string]solution.Container)
	//for _, container := range sol.Containers {
	//	figMap[container.Name] = container
	//}
	//out, err = yaml.Marshal(figMap)
	//if e != nil {
	//	err = e
	//	return
	//}
	//if len(out) == 0 {
	//	out = o
	//} else {
	//	out = append(out, "\n"...)
	//	out = append(out, o...)
	//}
	tmpSolution := sol
	ports = sol.Ports
	sol.Ports = localPort[]{1,2}

	out, err = yaml.Marshal(sol)
	return
}

func (port localPort) MarshalYAML() (ports interface{}, err error) {
	ports = strings.Join([]string{strconv.FormatUint(uint64(port.HostPort), 10),
		strconv.FormatUint(uint64(port.ContainerPort), 10)}, ":")
	return
}
