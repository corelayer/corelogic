package models/framework

type Placeholder struct {
	Name string `yaml:name`
	Expression string `yaml:expression`
}

type Dependency struct {
	Name string `yaml:name`
	Reference string `yaml:reference`
}

type Expression struct {
	Install string `yaml:install`
	Uninstall string `yaml:uninstall`
}

type Element struct {
	Name string `yaml:name`
	Dependencies []Dependency `yaml:dependencies`
	Expressions Expression `yaml:expressions`
}

type Section struct {
	Name string `yaml:name`
	Elements []Element `yaml:elements`
}

type SubModule struct {
	Name string `yaml:name`
	Dependencies []string `yaml:dependencies`
	Placeholders []Placeholder `yaml:placeholders`
	Sections []Section `yaml:sections`
}