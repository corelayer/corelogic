package models

// import (
// 	"fmt"
// 	"testing"
// )

// func TestElement_GetFullName(t *testing.T) {
// 	e := Element{
// 		Name: "TRUSTED_FULL",
// 		Fields: []Field{
// 			{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 			{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 			{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 			{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 		},
// 		Expressions: Expression{
// 			Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 			Uninstall: "rm cs policy <<name>>",
// 		},
// 	}

// 	output := e.GetFullName("prefix")
// 	expectedOutput := "prefix.TRUSTED_FULL"

// 	if output != expectedOutput {
// 		t.Errorf("Output string is incorrect, got: %q, want: %q", output, expectedOutput)
// 	}
// }

// func TestElement_GetFullyQualifiedExpression(t *testing.T) {
// 	e := Element{
// 		Name: "TRUSTED_FULL",
// 		Fields: []Field{
// 			{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 			{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 			{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 		},
// 		Expressions: Expression{
// 			Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 			Uninstall: "rm cs policy <<name>>",
// 		},
// 	}

// 	output, _ := e.GetFullyQualifiedExpression(e.Expressions.Install, "packageName.moduleName")
// 	expectedOutput := "add cs policy <<prefix>>TRUSTED_FULL q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>} <<core.placeholders.csa_trusted_full>>"

// 	if output != expectedOutput {
// 		t.Errorf("Output string is incorrect, got: %q, want: %q", output, expectedOutput)
// 	}

// }

// func TestSection_GetFullName(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 	}

// 	output := s.GetFullName("prefix")
// 	expectedOutput := "prefix.trafficmanagement.contentswitching.policies"

// 	if output != expectedOutput {
// 		t.Errorf("Output string is incorrect, got: %q, want: %q", output, expectedOutput)
// 	}
// }

// func TestSection_ExpandSectionPrefix(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.loadbalancing.servers",
// 	}

// 	output := s.expandSectionPrefix("<<prefix>>expression")
// 	expectedOutput := "<<trafficmanagement.loadbalancing.servers>>expression"
// 	if output != expectedOutput {
// 		t.Errorf("Output string is incorrect, got: %s, want: %s", output, expectedOutput)
// 	}
// }

