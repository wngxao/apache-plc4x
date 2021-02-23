//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type FirmwareType uint16

type IFirmwareType interface {
	Serialize(io utils.WriteBuffer) error
}

const (
	FirmwareType_SYSTEM_1                   FirmwareType = 0x0010
	FirmwareType_SYSTEM_2                   FirmwareType = 0x0020
	FirmwareType_SYSTEM_300                 FirmwareType = 0x0300
	FirmwareType_SYSTEM_7                   FirmwareType = 0x0700
	FirmwareType_SYSTEM_B                   FirmwareType = 0x07B0
	FirmwareType_IR_DECODER                 FirmwareType = 0x0810
	FirmwareType_COUPLER                    FirmwareType = 0x0910
	FirmwareType_NONE                       FirmwareType = 0x0AF0
	FirmwareType_SYSTEM_1_PL110             FirmwareType = 0x10B0
	FirmwareType_SYSTEM_B_PL110             FirmwareType = 0x17B0
	FirmwareType_MEDIA_COUPLER_PL_TP        FirmwareType = 0x1900
	FirmwareType_RF_BI_DIRECTIONAL_DEVICES  FirmwareType = 0x2000
	FirmwareType_RF_UNI_DIRECTIONAL_DEVICES FirmwareType = 0x2100
	FirmwareType_SYSTEM_1_TP0               FirmwareType = 0x3000
	FirmwareType_SYSTEM_1_PL132             FirmwareType = 0x4000
	FirmwareType_SYSTEM_7_KNX_NET_IP        FirmwareType = 0x5700
)

func FirmwareTypeByValue(value uint16) FirmwareType {
	switch value {
	case 0x0010:
		return FirmwareType_SYSTEM_1
	case 0x0020:
		return FirmwareType_SYSTEM_2
	case 0x0300:
		return FirmwareType_SYSTEM_300
	case 0x0700:
		return FirmwareType_SYSTEM_7
	case 0x07B0:
		return FirmwareType_SYSTEM_B
	case 0x0810:
		return FirmwareType_IR_DECODER
	case 0x0910:
		return FirmwareType_COUPLER
	case 0x0AF0:
		return FirmwareType_NONE
	case 0x10B0:
		return FirmwareType_SYSTEM_1_PL110
	case 0x17B0:
		return FirmwareType_SYSTEM_B_PL110
	case 0x1900:
		return FirmwareType_MEDIA_COUPLER_PL_TP
	case 0x2000:
		return FirmwareType_RF_BI_DIRECTIONAL_DEVICES
	case 0x2100:
		return FirmwareType_RF_UNI_DIRECTIONAL_DEVICES
	case 0x3000:
		return FirmwareType_SYSTEM_1_TP0
	case 0x4000:
		return FirmwareType_SYSTEM_1_PL132
	case 0x5700:
		return FirmwareType_SYSTEM_7_KNX_NET_IP
	}
	return 0
}

func FirmwareTypeByName(value string) FirmwareType {
	switch value {
	case "SYSTEM_1":
		return FirmwareType_SYSTEM_1
	case "SYSTEM_2":
		return FirmwareType_SYSTEM_2
	case "SYSTEM_300":
		return FirmwareType_SYSTEM_300
	case "SYSTEM_7":
		return FirmwareType_SYSTEM_7
	case "SYSTEM_B":
		return FirmwareType_SYSTEM_B
	case "IR_DECODER":
		return FirmwareType_IR_DECODER
	case "COUPLER":
		return FirmwareType_COUPLER
	case "NONE":
		return FirmwareType_NONE
	case "SYSTEM_1_PL110":
		return FirmwareType_SYSTEM_1_PL110
	case "SYSTEM_B_PL110":
		return FirmwareType_SYSTEM_B_PL110
	case "MEDIA_COUPLER_PL_TP":
		return FirmwareType_MEDIA_COUPLER_PL_TP
	case "RF_BI_DIRECTIONAL_DEVICES":
		return FirmwareType_RF_BI_DIRECTIONAL_DEVICES
	case "RF_UNI_DIRECTIONAL_DEVICES":
		return FirmwareType_RF_UNI_DIRECTIONAL_DEVICES
	case "SYSTEM_1_TP0":
		return FirmwareType_SYSTEM_1_TP0
	case "SYSTEM_1_PL132":
		return FirmwareType_SYSTEM_1_PL132
	case "SYSTEM_7_KNX_NET_IP":
		return FirmwareType_SYSTEM_7_KNX_NET_IP
	}
	return 0
}

func CastFirmwareType(structType interface{}) FirmwareType {
	castFunc := func(typ interface{}) FirmwareType {
		if sFirmwareType, ok := typ.(FirmwareType); ok {
			return sFirmwareType
		}
		return 0
	}
	return castFunc(structType)
}

func (m FirmwareType) LengthInBits() uint16 {
	return 16
}

func (m FirmwareType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmwareTypeParse(io *utils.ReadBuffer) (FirmwareType, error) {
	val, err := io.ReadUint16(16)
	if err != nil {
		return 0, nil
	}
	return FirmwareTypeByValue(val), nil
}

func (e FirmwareType) Serialize(io utils.WriteBuffer) error {
	err := io.WriteUint16(16, uint16(e))
	return err
}

func (e FirmwareType) String() string {
	switch e {
	case FirmwareType_SYSTEM_1:
		return "SYSTEM_1"
	case FirmwareType_SYSTEM_2:
		return "SYSTEM_2"
	case FirmwareType_SYSTEM_300:
		return "SYSTEM_300"
	case FirmwareType_SYSTEM_7:
		return "SYSTEM_7"
	case FirmwareType_SYSTEM_B:
		return "SYSTEM_B"
	case FirmwareType_IR_DECODER:
		return "IR_DECODER"
	case FirmwareType_COUPLER:
		return "COUPLER"
	case FirmwareType_NONE:
		return "NONE"
	case FirmwareType_SYSTEM_1_PL110:
		return "SYSTEM_1_PL110"
	case FirmwareType_SYSTEM_B_PL110:
		return "SYSTEM_B_PL110"
	case FirmwareType_MEDIA_COUPLER_PL_TP:
		return "MEDIA_COUPLER_PL_TP"
	case FirmwareType_RF_BI_DIRECTIONAL_DEVICES:
		return "RF_BI_DIRECTIONAL_DEVICES"
	case FirmwareType_RF_UNI_DIRECTIONAL_DEVICES:
		return "RF_UNI_DIRECTIONAL_DEVICES"
	case FirmwareType_SYSTEM_1_TP0:
		return "SYSTEM_1_TP0"
	case FirmwareType_SYSTEM_1_PL132:
		return "SYSTEM_1_PL132"
	case FirmwareType_SYSTEM_7_KNX_NET_IP:
		return "SYSTEM_7_KNX_NET_IP"
	}
	return ""
}
