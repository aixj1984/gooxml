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

type CT_Color struct {
	// Run Content Color
	ValAttr ST_HexColor
	// Run Content Theme Color
	ThemeColorAttr ST_ThemeColor
	// Run Content Theme Color Tint
	ThemeTintAttr *string
	// Run Content Theme Color Shade
	ThemeShadeAttr *string
}

func NewCT_Color() *CT_Color {
	ret := &CT_Color{}
	return ret
}

func (m *CT_Color) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr),
	})
	if m.ThemeColorAttr != ST_ThemeColorUnset {
		attr, err := m.ThemeColorAttr.MarshalXMLAttr(xml.Name{Local: "w:themeColor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ThemeTintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:themeTint"},
			Value: fmt.Sprintf("%v", *m.ThemeTintAttr),
		})
	}
	if m.ThemeShadeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:themeShade"},
			Value: fmt.Sprintf("%v", *m.ThemeShadeAttr),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Color) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_HexColor(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
		if attr.Name.Local == "themeColor" {
			m.ThemeColorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "themeTint" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeTintAttr = &parsed
			continue
		}
		if attr.Name.Local == "themeShade" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeShadeAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Color: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Color and its children
func (m *CT_Color) Validate() error {
	return m.ValidateWithPath("CT_Color")
}

// ValidateWithPath validates the CT_Color and its children, prefixing error messages with path
func (m *CT_Color) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	if err := m.ThemeColorAttr.ValidateWithPath(path + "/ThemeColorAttr"); err != nil {
		return err
	}
	return nil
}
