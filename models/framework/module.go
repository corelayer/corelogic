package models/framework

type Module struct {
	Name string `yaml:string`
	Dependencies []string `yaml:dependencies`
}