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

type CT_ExternalDefinedName struct {
	// Defined Name
	NameAttr string
	// Refers To
	RefersToAttr *string
	// Sheet Id
	SheetIdAttr *uint32
}

func NewCT_ExternalDefinedName() *CT_ExternalDefinedName {
	ret := &CT_ExternalDefinedName{}
	return ret
}

func (m *CT_ExternalDefinedName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr),
	})
	if m.RefersToAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "refersTo"},
			Value: fmt.Sprintf("%v", *m.RefersToAttr),
		})
	}
	if m.SheetIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "sheetId"},
			Value: fmt.Sprintf("%v", *m.SheetIdAttr),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExternalDefinedName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "refersTo" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefersToAttr = &parsed
			continue
		}
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SheetIdAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ExternalDefinedName: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ExternalDefinedName and its children
func (m *CT_ExternalDefinedName) Validate() error {
	return m.ValidateWithPath("CT_ExternalDefinedName")
}

// ValidateWithPath validates the CT_ExternalDefinedName and its children, prefixing error messages with path
func (m *CT_ExternalDefinedName) ValidateWithPath(path string) error {
	return nil
}