// func TestSection_GetFields(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 		Elements: []Element{
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 			{
// 				Name: "UNTRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 					{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = s.GetFields("packageName.moduleName")

// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL/name"
// 	expectedOutputValue1 := "<<trafficmanagement.contentswitching.policies>>TRUSTED_FULL"

// 	if err != nil {
// 		t.Errorf("Unexpected duplicate key.")
// 	}
// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)

// 		}
// 	} else {
// 		t.Errorf("Output key does not exist %s", expectedOutputKey1)
// 	}
// }

// func TestSection_GetField2(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 		Elements: []Element{
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 			{
// 				Name: "UNTRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 					{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = s.GetFields("packageName.moduleName")

// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL/name"
// 	expectedOutputValue1 := "<<trafficmanagement.contentswitching.policies>>TRUSTED_FULL"

// 	if err == nil {
// 		t.Errorf("Expected duplicate key.")
// 	}
// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)

// 		}
// 	} else {
// 		t.Errorf("Output key does not exist %s", expectedOutputKey1)
// 	}
// }

// func TestSection_GetInstallExpressions(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 		Elements: []Element{
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 			{
// 				Name: "UNTRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 					{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = s.GetInstallExpressions("packageName.moduleName")

// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}
// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
// 	expectedOutputValue1 := "add cs policy <<trafficmanagement.contentswitching.policies>>TRUSTED_FULL q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>} <<core.placeholders.csa_trusted_full>>"

// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
// 		}
// 	} else {
// 		t.Errorf("Output key does not exist %s", expectedOutputKey1)
// 	}

// 	expectedOutputKey2 := "packageName.moduleName.trafficmanagement.contentswitching.policies.UNTRUSTED_FULL"
// 	expectedOutputValue2 := "add cs policy <<trafficmanagement.contentswitching.policies>>UNTRUSTED_FULL q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>} <<core.placeholders.csa_untrusted_full>>"

// 	if _, isMapContainsKey := output[expectedOutputKey2]; isMapContainsKey {
// 		if output[expectedOutputKey2] != expectedOutputValue2 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey2], expectedOutputKey2, expectedOutputValue2)
// 		}
// 	} else {
// 		t.Errorf("Output key does not exist %s", expectedOutputKey2)
// 	}
// }

// func TestSection_GetInstallExpressions2(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 		Elements: []Element{
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 			{
// 				Name: "UNTRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 					{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 		},
// 	}

// 	var err error
// 	_, err = s.GetInstallExpressions("packageName.moduleName")

// 	if err == nil {
// 		t.Errorf("Expected duplicate key in fields")
// 	}
// }

// func TestSection_GetInstallExpressions3(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 		Elements: []Element{
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 					{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 		},
// 	}

// 	var err error
// 	_, err = s.GetInstallExpressions("packageName.moduleName")

// 	if err == nil {
// 		t.Errorf("Expected duplicate key in fields")
// 	}
// }

// func TestSection_GetUninstallExpressions(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 		Elements: []Element{
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 			{
// 				Name: "UNTRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 					{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = s.GetUninstallExpressions("packageName.moduleName")

// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}
// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
// 	expectedOutputValue1 := "rm cs policy <<trafficmanagement.contentswitching.policies>>TRUSTED_FULL"

// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
// 		}
// 	} else {
// 		t.Errorf("Output key does not exist %s", expectedOutputKey1)
// 	}

// 	expectedOutputKey2 := "packageName.moduleName.trafficmanagement.contentswitching.policies.UNTRUSTED_FULL"
// 	expectedOutputValue2 := "rm cs policy <<trafficmanagement.contentswitching.policies>>UNTRUSTED_FULL"

// 	if _, isMapContainsKey := output[expectedOutputKey2]; isMapContainsKey {
// 		if output[expectedOutputKey2] != expectedOutputValue2 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey2], expectedOutputKey2, expectedOutputValue2)
// 		}
// 	} else {
// 		t.Errorf("Output key does not exist %s", expectedOutputKey2)
// 	}
// }

// func TestSection_GetUninstallExpressions2(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 		Elements: []Element{
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 			{
// 				Name: "UNTRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 					{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 		},
// 	}

// 	var err error

// 	_, err = s.GetUninstallExpressions("packageName.moduleName")

// 	if err == nil {
// 		t.Errorf("Expected duplicate key in fields")
// 	}
// }

// func TestSection_GetUninstallExpressions3(t *testing.T) {
// 	s := Section{
// 		Name: "trafficmanagement.contentswitching.policies",
// 		Elements: []Element{
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 					{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 			{
// 				Name: "TRUSTED_FULL",
// 				Fields: []Field{
// 					{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 					{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 					{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 				},
// 				Expressions: Expression{
// 					Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 					Uninstall: "rm cs policy <<name>>",
// 				},
// 			},
// 		},
// 	}

// 	var err error

// 	_, err = s.GetUninstallExpressions("packageName.moduleName")

// 	if err == nil {
// 		t.Errorf("Expected duplicate key in fields")
// 	}
// }

// func TestModule_GetFullModuleName(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 	}

// 	output := m.GetFullModuleName("packageName")
// 	expectedOutput := "packageName.moduleName"

// 	if output != expectedOutput {
// 		t.Errorf("Output string is incorrect, got: %s, want: %s", output, expectedOutput)
// 	}
// }

// func TestModule_GetFields(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 		Sections: []Section{
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Name: "trafficmanagement.contentswitching.actions",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "<<core.contentswitching.appexpert.expressions.CSA_TRUSTED_FULL/name>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs action <<name>> -targetVserverExpr <<expression>>",
// 							Uninstall: "rm cs action <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "<<core.contentswitching.appexpert.expressions.CSA_UNTRUSTED_FULL/name>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs action <<name>> -targetVserverExpr <<expression>>",
// 							Uninstall: "rm cs action <<name>>",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = m.GetFields("packageName")

// 	if err != nil {
// 		t.Errorf("Unexpected duplicate key")
// 	}

// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.actions.UNTRUSTED_FULL/name"
// 	expectedOutputValue1 := "<<trafficmanagement.contentswitching.actions>>UNTRUSTED_FULL"

// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
// 		}
// 	} else {
// 		for k := range output {
// 			fmt.Println(k)
// 		}
// 		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
// 	}
// }

// func TestModule_GetFields2(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 		Sections: []Section{
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var err error
// 	_, err = m.GetFields("packageName")

// 	if err == nil {
// 		t.Errorf("Expected duplicate key")
// 	}
// }

// // Correct Module definition
// func TestModule_GetInstallExpressions(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 		Sections: []Section{
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = m.GetInstallExpressions("packageName")

// 	if err != nil {
// 		t.Errorf("Unexpected duplicate key")
// 	}

// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
// 	expectedOutputValue1 := "add cs policy <<trafficmanagement.contentswitching.policies>>TRUSTED_FULL q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>} <<core.placeholders.csa_trusted_full>>"

// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
// 		}
// 	} else {
// 		for k := range output {
// 			fmt.Println(k)
// 		}
// 		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
// 	}
// }

// // Duplicate key between sections
// func TestModule_GetInstallExpressions2(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 		Sections: []Section{
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = m.GetInstallExpressions("packageName",)

// 	if err == nil {
// 		t.Errorf("Expected duplicate key.")
// 	}

// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
// 	expectedOutputValue1 := "add cs policy <<trafficmanagement.contentswitching.policies>>TRUSTED_FULL q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>} <<core.placeholders.csa_trusted_full>>"

// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
// 		}
// 	} else {
// 		for k := range output {
// 			fmt.Println(k)
// 		}
// 		t.Errorf("Output key does not exist, expected: %s", expectedOutputKey1)
// 	}
// }

// // Duplicate key in section
// func TestModule_GetInstallExpressions3(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 		Sections: []Section{
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Name: "trafficmanagement.contentswitching.actions",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var err error
// 	_, err = m.GetInstallExpressions("packageName")

// 	if err == nil {
// 		t.Errorf("Expected duplicate key.")
// 	}
// }

// // Correct Module definition
// func TestModule_GetUninstallExpressions(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 		Sections: []Section{
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = m.GetUninstallExpressions("packageName")

// 	if err != nil {
// 		t.Errorf("Unexpected duplicate key: %s", err)
// 	}

// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
// 	expectedOutputValue1 := "rm cs policy <<trafficmanagement.contentswitching.policies>>TRUSTED_FULL"

// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
// 		}
// 	}
// }

// // Duplicate key between sections
// func TestModule_GetUninstallExpressions2(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 		Sections: []Section{
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var output map[string]string
// 	var err error

// 	output, err = m.GetUninstallExpressions("packageName")

// 	if err == nil {
// 		t.Errorf("Expected duplicate key.")
// 	}

// 	expectedOutputKey1 := "packageName.moduleName.trafficmanagement.contentswitching.policies.TRUSTED_FULL"
// 	expectedOutputValue1 := "rm cs policy <<trafficmanagement.contentswitching.policies>>TRUSTED_FULL"

// 	if _, isMapContainsKey := output[expectedOutputKey1]; isMapContainsKey {
// 		if output[expectedOutputKey1] != expectedOutputValue1 {
// 			t.Errorf("Output string is incorrect, got: %q for key %q, want: %q", output[expectedOutputKey1], expectedOutputKey1, expectedOutputValue1)
// 		}
// 	}
// }

// // Duplicate key in section
// func TestModule_GetUninstallExpressions3(t *testing.T) {
// 	m := Module{
// 		Name: "moduleName",
// 		Sections: []Section{
// 			{
// 				Name: "trafficmanagement.contentswitching.policies",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 					{
// 						Name: "UNTRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>UNTRUSTED_FULL"},
// 							{Id: "expression", Data: "q<<{core.appexpert.expressions.contentswitching.policies.untrusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_untrusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Name: "trafficmanagement.contentswitching.actions",
// 				Elements: []Element{
// 					{
// 						Name: "TRUSTED_FULL",
// 						Fields: []Field{
// 							{Id: "name", Data: "<<prefix>>TRUSTED_FULL"},
// 							{Id: "expression", Data: "q{<<core.appexpert.expressions.contentswitching.policies.trusted_full/name>>}"},
// 							{Id: "action", Data: "<<core.placeholders.csa_trusted_full>>"},
// 						},
// 						Expressions: Expression{
// 							Install:   "add cs policy <<name>> <<expression>> <<action>>",
// 							Uninstall: "rm cs policy <<name>>",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var err error
// 	_, err = m.GetUninstallExpressions("packageName")

// 	if err == nil {
// 		t.Errorf("Expected duplicate key.")
// 	}

// }
