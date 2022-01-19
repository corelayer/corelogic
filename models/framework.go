package models

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/corelayer/corelogic/general"
)

type DataMapWriter interface {
	AppendData(source map[string]string, destination map[string]string) (map[string]string, error)
}

type Framework struct {
	Release  Release   `yaml:"release"`
	Prefixes []Prefix  `yaml:"prefixes"`
	Packages []Package `yaml:"packages"`

	Expressions     map[string]string
	Fields          map[string]string
	SortedFieldKeys []string

	SectionData map[string][]string
}

type FrameworkReader interface {
	getPrefixMap() map[string]string
	getPrefixWithVersion(sectionName string) string
	appendData(source map[string]string, destination map[string]string) (map[string]string, error)
	getFieldsFromPackages() (map[string]string, error)
	unfoldFields(fields map[string]string) map[string]string
	getFields() (map[string]string, error)
	setSortedFieldKeys(fields map[string]string)
	getInstallExpressionsFromPackages(tagFilter []string) (map[string]string, error)
	getUninstallExpressionsFromPackages(tagFilter []string) (map[string]string, error)
	getExpressions(kind string, tagFilter []string) (map[string]string, error)
	replacePrefixesInExpression(expression string) string
	replaceFieldsInExpression(expression string) string
	replaceDataInExpression(expression string) string
	unfoldExpression(elementName string, ch chan<- UnfoldedExpressionData, wg *sync.WaitGroup)
	collectExpressionsForSection(sectionName string, ch chan<- string, wg *sync.WaitGroup)
	unfoldedExpressionCollector(count int, ch <-chan UnfoldedExpressionData, wg *sync.WaitGroup)
	sectionExpressionCollector(sectionName string, ch <-chan string, wg *sync.WaitGroup)
	unfoldExpressions()
	collectExpressionsPerSection()
	sortPrefixes(prefixes []Prefix)
	GetOutput(kind string, tagFilter []string) ([]string, error)
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

func (f *Framework) getFieldsFromPackages() (map[string]string, error) {
	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " get fields from packages"))

	output := make(map[string]string)
	var err error

	// Get all fields in all packages
	for _, p := range f.Packages {
		var fields map[string]string

		fields, err = p.GetFields()
		if err != nil {
			log.Fatal(err)
			//break
		}

		output, err = f.appendData(fields, output)
		if err != nil {
			log.Fatal(err)
			//break
		}
	}
	return output, err
}

func (f *Framework) unfoldFields(fields map[string]string) map[string]string {
	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " unfold fields"))

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
	fields, err := f.getFieldsFromPackages()
	if err != nil {
		log.Fatal(err)
	}

	return f.unfoldFields(fields), err
}

func (f *Framework) setSortedFieldKeys(fields map[string]string) {
	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " set sorted field keys"))

	fieldKeys := make([]string, 0, len(fields))
	for f := range fields {
		fieldKeys = append(fieldKeys, f)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(fieldKeys)))

	f.SortedFieldKeys = fieldKeys
}

func (f *Framework) getInstallExpressionsFromPackages(tagFilter []string) (map[string]string, error) {
	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " get install expressions from packages"))

	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, p := range f.Packages {
		expressions, err = p.GetInstallExpressions(tagFilter)
		if err != nil {
			log.Fatal(err)
		} else {
			output, err = f.appendData(expressions, output)
		}
	}

	return output, err
}

func (f *Framework) getUninstallExpressionsFromPackages(tagFilter []string) (map[string]string, error) {
	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " get uninstall expressions from packages"))

	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, p := range f.Packages {
		expressions, err = p.GetUninstallExpressions(tagFilter)
		if err != nil {
			log.Fatal(err)
		} else {
			output, err = f.appendData(expressions, output)
		}
	}

	return output, err
}

func (f *Framework) getExpressions(kind string, tagFilter []string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	if kind == "install" {
		output, err = f.getInstallExpressionsFromPackages(tagFilter)
	} else if kind == "uninstall" {
		output, err = f.getUninstallExpressionsFromPackages(tagFilter)
	}

	return output, err
}

func (f *Framework) replacePrefixesInExpression(expression string) string {
	// Replace prefixes in expressions
	for p := range f.getPrefixMap() {
		expression = strings.ReplaceAll(expression, "<<"+p+">>", f.getPrefixWithVersion(p))
	}

	return expression
}

func (f *Framework) replaceFieldsInExpression(expression string) string {
	for _, e := range f.SortedFieldKeys {
		expression = strings.ReplaceAll(expression, "<<"+e+">>", f.Fields[e])
	}

	return expression
}

func (f *Framework) replaceDataInExpression(expression string) string {
	if expression != "" {
		expression = f.replaceFieldsInExpression(expression)
		expression = f.replacePrefixesInExpression(expression)
		expression = strings.TrimSuffix(expression, "\n")
	}

	return expression
}

func (f *Framework) unfoldExpression(elementName string, ch chan<- UnfoldedExpressionData, wg *sync.WaitGroup) {
	output := UnfoldedExpressionData{
		key:   elementName,
		value: f.Expressions[elementName],
	}

	output.value = f.replaceDataInExpression(output.value)
	ch <- output
	wg.Done()
}

func (f *Framework) collectExpressionsForSection(sectionName string, ch chan<- string, wg *sync.WaitGroup) {
	for k := range f.Expressions {
		if strings.Contains(k, sectionName) {
			if f.Expressions[k] != "" {
				ch <- f.Expressions[k]
			}
		}
	}
	close(ch)
	wg.Done()
}

