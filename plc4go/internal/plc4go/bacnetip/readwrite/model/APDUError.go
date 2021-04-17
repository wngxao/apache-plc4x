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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type APDUError struct {
	OriginalInvokeId uint8
	Error            *BACnetError
	Parent           *APDU
}

// The corresponding interface
type IAPDUError interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUError) ApduType() uint8 {
	return 0x5
}

func (m *APDUError) InitializeParent(parent *APDU) {
}

func NewAPDUError(originalInvokeId uint8, error *BACnetError) *APDU {
	child := &APDUError{
		OriginalInvokeId: originalInvokeId,
		Error:            error,
		Parent:           NewAPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAPDUError(structType interface{}) *APDUError {
	castFunc := func(typ interface{}) *APDUError {
		if casted, ok := typ.(APDUError); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUError); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUError(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUError(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUError) GetTypeName() string {
	return "APDUError"
}

func (m *APDUError) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *APDUError) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (error)
	lengthInBits += m.Error.LengthInBits()

	return lengthInBits
}

func (m *APDUError) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUErrorParse(io *utils.ReadBuffer) (*APDU, error) {

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8(4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	originalInvokeId, _originalInvokeIdErr := io.ReadUint8(8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}

	// Simple Field (error)
	error, _errorErr := BACnetErrorParse(io)
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field")
	}

	// Create a partially initialized instance
	_child := &APDUError{
		OriginalInvokeId: originalInvokeId,
		Error:            error,
		Parent:           &APDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *APDUError) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Reserved Field (reserved)
		{
			_err := io.WriteUint8(4, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := io.WriteUint8(8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (error)
		_errorErr := m.Error.Serialize(io)
		if _errorErr != nil {
			return errors.Wrap(_errorErr, "Error serializing 'error' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *APDUError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "originalInvokeId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.OriginalInvokeId = data
			case "error":
				var dt *BACnetError
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Error = dt
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *APDUError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.OriginalInvokeId, xml.StartElement{Name: xml.Name{Local: "originalInvokeId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Error, xml.StartElement{Name: xml.Name{Local: "error"}}); err != nil {
		return err
	}
	return nil
}

func (m APDUError) String() string {
	return string(m.Box("", 120))
}

func (m APDUError) Box(name string, width int) utils.AsciiBox {
	boxName := "APDUError"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		boxes = append(boxes, utils.BoxAnything("OriginalInvokeId", m.OriginalInvokeId, width-2))
		boxes = append(boxes, utils.BoxAnything("Error", m.Error, width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
