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
	"github.com/corelayer/corelogic/controllers"
	"github.com/corelayer/corelogic/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sort"
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
				if m.IsDir() == false {
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
	fmt.Printf(string(output))


	//fmt.Println("================================================")
	//
	//fields, err8 := controller.Framework.GetFields()
	//if err8 != nil {
	//	fmt.Println(err8)
	//}
	//for k,v := range fields {
	//	fmt.Println(k, v)
	//}


	fmt.Println("================================================")

	install, err9 := controller.Framework.GetInstallExpressions()
	if err9 != nil {
		fmt.Println(err9)
	}
	for k,v := range install {
		fmt.Println(k, "\t", v)
	}


	fmt.Println("================================================")
	fmt.Println("================================================")

	dependencies := controller.Framework.GetDependencyList(install)
	for _ ,v := range dependencies {
		fmt.Println(v.Name, v.Count)
	}
	fmt.Println("================================================")
	sort.Sort(sort.Reverse(dependencies))
	for _ ,v := range dependencies {
		//fmt.Println(v.Name, v.Count)
		fmt.Println(install[v.Name])
	}

}
