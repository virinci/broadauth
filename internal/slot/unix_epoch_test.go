package slot_test

import (
	"testing"
	"time"

	"github.com/virinci/broadauth/internal/slot"
)

func TestUnixEpochSlotSource(t *testing.T) {
	slotSource := slot.NewUnixEpochSlotSource(2000)
	slotNum, err := slotSource.GetSlot()
	if err != nil {
		t.Fatalf("failed to get slot: %v", err)
	}

	if slotNum != slot.Slot(time.Now().UnixMilli()/2000) {
		t.Fatalf("slot is not the current unix epoch")
	}
}
