// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_SDCategory struct {
	TypeAttr string
	PriAttr  uint32
}

func NewCT_SDCategory() *CT_SDCategory {
	ret := &CT_SDCategory{}
	return ret
}

func (m *CT_SDCategory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "type"},
		Value: fmt.Sprintf("%v", m.TypeAttr),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "pri"},
		Value: fmt.Sprintf("%v", m.PriAttr),
	})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SDCategory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = parsed
			continue
		}
		if attr.Name.Local == "pri" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.PriAttr = uint32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SDCategory: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SDCategory and its children
func (m *CT_SDCategory) Validate() error {
	return m.ValidateWithPath("CT_SDCategory")
}

// ValidateWithPath validates the CT_SDCategory and its children, prefixing error messages with path
func (m *CT_SDCategory) ValidateWithPath(path string) error {
	return nil
}
