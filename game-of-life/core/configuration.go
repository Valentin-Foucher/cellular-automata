package gameoflife

type Configuration struct {
	M     int `yaml:"m"`
	N     int `yaml:"n"`
	Steps int `yaml:"steps"`
}
