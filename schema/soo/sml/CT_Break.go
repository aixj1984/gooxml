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

type CT_Break struct {
	// Id
	IdAttr *uint32
	// Minimum
	MinAttr *uint32
	// Maximum
	MaxAttr *uint32
	// Manual Page Break
	ManAttr *bool
	// Pivot-Created Page Break
	PtAttr *bool
}

func NewCT_Break() *CT_Break {
	ret := &CT_Break{}
	return ret
}

func (m *CT_Break) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr),
		})
	}
	if m.MinAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "min"},
			Value: fmt.Sprintf("%v", *m.MinAttr),
		})
	}
	if m.MaxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "max"},
			Value: fmt.Sprintf("%v", *m.MaxAttr),
		})
	}
	if m.ManAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "man"},
			Value: fmt.Sprintf("%d", b2i(*m.ManAttr)),
		})
	}
	if m.PtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "pt"},
			Value: fmt.Sprintf("%d", b2i(*m.PtAttr)),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Break) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IdAttr = &pt
			continue
		}
		if attr.Name.Local == "min" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MinAttr = &pt
			continue
		}
		if attr.Name.Local == "max" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MaxAttr = &pt
			continue
		}
		if attr.Name.Local == "man" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ManAttr = &parsed
			continue
		}
		if attr.Name.Local == "pt" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PtAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Break: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Break and its children
func (m *CT_Break) Validate() error {
	return m.ValidateWithPath("CT_Break")
}

// ValidateWithPath validates the CT_Break and its children, prefixing error messages with path
func (m *CT_Break) ValidateWithPath(path string) error {
	return nil
}
