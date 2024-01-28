// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux

package tstun

import "github.com/amnezia-vpn/amnezia-wg/tun"

func setLinkAttrs(iface tun.Device) error {
	return nil
}
