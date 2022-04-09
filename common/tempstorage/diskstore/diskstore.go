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

// Package diskstore implements tempStorage interface
// by using disk as a storage
package diskstore ;import (_c "github.com/unidoc/unioffice/common/tempstorage";_b "io/ioutil";_gb "os";_a "strings";);

// TempFile creates a new temp file by calling ioutil TempFile
func (_cc diskStorage )TempFile (dir ,pattern string )(_c .File ,error ){return _b .TempFile (dir ,pattern );};

// TempFile creates a new temp directory by calling ioutil TempDir
func (_e diskStorage )TempDir (pattern string )(string ,error ){return _b .TempDir ("",pattern )};

// Add is not applicable in the diskstore implementation
func (_ed diskStorage )Add (path string )error {return nil };type diskStorage struct{};

// Open opens file from disk according to a path
func (_d diskStorage )Open (path string )(_c .File ,error ){return _gb .OpenFile (path ,_gb .O_RDWR ,0644)};

// RemoveAll removes all files in the directory
func (_da diskStorage )RemoveAll (dir string )error {if _a .HasPrefix (dir ,_gb .TempDir ()){return _gb .RemoveAll (dir );};return nil ;};

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage (){_f :=diskStorage {};_c .SetAsStorage (&_f )};