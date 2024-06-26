// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"
)

type AG_Type struct {
	TypeAttr *string
}

func NewAG_Type() *AG_Type {
	ret := &AG_Type{}
	return ret
}

func (m *AG_Type) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "type"},
			Value: fmt.Sprintf("%v", *m.TypeAttr),
		})
	}
	return nil
}

func (m *AG_Type) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Type: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Type and its children
func (m *AG_Type) Validate() error {
	return m.ValidateWithPath("AG_Type")
}

// ValidateWithPath validates the AG_Type and its children, prefixing error messages with path
func (m *AG_Type) ValidateWithPath(path string) error {
	return nil
}
