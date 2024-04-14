package component

type Project struct {
	Config Config
	Path   string
	App    *Application
}

type Application struct {
	Modules []Module
	Files   []SourceFile
}

func NewApplication(modules []Module, files []SourceFile) (*Application, error) {
	return &Application{
		Modules: modules,
		Files:   files,
	}, nil
}

type Module struct {
	Name string
	File []SourceFile
}

func NewModule(name string, files []SourceFile) Module {
	return Module{
		Name: name,
		File: files,
	}
}
