package models

import "testing"

func TestRelease_GetVersionAsString(t *testing.T) {
	release := Release{
		Major: 10,
		Minor: 1,
	}

	result := release.GetVersionAsString()
	if result != "CL10_01" {
		t.Errorf("Output string is incorrect, got: %s, want: %s.", result, "10_01")
	}
}

func TestFramework_GetPrefixMap(t *testing.T) {
	framework := Framework{
		Prefixes: []Prefix{
			Prefix{Section: "AppExpert.Stringmaps", Prefix: "PSM"},
		},
	}

	prefixMap := framework.GetPrefixMap()

	if prefixMap["AppExpert.Stringmaps"] != "PSM" {
		t.Errorf("Output string is incorrect, got: %s for key: %s, want: %s", prefixMap["AppExpert.Stringmaps"], "AppExpert.Stringmaps", "PSM")
	}
}


func TestFramework_GetPrefixWithVersion(t *testing.T) {
	framework := Framework{
		Release: Release{
			Major: 10,
			Minor: 2,
		},
		Prefixes: []Prefix{
			Prefix{Section: "AppExpert.Stringmaps", Prefix: "PSM"},
		},
	}

	result := framework.GetPrefixWithVersion("AppExpert.Stringmaps")

	if result != "PSM_CL10_02" {
		t.Errorf("Output string is incorrect, got: %s, want: %s", result, "PSM_CL10_02")
	}
}