package main

import (
	"testing"
)

func TestIsBatteryLow(t *testing.T) {
	testCases := []struct {
		name     string
		vehicle  Vehicle
		expected bool
	}{
		{
			name:     "BatteryNotLow-1",
			vehicle:  Vehicle{"V-100", 99.0, 42, 100, "Austin", "ACTIVE"},
			expected: false,
		},
		{
			name:     "BatteryNotLow-2",
			vehicle:  Vehicle{"V-101", 58.2, 76, 70, "Lexington", "CHARGING"},
			expected: false,
		},
		{
			name:     "BatteryNotLow-3",
			vehicle:  Vehicle{"V-102", 76.9, 95, 80, "Alabama", "OFFLINE"},
			expected: false,
		},
		{
			name:     "BatteryLow-1",
			vehicle:  Vehicle{"V-103", 28.23, 26, 90, "New York", "ACTIVE"},
			expected: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsBatteryLow(tc.vehicle)
			if result != tc.expected {
				t.Fatalf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestIsOverheating(t *testing.T) {}

func TestGetBatteryStatus(t *testing.T) {}

func TestVehicleStatus(t *testing.T) {}
