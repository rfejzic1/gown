package component

type Config struct {
	Project ConfigProject `toml:"project"`
}

type ConfigProject struct {
	Name string `toml:"name"`
}
