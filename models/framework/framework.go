package models/framework

type Release struct {
	Major int `yaml:major`
	Minor int `yaml:minor`
}

type Module struct {
	Name string `yaml:string`
	Dependencies []string `yaml:dependencies`
}

type Prefix struct {
	Section string `yaml:string`
	Prefix string `yaml:prefix`
}

type Framework struct {
	Name string `yaml:name`
	Release Release `yaml:release`
	Modules []Module `yaml:modules`
	Prefixes []Prefix `yaml:prefixes`
}