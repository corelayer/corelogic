package models

// import (
// 	"testing"
// )

// func TestRelease_GetVersionAsString(t *testing.T) {
// 	r := Release{
// 		Major: 10,
// 		Minor: 1,
// 	}

// 	output := r.GetVersionAsString()
// 	expectedOutput := "CL1001"
// 	if output != expectedOutput {
// 		t.Errorf("Output string is incorrect, got: %s, want: %s.", output, expectedOutput)
// 	}
// }

// func TestPackage_AppendData(t *testing.T) {
// 	p := Package{
// 		Name: "core",
// 		Modules: []Module{
// 			{
// 				Name: "cs",
// 				Sections: []Section{
// 					{
// 						Name: "trafficmanagement.contentswitching.policies",
// 						Elements: []Element{
// 							{
// 								Name: "trusted_full",
// 								Fields: []Field{
// 									{Id: "name", Data: "{{prefix}}trusted_full"},
// 									{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 									{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 								},
// 								Expressions: Expression{
// 									Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 									Uninstall: "rm cs policy {{name}}",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Name: "sm",
// 				Sections: []Section{{
// 					Name: "appexpert.stringmaps",
// 					Elements: []Element{
// 						{
// 							Name: "csv_control",
// 							Fields: []Field{
// 								{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								},
// 							},
// 							Expressions: Expression{
// 								Install:   "add policy stringmap {{name}}",
// 								Uninstall: "rm policy stringmap {{name}}",
// 							},
// 						},
// 					},
// 				},
// 				},
// 			},
// 		},
// 	}

// 	input1, _ := p.GetInstallExpressions()
// 	input2 := input1

// 	_, err := p.AppendData(input2, input1)
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in package")
// 	}

// }

// func TestPackage_GetFields(t *testing.T) {
// 	p := Package{
// 		Name: "core",
// 		Modules: []Module{
// 			{
// 				Name: "cs",
// 				Sections: []Section{
// 					{
// 						Name: "trafficmanagement.contentswitching.policies",
// 						Elements: []Element{
// 							{
// 								Name: "trusted_full",
// 								Fields: []Field{
// 									{Id: "name", Data: "{{prefix}}trusted_full"},
// 									{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 									{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 								},
// 								Expressions: Expression{
// 									Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 									Uninstall: "rm cs policy {{name}}",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Name: "sm",
// 				Sections: []Section{{
// 					Name: "appexpert.stringmaps",
// 					Elements: []Element{
// 						{
// 							Name: "csv_control",
// 							Fields: []Field{
// 								{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								},
// 							},
// 							Expressions: Expression{
// 								Install:   "add policy stringmap {{name}}",
// 								Uninstall: "rm policy stringmap {{name}}",
// 							},
// 						},
// 					},
// 				},
// 				},
// 			},
// 		},
// 	}

// 	output, err := p.GetFields()

// 	if err != nil {
// 		t.Errorf("%s", err)

// 	}

// 	expectedOutputKey := "core.cs.trafficmanagement.contentswitching.policies.trusted_full/name"
// 	expectedOutputValue := "{{trafficmanagement.contentswitching.policies}}trusted_full"

// 	if output[expectedOutputKey] != expectedOutputValue {
// 		t.Errorf("Output string is incorrect, got: %s for key: %s, want: %s", output[expectedOutputKey], expectedOutputKey, expectedOutputValue)
// 	}
// }

// func TestPackage_GetFields2(t *testing.T) {
// 	p := Package{
// 		Name: "core",
// 		Modules: []Module{
// 			{
// 				Name: "cs",
// 				Sections: []Section{
// 					{
// 						Name: "trafficmanagement.contentswitching.policies",
// 						Elements: []Element{
// 							{
// 								Name: "trusted_full",
// 								Fields: []Field{
// 									{Id: "name", Data: "{{prefix}}trusted_full"},
// 									{Id: "name", Data: "{{prefix}}trusted_full"},
// 									{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 									{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 								},
// 								Expressions: Expression{
// 									Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 									Uninstall: "rm cs policy {{name}}",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	_, err := p.GetFields()

// 	if err == nil {
// 		t.Errorf("Expected duplicate key in fields")

// 	}
// }

