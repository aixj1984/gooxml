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

type CT_CompatSetting struct {
	// Name of Setting
	NameAttr *string
	// Namespace of Setting
	UriAttr *string
	// Value of Setting
	ValAttr *string
}

func NewCT_CompatSetting() *CT_CompatSetting {
	ret := &CT_CompatSetting{}
	return ret
}

func (m *CT_CompatSetting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:name"},
			Value: fmt.Sprintf("%v", *m.NameAttr),
		})
	}
	if m.UriAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:uri"},
			Value: fmt.Sprintf("%v", *m.UriAttr),
		})
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CompatSetting) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CompatSetting: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CompatSetting and its children
func (m *CT_CompatSetting) Validate() error {
	return m.ValidateWithPath("CT_CompatSetting")
}

// ValidateWithPath validates the CT_CompatSetting and its children, prefixing error messages with path
func (m *CT_CompatSetting) ValidateWithPath(path string) error {
	return nil
}
