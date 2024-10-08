package gameoflife

type Configuration struct {
	M            int     `yaml:"m"`
	N            int     `yaml:"n"`
	Steps        int     `yaml:"steps"`
	Display      string  `yaml:"display"`
	Distribution float32 `yaml:"distribution"`
	SleepTime    float32 `yaml:"sleep_time"`
}
