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

type CT_HierarchyUsage struct {
	// Hierarchy Usage
	HierarchyUsageAttr int32
}

func NewCT_HierarchyUsage() *CT_HierarchyUsage {
	ret := &CT_HierarchyUsage{}
	return ret
}

func (m *CT_HierarchyUsage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "hierarchyUsage"},
		Value: fmt.Sprintf("%v", m.HierarchyUsageAttr),
	})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_HierarchyUsage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "hierarchyUsage" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.HierarchyUsageAttr = int32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_HierarchyUsage: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_HierarchyUsage and its children
func (m *CT_HierarchyUsage) Validate() error {
	return m.ValidateWithPath("CT_HierarchyUsage")
}

// ValidateWithPath validates the CT_HierarchyUsage and its children, prefixing error messages with path
func (m *CT_HierarchyUsage) ValidateWithPath(path string) error {
	return nil
}
