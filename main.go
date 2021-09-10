/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/corelayer/corelogic/controllers"
	"github.com/corelayer/corelogic/models"
	"gopkg.in/yaml.v2"
)

func main() {
	rootDir := "assets/framework/11.0"
	source, err := ioutil.ReadFile(rootDir + "/framework.yaml")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(source))

	controller := controllers.FrameworkController{}
	framework := &models.Framework{}
	err2 := yaml.Unmarshal(source, framework)
	if err2 != nil {
		log.Fatal(err2)
	}
	controller.Framework = *framework
	controller.Framework.Packages = []models.Package{}

	subDirs, err3 := ioutil.ReadDir(rootDir + "/packages")
	if err3 != nil {
		log.Fatal(err3)
	}
	for _, v := range subDirs {
		if v.IsDir() {
			myPackage := models.Package{
				Name:    v.Name(),
				Modules: []models.Module{},
			}

			files, err4 := ioutil.ReadDir(rootDir + "/packages/" + v.Name())
			if err4 != nil {
				log.Fatal(err4)
			}
			for _, m := range files {
				if !m.IsDir() {
					fmt.Println(m.Name())
					moduleSource, err5 := ioutil.ReadFile(rootDir + "/packages/" + v.Name() + "/" + m.Name())
					if err5 != nil {
						log.Fatal(err5)
					}

					module := &models.Module{}
					err6 := yaml.Unmarshal(moduleSource, module)
					if err6 != nil {
						log.Fatal(err6)
					}
					myPackage.Modules = append(myPackage.Modules, *module)
				}
			}
			controller.Framework.Packages = append(controller.Framework.Packages, myPackage)
		}
	}

	output, err7 := json.MarshalIndent(controller.Framework, "", "\t")
	if err7 != nil {
		log.Fatal(err7)
	}
	fmt.Print(string(output))
	fmt.Println("")
	fmt.Println("===================== END UNMARSHAL ===========================")

	fmt.Println("")
	fmt.Println("===================== BEGIN FIELDS ===========================")
	fields, err8 := controller.Framework.GetFields()
	if err8 != nil {
		log.Fatal(err8)
	}
	for k, v := range fields {
		fmt.Println(k, v)
	}
	fmt.Println("")
	fmt.Println("===================== END FIELDS ===========================")

	fmt.Println("")
	fmt.Println("===================== BEGIN INSTALL EXPRESSIONS ===========================")
	install, err9 := controller.Framework.GetInstallExpressions()
	if err9 != nil {
		log.Fatal(err9)
	}
	for k, v := range install {
		fmt.Println(k, "\t", v)
	}
	fmt.Println("")
	fmt.Println("===================== END INSTALL EXPRESSIONS ===========================")

	fmt.Println("")
	fmt.Println("===================== BEGIN DEPENDENCY COUNT ===========================")
	sortedExpressions := controller.Framework.GetDependencyList(install)
	// sort.Sort(sort.Reverse(sortedExpressions))

	for _, v := range sortedExpressions {
		fmt.Println(v.Name, v.Count)
	}
	fmt.Println("")
	fmt.Println("===================== END DEPENDENCY COUNT ===========================")

	fmt.Println("")
	fmt.Println("===================== BEGIN INSTALL EXPRESSIONS ===========================")
	// sort.Sort(sort.Reverse(sortedExpressions))
	for _, v := range sortedExpressions {
		sortedFieldKeys := make([]string, 0, len(fields))
		for f := range fields {
			sortedFieldKeys = append(sortedFieldKeys, f)
		}
		sort.Sort(sort.Reverse(sort.StringSlice(sortedFieldKeys)))

		// fmt.Println(v.Name, v.Count)
		//count := strings.Count(install[v.Name], "\n")

		if install[v.Name] != "" {
			// Replace fields referenced in expressions
			for _, e := range sortedFieldKeys {

				install[v.Name] = strings.ReplaceAll(install[v.Name], "<<"+e+">>", fields[e])
			}

			// Replace fields referenced in fields
			for _, e := range sortedFieldKeys {

				install[v.Name] = strings.ReplaceAll(install[v.Name], "<<"+e+">>", fields[e])
			}

			for k := range framework.GetPrefixMap() {
				//fmt.Println("1", k, "2", p, "3")
				install[v.Name] = strings.ReplaceAll(install[v.Name], "<<"+k+">>", framework.GetPrefixWithVersion(k))
			}
			// fmt.Println(v.Name)
			fmt.Println(strings.TrimSuffix(install[v.Name], "\n"))
		}
	}
	fmt.Println("")
	fmt.Println("===================== END INSTALL EXPRESSIONS ===========================")
}
