package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/corelayer/corelogic/general"

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
	defer general.FinishTimer(general.StartTimer("Loading framework " + version))
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

	for _, d := range subDirs {
		if d.IsDir() {
			var p models.Package
			p, err = c.GetPackagesFromDirectory(rootDir, d.Name())
			if err != nil {
				return err
			}
			c.Framework.Packages = append(c.Framework.Packages, p)
		}
	}

	return err
}

func (c *FrameworkController) GetPackagesFromDirectory(rootDir string, directoryName string) (models.Package, error) {
	// defer general.FinishTimer(general.StartTimer("GetPackagesFromDirectory " + rootDir + "/packages/" + directoryName))

	myPackage := models.Package{
		Name:    directoryName,
		Modules: []models.Module{},
	}

	files, err := ioutil.ReadDir(rootDir + "/packages/" + myPackage.Name)
	if err != nil {
		log.Fatal(err)
		return myPackage, err
	}

	for _, f := range files {
		if !f.IsDir() {
			if filepath.Ext(f.Name()) == ".yaml" {
				// log.Println(f.Name())
				var module models.Module
				module, err = c.GetModuleFromFile(rootDir + "/packages/" + myPackage.Name + "/" + f.Name())
				if err != nil {
					return myPackage, err
				}
				myPackage.Modules = append(myPackage.Modules, module)
			}
		} else {
			var modules []models.Module
			modules, err = c.GetModulesFromDirectory(rootDir + "/packages/" + myPackage.Name + "/" + f.Name())
			if err != nil {
				return myPackage, err
			}
			myPackage.Modules = append(myPackage.Modules, modules...)
		}
	}
	return myPackage, err
}

func (c *FrameworkController) GetModuleFromFile(filePath string) (models.Module, error) {
	// defer general.FinishTimer(general.StartTimer("GetModuleFromFile " + filePath))

	module := models.Module{}

	moduleSource, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return module, err
	}

	err = yaml.Unmarshal(moduleSource, &module)
	if err != nil {
		log.Fatal(err)
	}

	return module, err
}

func (c *FrameworkController) GetModulesFromDirectory(filePath string) ([]models.Module, error) {
	// defer general.FinishTimer(general.StartTimer("GetModulesFromDirectory " + filePath))

	var modules []models.Module

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
		return modules, err
	}

	for _, f := range files {
		if !f.IsDir() {
			if filepath.Ext(f.Name()) == ".yaml" {
				// log.Println(f.Name())
				module, err := c.GetModuleFromFile(filePath + "/" + f.Name())
				if err != nil {
					log.Fatal(err)
					return modules, err
				}
				modules = append(modules, module)
			}
		}
	}

	return modules, err
}

func (c *FrameworkController) GetFrameworkAsJson(indent string) string {
	defer general.FinishTimer(general.StartTimer("GetFrameworkAsJson"))

	output, err := json.MarshalIndent(c.Framework, "", indent)
	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}
