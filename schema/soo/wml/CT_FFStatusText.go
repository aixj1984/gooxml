// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_FFStatusText struct {
	// Status Text Type
	TypeAttr ST_InfoTextType
	// Status Text Value
	ValAttr *string
}

func NewCT_FFStatusText() *CT_FFStatusText {
	ret := &CT_FFStatusText{}
	return ret
}

func (m *CT_FFStatusText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_InfoTextTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "w:type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FFStatusText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FFStatusText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FFStatusText and its children
func (m *CT_FFStatusText) Validate() error {
	return m.ValidateWithPath("CT_FFStatusText")
}

// ValidateWithPath validates the CT_FFStatusText and its children, prefixing error messages with path
func (m *CT_FFStatusText) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
