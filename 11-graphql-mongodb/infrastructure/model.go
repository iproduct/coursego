package infrastructure

type app struct {
	Appname     string `yaml:"name"`
	Debug       bool   `yaml:"debug"`
	Port        string `yaml:"port"`
	Service     string `yaml:"service"`
	Certificate string `yaml:"certificate"`
	Pem_key     string `yaml:"pem_key"`
	Host        string `yaml:"host"`
	Stage       string `yaml:"stage"`
}

type database struct {
	Name       string `yaml:"name"`
	Connection string `yaml:"connection"`
}

type Environment struct {
	App       app                 `yaml:"app"`
	Databases map[string]database `yaml:"databases"`
	path      string
}
