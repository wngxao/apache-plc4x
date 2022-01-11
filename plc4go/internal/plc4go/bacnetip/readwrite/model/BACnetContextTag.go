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

// The data-structure of this message
type BACnetContextTag struct {
	TagNumber       uint8
	TagClass        TagClass
	LengthValueType uint8
	ExtTagNumber    *uint8
	ExtLength       *uint8
	ExtExtLength    *uint16
	ExtExtExtLength *uint32
	ActualTagNumber uint8
	ActualLength    uint32
	Child           IBACnetContextTagChild
}

// The corresponding interface
type IBACnetContextTag interface {
	DataType() BACnetDataType
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetContextTagParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetContextTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetContextTagChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetContextTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32)
	GetTypeName() string
	IBACnetContextTag
}

func NewBACnetContextTag(tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) *BACnetContextTag {
	return &BACnetContextTag{TagNumber: tagNumber, TagClass: tagClass, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength, ExtExtLength: extExtLength, ExtExtExtLength: extExtExtLength, ActualTagNumber: actualTagNumber, ActualLength: actualLength}
}

func CastBACnetContextTag(structType interface{}) *BACnetContextTag {
	castFunc := func(typ interface{}) *BACnetContextTag {
		if casted, ok := typ.(BACnetContextTag); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTag) GetTypeName() string {
	return "BACnetContextTag"
}

func (m *BACnetContextTag) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTag) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetContextTag) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	// Optional Field (extExtLength)
	if m.ExtExtLength != nil {
		lengthInBits += 16
	}

	// Optional Field (extExtExtLength)
	if m.ExtExtExtLength != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTag) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTag"); pullErr != nil {
		return nil, pullErr
	}

	// Assert Field (tagNumber) (Can be skipped, if a given expression evaluates to false)
	tagNumber, _err := readBuffer.ReadUint8("tagNumber", 4)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'tagNumber' field")
	}
	if tagNumber != tagNumberArgument {
		return nil, utils.ParseAssertError
	}

	// Assert Field (tagClass) (Can be skipped, if a given expression evaluates to false)
	tagClass, _err := TagClassParse(readBuffer)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'tagClass' field")
	}
	if tagClass != TagClass_CONTEXT_SPECIFIC_TAGS {
		return nil, utils.ParseAssertError
	}

	// Simple Field (lengthValueType)
	_lengthValueType, _lengthValueTypeErr := readBuffer.ReadUint8("lengthValueType", 3)
	if _lengthValueTypeErr != nil {
		return nil, errors.Wrap(_lengthValueTypeErr, "Error parsing 'lengthValueType' field")
	}
	lengthValueType := _lengthValueType

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((tagNumber) == (15)) {
		_val, _err := readBuffer.ReadUint8("extTagNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extTagNumber' field")
		}
		extTagNumber = &_val
	}

	// Virtual field
	_actualTagNumber := utils.InlineIf(bool((tagNumber) < (15)), func() interface{} { return uint8(tagNumber) }, func() interface{} { return uint8((*extTagNumber)) }).(uint8)
	actualTagNumber := uint8(_actualTagNumber)

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool((lengthValueType) == (5)) {
		_val, _err := readBuffer.ReadUint8("extLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extLength' field")
		}
		extLength = &_val
	}

	// Optional Field (extExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtLength *uint16 = nil
	if bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (254))) {
		_val, _err := readBuffer.ReadUint16("extExtLength", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtLength' field")
		}
		extExtLength = &_val
	}

	// Optional Field (extExtExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtExtLength *uint32 = nil
	if bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (255))) {
		_val, _err := readBuffer.ReadUint32("extExtExtLength", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtExtLength' field")
		}
		extExtExtLength = &_val
	}

	// Virtual field
	_actualLength := utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (255))), func() interface{} { return uint32((*extExtExtLength)) }, func() interface{} {
		return uint32(uint32(utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (254))), func() interface{} { return uint32((*extExtLength)) }, func() interface{} {
			return uint32(uint32(utils.InlineIf(bool((lengthValueType) == (5)), func() interface{} { return uint32((*extLength)) }, func() interface{} { return uint32(lengthValueType) }).(uint32)))
		}).(uint32)))
	}).(uint32)
	actualLength := uint32(_actualLength)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetContextTag
	var typeSwitchError error
	switch {
	case dataType == BACnetDataType_BOOLEAN: // BACnetContextTagBoolean
		_parent, typeSwitchError = BACnetContextTagBooleanParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_UNSIGNED_INTEGER: // BACnetContextTagUnsignedInteger
		_parent, typeSwitchError = BACnetContextTagUnsignedIntegerParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_SIGNED_INTEGER: // BACnetContextTagSignedInteger
		_parent, typeSwitchError = BACnetContextTagSignedIntegerParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_REAL: // BACnetContextTagReal
		_parent, typeSwitchError = BACnetContextTagRealParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_DOUBLE: // BACnetContextTagDouble
		_parent, typeSwitchError = BACnetContextTagDoubleParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_OCTET_STRING: // BACnetContextTagOctetString
		_parent, typeSwitchError = BACnetContextTagOctetStringParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_CHARACTER_STRING: // BACnetContextTagCharacterString
		_parent, typeSwitchError = BACnetContextTagCharacterStringParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_BIT_STRING: // BACnetContextTagBitString
		_parent, typeSwitchError = BACnetContextTagBitStringParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_ENUMERATED: // BACnetContextTagEnumerated
		_parent, typeSwitchError = BACnetContextTagEnumeratedParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_DATE: // BACnetContextTagDate
		_parent, typeSwitchError = BACnetContextTagDateParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_TIME: // BACnetContextTagTime
		_parent, typeSwitchError = BACnetContextTagTimeParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_BACNET_OBJECT_IDENTIFIER: // BACnetContextTagObjectIdentifier
		_parent, typeSwitchError = BACnetContextTagObjectIdentifierParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_BACNET_PROPERTY_IDENTIFIER: // BACnetContextTagPropertyIdentifier
		_parent, typeSwitchError = BACnetContextTagPropertyIdentifierParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_EVENT_TYPE: // BACnetContextTagEventType
		_parent, typeSwitchError = BACnetContextTagEventTypeParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_EVENT_STATE: // BACnetContextTagEventState
		_parent, typeSwitchError = BACnetContextTagEventStateParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_NOTIFY_TYPE: // BACnetContextTagNotifyType
		_parent, typeSwitchError = BACnetContextTagNotifyTypeParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_BACNET_DEVICE_STATE: // BACnetContextTagDeviceState
		_parent, typeSwitchError = BACnetContextTagDeviceStateParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_OPENING_TAG: // BACnetOpeningTag
		_parent, typeSwitchError = BACnetOpeningTagParse(readBuffer, tagNumberArgument, dataType, lengthValueType)
	case dataType == BACnetDataType_CLOSING_TAG: // BACnetClosingTag
		_parent, typeSwitchError = BACnetClosingTagParse(readBuffer, tagNumberArgument, dataType, lengthValueType)
	case true: // BACnetContextTagEmpty
		_parent, typeSwitchError = BACnetContextTagEmptyParse(readBuffer, tagNumberArgument, dataType)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTag"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength, actualTagNumber, actualLength)
	return _parent, nil
}

