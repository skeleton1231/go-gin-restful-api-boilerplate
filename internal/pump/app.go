// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package pump does all the work necessary to create an iam pump server.
package pump

import (
	genericapiserver "github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pkg/server"

	"github.com/marmotedu/log"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pump/config"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pump/options"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/pkg/app"
)

const commandDesc = `IAM Pump is a pluggable analytics purger to move Analytics generated by your iam nodes to any back-end.

Find more iam-pump information at:
    https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iam-pump.md`

// NewApp creates an App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("IAM analytics server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		stopCh := genericapiserver.SetupSignalHandler()

		return Run(cfg, stopCh)
	}
}
