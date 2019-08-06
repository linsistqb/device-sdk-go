// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/edgexfoundry/device-sdk-go"
	"github.com/edgexfoundry/device-sdk-go/example/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "device-random"
)

func main() {
	d := driver.RandomDriver{}
	startup.Bootstrap(serviceName, device.Version, &d)
}
