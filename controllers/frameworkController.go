package controllers

import "github.com/corelayer/corelogic/models"

type FrameworkController struct {
	Framework models.Framework
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