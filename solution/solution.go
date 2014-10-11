package solution

type Solution map[string]Container

type Container struct {
	Image       string
	Ports       []Port
	Environment map[string]string
}
type Port struct {
	HostPort      uint16
	ContainerPort uint16
}
