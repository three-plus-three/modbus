// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package modbus

import (
	"io"
	"net"
	"time"
)

const (
	// Default timeout
	serialOverTCPTimeout     = 5 * time.Second
	serialOverTCPIdleTimeout = 60 * time.Second
)

type SerialOverTCPConfig struct {
	Address string
	Timeout time.Duration
}

func (cfg *SerialOverTCPConfig) Dail() (io.ReadWriteCloser, error) {
	dialer := net.Dialer{Timeout: cfg.Timeout}
	conn, err := dialer.Dial("tcp", cfg.Address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
