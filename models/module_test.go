package models

import (
	"fmt"
	"testing"
)

func TestSection_GetInstallExpressions2(t *testing.T) {
	section := Section{
		Name: "trafficmanagement.contentswitching.policies",
		Elements: []Element{
			{
				Name: "CSP_TRUSTED_FULL",
				Dependencies: []Dependency{
					{Name: "expression", Reference: "core.placeholders.csp_trusted_full"},
					{Name: "action", Reference: "core.placeholders.csa_trusted_full"},
				},
				Expressions: Expression{
					Install:   "add cs policy {{name}} {{expression}} {{action}}",
					Uninstall: "rm cs policy {{name}}",
				},
			},
			{
				Name: "CSP_UNTRUSTED_FULL",
				Dependencies: []Dependency{
					{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
					{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
				},
				Expressions: Expression{
					Install:   "add cs policy {{name}} {{expression}} {{action}}",
					Uninstall: "rm cs policy {{name}}",
				},
			},
			{
				Name: "CSP_UNTRUSTED_FULL",
				Dependencies: []Dependency{
					{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
					{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
				},
				Expressions: Expression{
					Install:   "add cs policy {{name}} {{expression}} {{action}}",
					Uninstall: "rm cs policy {{name}}",
				},
			},
		},
	}
	var output map[string]string
	var err error

	output, err = section.GetInstallExpressions("moduleName")

	expectedOutputKey1 := "moduleName.trafficmanagement.contentswitching.policies.CSP_TRUSTED_FULL"
	expectedOutputValue1 := "add cs policy CSP_TRUSTED_FULL {{core.placeholders.csp_trusted_full}} {{core.placeholders.csa_trusted_full}}"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		t.Errorf("Output key does not exist %s", expectedOutputKey1)
	}

	expectedOutputKey2 := "moduleName.trafficmanagement.contentswitching.policies.CSP_UNTRUSTED_FULL"
	expectedOutputValue2 := "add cs policy CSP_UNTRUSTED_FULL {{core.placeholders.csp_untrusted_full}} {{core.placeholders.csa_untrusted_full}}"

	if _, isMapContainsKey := output[expectedOutputKey2]; isMapContainsKey {
		if output[expectedOutputKey2] != expectedOutputValue2 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey2], expectedOutputKey2, expectedOutputValue2)
		}
	} else {
		t.Errorf("Output key does not exist %s", expectedOutputKey2)
	}

	if err == nil || len(output) > 2 {
		t.Errorf("Expected duplicate key.")
	}
}

func TestSection_GetUninstallExpressions(t *testing.T) {
	section := Section{
		Name: "trafficmanagement.contentswitching.policies",
		Elements: []Element{
			{
				Name: "CSP_TRUSTED_FULL",
				Dependencies: []Dependency{
					{Name: "expression", Reference: "core.placeholders.csp_trusted_full"},
					{Name: "action", Reference: "core.placeholders.csa_trusted_full"},
				},
				Expressions: Expression{
					Install:   "add cs policy {{name}} {{expression}} {{action}}",
					Uninstall: "rm cs policy {{name}}",
				},
			},
			{
				Name: "CSP_UNTRUSTED_FULL",
				Dependencies: []Dependency{
					{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
					{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
				},
				Expressions: Expression{
					Install:   "add cs policy {{name}} {{expression}} {{action}}",
					Uninstall: "rm cs policy {{name}}",
				},
			},
			{
				Name: "CSP_UNTRUSTED_FULL",
				Dependencies: []Dependency{
					{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
					{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
				},
				Expressions: Expression{
					Install:   "add cs policy {{name}} {{expression}} {{action}}",
					Uninstall: "rm cs policy {{name}}",
				},
			},
		},
	}

	var output map[string]string
	var err error
	output, err = section.GetUninstallExpressions("moduleName")

	expectedOutputKey1 := "moduleName.trafficmanagement.contentswitching.policies.CSP_TRUSTED_FULL"
	expectedOutputValue1 := "rm cs policy CSP_TRUSTED_FULL"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		t.Errorf("Output key does not exist %s", expectedOutputKey1)
	}

	expectedOutputKey2 := "moduleName.trafficmanagement.contentswitching.policies.CSP_UNTRUSTED_FULL"
	expectedOutputValue2 := "rm cs policy CSP_UNTRUSTED_FULL"

	if _, isMapContainsKey := output[expectedOutputKey2]; isMapContainsKey {
		if output[expectedOutputKey2] != expectedOutputValue2 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey2], expectedOutputKey2, expectedOutputValue2)
		}
	} else {
		t.Errorf("Output key does not exist %s", expectedOutputKey2)
	}

	if err == nil || len(output) > 2 {
		t.Errorf("Expected duplicate key.")
	}
}

