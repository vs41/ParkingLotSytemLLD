package model

import (
	"testing"
	"time"
)

func TestTicketCanOnlyBeCompletedOnce(t *testing.T) {
	parkedAt := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	ticket := NewTicket("T1", "V1", "S1", "F1", "L1", parkedAt, Car)
	exitedAt := parkedAt.Add(time.Hour)
	if err := ticket.Complete(exitedAt); err != nil {
		t.Fatalf("Complete() error = %v", err)
	}
	if ticket.GetStatus() != TicketInactive {
		t.Fatalf("ticket status = %q, want %q", ticket.GetStatus(), TicketInactive)
	}
	if err := ticket.Complete(exitedAt); err == nil {
		t.Fatal("completed ticket must reject a second completion")
	}
}
