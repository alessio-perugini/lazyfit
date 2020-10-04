package lazyfit

type API struct {
	Base        string `yaml:"base"`
	BaseBooking string `yaml:"baseBooking"`
	BaseAccount string `yaml:"baseAccount"`
	Login       string `yaml:"login"`
	NewBook     string `yaml:"newBook"`
	Book        string `yaml:"book"`
	Delete      string `yaml:"delete"`
	TimeTable   string `yaml:"timeTable"`
	Courses     string `yaml:"courses"`
	Schedules   string `yaml:"schedules"`
}

func NewApi() *API {
	return &Conf.API
}
