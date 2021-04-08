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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7ParameterWriteVarRequest struct {
	Items  []*S7VarRequestParameterItem
	Parent *S7Parameter
}

// The corresponding interface
type IS7ParameterWriteVarRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7ParameterWriteVarRequest) ParameterType() uint8 {
	return 0x05
}

func (m *S7ParameterWriteVarRequest) MessageType() uint8 {
	return 0x01
}

func (m *S7ParameterWriteVarRequest) InitializeParent(parent *S7Parameter) {
}

func NewS7ParameterWriteVarRequest(items []*S7VarRequestParameterItem) *S7Parameter {
	child := &S7ParameterWriteVarRequest{
		Items:  items,
		Parent: NewS7Parameter(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7ParameterWriteVarRequest(structType interface{}) *S7ParameterWriteVarRequest {
	castFunc := func(typ interface{}) *S7ParameterWriteVarRequest {
		if casted, ok := typ.(S7ParameterWriteVarRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*S7ParameterWriteVarRequest); ok {
			return casted
		}
		if casted, ok := typ.(S7Parameter); ok {
			return CastS7ParameterWriteVarRequest(casted.Child)
		}
		if casted, ok := typ.(*S7Parameter); ok {
			return CastS7ParameterWriteVarRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7ParameterWriteVarRequest) GetTypeName() string {
	return "S7ParameterWriteVarRequest"
}

func (m *S7ParameterWriteVarRequest) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}

func (m *S7ParameterWriteVarRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7ParameterWriteVarRequestParse(io *utils.ReadBuffer) (*S7Parameter, error) {

	// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numItems, _numItemsErr := io.ReadUint8(8)
	_ = numItems
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field")
	}

	// Array field (items)
	// Count array
	items := make([]*S7VarRequestParameterItem, numItems)
	for curItem := uint16(0); curItem < uint16(numItems); curItem++ {
		_item, _err := S7VarRequestParameterItemParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &S7ParameterWriteVarRequest{
		Items:  items,
		Parent: &S7Parameter{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7ParameterWriteVarRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numItems := uint8(uint8(len(m.Items)))
		_numItemsErr := io.WriteUint8(8, (numItems))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		// Array Field (items)
		if m.Items != nil {
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *S7ParameterWriteVarRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "items":
			arrayLoop:
				for {
					token, err = d.Token()
					switch token.(type) {
					case xml.StartElement:
						tok := token.(xml.StartElement)
						var dt *S7VarRequestParameterItem
						if err := d.DecodeElement(&dt, &tok); err != nil {
							return err
						}
						m.Items = append(m.Items, dt)
					default:
						break arrayLoop
					}
				}
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (m *S7ParameterWriteVarRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, arrayElement := range m.Items {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
	}
	return nil
}

func (m S7ParameterWriteVarRequest) String() string {
	return string(m.Box("S7ParameterWriteVarRequest", utils.DefaultWidth*2))
}

func (m S7ParameterWriteVarRequest) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "S7ParameterWriteVarRequest"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Items", m.Items, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
