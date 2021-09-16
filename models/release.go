package models

import "fmt"

type Release struct {
	Major int `yaml:major`
	Minor int `yaml:minor`
}

type ReleaseReader interface {
	GetVersionAsString() string
}

func (r *Release) GetVersionAsString() string {
	return fmt.Sprintf("CL%02d%02d", r.Major, r.Minor)
}
