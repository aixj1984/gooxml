// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_CellProtection struct {
	// Cell Locked
	LockedAttr *bool
	// Hidden Cell
	HiddenAttr *bool
}

func NewCT_CellProtection() *CT_CellProtection {
	ret := &CT_CellProtection{}
	return ret
}

func (m *CT_CellProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LockedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "locked"},
			Value: fmt.Sprintf("%d", b2i(*m.LockedAttr)),
		})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenAttr)),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellProtection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "locked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = &parsed
			continue
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CellProtection: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CellProtection and its children
func (m *CT_CellProtection) Validate() error {
	return m.ValidateWithPath("CT_CellProtection")
}

// ValidateWithPath validates the CT_CellProtection and its children, prefixing error messages with path
func (m *CT_CellProtection) ValidateWithPath(path string) error {
	return nil
}
