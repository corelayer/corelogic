package models

import (
	"fmt"
	"testing"
)

var eCspTrustedFull = Element{
	Name: "TRUSTED_FULL",
	Fields: []Field{
		{Id: "name", Data: "{{prefix}}TRUSTED_FULL"},
		{Id: "name", Data: "{{prefix}}TRUSTED_FULL"},
		{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.trusted_full/name}}}"},
		{Id: "action", Data: "{{core.placeholders.csa_trusted_full}}"},
	},
	Expressions: Expression{
		Install:   "add cs policy {{name}} {{expression}} {{action}}",
		Uninstall: "rm cs policy {{name}}",
	},
}

var eCspUntrustedFull = Element{
	Name: "UNTRUSTED_FULL",
	Fields: []Field{
		{Id: "name", Data: "UNTRUSTED_FULL"},
		{Id: "expression", Data: "q{{{core.appexpert.expressions.contentswitching.policies.untrusted_full/name}}}"},
		{Id: "action", Data: "{{core.placeholders.csa_untrusted_full}}"},
	},
	Expressions: Expression{
		Install:   "add cs policy {{name}} {{expression}} {{action}}",
		Uninstall: "rm cs policy {{name}}",
	},
}

var eCsaTrustedFull = Element{
	Name: "TRUSTED_FULL",
	Fields: []Field{
		{Id: "name", Data: "TRUSTED_FULL"},
		{Id: "expression", Data: "{{core.contentswitching.appexpert.expressions.CSA_TRUSTED_FULL/name}}"},
	},
	Expressions: Expression{
		Install:   "add cs action {{name}} -targetVserverExpr {{expression}}",
		Uninstall: "rm cs action {{name}}",
	},
}

var eCsaUntrustedFull = Element{
	Name: "UNTRUSTED_FULL",
	Fields: []Field{
		{Id: "name", Data: "UNTRUSTED_FULL"},
		{Id: "expression", Data: "{{core.contentswitching.appexpert.expressions.CSA_UNTRUSTED_FULL/name}}"},
	},
	Expressions: Expression{
		Install:   "add cs action {{name}} -targetVserverExpr {{expression}}",
		Uninstall: "rm cs action {{name}}",
	},
}

var sTrafficManagementContentSwitchingPolicies = Section{
	Name: "trafficmanagement.contentswitching.policies",
	Elements: []Element{
		eCspTrustedFull,
		eCspUntrustedFull,
		eCspUntrustedFull,
	},
}

var sTrafficManagementContentSwitchingAction = Section{
	Name: "trafficmanagement.contentswitching.actions",
	Elements: []Element{
		eCsaTrustedFull,
		eCsaUntrustedFull,
		eCsaUntrustedFull,
	},
}


func TestElement_GetFullName(t *testing.T) {
	output := eCspTrustedFull.GetFullName("prefix")
	expectedOutput := "prefix.TRUSTED_FULL"

	if output != expectedOutput {
		t.Errorf("Output string is incorrect, got: %q, want: %q", output, expectedOutput)
	}
}

func TestSection_GetFullName(t *testing.T) {
	output := sTrafficManagementContentSwitchingPolicies.GetFullName("prefix")
	expectedOutput := "prefix.trafficmanagement.contentswitching.policies"

	if output != expectedOutput {
		t.Errorf("Output string is incorrect, got: %q, want: %q", output, expectedOutput)
	}
}

func TestSection_GetFields(t *testing.T) {
	section := Section{
		Name: "trafficmanagement.contentswitching.policies",
		Elements: []Element{
			eCspTrustedFull,
			eCspUntrustedFull,
		},
	}

	var output map[string]string
	var err error

	output, err = section.GetFields("moduleName")

	expectedOutputKey1 := "moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL/name"
	expectedOutputValue1 := "{{prefix}}TRUSTED_FULL"

	if err == nil {
		t.Errorf("Expected duplicate key.")
	}
	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)

		}
	} else {
		t.Errorf("Output key does not exist %s", expectedOutputKey1)
	}
}

func TestSection_GetInstallExpressions(t *testing.T) {
	section := sTrafficManagementContentSwitchingPolicies

	var output map[string]string
	var err error

	output, err = section.GetInstallExpressions("moduleName")

	expectedOutputKey1 := "moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
	expectedOutputValue1 := "add cs policy {{name}} {{expression}} {{action}}"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		t.Errorf("Output key does not exist %s", expectedOutputKey1)
	}

	expectedOutputKey2 := "moduleName.trafficmanagement.contentswitching.policies.UNTRUSTED_FULL"
	expectedOutputValue2 := "add cs policy {{name}} {{expression}} {{action}}"

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
	section := sTrafficManagementContentSwitchingPolicies

	var output map[string]string
	var err error
	output, err = section.GetUninstallExpressions("moduleName")

	expectedOutputKey1 := "moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
	expectedOutputValue1 := "rm cs policy {{name}}"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		t.Errorf("Output key does not exist %s", expectedOutputKey1)
	}

	expectedOutputKey2 := "moduleName.trafficmanagement.contentswitching.policies.UNTRUSTED_FULL"
	expectedOutputValue2 := "rm cs policy {{name}}"

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

