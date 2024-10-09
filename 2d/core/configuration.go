package core2d

type Configuration struct {
	M            int     `yaml:"m"`
	N            int     `yaml:"n"`
	Steps        int     `yaml:"steps"`
	Display      string  `yaml:"display"`
	Automaton    string  `yaml:"automaton"`
	Distribution float32 `yaml:"distribution"`
	SleepTime    float32 `yaml:"sleep_time"`
}
