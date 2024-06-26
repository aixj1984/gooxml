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

type CT_ComplementTransform struct{}

func NewCT_ComplementTransform() *CT_ComplementTransform {
	ret := &CT_ComplementTransform{}
	return ret
}

func (m *CT_ComplementTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ComplementTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ComplementTransform: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ComplementTransform and its children
func (m *CT_ComplementTransform) Validate() error {
	return m.ValidateWithPath("CT_ComplementTransform")
}

// ValidateWithPath validates the CT_ComplementTransform and its children, prefixing error messages with path
func (m *CT_ComplementTransform) ValidateWithPath(path string) error {
	return nil
}
