package config_router

type RouterPath int

const (
	Main     = "Main"
	Settings = "Settings"
)

func (p RouterPath) String() string {
	return [...]string{"Main", "Settings"}[p]
}
