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
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type NLM struct {
	VendorId *uint16
	Child    INLMChild
}

// The corresponding interface
type INLM interface {
	MessageType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type INLMParent interface {
	SerializeParent(io utils.WriteBuffer, child INLM, serializeChildFunction func() error) error
	GetTypeName() string
}

type INLMChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *NLM, vendorId *uint16)
	GetTypeName() string
	INLM
}

func NewNLM(vendorId *uint16) *NLM {
	return &NLM{VendorId: vendorId}
}

func CastNLM(structType interface{}) *NLM {
	castFunc := func(typ interface{}) *NLM {
		if casted, ok := typ.(NLM); ok {
			return &casted
		}
		if casted, ok := typ.(*NLM); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *NLM) GetTypeName() string {
	return "NLM"
}

func (m *NLM) LengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 8

	// Optional Field (vendorId)
	if m.VendorId != nil {
		lengthInBits += 16
	}

	// Length of sub-type elements will be added by sub-type...
	lengthInBits += m.Child.LengthInBits()

	return lengthInBits
}

func (m *NLM) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func NLMParse(io *utils.ReadBuffer, apduLength uint16) (*NLM, error) {

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := io.ReadUint8(8)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field")
	}

	// Optional Field (vendorId) (Can be skipped, if a given expression evaluates to false)
	var vendorId *uint16 = nil
	if bool(bool(bool((messageType) >= (128)))) && bool(bool(bool((messageType) <= (255)))) {
		_val, _err := io.ReadUint16(16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'vendorId' field")
		}
		vendorId = &_val
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *NLM
	var typeSwitchError error
	switch {
	case messageType == 0x0: // NLMWhoIsRouterToNetwork
		_parent, typeSwitchError = NLMWhoIsRouterToNetworkParse(io, apduLength, messageType)
	case messageType == 0x1: // NLMIAmRouterToNetwork
		_parent, typeSwitchError = NLMIAmRouterToNetworkParse(io, apduLength, messageType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, vendorId)
	return _parent, nil
}

func (m *NLM) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *NLM) SerializeParent(io utils.WriteBuffer, child INLM, serializeChildFunction func() error) error {

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.MessageType())
	_messageTypeErr := io.WriteUint8(8, (messageType))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Optional Field (vendorId) (Can be skipped, if the value is null)
	var vendorId *uint16 = nil
	if m.VendorId != nil {
		vendorId = m.VendorId
		_vendorIdErr := io.WriteUint16(16, *(vendorId))
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	return nil
}

func (m *NLM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "vendorId":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.VendorId = &data
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of NLM")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.bacnetip.readwrite.NLMWhoIsRouterToNetwork":
					var dt *NLMWhoIsRouterToNetwork
					if m.Child != nil {
						dt = m.Child.(*NLMWhoIsRouterToNetwork)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.NLMIAmRouterToNetwork":
					var dt *NLMIAmRouterToNetwork
					if m.Child != nil {
						dt = m.Child.(*NLMIAmRouterToNetwork)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

func (m *NLM) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.bacnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.VendorId, xml.StartElement{Name: xml.Name{Local: "vendorId"}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m NLM) String() string {
	return string(m.Box("NLM", utils.DefaultWidth*2))
}

func (m NLM) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "NLM"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("VendorId", m.VendorId, width-2))
	boxes = append(boxes, utils.BoxAnything("", m.Child, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
