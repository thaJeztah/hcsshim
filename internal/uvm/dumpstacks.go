//go:build windows

package uvm

import (
	"context"
)

func (uvm *UtilityVM) DumpStacks(ctx context.Context) (string, error) {
	if uvm.gc == nil || !uvm.guestCaps.IsDumpStacksSupported() {
		return "", nil
	}

	return uvm.gc.DumpStacks(ctx)
}
