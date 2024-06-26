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

type CT_PageMargins struct {
	// Left Page Margin
	LeftAttr float64
	// Right Page Margin
	RightAttr float64
	// Top Page Margin
	TopAttr float64
	// Bottom Page Margin
	BottomAttr float64
	// Header Page Margin
	HeaderAttr float64
	// Footer Page Margin
	FooterAttr float64
}

func NewCT_PageMargins() *CT_PageMargins {
	ret := &CT_PageMargins{}
	return ret
}

func (m *CT_PageMargins) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "left"},
		Value: fmt.Sprintf("%v", m.LeftAttr),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "right"},
		Value: fmt.Sprintf("%v", m.RightAttr),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "top"},
		Value: fmt.Sprintf("%v", m.TopAttr),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "bottom"},
		Value: fmt.Sprintf("%v", m.BottomAttr),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "header"},
		Value: fmt.Sprintf("%v", m.HeaderAttr),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "footer"},
		Value: fmt.Sprintf("%v", m.FooterAttr),
	})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageMargins) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "left" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.LeftAttr = parsed
			continue
		}
		if attr.Name.Local == "right" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.RightAttr = parsed
			continue
		}
		if attr.Name.Local == "top" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.TopAttr = parsed
			continue
		}
		if attr.Name.Local == "bottom" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.BottomAttr = parsed
			continue
		}
		if attr.Name.Local == "header" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.HeaderAttr = parsed
			continue
		}
		if attr.Name.Local == "footer" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.FooterAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageMargins: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PageMargins and its children
func (m *CT_PageMargins) Validate() error {
	return m.ValidateWithPath("CT_PageMargins")
}

// ValidateWithPath validates the CT_PageMargins and its children, prefixing error messages with path
func (m *CT_PageMargins) ValidateWithPath(path string) error {
	return nil
}
