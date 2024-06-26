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

type CT_DefinedName struct {
	NameAttr              string
	CommentAttr           *string
	CustomMenuAttr        *string
	DescriptionAttr       *string
	HelpAttr              *string
	StatusBarAttr         *string
	LocalSheetIdAttr      *uint32
	HiddenAttr            *bool
	FunctionAttr          *bool
	VbProcedureAttr       *bool
	XlmAttr               *bool
	FunctionGroupIdAttr   *uint32
	ShortcutKeyAttr       *string
	PublishToServerAttr   *bool
	WorkbookParameterAttr *bool
	Content               string
}

func NewCT_DefinedName() *CT_DefinedName {
	ret := &CT_DefinedName{}
	return ret
}

func (m *CT_DefinedName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr),
	})
	if m.CommentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "comment"},
			Value: fmt.Sprintf("%v", *m.CommentAttr),
		})
	}
	if m.CustomMenuAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "customMenu"},
			Value: fmt.Sprintf("%v", *m.CustomMenuAttr),
		})
	}
	if m.DescriptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "description"},
			Value: fmt.Sprintf("%v", *m.DescriptionAttr),
		})
	}
	if m.HelpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "help"},
			Value: fmt.Sprintf("%v", *m.HelpAttr),
		})
	}
	if m.StatusBarAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "statusBar"},
			Value: fmt.Sprintf("%v", *m.StatusBarAttr),
		})
	}
	if m.LocalSheetIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "localSheetId"},
			Value: fmt.Sprintf("%v", *m.LocalSheetIdAttr),
		})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenAttr)),
		})
	}
	if m.FunctionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "function"},
			Value: fmt.Sprintf("%d", b2i(*m.FunctionAttr)),
		})
	}
	if m.VbProcedureAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "vbProcedure"},
			Value: fmt.Sprintf("%d", b2i(*m.VbProcedureAttr)),
		})
	}
	if m.XlmAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "xlm"},
			Value: fmt.Sprintf("%d", b2i(*m.XlmAttr)),
		})
	}
	if m.FunctionGroupIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "functionGroupId"},
			Value: fmt.Sprintf("%v", *m.FunctionGroupIdAttr),
		})
	}
	if m.ShortcutKeyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "shortcutKey"},
			Value: fmt.Sprintf("%v", *m.ShortcutKeyAttr),
		})
	}
	if m.PublishToServerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "publishToServer"},
			Value: fmt.Sprintf("%d", b2i(*m.PublishToServerAttr)),
		})
	}
	if m.WorkbookParameterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "workbookParameter"},
			Value: fmt.Sprintf("%d", b2i(*m.WorkbookParameterAttr)),
		})
	}
	e.EncodeElement(m.Content, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DefinedName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
			continue
		}
		if attr.Name.Local == "function" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FunctionAttr = &parsed
			continue
		}
		if attr.Name.Local == "comment" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CommentAttr = &parsed
			continue
		}
		if attr.Name.Local == "description" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DescriptionAttr = &parsed
			continue
		}
		if attr.Name.Local == "help" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HelpAttr = &parsed
			continue
		}
		if attr.Name.Local == "statusBar" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StatusBarAttr = &parsed
			continue
		}
		if attr.Name.Local == "localSheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LocalSheetIdAttr = &pt
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "customMenu" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CustomMenuAttr = &parsed
			continue
		}
		if attr.Name.Local == "vbProcedure" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VbProcedureAttr = &parsed
			continue
		}
		if attr.Name.Local == "xlm" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.XlmAttr = &parsed
			continue
		}
		if attr.Name.Local == "functionGroupId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FunctionGroupIdAttr = &pt
			continue
		}
		if attr.Name.Local == "shortcutKey" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ShortcutKeyAttr = &parsed
			continue
		}
		if attr.Name.Local == "publishToServer" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishToServerAttr = &parsed
			continue
		}
		if attr.Name.Local == "workbookParameter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.WorkbookParameterAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DefinedName: %s", err)
		}
		if cd, ok := tok.(xml.CharData); ok {
			m.Content = string(cd)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DefinedName and its children
func (m *CT_DefinedName) Validate() error {
	return m.ValidateWithPath("CT_DefinedName")
}

// ValidateWithPath validates the CT_DefinedName and its children, prefixing error messages with path
func (m *CT_DefinedName) ValidateWithPath(path string) error {
	return nil
}
