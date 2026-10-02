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

func TestGetBatteryStatus(t *testing.T) {
	testCases := []struct {
		name     string
		vehicle  Vehicle
		expected string
	}{
		{
			name:     "NormalBatteryStatus",
			vehicle:  Vehicle{"V-100", 99.0, 42, 100, "Austin", "ACTIVE"},
			expected: "NORMAL",
		},
		{
			name:     "LowBatteryStatus",
			vehicle:  Vehicle{"V-103", 28.23, 26, 90, "New York", "ACTIVE"},
			expected: "LOW_BATTERY",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetBatteryStatus(tc.vehicle)
			if tc.expected != result {
				t.Errorf("Expected %v, actual %v", tc.expected, result)
			}
		})
	}
}

func TestVehicleStatus(t *testing.T) {
	testCases := []struct {
		name     string
		vehicle  Vehicle
		expected string
	}{
		{
			name:     "LowBatteryStatus",
			vehicle:  Vehicle{"V-103", 28.23, 26, 90, "New York", "ACTIVE"},
			expected: "LOW_BATTERY",
		},
		{
			name:     "OverheatingVehicleStatus",
			vehicle:  Vehicle{"V-104", 80.0, 33, 59, "Oklahoma", "ACTIVE"},
			expected: "OVERHEATING",
		},
		{
			name:     "HealthyVehicleStatus",
			vehicle:  Vehicle{"V-104", 80.0, 33, 59, "Oklahoma", "ACTIVE"},
			expected: "HEALTHY",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := VehicleStatus(tc.vehicle)
			if result != tc.expected {
				t.Errorf("Expected %v, actual %v", tc.expected, result)
			}
		})
	}
}
