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
				} else {
					subfiles, err := ioutil.ReadDir(rootDir + "/packages/" + v.Name() + "/" + m.Name())
					if err != nil {
						log.Fatal(err)
						return err
					}

					for _, sm := range subfiles {
						if !sm.IsDir() {
							// fmt.Println(m.Name())
							moduleSource, err := ioutil.ReadFile(rootDir + "/packages/" + v.Name() + "/" + m.Name() + "/" + sm.Name())
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
