package models

type Field struct {
	Id     string `yaml:"id"`
	Data   string `yaml:"data"`
	Prefix bool   `yaml:"prefix"`
}
