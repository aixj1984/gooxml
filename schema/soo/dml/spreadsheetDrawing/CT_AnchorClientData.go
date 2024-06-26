// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_AnchorClientData struct {
	FLocksWithSheetAttr  *bool
	FPrintsWithSheetAttr *bool
}

func NewCT_AnchorClientData() *CT_AnchorClientData {
	ret := &CT_AnchorClientData{}
	return ret
}

func (m *CT_AnchorClientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FLocksWithSheetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "fLocksWithSheet"},
			Value: fmt.Sprintf("%d", b2i(*m.FLocksWithSheetAttr)),
		})
	}
	if m.FPrintsWithSheetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "fPrintsWithSheet"},
			Value: fmt.Sprintf("%d", b2i(*m.FPrintsWithSheetAttr)),
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AnchorClientData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fLocksWithSheet" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FLocksWithSheetAttr = &parsed
			continue
		}
		if attr.Name.Local == "fPrintsWithSheet" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FPrintsWithSheetAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AnchorClientData: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AnchorClientData and its children
func (m *CT_AnchorClientData) Validate() error {
	return m.ValidateWithPath("CT_AnchorClientData")
}

// ValidateWithPath validates the CT_AnchorClientData and its children, prefixing error messages with path
func (m *CT_AnchorClientData) ValidateWithPath(path string) error {
	return nil
}