func (f *Framework) unfoldedExpressionCollector(count int, ch <-chan UnfoldedExpressionData, wg *sync.WaitGroup) {
	completed := false

	var expressions = make(map[string]string)
	for !completed {
		select {
		case data, ok := <-ch:
			if !ok {
				completed = true
			}
			expressions[data.key] = data.value
			count--
		default:
			if count == 0 {
				completed = true
			}
		}
	}
	f.Expressions = expressions
	wg.Done()
}

func (f *Framework) sectionExpressionCollector(sectionName string, ch <-chan string, wg *sync.WaitGroup) {
	completed := false

	expressions := []string{}
	for !completed {
		select {
		case data, ok := <-ch:
			if !ok {
				completed = true
			} else {
				expressions = append(expressions, data)
			}
		}
	}
	f.SectionData[sectionName] = expressions
	wg.Done()
}

func (f *Framework) unfoldExpressions() {
	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " unfold expressions"))

	wg := &sync.WaitGroup{}
	ch := make(chan UnfoldedExpressionData)

	count := 0
	for k := range f.Expressions {
		wg.Add(1)
		count++
		go f.unfoldExpression(k, ch, wg)
	}
	wg.Add(1)
	go f.unfoldedExpressionCollector(count, ch, wg)

	wg.Wait()
	close(ch)
}

func (f *Framework) collectExpressionsPerSection() {
	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " collect expressions per section"))

	wg := &sync.WaitGroup{}

	// count := 0
	for _, p := range f.Prefixes {
		ch := make(chan string)

		wg.Add(1)
		go f.collectExpressionsForSection(p.Section, ch, wg)

		wg.Add(1)
		go f.sectionExpressionCollector(p.Section, ch, wg)
	}

	wg.Wait()
}

// func (f *Framework) countUniqueFields() int {
// 	counter := 0

// 	for u := range f.SortedFieldKeys {
// 		if strings.Contains(f.SortedFieldKeys[u], "/name") {
// 			counter++
// 		}
// 	}

// 	return counter
// }

func (f *Framework) sortPrefixes(prefixes []Prefix) {
	sort.Slice(prefixes, func(i, j int) bool {
		return prefixes[i].ProcessingOrder < prefixes[j].ProcessingOrder
	})
}

func (f *Framework) GetOutput(kind string, tagFilter []string) ([]string, error) {
	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " get " + kind + " output"))

	var output []string
	//f.Expressions = make(map[string]string)
	var err error
	f.Expressions, err = f.getExpressions(kind, tagFilter)
	if err != nil {
		log.Fatal(err)
		return output, err
	}

	f.Fields, err = f.getFields()
	if err != nil {
		log.Fatal(err)
		return output, err
	}

	f.setSortedFieldKeys(f.Fields)
	f.unfoldExpressions()
	f.sortPrefixes(f.Prefixes)
	f.SectionData = make(map[string][]string)
	f.collectExpressionsPerSection()

	for _, p := range f.Prefixes {
		output = append(output, "### "+p.Section)
		output = append(output, f.SectionData[p.Section]...)
		output = append(output, "##########################")
	}
	//uniqueElementNames := f.countUniqueFields()

	// dependencyList := make(DependencyList, f.countUniqueFields())
	// i := 0
	// for fieldKey := range f.SortedFieldKeys {
	// 	if strings.Contains(f.SortedFieldKeys[fieldKey], "/name") {
	// 		j := 0
	// 		for expression := range f.Expressions {
	// 			if expression != strings.TrimSuffix(f.SortedFieldKeys[fieldKey], "/name") {
	// 				if strings.Contains(f.Expressions[expression], f.Fields[f.SortedFieldKeys[fieldKey]]) {
	// 					// fmt.Println(expression, sortedFieldKeys[f])
	// 					re := regexp.MustCompile(f.Fields[f.SortedFieldKeys[fieldKey]])
	// 					count := len(re.FindAllString(f.Expressions[expression], -1))
	// 					j = j + count
	// 					// fmt.Println(fields[sortedFieldKeys[f]], count, j, "\n", expressions[expression])
	// 				}
	// 			}
	// 		}
	// 		dependencyList[i] = Dependency{
	// 			Name:  strings.TrimSuffix(f.SortedFieldKeys[fieldKey], "/name"),
	// 			Count: j,
	// 		}
	// 		i++
	// 	}
	// }

	// sort.Sort(sort.Reverse(dependencyList))
	// // fmt.Println("----------------------- COUNTER -----------------------")
	// for dependency := range dependencyList {
	// 	if f.Expressions[dependencyList[dependency].Name] != "" {
	// 		output = append(output, f.Expressions[dependencyList[dependency].Name])
	// 		// fmt.Println(dependencyList[dependency].Name, dependencyList[dependency].Count)
	// 	}
	// }
	// // fmt.Println("----------------------- COUNTER -----------------------")

	return output, err
}

// func (f *Framework) CountDependencies(search string, expressions map[string]string) int {
// 	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " count dependencies for " + search))

// 	j := 0
// 	for _, v := range expressions {
// 		if strings.Contains(v, search) {
// 			j++
// 		}
// 	}
// 	return j
// }

// func (f *Framework) GetDependencyList(expressions map[string]string) DependencyList {
// 	defer general.FinishTimer(general.StartTimer("Framework " + f.Release.GetVersionAsString() + " get dependency list"))

// 	output := make(DependencyList, len(expressions))

// 	i := 0
// 	for k := range expressions {

// 		output[i] = Dependency{
// 			Name:  k,
// 			Count: f.CountDependencies(k, expressions),
// 		}
// 		i++
// 	}

// 	return f.SortDependencyList(output)
// }

// func (f *Framework) SortDependencyList(input DependencyList) DependencyList {
// 	// output := make(DependencyList, len(input))
// 	sort.Sort(sort.Reverse(input))

// 	// maxCount := input[0].Count

// 	return input
// }
