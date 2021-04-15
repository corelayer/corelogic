package models

import "testing"


func TestModule_GetPlaceholderMap(t *testing.T) {
	module := Module{
		Name:         "dns",
		Package:      "core",
		Placeholders: []Placeholder{
			{
				Name:       "DUMMY1",
				Expression: "dummy1",
			},
			{
				Name:       "DUMMY2",
				Expression: "dummy2",
			},
		},
		Sections:     nil,
	}

	output := module.GetPlaceholderMap()

	if output["core.dns.DUMMY1"] != "dummy1" {
		t.Errorf("Output string is incorrect, got: %s for key: %s, want: %s", output["DUMMY1"], "DUMMY1", "dummy1")
	}
}

func TestModule_GetFullModuleName(t *testing.T) {
	module := Module{
		Name:         "dns",
		Package:      "core",
		//Dependencies: nil,
		Placeholders: nil,
		Sections:     nil,
	}

	output := module.GetFullModuleName()

	if output != "core.dns" {
		t.Errorf("Output string is incorrect, got: %s, want: %s", output, "core.dns")
	}
}
