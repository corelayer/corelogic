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
	"fmt"
	"log"
	"os"

	"github.com/corelayer/corelogic/general"

	"github.com/corelayer/corelogic/controllers"
)

func main() {
	defer general.FinishTimer(general.StartTimer("Program execution"))
	controller := controllers.FrameworkController{}
	err := controller.Load("11.0")

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(controller.GetFrameworkAsJson("\t"))
	tagFilter := []string{"fake"}
	var install []string
	install, err = controller.Framework.GetOutput("install", tagFilter)
	if err != nil {
		log.Fatal(err)
	}

	printLines("config.conf", install)
}

func printLines(filePath string, expressions []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range expressions {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}
