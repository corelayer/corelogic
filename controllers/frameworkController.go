package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/corelayer/corelogic/models"
	"gopkg.in/yaml.v2"
)

type FrameworkController struct {
	Framework models.Framework
}

type FrameworkControllerReader interface {
	Load(version string) error
	GetFrameworkAsJson(indent string) string
}

func (c *FrameworkController) Load(version string) error {
	rootDir := "assets/framework/" + version
	var source []byte
	var err error

	source, err = ioutil.ReadFile(rootDir + "/framework.yaml")
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = yaml.Unmarshal(source, &c.Framework)
	if err != nil {
		log.Fatal(err)
		return err
	}

	c.Framework.Packages = []models.Package{}

	subDirs, err := ioutil.ReadDir(rootDir + "/packages")
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, v := range subDirs {
		if v.IsDir() {
			myPackage := models.Package{
				Name:    v.Name(),
				Modules: []models.Module{},
			}

			files, err := ioutil.ReadDir(rootDir + "/packages/" + v.Name())
			if err != nil {
				log.Fatal(err)
				return err
			}

			for _, m := range files {
				if !m.IsDir() {
					// fmt.Println(m.Name())
					moduleSource, err := ioutil.ReadFile(rootDir + "/packages/" + v.Name() + "/" + m.Name())
					if err != nil {
						log.Fatal(err)
						return err
					}

					module := &models.Module{}
					err = yaml.Unmarshal(moduleSource, module)
					if err != nil {
						log.Fatal(err)
						return err
					}
					myPackage.Modules = append(myPackage.Modules, *module)
				}
			}
			c.Framework.Packages = append(c.Framework.Packages, myPackage)
		}
	}

	return err
}

func (c *FrameworkController) GetFrameworkAsJson(indent string) string {
	output, err := json.MarshalIndent(c.Framework, "", indent)
	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}

//
//func (f *FrameworkController) GetVersion() string {
//	return f.Release.GetVersion()
//}
//
//
//type ModuleDependency interface {
//	GetSortedModules() []models.Module
//	getModuleDependencyCounter() []moduleDependency
//	countModuleDependencies(moduleDependencyCounter []moduleDependency) []moduleDependency
//	sortModuleDependencyCounter(moduleDependencyCounter []moduleDependency) []moduleDependency
//}
//
//func (f *FrameworkController) GetSortedModules() []Module {
//	var result = []Module{}
//
//	dependencyCounter := f.getModuleDependencyCounter()
//	dependencyCounter = f.countModuleDependencies(dependencyCounter)
//	dependencyCounter = f.sortModuleDependencyCounter(dependencyCounter)
//
//	for _, v := range dependencyCounter {
//		append(result, f.Mo)
//	}
//	return result
//}
//
//type moduleDependency struct {
//	Name string
//	Counter int
//}
//
//func (f *FrameworkController) countModuleDependencies(moduleDependencyCounter []moduleDependency) []moduleDependency {
//	for _, v := range moduleDependencyCounter {
//		for _, m := range f.Modules {
//			for _, d := range m.Dependencies {
//				if v.Name == d {
//					v.Counter++
//				}
//			}
//		}
//	}
//	return moduleDependencyCounter
//}
//
//func (f *FrameworkController) getModuleDependencyCounter() []moduleDependency {
//	var moduleDependencyCounter []moduleDependency
//	for _, v := range f.Modules {
//		moduleDependencyCounter = append(
//			moduleDependencyCounter,
//			moduleDependency{v.Name, 0})
//	}
//	return moduleDependencyCounter
//}
//
//func (f *FrameworkController) sortModuleDependencyCounter(moduleDependencyCounter []moduleDependency) []moduleDependency {
//	sort.SliceStable(
//		moduleDependencyCounter,
//		func(i, j int) bool {
//			return moduleDependencyCounter[i].Counter < moduleDependencyCounter[j].Counter
//		})
//
//	return moduleDependencyCounter
//}
