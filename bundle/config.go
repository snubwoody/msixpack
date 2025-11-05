package bundle

type Config struct {
	Package     ConfigPackage
	Application ConfigApplication
}

type ConfigPackage struct {
	Version       string
	Name          string
	DisplayName   string
	Publisher     string
	PublisherName string
	Logo          string
	Resources     string
}

type ConfigApplication struct {
	Id          string
	Executable  string
	Name        string
	Description string
}
