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

type CT_TextNormalAutofit struct {
	FontScaleAttr      *ST_TextFontScalePercentOrPercentString
	LnSpcReductionAttr *ST_TextSpacingPercentOrPercentString
}

func NewCT_TextNormalAutofit() *CT_TextNormalAutofit {
	ret := &CT_TextNormalAutofit{}
	return ret
}

func (m *CT_TextNormalAutofit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FontScaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "fontScale"},
			Value: fmt.Sprintf("%v", *m.FontScaleAttr),
		})
	}
	if m.LnSpcReductionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "lnSpcReduction"},
			Value: fmt.Sprintf("%v", *m.LnSpcReductionAttr),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextNormalAutofit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fontScale" {
			parsed, err := ParseUnionST_TextFontScalePercentOrPercentString(attr.Value)
			if err != nil {
				return err
			}
			m.FontScaleAttr = &parsed
			continue
		}
		if attr.Name.Local == "lnSpcReduction" {
			parsed, err := ParseUnionST_TextSpacingPercentOrPercentString(attr.Value)
			if err != nil {
				return err
			}
			m.LnSpcReductionAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextNormalAutofit: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextNormalAutofit and its children
func (m *CT_TextNormalAutofit) Validate() error {
	return m.ValidateWithPath("CT_TextNormalAutofit")
}

// ValidateWithPath validates the CT_TextNormalAutofit and its children, prefixing error messages with path
func (m *CT_TextNormalAutofit) ValidateWithPath(path string) error {
	if m.FontScaleAttr != nil {
		if err := m.FontScaleAttr.ValidateWithPath(path + "/FontScaleAttr"); err != nil {
			return err
		}
	}
	if m.LnSpcReductionAttr != nil {
		if err := m.LnSpcReductionAttr.ValidateWithPath(path + "/LnSpcReductionAttr"); err != nil {
			return err
		}
	}
	return nil
}
