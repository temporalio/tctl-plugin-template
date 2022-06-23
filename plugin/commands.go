// The MIT License
//
// Copyright (c) 2021 Temporal Technologies Inc.  All rights reserved.
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

	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	{
		Name:  "hello",
		Usage: "Says hello",
		Action: func(c *cli.Context) error {
			return SayHello(c)
		},
	},
	{
		Name:  "set-hello",
		Usage: "Set hello value",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "value",
				Usage:    "Value to set",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return SetHelloValue(c)
		},
	},
}

// SayHello prints hello list
func SayHello(c *cli.Context) error {
	value, err := pluginConfig.GetByEnvironment(c, helloKey)

	if err != nil {
		return err
	}

	fmt.Printf("Hello %s\n", value)

	return nil
}

// SayHello prints hello list
func SetHelloValue(c *cli.Context) error {
	value := c.String("value")

	if err := pluginConfig.SetByEnvironment(c, helloKey, value); err != nil {
		return fmt.Errorf("unable to set property %s: %s", helloKey, err)
	}

	return nil
}
