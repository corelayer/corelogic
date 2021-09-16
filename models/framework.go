package models

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

type DataMapWriter interface {
	AppendData(source map[string]string, destination map[string]string) (map[string]string, error)
}

type Framework struct {
	Release  Release   `yaml:release`
	Prefixes []Prefix  `yaml:prefixes`
	Packages []Package `yaml:packages`
}

type FrameworkReader interface {
	GetPrefixMap() map[string]string
	GetPrefixWithVersion(sectionName string) string
	GetInstallExpressions() (map[string]string, error)
	GetUninstallExpressions() (map[string]string, error)
}

func (f *Framework) GetPrefixMap() map[string]string {
	result := make(map[string]string)

	for _, v := range f.Prefixes {
		result[v.Section] = v.Prefix
	}

	return result
}

func (f *Framework) GetPrefixWithVersion(sectionName string) string {
	return strings.Join([]string{f.GetPrefixMap()[sectionName], f.Release.GetVersionAsString()}, "_")
}

func (f *Framework) GetFields() (map[string]string, error) {
	output := make(map[string]string)
	var err error

	// var fields map[string]string
	for _, p := range f.Packages {
		fields, err := p.GetFields()
		if err != nil {
			log.Fatal(err)
		} else {
			output, err = f.AppendData(fields, output)
		}
	}

	re := regexp.MustCompile(`<<[a-zA-Z0-9_.]*/[a-zA-Z0-9_]*>>`)
	//var foundKeys []string
	for key, _ := range output {
		// fmt.Println("Current key", key)
		// fmt.Println("Current value", value)
		// fmt.Println("=========")
		loop := true
		for loop {
			foundKeys := re.FindAllString(output[key], -1)
			// fmt.Printf("%d\n", len(foundKeys))
			for _, foundKey := range foundKeys {
				// fmt.Println("Found key", foundKey)
				searchKey := strings.ReplaceAll(foundKey, "<<", "")
				searchKey = strings.ReplaceAll(searchKey, ">>", "")
				// fmt.Println("Search key", searchKey)
				// fmt.Println("Search value", output[searchKey])
				output[key] = strings.ReplaceAll(output[key], foundKey, output[searchKey])
				// fmt.Println("Temp output", output[key])
				// fmt.Println("----")

			}
			// fmt.Println("====")
			// fmt.Println("New output K", key)
			// fmt.Println("New output V", output[key])
			// fmt.Println("")
			// fmt.Println("")

			if re.MatchString(output[key]) {
				loop = true
			} else {
				loop = false
			}
		}
	}
	return output, err
}

func (f *Framework) GetInstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, p := range f.Packages {
		expressions, err = p.GetInstallExpressions()
		if err != nil {
			log.Fatal(err)
		} else {
			output, err = f.AppendData(expressions, output)
		}
	}

	return output, err
}

func (f *Framework) GetUninstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, p := range f.Packages {
		expressions, err = p.GetUninstallExpressions()
		if err != nil {
			log.Fatal(err)
		} else {
			output, err = f.AppendData(expressions, output)
		}
	}

	return output, err
}

func (f *Framework) AppendData(source map[string]string, destination map[string]string) (map[string]string, error) {
	var err error

	for k, v := range source {
		if _, isMapContainsKey := destination[k]; isMapContainsKey {
			err = fmt.Errorf("duplicate key %q found in framework", k)
			log.Fatal(err)
		} else {
			destination[k] = v
		}
	}

	return destination, err
}

func (f *Framework) CountDependencies(search string, expressions map[string]string) int {
	j := 0
	for _, v := range expressions {
		if strings.Contains(v, search) {
			j++
		}

		// if !strings.Contains(v, search) && strings.Contains(v, "ENDPOINT") {
		// 	fmt.Println(search, v)
		// }
	}
	return j
}

func (f *Framework) GetDependencyList(expressions map[string]string) DependencyList {
	output := make(DependencyList, len(expressions))

	i := 0
	for k := range expressions {

		output[i] = Dependency{
			Name:  k,
			Count: f.CountDependencies(k, expressions),
		}
		i++
	}

	return f.SortDependencyList(output)
}

func (f *Framework) SortDependencyList(input DependencyList) DependencyList {
	// output := make(DependencyList, len(input))
	sort.Sort(sort.Reverse(input))

	// maxCount := input[0].Count

	return input
}
