//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package update contains definitions needed for updating references after removing rows/columns.
package update ;

// UpdateQuery contains terms of how to update references after removing row/column.
type UpdateQuery struct{

// UpdateType is one of the update types like UpdateActionRemoveColumn.
UpdateType UpdateAction ;

// ColumnIdx is the index of the column removed.
ColumnIdx uint32 ;

// SheetToUpdate contains the name of the sheet on which removing happened.
SheetToUpdate string ;

// UpdateCurrentSheet is true if references without sheet prefix should be updated as well.
UpdateCurrentSheet bool ;};

// UpdateAction is the type for update types constants.
type UpdateAction byte ;const (UpdateActionRemoveColumn UpdateAction =iota ;);