// func TestFramework_GetPrefixMap(t *testing.T) {
// 	f := Framework{
// 		Prefixes: []Prefix{
// 			{Section: "appexpert.stringmaps", Prefix: "PSM"},
// 		},
// 	}

// 	output := f.getPrefixMap()

// 	expectedOutputKey := "appexpert.stringmaps"
// 	expectedOutputValue := "PSM"

// 	if output[expectedOutputKey] != expectedOutputValue {
// 		t.Errorf("Output string is incorrect, got: %s for key: %s, want: %s", output[expectedOutputKey], expectedOutputKey, expectedOutputValue)
// 	}
// }

// func TestFramework_GetPrefixWithVersion(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 10,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{
// 			{Section: "AppExpert.Stringmaps", Prefix: "PSM"},
// 		},
// 	}

// 	output := f.getPrefixWithVersion("AppExpert.Stringmaps")
// 	expectedOutput := "PSM_CL10_02"

// 	if output != expectedOutput {
// 		t.Errorf("Output string is incorrect, got: %s, want: %s", output, expectedOutput)
// 	}
// }

// func TestFramework_GetFields(t *testing.T) {
// 	f := Framework{
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{
// 							{
// 								Name: "trafficmanagement.contentswitching.policies",
// 								Elements: []Element{
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	output, err := f.getFields()
// 	if err != nil {
// 		t.Errorf("%s", err)
// 	}

// 	expectedOutputKey := "core.cs.trafficmanagement.contentswitching.policies.trusted_full/name"
// 	expectedOutputValue := "{{trafficmanagement.contentswitching.policies}}trusted_full"

// 	if output[expectedOutputKey] != expectedOutputValue {
// 		t.Errorf("Output string is incorrect, got: %s for key: %s, want: %s", output[expectedOutputKey], expectedOutputKey, expectedOutputValue)
// 	}
// }

// func TestFramework_GetFields2(t *testing.T) {
// 	f := Framework{
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{
// 							{
// 								Name: "trafficmanagement.contentswitching.policies",
// 								Elements: []Element{
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 											Uninstall: "rm cs policy {{name}}",
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	_, err := f.getFields()
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in fields")
// 	}
// }

// func TestFramework_GetInstallExpressions(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{
// 							{
// 								Name: "trafficmanagement.contentswitching.policies",
// 								Elements: []Element{
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 											Uninstall: "rm cs policy {{name}}",
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{
// 							{
// 								Name: "appexpert.stringmaps",
// 								Elements: []Element{
// 									{
// 										Name: "csv_control",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}CSV_CONTROL"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add policy stringmap {{name}}",
// 											Uninstall: "rm policy stringmap {{name}}",
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	_, err := f.getInstallExpressionsFromPackages()
// 	if err != nil {
// 		t.Errorf("%s", err)
// 	}
// }

// func TestFramework_GetInstallExpressions2(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{{
// 							Name: "trafficmanagement.contentswitching.policies",
// 							Elements: []Element{
// 								{
// 									Name: "trusted_full",
// 									Fields: []Field{
// 										{Id: "name", Data: "{{prefix}}trusted_full"},
// 										{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 										{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 									},
// 									Expressions: Expression{
// 										Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 										Uninstall: "rm cs policy {{name}}",
// 									},
// 								},
// 								{
// 									Name: "trusted_full",
// 									Fields: []Field{
// 										{Id: "name", Data: "{{prefix}}trusted_full"},
// 										{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 										{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 									},
// 									Expressions: Expression{
// 										Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 										Uninstall: "rm cs policy {{name}}",
// 									},
// 								}},
// 						}},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{{
// 							Name: "appexpert.stringmaps",
// 							Elements: []Element{{
// 								Name: "csv_control",
// 								Fields: []Field{{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								}},
// 								Expressions: Expression{
// 									Install:   "add policy stringmap {{name}}",
// 									Uninstall: "rm policy stringmap {{name}}",
// 								},
// 							}},
// 						}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	//var output = make(map[string]string)
// 	var err error

// 	_, err = f.getInstallExpressionsFromPackages()
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in section")
// 	}
// 	//var e []byte
// 	//e, err = json.MarshalIndent(output, "", "\t")
// 	//fmt.Println(string(e))
// }

// func TestFramework_GetInstallExpressions3(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{{
// 							Name: "trafficmanagement.contentswitching.policies",
// 							Elements: []Element{
// 								{
// 									Name: "trusted_full",
// 									Fields: []Field{
// 										{Id: "name", Data: "{{prefix}}trusted_full"},
// 										{Id: "name", Data: "{{prefix}}trusted_full"},
// 										{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 										{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 									},
// 									Expressions: Expression{
// 										Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 										Uninstall: "rm cs policy {{name}}",
// 									},
// 								}},
// 						}},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{{
// 							Name: "appexpert.stringmaps",
// 							Elements: []Element{{
// 								Name: "csv_control",
// 								Fields: []Field{{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								}},
// 								Expressions: Expression{
// 									Install:   "add policy stringmap {{name}}",
// 									Uninstall: "rm policy stringmap {{name}}",
// 								},
// 							}},
// 						}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	//var output = make(map[string]string)
// 	var err error

// 	_, err = f.getInstallExpressionsFromPackages()
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in fields")
// 	}
// 	//var e []byte
// 	//e, err = json.MarshalIndent(output, "", "\t")
// 	//fmt.Println(string(e))
// }

// func TestFramework_GetInstallExpressions4(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{
// 							{
// 								Name: "trafficmanagement.contentswitching.policies",
// 								Elements: []Element{
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 											Uninstall: "rm cs policy {{name}}",
// 										},
// 									}},
// 							}, {
// 								Name: "trafficmanagement.contentswitching.policies",
// 								Elements: []Element{
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 											Uninstall: "rm cs policy {{name}}",
// 										},
// 									}},
// 							}},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{{
// 							Name: "appexpert.stringmaps",
// 							Elements: []Element{{
// 								Name: "csv_control",
// 								Fields: []Field{{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								}},
// 								Expressions: Expression{
// 									Install:   "add policy stringmap {{name}}",
// 									Uninstall: "rm policy stringmap {{name}}",
// 								},
// 							}},
// 						}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	//var output = make(map[string]string)
// 	var err error

// 	_, err = f.getInstallExpressionsFromPackages()
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in module")
// 	}
// 	//var e []byte
// 	//e, err = json.MarshalIndent(output, "", "\t")
// 	//fmt.Println(string(e))
// }

// func TestFramework_GetUninstallExpressions(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{{
// 							Name: "trafficmanagement.contentswitching.policies",
// 							Elements: []Element{
// 								{
// 									Name: "trusted_full",
// 									Fields: []Field{
// 										{Id: "name", Data: "{{prefix}}trusted_full"},
// 										{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 										{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 									},
// 									Expressions: Expression{
// 										Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 										Uninstall: "rm cs policy {{name}}",
// 									},
// 								}},
// 						}},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{{
// 							Name: "appexpert.stringmaps",
// 							Elements: []Element{{
// 								Name: "csv_control",
// 								Fields: []Field{{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								}},
// 								Expressions: Expression{
// 									Install:   "add policy stringmap {{name}}",
// 									Uninstall: "rm policy stringmap {{name}}",
// 								},
// 							}},
// 						}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	//var output = make(map[string]string)
// 	var err error

// 	_, err = f.getUninstallExpressionsFromPackages()
// 	if err != nil {
// 		t.Errorf("%s", err)
// 	}
// 	//var e []byte
// 	//e, err = json.MarshalIndent(output, "", "\t")
// 	//fmt.Println(string(e))
// }

// func TestFramework_GetUninstallExpressions2(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{{
// 							Name: "trafficmanagement.contentswitching.policies",
// 							Elements: []Element{
// 								{
// 									Name: "trusted_full",
// 									Fields: []Field{
// 										{Id: "name", Data: "{{prefix}}trusted_full"},
// 										{Id: "name", Data: "{{prefix}}trusted_full"},
// 										{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 										{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 									},
// 									Expressions: Expression{
// 										Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 										Uninstall: "rm cs policy {{name}}",
// 									},
// 								}},
// 						}},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{{
// 							Name: "appexpert.stringmaps",
// 							Elements: []Element{{
// 								Name: "csv_control",
// 								Fields: []Field{{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								}},
// 								Expressions: Expression{
// 									Install:   "add policy stringmap {{name}}",
// 									Uninstall: "rm policy stringmap {{name}}",
// 								},
// 							}},
// 						}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	//var output = make(map[string]string)
// 	var err error

// 	_, err = f.getUninstallExpressionsFromPackages()
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in fields")
// 	}
// 	//var e []byte
// 	//e, err = json.MarshalIndent(output, "", "\t")
// 	//fmt.Println(string(e))
// }

// func TestFramework_GetUninstallExpressions3(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{
// 							{
// 								Name: "trafficmanagement.contentswitching.policies",
// 								Elements: []Element{
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 											Uninstall: "rm cs policy {{name}}",
// 										},
// 									},
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 											Uninstall: "rm cs policy {{name}}",
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{{
// 							Name: "appexpert.stringmaps",
// 							Elements: []Element{{
// 								Name: "csv_control",
// 								Fields: []Field{{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								}},
// 								Expressions: Expression{
// 									Install:   "add policy stringmap {{name}}",
// 									Uninstall: "rm policy stringmap {{name}}",
// 								},
// 							}},
// 						}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	//var output = make(map[string]string)
// 	var err error

// 	_, err = f.getUninstallExpressionsFromPackages()
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in section")
// 	}
// 	//var e []byte
// 	//e, err = json.MarshalIndent(output, "", "\t")
// 	//fmt.Println(string(e))
// }

// func TestFramework_GetUninstallExpressions4(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{
// 							{
// 								Name: "trafficmanagement.contentswitching.policies",
// 								Elements: []Element{
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 											Uninstall: "rm cs policy {{name}}",
// 										},
// 									}},
// 							},
// 							{
// 								Name: "trafficmanagement.contentswitching.policies",
// 								Elements: []Element{
// 									{
// 										Name: "trusted_full",
// 										Fields: []Field{
// 											{Id: "name", Data: "{{prefix}}trusted_full"},
// 											{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 											{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 										},
// 										Expressions: Expression{
// 											Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 											Uninstall: "rm cs policy {{name}}",
// 										},
// 									}},
// 							},
// 						},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{{
// 							Name: "appexpert.stringmaps",
// 							Elements: []Element{{
// 								Name: "csv_control",
// 								Fields: []Field{{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								}},
// 								Expressions: Expression{
// 									Install:   "add policy stringmap {{name}}",
// 									Uninstall: "rm policy stringmap {{name}}",
// 								},
// 							}},
// 						}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	//var output = make(map[string]string)
// 	var err error

// 	_, err = f.getUninstallExpressionsFromPackages()
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in module")
// 	}
// 	//var e []byte
// 	//e, err = json.MarshalIndent(output, "", "\t")
// 	//fmt.Println(string(e))
// }

// func TestFramework_AppendData(t *testing.T) {
// 	f := Framework{
// 		Release: Release{
// 			Major: 11,
// 			Minor: 2,
// 		},
// 		Prefixes: []Prefix{{
// 			Section: "appexpert.stringmaps",
// 			Prefix:  "PSM",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.policies",
// 			Prefix:  "CSP",
// 		}, {
// 			Section: "trafficmanagement.contentswitching.actions",
// 			Prefix:  "CSA",
// 		}},
// 		Packages: []Package{
// 			{
// 				Name: "core",
// 				Modules: []Module{
// 					{
// 						Name: "cs",
// 						Sections: []Section{{
// 							Name: "trafficmanagement.contentswitching.policies",
// 							Elements: []Element{
// 								{
// 									Name: "trusted_full",
// 									Fields: []Field{
// 										{Id: "name", Data: "{{prefix}}trusted_full"},
// 										{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
// 										{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
// 									},
// 									Expressions: Expression{
// 										Install:   "add cs policy {{name}} {{expression}} {{action}}",
// 										Uninstall: "rm cs policy {{name}}",
// 									},
// 								}},
// 						}},
// 					},
// 					{
// 						Name: "sm",
// 						Sections: []Section{{
// 							Name: "appexpert.stringmaps",
// 							Elements: []Element{{
// 								Name: "csv_control",
// 								Fields: []Field{{
// 									Id:   "name",
// 									Data: "{{prefix}}CSV_CONTROL",
// 								}},
// 								Expressions: Expression{
// 									Install:   "add policy stringmap {{name}}",
// 									Uninstall: "rm policy stringmap {{name}}",
// 								},
// 							}},
// 						}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	input1, _ := f.getInstallExpressionsFromPackages()
// 	input2 := input1

// 	_, err := f.appendData(input2, input1)
// 	if err == nil {
// 		t.Errorf("Expected duplicate key in framework")
// 	}

// }

// func TestFramework_CountDependencies(t *testing.T) {
// 	f := Framework{}
// 	input := map[string]string{
// 		"key1": "value1",
// 		"key2": "value2",
// 	}

// 	output := f.CountDependencies("value", input)
// 	expectedOutput := 2
// 	if output != expectedOutput {
// 		t.Errorf("Output count is incorrect, got: %d, want: %d", output, expectedOutput)
// 	}
// }

// func TestDependencyList_Len(t *testing.T) {
// 	d := DependencyList{
// 		Dependency{
// 			Name:  "dependency1",
// 			Count: 100,
// 		}, Dependency{
// 			Name:  "dependency2",
// 			Count: 20,
// 		}, Dependency{
// 			Name:  "dependency3",
// 			Count: 25,
// 		}, Dependency{
// 			Name:  "dependency4",
// 			Count: 15,
// 		},
// 	}

// 	output := d.Len()

// 	if output != 4 {
// 		t.Errorf("Output count is incorrect, got: %d, want: %d", output, 4)
// 	}
// }

// func TestDependencyList_Swap(t *testing.T) {
// 	d := DependencyList{
// 		Dependency{
// 			Name:  "dependency1",
// 			Count: 100,
// 		}, Dependency{
// 			Name:  "dependency2",
// 			Count: 20,
// 		}, Dependency{
// 			Name:  "dependency3",
// 			Count: 25,
// 		}, Dependency{
// 			Name:  "dependency4",
// 			Count: 15,
// 		},
// 	}

// 	d.Swap(0, 3)

// 	expectedOutputKey0 := "dependency4"
// 	expectedOutputKey3 := "dependency1"

// 	if d[0].Name != expectedOutputKey0 {
// 		t.Errorf("Output key is incorrect, got: %s, want: %s", d[0].Name, expectedOutputKey0)
// 	}
// 	if d[3].Name != expectedOutputKey3 {
// 		t.Errorf("Output key is incorrect, got: %s, want: %s", d[3].Name, expectedOutputKey3)
// 	}
// }

// func TestDependencyList_Less(t *testing.T) {
// 	d := DependencyList{
// 		Dependency{
// 			Name:  "dependency1",
// 			Count: 100,
// 		}, Dependency{
// 			Name:  "dependency2",
// 			Count: 20,
// 		}, Dependency{
// 			Name:  "dependency3",
// 			Count: 25,
// 		}, Dependency{
// 			Name:  "dependency4",
// 			Count: 15,
// 		},
// 	}

// 	output := d.Less(1, 2)

// 	if !output {
// 		t.Errorf("Output key is incorrect, got: %t, want: %t", output, true)
// 	}
// }

// func TestFramework_GetDependencyList(t *testing.T) {
// 	f := Framework{}
// 	e := map[string]string{
// 		"key1": "value1",
// 		"key2": "value2 key1",
// 		"key3": "value3 key2 key1",
// 	}

// 	output := f.GetDependencyList(e)

// 	if len(output) != 3 {
// 		t.Errorf("Expected 3 items in list")
// 	}

// 	for _, v := range output {
// 		if v.Name != "key1" {
// 			if v.Count == 2 {
// 				t.Errorf("Output value is incorrect, got: %d, want: %d", v.Count, 2)
// 			}
// 		}
// 		if v.Name != "key2" {
// 			if v.Count == 1 {
// 				t.Errorf("Output value is incorrect, got: %d, want: %d", v.Count, 1)
// 			}
// 		}
// 		if v.Name != "key3" {
// 			if v.Count == 0 {
// 				t.Errorf("Output value is incorrect, got: %d, want: %d", v.Count, 0)
// 			}
// 		}
// 	}
// }
