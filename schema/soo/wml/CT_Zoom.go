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

type CT_Zoom struct {
	// Zoom Type
	ValAttr ST_Zoom
	// Zoom Percentage
	PercentAttr ST_DecimalNumberOrPercent
}

func NewCT_Zoom() *CT_Zoom {
	ret := &CT_Zoom{}
	return ret
}

func (m *CT_Zoom) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_ZoomUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "w:percent"},
		Value: fmt.Sprintf("%v", m.PercentAttr),
	})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Zoom) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "percent" {
			parsed, err := ParseUnionST_DecimalNumberOrPercent(attr.Value)
			if err != nil {
				return err
			}
			m.PercentAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Zoom: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Zoom and its children
func (m *CT_Zoom) Validate() error {
	return m.ValidateWithPath("CT_Zoom")
}

// ValidateWithPath validates the CT_Zoom and its children, prefixing error messages with path
func (m *CT_Zoom) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	if err := m.PercentAttr.ValidateWithPath(path + "/PercentAttr"); err != nil {
		return err
	}
	return nil
}
