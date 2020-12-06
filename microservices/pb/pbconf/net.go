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
	return s.Name "_svc"
}

var (
Films    = &service{Name: "films"}
Sync       = &service{Name: "sync"}
)

var devPorts = map[string]string{
	"films":    ":9900",
	"sync":     ":9991",
}
