package models

type Dependency struct {
	Name  string
	Count int
}

type DependencyList []Dependency

func (d DependencyList) Len() int {
	return len(d)
}

func (d DependencyList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d DependencyList) Less(i, j int) bool {
	return d[i].Count < d[j].Count
}
