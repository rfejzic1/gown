package component

type Project struct {
	Path        string
	Config      Config
	Application *Application
	Setup       *Setup
	Web         *Web
	Commands    []Command
}

func NewProject(path string, config Config) *Project {
	return &Project{
		Path:        path,
		Config:      config,
		Application: NewApplication(),
		Setup:       &Setup{},
		Web:         &Web{},
		Commands:    []Command{},
	}
}

func (p *Project) AddCommand(cmd Command) {
	p.Commands = append(p.Commands, cmd)
}

type Application struct {
	Modules []Module
	Sources []SourceFile
}

func NewApplication() *Application {
	return &Application{
		Modules: []Module{},
		Sources: []SourceFile{},
	}
}

func (a *Application) AddModule(module Module) {
	a.Modules = append(a.Modules, module)
}

type Module struct {
	Name    string
	Sources []SourceFile
}

type Setup struct {
	Sources []SourceFile
}

type Web struct {
	Sources []SourceFile
}

type Command struct {
	Sources []SourceFile
}
