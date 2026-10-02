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

func TestIsOverheating(t *testing.T) {
	testCases := []struct {
		name     string
		vehicle  Vehicle
		expected bool
	}{
		{
			name:     "IsNotOverheating",
			vehicle:  Vehicle{"V-100", 100.0, 86.0, 90.0, "Dallas", "CHARGING"},
			expected: false,
		},
		{
			name:     "IsOverheating",
			vehicle:  Vehicle{"V-101", 92.0, 75.0, 150, "Arizona", "OFFLINE"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsOverheating(tc.vehicle)
			if result != tc.expected {
				t.Errorf("Expected %v, actual %v", tc.expected, result)
			}
		})

	}

}

func TestGetBatteryStatus(t *testing.T) {}

func TestVehicleStatus(t *testing.T) {}
