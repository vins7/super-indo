package server

type Server struct {
	Port   string `yaml:"port"`
	Secret string `yaml:"secret"`
}
