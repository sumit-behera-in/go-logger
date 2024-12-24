package goLogger

import (
	"testing"
)

func TestLoadTimeZone(t *testing.T) {
	tests := []struct {
		timeZoneCode string
	}{
		{"ACST"}, {"AEST"}, {"AKST"}, {"AST"}, {"AWST"}, {"BST"},
		{"CCT"}, {"CDT"}, {"CET"}, {"CST"}, {"EAT"}, {"EDT"},
		{"EET"}, {"EST"}, {"GMT"}, {"HKT"}, {"HST"}, {"IST"},
		{"JST"}, {"KST"}, {"MDT"}, {"MSK"}, {"MST"}, {"NZST"},
		{"PDT"}, {"PST"}, {"SAST"}, {"SGT"}, {"UTC"}, {"WAT"},
	}

	for _, tt := range tests {
		location, _ := loadTimeZone(tt.timeZoneCode)
		if location == nil {
			t.Errorf("Expected valid location for timezone %s, got nil", tt.timeZoneCode)
		}

	}
}
