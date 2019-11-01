package tracker

import (
	"testing"
)

func TestStatsTracker_TrackPasswordsSent_CorrectValue(t *testing.T) {
	expected := uint64(42)

	statsTracker := NewStatsTracker()
	statsTracker.TrackPasswordsSent(expected)

	actual := statsTracker.passwordsSent
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestStatsTracker_TrackHashesComputed_CorrectValue(t *testing.T) {
	expected := uint64(42)

	statsTracker := NewStatsTracker()
	statsTracker.TrackHashesComputed(expected)

	actual := statsTracker.hashesComputed
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}


func TestStatsTracker_TrackHashesCracked_CorrectValue(t *testing.T) {
	expected := uint64(42)

	statsTracker := NewStatsTracker()
	statsTracker.TrackHashesCracked(expected)

	actual := statsTracker.hashesCracked
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestStatsTracker_TrackHashMatchAttempt_CorrectValue(t *testing.T) {
	expected := uint64(1)

	statsTracker := NewStatsTracker()
	statsTracker.TrackHashMatchAttempt()

	actual := statsTracker.hashesCracked
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestStatsTracker_GetPasswordsSent_CorrectValue(t *testing.T) {
	expected := uint64(42)

	statsTracker := StatsTracker{
		passwordsSent:  expected,
		hashesComputed: 0,
		hashesCracked:  0,
	}

	actual := statsTracker.GetPasswordsSent()
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestStatsTracker_GetHashesComputed_CorrectValue(t *testing.T) {
	expected := uint64(42)

	statsTracker := StatsTracker{
		passwordsSent:  0,
		hashesComputed: expected,
		hashesCracked:  0,
	}

	actual := statsTracker.GetHashesComputed()
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestStatsTracker_GetHashesCracked_CorrectValue(t *testing.T) {
	expected := uint64(42)

	statsTracker := StatsTracker{
		passwordsSent:  0,
		hashesComputed: 0,
		hashesCracked:  expected,
	}

	actual := statsTracker.GetHashesCracked()
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestStatsTracker_GetHashMatchAttempts_CorrectValue(t *testing.T) {
	expected := uint64(42)

	statsTracker := StatsTracker{
		passwordsSent:  0,
		hashesComputed: 0,
		hashesCracked:  expected,
	}

	actual := statsTracker.GetHashMatchAttempts()
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}
