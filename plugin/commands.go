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
	"context"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/temporalio/tctl-kit/pkg/output"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/server/common/collection"
	"go.temporal.io/server/common/rpc"
)

var commands = []*cli.Command{
	{
		Name:  "hello",
		Usage: "Say hello",
		Action: func(c *cli.Context) error {
			return SayHello(c)
		},
	},
	{
		Name:  "set-hello",
		Usage: "Set hello value using config feature",
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
	{
		Name:  "list",
		Usage: "Print namespaces",
		Action: func(c *cli.Context) error {
			return List(c)
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

func List(c *cli.Context) error {
	client := temporalFactory.SDKClient(c, "default")

	paginationFunc := func(npt []byte) ([]interface{}, []byte, error) {
		var items []interface{}
		var err error
		req := &workflowservice.ListNamespacesRequest{
			NextPageToken: npt,
		}
		ctx, cancel := newContext(c)
		defer cancel()
		res, err := client.WorkflowService().ListNamespaces(ctx, req)
		if err != nil {
			return nil, nil, err
		}

		return items, res.GetNextPageToken(), nil
	}

	iter := collection.NewPagingIterator(paginationFunc)
	opts := &output.PrintOptions{
		Fields:     []string{"Execution.WorkflowId", "Execution.RunId", "StartTime"},
		FieldsLong: []string{"Type.Name", "TaskQueue", "ExecutionTime", "CloseTime"},
	}
	return output.Pager(c, iter, opts)
}

// SayHello prints hello list
func SetHelloValue(c *cli.Context) error {
	value := c.String("value")

	if err := pluginConfig.SetByEnvironment(c, helloKey, value); err != nil {
		return fmt.Errorf("unable to set property %s: %s", helloKey, err)
	}

	return nil
}

func newContext(c *cli.Context) (context.Context, context.CancelFunc) {
	timeout := time.Duration(5) * time.Second

	return rpc.NewContextWithTimeoutAndCLIHeaders(timeout)
}
