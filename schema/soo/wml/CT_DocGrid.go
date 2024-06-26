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
	"strconv"
)

type CT_DocGrid struct {
	// Document Grid Type
	TypeAttr ST_DocGrid
	// Document Grid Line Pitch
	LinePitchAttr *int64
	// Document Grid Character Pitch
	CharSpaceAttr *int64
}

func NewCT_DocGrid() *CT_DocGrid {
	ret := &CT_DocGrid{}
	return ret
}

func (m *CT_DocGrid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_DocGridUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "w:type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.LinePitchAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:linePitch"},
			Value: fmt.Sprintf("%v", *m.LinePitchAttr),
		})
	}
	if m.CharSpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:charSpace"},
			Value: fmt.Sprintf("%v", *m.CharSpaceAttr),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocGrid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "linePitch" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.LinePitchAttr = &parsed
			continue
		}
		if attr.Name.Local == "charSpace" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.CharSpaceAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DocGrid: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DocGrid and its children
func (m *CT_DocGrid) Validate() error {
	return m.ValidateWithPath("CT_DocGrid")
}

// ValidateWithPath validates the CT_DocGrid and its children, prefixing error messages with path
func (m *CT_DocGrid) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
