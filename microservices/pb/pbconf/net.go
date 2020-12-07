package pbconf

type service struct {
	Name string
}

func (s *service) GetAddr() string {
	return s.GetHost() + s.GetPort()
}

func (s *service) GetPort() (port string) {
	return devPorts[s.Name]
}

func (s *service) GetHost() (h string) {
	return s.Name + "_svc"
}

var (
	Films = &service{Name: "films"}
)

var devPorts = map[string]string{
	"films": ":9990",
}
