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
	getPrefixMap() map[string]string
	getPrefixWithVersion(sectionName string) string
	retrieveFieldsFromPackages() (map[string]string, error)
	unfoldFields(fields map[string]string) map[string]string
	appendData(source map[string]string, destination map[string]string) (map[string]string, error)
	getInstallExpressions() (map[string]string, error)
	getUninstallExpressions() (map[string]string, error)
	getExpressions() (map[string]string, error)

	getFields() (map[string]string, error)

	GetOutput(kind string) (map[string]string, error)
}

func (f *Framework) getPrefixMap() map[string]string {
	result := make(map[string]string)

	for _, v := range f.Prefixes {
		result[v.Section] = v.Prefix
	}

	return result
}

func (f *Framework) getPrefixWithVersion(sectionName string) string {
	return strings.Join([]string{f.getPrefixMap()[sectionName], f.Release.GetVersionAsString()}, "_")
}

func (f *Framework) appendData(source map[string]string, destination map[string]string) (map[string]string, error) {
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

func (f *Framework) retrieveFieldsFromPackages() (map[string]string, error) {
	output := make(map[string]string)
	var err error

	// Get all fields in all packages
	for _, p := range f.Packages {
		var fields map[string]string

		fields, err = p.GetFields()
		if err != nil {
			log.Fatal(err)
			break
		}

		output, err = f.appendData(fields, output)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
	return output, err
}

func (f *Framework) unfoldFields(fields map[string]string) map[string]string {
	re := regexp.MustCompile(`<<[a-zA-Z0-9_.]*/[a-zA-Z0-9_]*>>`)
	for key := range fields {
		loop := true
		for loop {
			foundKeys := re.FindAllString(fields[key], -1)
			for _, foundKey := range foundKeys {
				searchKey := strings.ReplaceAll(foundKey, "<<", "")
				searchKey = strings.ReplaceAll(searchKey, ">>", "")
				fields[key] = strings.ReplaceAll(fields[key], foundKey, fields[searchKey])
			}

			if !re.MatchString(fields[key]) {
				loop = false
			}
		}

		for k := range f.getPrefixMap() {
			fields[key] = strings.ReplaceAll(fields[key], "<<"+k+">>", f.getPrefixWithVersion(k))
		}
	}

	return fields
}

func (f *Framework) getFields() (map[string]string, error) {
	fields, err := f.retrieveFieldsFromPackages()
	if err != nil {
		log.Fatal(err)
	}

	return f.unfoldFields(fields), err
}

func (f *Framework) getSortedFieldKeys(fields map[string]string) []string {
	fieldKeys := make([]string, 0, len(fields))
	for f := range fields {
		fieldKeys = append(fieldKeys, f)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(fieldKeys)))

	return fieldKeys
}

func (f *Framework) getInstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, p := range f.Packages {
		expressions, err = p.GetInstallExpressions()
		if err != nil {
			log.Fatal(err)
		} else {
			output, err = f.appendData(expressions, output)
		}
	}

	return output, err
}

func (f *Framework) getUninstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, p := range f.Packages {
		expressions, err = p.GetUninstallExpressions()
		if err != nil {
			log.Fatal(err)
		} else {
			output, err = f.appendData(expressions, output)
		}
	}

	return output, err
}

func (f *Framework) getExpressions(kind string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	if kind == "install" {
		output, err = f.getInstallExpressions()
	} else if kind == "uninstall" {
		output, err = f.getUninstallExpressions()
	}

	return output, err
}

func (f *Framework) GetOutput(kind string) ([]string, error) {
	var output []string
	expressions, err := f.getExpressions(kind)
	if err != nil {
		log.Fatal(err)
		return output, err
	}

	fields, err := f.getFields()
	if err != nil {
		log.Fatal(err)
		return output, err
	}

	fieldKeys := f.getSortedFieldKeys(fields)

	for k := range expressions {
		if expressions[k] != "" {
			// Replace fields referenced in expressions
			for _, e := range fieldKeys {
				if strings.Contains(expressions[k], e) {
					expressions[k] = strings.ReplaceAll(expressions[k], "<<"+e+">>", fields[e])
				}
			}

			// Replace prefixes in expressions
			for p := range f.getPrefixMap() {
				expressions[k] = strings.ReplaceAll(expressions[k], "<<"+p+">>", f.getPrefixWithVersion(p))
			}
		}
	}

	uniqueElementNames := 0
	for u := range fieldKeys {
		if strings.Contains(fieldKeys[u], "/name") {
			uniqueElementNames++
		}
	}

	// uniqueElementNames := len(fieldKeys)

	counter := make(DependencyList, uniqueElementNames)
	i := 0
	for f := range fieldKeys {
		if strings.Contains(fieldKeys[f], "/name") {
			j := 0
			for e := range expressions {
				if e != strings.TrimSuffix(fieldKeys[f], "/name") {
					if strings.Contains(expressions[e], fields[fieldKeys[f]]) {
						// fmt.Println(e, fieldKeys[f])
						re := regexp.MustCompile(fields[fieldKeys[f]])
						count := len(re.FindAllString(expressions[e], -1))
						j = j + count

						// fmt.Println(fields[fieldKeys[f]], count, j, "\n", expressions[e])
					}
				}
			}
			counter[i] = Dependency{
				Name:  strings.TrimSuffix(fieldKeys[f], "/name"),
				Count: j,
			}
			i++
			fmt.Println()
		}
	}

	sort.Sort(sort.Reverse(counter))
	fmt.Println("----------------------- COUNTER -----------------------")
	for o := range counter {
		output = append(output, expressions[counter[o].Name])
		fmt.Println(counter[o].Name, counter[o].Count)
	}
	fmt.Println("----------------------- COUNTER -----------------------")

	return output, err
}

func (f *Framework) CountDependencies(search string, expressions map[string]string) int {
	j := 0
	for _, v := range expressions {
		if strings.Contains(v, search) {
			j++
		}
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
