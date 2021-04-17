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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetTagApplicationBoolean struct {
	Parent *BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationBoolean interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTagApplicationBoolean) ContextSpecificTag() uint8 {
	return 0
}

func (m *BACnetTagApplicationBoolean) InitializeParent(parent *BACnetTag, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) {
	m.Parent.TypeOrTagNumber = typeOrTagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
}

func NewBACnetTagApplicationBoolean(typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) *BACnetTag {
	child := &BACnetTagApplicationBoolean{
		Parent: NewBACnetTag(typeOrTagNumber, lengthValueType, extTagNumber, extLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetTagApplicationBoolean(structType interface{}) *BACnetTagApplicationBoolean {
	castFunc := func(typ interface{}) *BACnetTagApplicationBoolean {
		if casted, ok := typ.(BACnetTagApplicationBoolean); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagApplicationBoolean); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTag); ok {
			return CastBACnetTagApplicationBoolean(casted.Child)
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return CastBACnetTagApplicationBoolean(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagApplicationBoolean) GetTypeName() string {
	return "BACnetTagApplicationBoolean"
}

func (m *BACnetTagApplicationBoolean) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTagApplicationBoolean) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetTagApplicationBoolean) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagApplicationBooleanParse(io *utils.ReadBuffer) (*BACnetTag, error) {

	// Create a partially initialized instance
	_child := &BACnetTagApplicationBoolean{
		Parent: &BACnetTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetTagApplicationBoolean) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetTagApplicationBoolean) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BACnetTagApplicationBoolean) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BACnetTagApplicationBoolean) String() string {
	return string(m.Box("", 120))
}

func (m BACnetTagApplicationBoolean) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetTagApplicationBoolean"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
