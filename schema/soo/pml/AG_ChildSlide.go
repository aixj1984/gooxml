// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type AG_ChildSlide struct {
	ShowMasterSpAttr     *bool
	ShowMasterPhAnimAttr *bool
}

func NewAG_ChildSlide() *AG_ChildSlide {
	ret := &AG_ChildSlide{}
	return ret
}

func (m *AG_ChildSlide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ShowMasterSpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "showMasterSp"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowMasterSpAttr)),
		})
	}
	if m.ShowMasterPhAnimAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "showMasterPhAnim"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowMasterPhAnimAttr)),
		})
	}
	return nil
}

func (m *AG_ChildSlide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "showMasterSp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMasterSpAttr = &parsed
			continue
		}
		if attr.Name.Local == "showMasterPhAnim" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMasterPhAnimAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_ChildSlide: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_ChildSlide and its children
func (m *AG_ChildSlide) Validate() error {
	return m.ValidateWithPath("AG_ChildSlide")
}

// ValidateWithPath validates the AG_ChildSlide and its children, prefixing error messages with path
func (m *AG_ChildSlide) ValidateWithPath(path string) error {
	return nil
}
