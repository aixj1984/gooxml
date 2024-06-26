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

type CT_IconFilter struct {
	// Icon Set
	IconSetAttr ST_IconSetType
	// Icon Id
	IconIdAttr *uint32
}

func NewCT_IconFilter() *CT_IconFilter {
	ret := &CT_IconFilter{}
	ret.IconSetAttr = ST_IconSetType(1)
	return ret
}

func (m *CT_IconFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.IconSetAttr.MarshalXMLAttr(xml.Name{Local: "iconSet"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.IconIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "iconId"},
			Value: fmt.Sprintf("%v", *m.IconIdAttr),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_IconFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.IconSetAttr = ST_IconSetType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "iconSet" {
			m.IconSetAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "iconId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IconIdAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_IconFilter: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_IconFilter and its children
func (m *CT_IconFilter) Validate() error {
	return m.ValidateWithPath("CT_IconFilter")
}

// ValidateWithPath validates the CT_IconFilter and its children, prefixing error messages with path
func (m *CT_IconFilter) ValidateWithPath(path string) error {
	if m.IconSetAttr == ST_IconSetTypeUnset {
		return fmt.Errorf("%s/IconSetAttr is a mandatory field", path)
	}
	if err := m.IconSetAttr.ValidateWithPath(path + "/IconSetAttr"); err != nil {
		return err
	}
	return nil
}
