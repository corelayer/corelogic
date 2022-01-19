package models

type Prefix struct {
	Section         string `yaml:"section"`
	Prefix          string `yaml:"prefix"`
	ProcessingOrder int
}
