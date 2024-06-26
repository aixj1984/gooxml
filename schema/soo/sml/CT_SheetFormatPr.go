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

type CT_SheetFormatPr struct {
	// Base Column Width
	BaseColWidthAttr *uint32
	// Default Column Width
	DefaultColWidthAttr *float64
	// Default Row Height
	DefaultRowHeightAttr float64
	// Custom Height
	CustomHeightAttr *bool
	// Hidden By Default
	ZeroHeightAttr *bool
	// Thick Top Border
	ThickTopAttr *bool
	// Thick Bottom Border
	ThickBottomAttr *bool
	// Maximum Outline Row
	OutlineLevelRowAttr *uint8
	// Column Outline Level
	OutlineLevelColAttr *uint8
}

func NewCT_SheetFormatPr() *CT_SheetFormatPr {
	ret := &CT_SheetFormatPr{}
	return ret
}

func (m *CT_SheetFormatPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BaseColWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "baseColWidth"},
			Value: fmt.Sprintf("%v", *m.BaseColWidthAttr),
		})
	}
	if m.DefaultColWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "defaultColWidth"},
			Value: fmt.Sprintf("%v", *m.DefaultColWidthAttr),
		})
	}
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "defaultRowHeight"},
		Value: fmt.Sprintf("%v", m.DefaultRowHeightAttr),
	})
	if m.CustomHeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "customHeight"},
			Value: fmt.Sprintf("%d", b2i(*m.CustomHeightAttr)),
		})
	}
	if m.ZeroHeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "zeroHeight"},
			Value: fmt.Sprintf("%d", b2i(*m.ZeroHeightAttr)),
		})
	}
	if m.ThickTopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "thickTop"},
			Value: fmt.Sprintf("%d", b2i(*m.ThickTopAttr)),
		})
	}
	if m.ThickBottomAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "thickBottom"},
			Value: fmt.Sprintf("%d", b2i(*m.ThickBottomAttr)),
		})
	}
	if m.OutlineLevelRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "outlineLevelRow"},
			Value: fmt.Sprintf("%v", *m.OutlineLevelRowAttr),
		})
	}
	if m.OutlineLevelColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "outlineLevelCol"},
			Value: fmt.Sprintf("%v", *m.OutlineLevelColAttr),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SheetFormatPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "baseColWidth" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BaseColWidthAttr = &pt
			continue
		}
		if attr.Name.Local == "defaultColWidth" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.DefaultColWidthAttr = &parsed
			continue
		}
		if attr.Name.Local == "defaultRowHeight" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.DefaultRowHeightAttr = parsed
			continue
		}
		if attr.Name.Local == "customHeight" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomHeightAttr = &parsed
			continue
		}
		if attr.Name.Local == "zeroHeight" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ZeroHeightAttr = &parsed
			continue
		}
		if attr.Name.Local == "thickTop" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ThickTopAttr = &parsed
			continue
		}
		if attr.Name.Local == "thickBottom" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ThickBottomAttr = &parsed
			continue
		}
		if attr.Name.Local == "outlineLevelRow" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.OutlineLevelRowAttr = &pt
			continue
		}
		if attr.Name.Local == "outlineLevelCol" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.OutlineLevelColAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SheetFormatPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SheetFormatPr and its children
func (m *CT_SheetFormatPr) Validate() error {
	return m.ValidateWithPath("CT_SheetFormatPr")
}

// ValidateWithPath validates the CT_SheetFormatPr and its children, prefixing error messages with path
func (m *CT_SheetFormatPr) ValidateWithPath(path string) error {
	return nil
}