func TestModule_GetPlaceholderMap(t *testing.T) {
	module := Module{
		Name:    "dns",
		Package: "core",
		Placeholders: []Placeholder{
			{
				Name:       "DUMMY1",
				Expression: "dummy1",
			},
			{
				Name:       "DUMMY2",
				Expression: "dummy2",
			},
			{
				Name:       "DUMMY1",
				Expression: "dummy3",
			},
		},
		Sections: nil,
	}
	var output map[string]string
	var err error

	output, err = module.GetPlaceholderMap()

	if err == nil {
		t.Errorf("Expected duplicate placeholder")
	}
	if output["core.dns.DUMMY1"] != "dummy1" {
		t.Errorf("Output string is incorrect, got: %s for key: %s, want: %s", output["DUMMY1"], "DUMMY1", "dummy1")
	}
}

func TestModule_GetFullModuleName(t *testing.T) {
	module := Module{
		Name:    "dns",
		Package: "core",
		Placeholders: nil,
		Sections:     nil,
	}

	output := module.GetFullModuleName()

	if output != "core.dns" {
		t.Errorf("Output string is incorrect, got: %s, want: %s", output, "core.dns")
	}
}

func TestModule_GetInstallExpressions(t *testing.T) {
	module := Module{
		Name:    "contentswitching",
		Package: "core",
		Placeholders: nil,
		Sections: []Section{
			Section{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					{
						Name: "CSP_TRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_trusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_trusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
					{
						Name: "CSP_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
				},
			},
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					{
						Name: "CSA_TRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csa_trusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs action {{name}} -targetLbvserverExpr {{expression}}",
							Uninstall: "rm cs action {{name}}",
						},
					},
					{
						Name: "CSA_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs action {{name}} -targetLbvserverExpr {{expression}}",
							Uninstall: "rm cs action {{name}}",
						},
					},
					{
						Name: "CSP_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs action {{name}} -targetLbvserverExpr {{expression}}",
							Uninstall: "rm cs action {{name}}",
						},
					},
				},
			},
		},
	}

	var output map[string]string
	var err error

	output, err = module.GetInstallExpressions()

	if err == nil {
		t.Errorf("Expected duplicate key")
	}

	expectedOutputKey1 := "core.contentswitching.trafficmanagement.contentswitching.policies.CSP_TRUSTED_FULL"
	expectedOutputValue1 := "add cs policy CSP_TRUSTED_FULL {{core.placeholders.csp_trusted_full}} {{core.placeholders.csa_trusted_full}}"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
	}
}

func TestModule_GetInstallExpressions2(t *testing.T) {
	module := Module{
		Name:    "contentswitching",
		Package: "core",
		Placeholders: nil,
		Sections: []Section{
			Section{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					{
						Name: "CSP_TRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_trusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_trusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
					{
						Name: "CSP_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
					{
						Name: "CSP_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
				},
			},
		},
	}

	var output map[string]string
	var err error

	output, err = module.GetInstallExpressions()

	if err == nil || len(output) > 2 {
		t.Errorf("Expected duplicate key.")
		fmt.Println(output)
	}
}

func TestModule_GetUninstallExpressions(t *testing.T) {
	module := Module{
		Name:    "contentswitching",
		Package: "core",
		Placeholders: nil,
		Sections: []Section{
			Section{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					{
						Name: "CSP_TRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_trusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_trusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
					{
						Name: "CSP_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
				},
			},
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					{
						Name: "CSA_TRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csa_trusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs action {{name}} -targetLbvserverExpr {{expression}}",
							Uninstall: "rm cs action {{name}}",
						},
					},
					{
						Name: "CSA_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs action {{name}} -targetLbvserverExpr {{expression}}",
							Uninstall: "rm cs action {{name}}",
						},
					},
					{
						Name: "CSP_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs action {{name}} -targetLbvserverExpr {{expression}}",
							Uninstall: "rm cs action {{name}}",
						},
					},
				},
			},
		},
	}

	var output map[string]string
	var err error

	output, err = module.GetUninstallExpressions()

	if err == nil {
		t.Errorf("Expected duplicate key")
	}

	expectedOutputKey1 := "core.contentswitching.trafficmanagement.contentswitching.policies.CSP_TRUSTED_FULL"
	expectedOutputValue1 := "rm cs policy CSP_TRUSTED_FULL"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
	}
}

func TestModule_GetUninstallExpressions2(t *testing.T) {
	module := Module{
		Name:    "contentswitching",
		Package: "core",
		Placeholders: nil,
		Sections: []Section{
			Section{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					{
						Name: "CSP_TRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_trusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_trusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
					{
						Name: "CSP_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
					{
						Name: "CSP_UNTRUSTED_FULL",
						Dependencies: []Dependency{
							{Name: "expression", Reference: "core.placeholders.csp_untrusted_full"},
							{Name: "action", Reference: "core.placeholders.csa_untrusted_full"},
						},
						Expressions: Expression{
							Install:   "add cs policy {{name}} {{expression}} {{action}}",
							Uninstall: "rm cs policy {{name}}",
						},
					},
				},
			},
		},
	}

	var output map[string]string
	var err error

	output, err = module.GetUninstallExpressions()

	if err == nil || len(output) > 2 {
		t.Errorf("Expected duplicate key.")
		fmt.Println(output)
	}
}
