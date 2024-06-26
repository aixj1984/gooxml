// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_Vector3D struct {
	DxAttr ST_Coordinate
	DyAttr ST_Coordinate
	DzAttr ST_Coordinate
}

func NewCT_Vector3D() *CT_Vector3D {
	ret := &CT_Vector3D{}
	return ret
}

func (m *CT_Vector3D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "dx"},
		Value: fmt.Sprintf("%v", m.DxAttr),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "dy"},
		Value: fmt.Sprintf("%v", m.DyAttr),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "dz"},
		Value: fmt.Sprintf("%v", m.DzAttr),
	})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Vector3D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dx" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.DxAttr = parsed
			continue
		}
		if attr.Name.Local == "dy" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.DyAttr = parsed
			continue
		}
		if attr.Name.Local == "dz" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.DzAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Vector3D: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Vector3D and its children
func (m *CT_Vector3D) Validate() error {
	return m.ValidateWithPath("CT_Vector3D")
}

// ValidateWithPath validates the CT_Vector3D and its children, prefixing error messages with path
func (m *CT_Vector3D) ValidateWithPath(path string) error {
	if err := m.DxAttr.ValidateWithPath(path + "/DxAttr"); err != nil {
		return err
	}
	if err := m.DyAttr.ValidateWithPath(path + "/DyAttr"); err != nil {
		return err
	}
	if err := m.DzAttr.ValidateWithPath(path + "/DzAttr"); err != nil {
		return err
	}
	return nil
}
