// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package tstun

import (
	"time"

	"github.com/amnezia-vpn/amnezia-wg/tun"
	"tailscale.com/types/logger"
)

// Dummy implementation that does nothing.
func waitInterfaceUp(iface tun.Device, timeout time.Duration, logf logger.Logf) error {
	return nil
}
