// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package plugin

import (
	"fmt"

	"github.com/temporalio/tctl-kit/pkg/color"
	"github.com/urfave/cli/v2"
)

// Flags used to specify cli command line arguments
var (
	FlagAddress                    = "address"
	FlagTLSCertPath                = "tls-cert-path"
	FlagTLSKeyPath                 = "tls-key-path"
	FlagTLSCaPath                  = "tls-ca-path"
	FlagTLSDisableHostVerification = "tls-disable-host-verification"
	FlagTLSServerName              = "tls-server-name"
)

var globalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  FlagAddress,
		Value: "",
		Usage: "host:port for Temporal frontend service",
	},
	&cli.StringFlag{
		Name:  FlagTLSCertPath,
		Value: "",
		Usage: "Path to x509 certificate",
	},
	&cli.StringFlag{
		Name:  FlagTLSKeyPath,
		Value: "",
		Usage: "Path to private key",
	},
	&cli.StringFlag{
		Name:  FlagTLSCaPath,
		Value: "",
		Usage: "Path to server CA certificate",
	},
	&cli.BoolFlag{
		Name:  FlagTLSDisableHostVerification,
		Usage: "Disable tls host name verification (tls must be enabled)",
	},
	&cli.StringFlag{
		Name:  FlagTLSServerName,
		Value: "",
		Usage: "Override for target server name",
	},
	&cli.StringFlag{
		Name:  color.FlagColor,
		Usage: fmt.Sprintf("when to use color: %v, %v, %v.", color.Auto, color.Always, color.Never),
		Value: string(color.Auto),
	},
}
