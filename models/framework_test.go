package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRelease_GetVersionAsString(t *testing.T) {
	r := Release{
		Major: 10,
		Minor: 1,
	}

	output := r.GetVersionAsString()
	expectedOutput := "CL10_01"
	if output != expectedOutput {
		t.Errorf("Output string is incorrect, got: %s, want: %s.", output, expectedOutput)
	}
}

func TestFramework_GetPrefixMap(t *testing.T) {
	f := Framework{
		Prefixes: []SectionPrefix{
			SectionPrefix{Name: "appexpert.stringmaps", Prefix: "PSM"},
		},
	}

	output := f.GetPrefixMap()

	expectedOutputKey := "appexpert.stringmaps"
	expectedOutputValue := "PSM"

	if output[expectedOutputKey] != expectedOutputValue {
		t.Errorf("Output string is incorrect, got: %s for key: %s, want: %s", output[expectedOutputKey], expectedOutputKey, expectedOutputValue)
	}
}

func TestFramework_GetPrefixWithVersion(t *testing.T) {
	f := Framework{
		Release: Release{
			Major: 10,
			Minor: 2,
		},
		Prefixes: []SectionPrefix{
			SectionPrefix{Name: "AppExpert.Stringmaps", Prefix: "PSM"},
		},
	}

	output := f.GetPrefixWithVersion("AppExpert.Stringmaps")
	expectedOutput := "PSM_CL10_02"

	if output != expectedOutput {
		t.Errorf("Output string is incorrect, got: %s, want: %s", output, expectedOutput)
	}
}

func TestFramework_GetInstallExpressions(t *testing.T) {
	f := Framework{
		Release: Release{
			Major: 11,
			Minor: 2,
		},
		Prefixes: []SectionPrefix{{
			Name:   "appexpert.stringmaps",
			Prefix: "PSM",
		}, {
			Name:   "trafficmanagement.contentswitching.policies",
			Prefix: "CSP",
		}, {
			Name:   "trafficmanagement.contentswitching.actions",
			Prefix: "CSA",
		}},
		Packages: []Package{
			{
				Name: "core",
				Modules: []Module{
					{
						Name: "cs",
						Sections: []Section{{
							Name: "trafficmanagement.contentswitching.policies",
							Elements: []Element{
								{
									Name: "trusted_full",
									Fields: []Field{
										{Id: "name", Data: "{{prefix}}trusted_full"},
										{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
										{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
									},
									Expressions: Expression{
										Install:   "add cs policy {{name}} {{expression}} {{action}}",
										Uninstall: "rm cs policy {{name}}",
									},
								}},
						}},
					},
					{
						Name: "sm",
						Sections: []Section{{
							Name: "appexpert.stringmaps",
							Elements: []Element{{
								Name: "cs_control",
								Fields: []Field{{
									Id:   "name",
									Data: "{{prefix}}CS_CONTROL",
								}},
								Expressions: Expression{
									Install:   "add policy stringmap {{name}}",
									Uninstall: "rm policy stringmap {{name}}",
								},
							}},
						}},
					},
				},
			},
		},
	}

	var output = make(map[string]string)
	var err error

	output, err = f.GetInstallExpressions()
	if err != nil {
		t.Errorf("%s", err)
	}
	var e []byte
	e, err = json.MarshalIndent(output, "", "\t")
	fmt.Println(string(e))
}

func TestFramework_GetUninstallExpressions(t *testing.T) {
	f := Framework{
		Release: Release{
			Major: 11,
			Minor: 2,
		},
		Prefixes: []SectionPrefix{{
			Name:   "appexpert.stringmaps",
			Prefix: "PSM",
		}, {
			Name:   "trafficmanagement.contentswitching.policies",
			Prefix: "CSP",
		}, {
			Name:   "trafficmanagement.contentswitching.actions",
			Prefix: "CSA",
		}},
		Packages: []Package{
			{
				Name: "core",
				Modules: []Module{
					{
						Name: "cs",
						Sections: []Section{{
							Name: "trafficmanagement.contentswitching.policies",
							Elements: []Element{
								{
									Name: "trusted_full",
									Fields: []Field{
										{Id: "name", Data: "{{prefix}}trusted_full"},
										{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
										{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
									},
									Expressions: Expression{
										Install:   "add cs policy {{name}} {{expression}} {{action}}",
										Uninstall: "rm cs policy {{name}}",
									},
								}},
						}},
					},
					{
						Name: "sm",
						Sections: []Section{{
							Name: "appexpert.stringmaps",
							Elements: []Element{{
								Name: "cs_control",
								Fields: []Field{{
									Id:   "name",
									Data: "{{prefix}}CS_CONTROL",
								}},
								Expressions: Expression{
									Install:   "add policy stringmap {{name}}",
									Uninstall: "rm policy stringmap {{name}}",
								},
							}},
						}},
					},
				},
			},
		},
	}

	var output = make(map[string]string)
	var err error

	output, err = f.GetUninstallExpressions()
	if err != nil {
		t.Errorf("%s", err)
	}
	var e []byte
	e, err = json.MarshalIndent(output, "", "\t")
	fmt.Println(string(e))
}
