// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/pml"
)

func TestCT_NotesSlideConstructor(t *testing.T) {
	v := pml.NewCT_NotesSlide()
	if v == nil {
		t.Errorf("pml.NewCT_NotesSlide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesSlide should validate: %s", err)
	}
}

func TestCT_NotesSlideMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesSlide()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesSlide()
	xml.Unmarshal(buf, v2)
}
