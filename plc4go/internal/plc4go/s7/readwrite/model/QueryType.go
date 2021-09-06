/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type QueryType uint8

type IQueryType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	QueryType_BYALARMTYPE QueryType = 0x01
	QueryType_ALARM_8     QueryType = 0x02
	QueryType_ALARM_S     QueryType = 0x04
)

var QueryTypeValues []QueryType

func init() {
	_ = errors.New
	QueryTypeValues = []QueryType{
		QueryType_BYALARMTYPE,
		QueryType_ALARM_8,
		QueryType_ALARM_S,
	}
}

func QueryTypeByValue(value uint8) QueryType {
	switch value {
	case 0x01:
		return QueryType_BYALARMTYPE
	case 0x02:
		return QueryType_ALARM_8
	case 0x04:
		return QueryType_ALARM_S
	}
	return 0
}

func QueryTypeByName(value string) QueryType {
	switch value {
	case "BYALARMTYPE":
		return QueryType_BYALARMTYPE
	case "ALARM_8":
		return QueryType_ALARM_8
	case "ALARM_S":
		return QueryType_ALARM_S
	}
	return 0
}

func CastQueryType(structType interface{}) QueryType {
	castFunc := func(typ interface{}) QueryType {
		if sQueryType, ok := typ.(QueryType); ok {
			return sQueryType
		}
		return 0
	}
	return castFunc(structType)
}

func (m QueryType) LengthInBits() uint16 {
	return 8
}

func (m QueryType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func QueryTypeParse(readBuffer utils.ReadBuffer) (QueryType, error) {
	val, err := readBuffer.ReadUint8("QueryType", 8)
	if err != nil {
		return 0, nil
	}
	return QueryTypeByValue(val), nil
}

func (e QueryType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("QueryType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e QueryType) name() string {
	switch e {
	case QueryType_BYALARMTYPE:
		return "BYALARMTYPE"
	case QueryType_ALARM_8:
		return "ALARM_8"
	case QueryType_ALARM_S:
		return "ALARM_S"
	}
	return ""
}

func (e QueryType) String() string {
	return e.name()
}