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

package vmldrawing ;import (_e "encoding/xml";_b "fmt";_c "github.com/unidoc/unioffice";_ed "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes";_a "github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/office/excel";_ag "github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml";);func NewContainer ()*Container {return &Container {}};type Container struct{Layout *_ag .OfcShapelayout ;ShapeType *_ag .Shapetype ;Shape []*_ag .Shape ;};func (_fd *Container )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0076"},Value :"\u0075\u0072n\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d:v\u006d\u006c"});start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u006f"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006di\u0063\u0072\u006f\u0073\u006f\u0066t\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u006ff\u0066\u0069\u0063\u0065"});start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"});start .Name .Local ="\u0078\u006d\u006c";e .EncodeToken (start );if _fd .Layout !=nil {_g :=_e .StartElement {Name :_e .Name {Local :"\u006f\u003a\u0073\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074"}};e .EncodeElement (_fd .Layout ,_g );};if _fd .ShapeType !=nil {_gc :=_e .StartElement {Name :_e .Name {Local :"v\u003a\u0073\u0068\u0061\u0070\u0065\u0074\u0079\u0070\u0065"}};e .EncodeElement (_fd .ShapeType ,_gc );};for _ ,_dg :=range _fd .Shape {_fc :=_e .StartElement {Name :_e .Name {Local :"\u0076:\u0073\u0068\u0061\u0070\u0065"}};e .EncodeElement (_dg ,_fc );};return e .EncodeToken (_e .EndElement {Name :start .Name });};

// NewCommentShape creates a new comment shape for a given cell index.  The
// indices here are zero based.
func NewCommentShape (col ,row int64 )*_ag .Shape {_ae :=_ag .NewShape ();_ae .IdAttr =_c .String (_b .Sprintf ("\u0063\u0073\u005f\u0025\u0064\u005f\u0025\u0064",col ,row ));_ae .TypeAttr =_c .String ("\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032");_ae .StyleAttr =_c .String ("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006cu\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074:\u0038\u0030\u0070\u0074;\u006d\u0061\u0072\u0067\u0069n-\u0074o\u0070\u003a\u0032pt\u003b\u0077\u0069\u0064\u0074\u0068\u003a1\u0030\u0034\u0070\u0074\u003b\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0037\u0036\u0070\u0074\u003b\u007a\u002d\u0069\u006e\u0064\u0065x\u003a\u0031\u003bv\u0069\u0073\u0069\u0062\u0069\u006c\u0069t\u0079\u003a\u0068\u0069\u0064\u0064\u0065\u006e");_ae .FillcolorAttr =_c .String ("\u0023f\u0062\u0066\u0036\u0064\u0036");_ae .StrokecolorAttr =_c .String ("\u0023e\u0064\u0065\u0061\u0061\u0031");_d :=_ag .NewEG_ShapeElements ();_d .Fill =_ag .NewFill ();_d .Fill .Color2Attr =_c .String ("\u0023f\u0062\u0066\u0065\u0038\u0032");_d .Fill .AngleAttr =_c .Float64 (-180);_d .Fill .TypeAttr =_ag .ST_FillTypeGradient ;_d .Fill .Fill =_ag .NewOfcFill ();_d .Fill .Fill .ExtAttr =_ag .ST_ExtView ;_d .Fill .Fill .TypeAttr =_ag .OfcST_FillTypeGradientUnscaled ;_ae .EG_ShapeElements =append (_ae .EG_ShapeElements ,_d );_fb :=_ag .NewEG_ShapeElements ();_fb .Shadow =_ag .NewShadow ();_fb .Shadow .OnAttr =_ed .ST_TrueFalseT ;_fb .Shadow .ObscuredAttr =_ed .ST_TrueFalseT ;_ae .EG_ShapeElements =append (_ae .EG_ShapeElements ,_fb );_df :=_ag .NewEG_ShapeElements ();_df .Path =_ag .NewPath ();_df .Path .ConnecttypeAttr =_ag .OfcST_ConnectTypeNone ;_ae .EG_ShapeElements =append (_ae .EG_ShapeElements ,_df );_fe :=_ag .NewEG_ShapeElements ();_fe .Textbox =_ag .NewTextbox ();_fe .Textbox .StyleAttr =_c .String ("\u006d\u0073\u006f\u002ddi\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u002d\u0061\u006c\u0074\u003a\u0061\u0075t\u006f");_ae .EG_ShapeElements =append (_ae .EG_ShapeElements ,_fe );_agc :=_ag .NewEG_ShapeElements ();_agc .ClientData =_a .NewClientData ();_agc .ClientData .ObjectTypeAttr =_a .ST_ObjectTypeNote ;_agc .ClientData .MoveWithCells =_ed .ST_TrueFalseBlankT ;_agc .ClientData .SizeWithCells =_ed .ST_TrueFalseBlankT ;_agc .ClientData .Anchor =_c .String ("\u0031,\u0020\u0031\u0035\u002c\u0020\u0030\u002c\u0020\u0032\u002c\u00202\u002c\u0020\u0035\u0034\u002c\u0020\u0035\u002c\u0020\u0033");_agc .ClientData .AutoFill =_ed .ST_TrueFalseBlankFalse ;_agc .ClientData .Row =_c .Int64 (row );_agc .ClientData .Column =_c .Int64 (col );_ae .EG_ShapeElements =append (_ae .EG_ShapeElements ,_agc );return _ae ;};func (_ee *Container )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_ee .Shape =nil ;_da :for {_ca ,_cc :=d .Token ();if _cc !=nil {return _cc ;};switch _be :=_ca .(type ){case _e .StartElement :switch _be .Name .Local {case "s\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074":_ee .Layout =_ag .NewOfcShapelayout ();if _eg :=d .DecodeElement (_ee .Layout ,&_be );_eg !=nil {return _eg ;};case "\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e":_ee .ShapeType =_ag .NewShapetype ();if _cf :=d .DecodeElement (_ee .ShapeType ,&_be );_cf !=nil {return _cf ;};case "\u0073\u0068\u0061p\u0065":_bd :=_ag .NewShape ();if _feab :=d .DecodeElement (_bd ,&_be );_feab !=nil {return _feab ;};_ee .Shape =append (_ee .Shape ,_bd );};case _e .EndElement :break _da ;};};return nil ;};

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing ()*Container {_aga :=NewContainer ();_aga .Layout =_ag .NewOfcShapelayout ();_aga .Layout .ExtAttr =_ag .ST_ExtEdit ;_aga .Layout .Idmap =_ag .NewOfcCT_IdMap ();_aga .Layout .Idmap .DataAttr =_c .String ("\u0031");_aga .Layout .Idmap .ExtAttr =_ag .ST_ExtEdit ;_aga .ShapeType =_ag .NewShapetype ();_aga .ShapeType .IdAttr =_c .String ("_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032");_aga .ShapeType .CoordsizeAttr =_c .String ("2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030");_aga .ShapeType .SptAttr =_c .Float32 (202);_aga .ShapeType .PathAttr =_c .String ("\u006d\u0030\u002c0l\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u00321\u00360\u0030,\u00321\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u0030\u0078\u0065");_eb :=_ag .NewEG_ShapeElements ();_aga .ShapeType .EG_ShapeElements =append (_aga .ShapeType .EG_ShapeElements ,_eb );_eb .Path =_ag .NewPath ();_eb .Path .GradientshapeokAttr =_ed .ST_TrueFalseT ;_eb .Path .ConnecttypeAttr =_ag .OfcST_ConnectTypeRect ;return _aga ;};