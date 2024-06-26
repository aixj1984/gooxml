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
	"time"
)

type CT_TrackChangeRange struct {
	DisplacedByCustomXmlAttr ST_DisplacedByCustomXml
	AuthorAttr               string
	DateAttr                 *time.Time
	// Annotation Identifier
	IdAttr int64
}

func NewCT_TrackChangeRange() *CT_TrackChangeRange {
	ret := &CT_TrackChangeRange{}
	return ret
}

func (m *CT_TrackChangeRange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DisplacedByCustomXmlAttr != ST_DisplacedByCustomXmlUnset {
		attr, err := m.DisplacedByCustomXmlAttr.MarshalXMLAttr(xml.Name{Local: "w:displacedByCustomXml"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr),
	})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr),
		})
	}
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr),
	})
	start.Name.Local = "w:CT_TrackChangeRange"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TrackChangeRange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "displacedByCustomXml" {
			m.DisplacedByCustomXmlAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
			continue
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TrackChangeRange: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TrackChangeRange and its children
func (m *CT_TrackChangeRange) Validate() error {
	return m.ValidateWithPath("CT_TrackChangeRange")
}

// ValidateWithPath validates the CT_TrackChangeRange and its children, prefixing error messages with path
func (m *CT_TrackChangeRange) ValidateWithPath(path string) error {
	if err := m.DisplacedByCustomXmlAttr.ValidateWithPath(path + "/DisplacedByCustomXmlAttr"); err != nil {
		return err
	}
	return nil
}
