package service

type Flag string

const (
	COVID19            Flag = "COVID-19"
	Attractive         Flag = "Attractive"
	SleepApneaSyndrome Flag = "SleepApneaSyndrome"
)

type Config map[Flag]bool

func (c *Config) IsEnabled(flag Flag) bool {
	return (*c)[flag]
}

func (c *Config) ToggleOn(flag Flag) {
	(*c)[flag] = true
}

func (c *Config) ToggleOff(flag Flag) {
	(*c)[flag] = false
}