func (m *BACnetContextTag) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetContextTag) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetContextTag, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetContextTag"); pushErr != nil {
		return pushErr
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.LengthValueType)
	_lengthValueTypeErr := writeBuffer.WriteUint8("lengthValueType", 3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.Wrap(_lengthValueTypeErr, "Error serializing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.ExtTagNumber != nil {
		extTagNumber = m.ExtTagNumber
		_extTagNumberErr := writeBuffer.WriteUint8("extTagNumber", 8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.Wrap(_extTagNumberErr, "Error serializing 'extTagNumber' field")
		}
	}
	// Virtual field
	if _actualTagNumberErr := writeBuffer.WriteVirtual("actualTagNumber", m.ActualTagNumber); _actualTagNumberErr != nil {
		return errors.Wrap(_actualTagNumberErr, "Error serializing 'actualTagNumber' field")
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.ExtLength != nil {
		extLength = m.ExtLength
		_extLengthErr := writeBuffer.WriteUint8("extLength", 8, *(extLength))
		if _extLengthErr != nil {
			return errors.Wrap(_extLengthErr, "Error serializing 'extLength' field")
		}
	}

	// Optional Field (extExtLength) (Can be skipped, if the value is null)
	var extExtLength *uint16 = nil
	if m.ExtExtLength != nil {
		extExtLength = m.ExtExtLength
		_extExtLengthErr := writeBuffer.WriteUint16("extExtLength", 16, *(extExtLength))
		if _extExtLengthErr != nil {
			return errors.Wrap(_extExtLengthErr, "Error serializing 'extExtLength' field")
		}
	}

	// Optional Field (extExtExtLength) (Can be skipped, if the value is null)
	var extExtExtLength *uint32 = nil
	if m.ExtExtExtLength != nil {
		extExtExtLength = m.ExtExtExtLength
		_extExtExtLengthErr := writeBuffer.WriteUint32("extExtExtLength", 32, *(extExtExtLength))
		if _extExtExtLengthErr != nil {
			return errors.Wrap(_extExtExtLengthErr, "Error serializing 'extExtExtLength' field")
		}
	}
	// Virtual field
	if _actualLengthErr := writeBuffer.WriteVirtual("actualLength", m.ActualLength); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetContextTag"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetContextTag) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
