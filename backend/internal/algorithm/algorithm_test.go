package algorithm

import (
	"testing"
)

func TestAddBalance(t *testing.T) {
	a := NewAlgorithm(100)
	a.AddBalance(50)

	if a.Balance != 150 {
		t.Errorf("Expected balance 150, got %.2f", a.Balance)
	}

	logs := a.GetLogs()
	if len(logs) != 1 || logs[0] != "Added 50.00 to balance" {
		t.Errorf("Unexpected logs: %v", logs)
	}
}

func TestRemoveBalance(t *testing.T) {
	a := NewAlgorithm(100)
	err := a.RemoveBalance(30)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if a.Balance != 70 {
		t.Errorf("Expected balance 70, got %.2f", a.Balance)
	}

	logs := a.GetLogs()
	if len(logs) != 1 || logs[0] != "Removed 30.00 from balance" {
		t.Errorf("Unexpected logs: %v", logs)
	}
}

func TestRemoveBalanceInsufficient(t *testing.T) {
	a := NewAlgorithm(50)
	err := a.RemoveBalance(100)
	if err == nil {
		t.Errorf("Expected error for insufficient balance")
	}

	if a.Balance != 50 {
		t.Errorf("Balance should remain 50, got %.2f", a.Balance)
	}

	logs := a.GetLogs()
	if len(logs) != 1 || logs[0] != "insufficient balance to remove 100.00" {
		t.Errorf("Unexpected logs: %v", logs)
	}
}

func TestAddLog(t *testing.T) {
	a := NewAlgorithm(0)
	a.AddLog("Custom log")
	logs := a.GetLogs()
	if len(logs) != 1 || logs[0] != "Custom log" {
		t.Errorf("Unexpected logs: %v", logs)
	}
}
