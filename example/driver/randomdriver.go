// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a implementation of a ProtocolDriver interface.
//
package driver

import (
	"fmt"
	"sync"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

type RandomDriver struct {
	lc            logger.LoggingClient
	asyncCh       chan<- *dsModels.AsyncValues
	randomDevices sync.Map
}

func (d *RandomDriver) DisconnectDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Info(fmt.Sprintf("RandomDriver.DisconnectDevice: device-random driver is disconnecting to %s", deviceName))
	return nil
}

func (d *RandomDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues) error {
	d.lc = lc
	d.asyncCh = asyncCh
	return nil
}

func (d *RandomDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	rd := d.retrieveRandomDevice(deviceName)
	fmt.Println("hello HandleReadCommands!!!!!!!!!!!!!!!!")
	res = make([]*dsModels.CommandValue, len(reqs))
	now := time.Now().UnixNano()

	for i, req := range reqs {
		fmt.Println(req)
		fmt.Println(req.Type)
		t := req.Type
		v, err := rd.value(t)
		if err != nil {
			return nil, err
		}
		var cv *dsModels.CommandValue
		switch t {
		case dsModels.Int8:
			cv, _ = dsModels.NewInt8Value(req.DeviceResourceName, now, int8(v))
		case dsModels.Int16:
			cv, _ = dsModels.NewInt16Value(req.DeviceResourceName, now, int16(v))
		case dsModels.Int32:
			cv, _ = dsModels.NewInt32Value(req.DeviceResourceName, now, int32(v))
		}
		res[i] = cv
	}

	return res, nil
}

func (d *RandomDriver) retrieveRandomDevice(deviceName string) (rdv *randomDevice) {
	rd, ok := d.randomDevices.LoadOrStore(deviceName, newRandomDevice())
	if rdv, ok = rd.(*randomDevice); !ok {
		panic("The value in randomDevices has to be a reference of randomDevice")
	}
	return rdv
}

func (d *RandomDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {
	rd := d.retrieveRandomDevice(deviceName)

	for _, param := range params {
		switch param.DeviceResourceName {
		case "Min_Int8":
			v, err := param.Int8Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v < defMinInt8 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: minimum value %d of %T must be int between %d ~ %d", v, v, defMinInt8, defMaxInt8)
			}

			rd.minInt8 = int64(v)
		case "Max_Int8":
			v, err := param.Int8Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v > defMaxInt8 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: maximum value %d of %T must be int between %d ~ %d", v, v, defMinInt8, defMaxInt8)
			}

			rd.maxInt8 = int64(v)
		case "Min_Int16":
			v, err := param.Int16Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v < defMinInt16 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: minimum value %d of %T must be int between %d ~ %d", v, v, defMinInt16, defMaxInt16)
			}

			rd.minInt16 = int64(v)
		case "Max_Int16":
			v, err := param.Int16Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v > defMaxInt16 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: maximum value %d of %T must be int between %d ~ %d", v, v, defMinInt16, defMaxInt16)
			}

			rd.maxInt16 = int64(v)
		case "Min_Int32":
			v, err := param.Int32Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v < defMinInt32 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: minimum value %d of %T must be int between %d ~ %d", v, v, defMinInt32, defMaxInt32)
			}

			rd.minInt32 = int64(v)
		case "Max_Int32":
			v, err := param.Int32Value()
			if err != nil {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
			}
			if v > defMaxInt32 {
				return fmt.Errorf("RandomDriver.HandleWriteCommands: maximum value %d of %T must be int between %d ~ %d", v, v, defMinInt32, defMaxInt32)
			}

			rd.maxInt32 = int64(v)
		default:
			return fmt.Errorf("RandomDriver.HandleWriteCommands: there is no matched device resource for %s", param.String())
		}
	}

	return nil
}

func (d *RandomDriver) Stop(force bool) error {
	d.lc.Info("RandomDriver.Stop: device-random driver is stopping...")
	return nil
}