func TestModule_GetFullModuleName(t *testing.T) {
	module := Module{
		Name:     "dns",
		Package:  "core",
		Sections: nil,
	}

	output := module.GetFullModuleName()

	if output != "core.dns" {
		t.Errorf("Output string is incorrect, got: %s, want: %s", output, "core.dns")
	}
}

// Correct Module definition
func TestModule_GetInstallExpressions(t *testing.T) {
	m := Module{
		Name:    "moduleName",
		Package: "packageName",
		Sections: []Section{
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					eCspTrustedFull,
					eCspUntrustedFull,
				},
			},
			{
				Name: "trafficmanagement.contentswitching.actions",
				Elements: []Element{
					eCsaTrustedFull,
					eCsaUntrustedFull,
				},
			},
		},
	}

	var output map[string]string
	var err error

	output, err = m.GetInstallExpressions()

	if err != nil {
		t.Errorf("Unexpected duplicate key")
	}

	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
	expectedOutputValue1 := "add cs policy {{name}} {{expression}} {{action}}"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		for k, _ := range output {
			fmt.Println(k)
		}
		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
	}
}

// Duplicate key between sections
func TestModule_GetInstallExpressions2(t *testing.T) {
	m := Module{
		Name:    "moduleName",
		Package: "packageName",
		Sections: []Section{
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					eCspTrustedFull,
					eCspUntrustedFull,
				},
			},
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					eCsaTrustedFull,
				},
			},
		},
	}

	var output map[string]string
	var err error

	output, err = m.GetInstallExpressions()

	if err == nil {
		t.Errorf("Unexpected duplicate key with error %s", err)
	}

	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
	expectedOutputValue1 := "add cs policy {{name}} {{expression}} {{action}}"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		for k, _ := range output {
			fmt.Println(k)
		}
		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
	}
}


// Duplicate key in section
func TestModule_GetInstallExpressions3(t *testing.T) {
	module  := Module{
		Name:    "moduleName",
		Package: "packageName",
		Sections: []Section{
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					eCspTrustedFull,
					eCspUntrustedFull,
					eCspUntrustedFull,
				},
			},
			{
				Name: "trafficmanagement.contentswitching.actions",
				Elements: []Element{
					eCsaTrustedFull,
					eCsaUntrustedFull,
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





// Correct Module definition
func TestModule_GetUninstallExpressions(t *testing.T) {
	m := Module{
		Name:    "moduleName",
		Package: "packageName",
		Sections: []Section{
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					eCspTrustedFull,
					eCspUntrustedFull,
				},
			},
			{
				Name: "trafficmanagement.contentswitching.actions",
				Elements: []Element{
					eCsaTrustedFull,
					eCsaUntrustedFull,
				},
			},
		},
	}

	var output map[string]string
	var err error

	output, err = m.GetUninstallExpressions()

	if err != nil {
		t.Errorf("Unexpected duplicate key")
	}

	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
	expectedOutputValue1 := "rm cs policy {{name}}"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		for k, _ := range output {
			fmt.Println(k)
		}
		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
	}
}

// Duplicate key between sections
func TestModule_GetUninstallExpressions2(t *testing.T) {
	m := Module{
		Name:    "moduleName",
		Package: "packageName",
		Sections: []Section{
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					eCspTrustedFull,
					eCspUntrustedFull,
				},
			},
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					eCsaTrustedFull,
				},
			},
		},
	}

	var output map[string]string
	var err error

	output, err = m.GetUninstallExpressions()

	if err == nil {
		t.Errorf("Unexpected duplicate key with error %s", err)
	}

	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
	expectedOutputValue1 := "rm cs policy {{name}}"

	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
		if output[expectedOutputKey1] != expectedOutputValue1 {
			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
		}
	} else {
		for k, _ := range output {
			fmt.Println(k)
		}
		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
	}
}


// Duplicate key in section
func TestModule_GetUninstallExpressions3(t *testing.T) {
	module  := Module{
		Name:    "moduleName",
		Package: "packageName",
		Sections: []Section{
			{
				Name: "trafficmanagement.contentswitching.policies",
				Elements: []Element{
					eCspTrustedFull,
					eCspUntrustedFull,
					eCspUntrustedFull,
				},
			},
			{
				Name: "trafficmanagement.contentswitching.actions",
				Elements: []Element{
					eCsaTrustedFull,
					eCsaUntrustedFull,
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