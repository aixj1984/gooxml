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

type CT_CalcCell struct {
	// Cell Reference
	RAttr   *string
	RefAttr *string
	// Sheet Id
	IAttr *int32
	// Child Chain
	SAttr *bool
	// New Dependency Level
	LAttr *bool
	// New Thread
	TAttr *bool
	// Array
	AAttr *bool
}

func NewCT_CalcCell() *CT_CalcCell {
	ret := &CT_CalcCell{}
	return ret
}

func (m *CT_CalcCell) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "r"},
			Value: fmt.Sprintf("%v", *m.RAttr),
		})
	}
	if m.RefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "ref"},
			Value: fmt.Sprintf("%v", *m.RefAttr),
		})
	}
	if m.IAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "i"},
			Value: fmt.Sprintf("%v", *m.IAttr),
		})
	}
	if m.SAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "s"},
			Value: fmt.Sprintf("%d", b2i(*m.SAttr)),
		})
	}
	if m.LAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "l"},
			Value: fmt.Sprintf("%d", b2i(*m.LAttr)),
		})
	}
	if m.TAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "t"},
			Value: fmt.Sprintf("%d", b2i(*m.TAttr)),
		})
	}
	if m.AAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "a"},
			Value: fmt.Sprintf("%d", b2i(*m.AAttr)),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CalcCell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = &parsed
			continue
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = &parsed
			continue
		}
		if attr.Name.Local == "i" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.IAttr = &pt
			continue
		}
		if attr.Name.Local == "s" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SAttr = &parsed
			continue
		}
		if attr.Name.Local == "l" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LAttr = &parsed
			continue
		}
		if attr.Name.Local == "t" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TAttr = &parsed
			continue
		}
		if attr.Name.Local == "a" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CalcCell: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CalcCell and its children
func (m *CT_CalcCell) Validate() error {
	return m.ValidateWithPath("CT_CalcCell")
}

// ValidateWithPath validates the CT_CalcCell and its children, prefixing error messages with path
func (m *CT_CalcCell) ValidateWithPath(path string) error {
	return nil
}
