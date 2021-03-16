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

package schemaLibrary ;import (_a "encoding/xml";_b "fmt";_fd "github.com/unidoc/unioffice";);func (_db *SchemaLibrary )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_db .CT_SchemaLibrary =*NewCT_SchemaLibrary ();_gcb :for {_gdd ,_gf :=d .Token ();if _gf !=nil {return _gf ;};switch _dc :=_gdd .(type ){case _a .StartElement :switch _dc .Name {case _a .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_ddd :=NewCT_Schema ();if _gg :=d .DecodeElement (_ddd ,&_dc );_gg !=nil {return _gg ;};_db .Schema =append (_db .Schema ,_ddd );default:_fd .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076",_dc .Name );if _gcbg :=d .Skip ();_gcbg !=nil {return _gcbg ;};};case _a .EndElement :break _gcb ;case _a .CharData :};};return nil ;};type CT_Schema struct{UriAttr *string ;ManifestLocationAttr *string ;SchemaLocationAttr *string ;SchemaLanguageAttr *string ;};type SchemaLibrary struct{CT_SchemaLibrary };func (_ac *CT_Schema )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {for _ ,_bg :=range start .Attr {if _bg .Name .Local =="\u0075\u0072\u0069"{_ag ,_e :=_bg .Value ,error (nil );if _e !=nil {return _e ;};_ac .UriAttr =&_ag ;continue ;};if _bg .Name .Local =="\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"{_d ,_aa :=_bg .Value ,error (nil );if _aa !=nil {return _aa ;};_ac .ManifestLocationAttr =&_d ;continue ;};if _bg .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"{_eb ,_fda :=_bg .Value ,error (nil );if _fda !=nil {return _fda ;};_ac .SchemaLocationAttr =&_eb ;continue ;};if _bg .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"{_bc ,_eba :=_bg .Value ,error (nil );if _eba !=nil {return _eba ;};_ac .SchemaLanguageAttr =&_bc ;continue ;};};for {_cb ,_fe :=d .Token ();if _fe !=nil {return _b .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073",_fe );};if _af ,_gb :=_cb .(_a .EndElement );_gb &&_af .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Schema and its children
func (_fa *CT_Schema )Validate ()error {return _fa .ValidateWithPath ("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da");};

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_gd *CT_SchemaLibrary )ValidateWithPath (path string )error {for _ab ,_ebe :=range _gd .Schema {if _ge :=_ebe .ValidateWithPath (_b .Sprintf ("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d",path ,_ab ));_ge !=nil {return _ge ;};};return nil ;};type CT_SchemaLibrary struct{Schema []*CT_Schema ;};

// Validate validates the CT_SchemaLibrary and its children
func (_fdb *CT_SchemaLibrary )Validate ()error {return _fdb .ValidateWithPath ("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_ba *CT_Schema )ValidateWithPath (path string )error {return nil };func (_bda *CT_SchemaLibrary )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_cg :for {_bdf ,_bf :=d .Token ();if _bf !=nil {return _bf ;};switch _ga :=_bdf .(type ){case _a .StartElement :switch _ga .Name {case _a .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_dd :=NewCT_Schema ();if _acd :=d .DecodeElement (_dd ,&_ga );_acd !=nil {return _acd ;};_bda .Schema =append (_bda .Schema ,_dd );default:_fd .Log ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v",_ga .Name );if _cgd :=d .Skip ();_cgd !=nil {return _cgd ;};};case _a .EndElement :break _cg ;case _a .CharData :};};return nil ;};func (_c *CT_Schema )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {if _c .UriAttr !=nil {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u006d\u0061\u003a\u0075\u0072\u0069"},Value :_b .Sprintf ("\u0025\u0076",*_c .UriAttr )});};if _c .ManifestLocationAttr !=nil {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"},Value :_b .Sprintf ("\u0025\u0076",*_c .ManifestLocationAttr )});};if _c .SchemaLocationAttr !=nil {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"},Value :_b .Sprintf ("\u0025\u0076",*_c .SchemaLocationAttr )});};if _c .SchemaLanguageAttr !=nil {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"},Value :_b .Sprintf ("\u0025\u0076",*_c .SchemaLanguageAttr )});};e .EncodeToken (start );e .EncodeToken (_a .EndElement {Name :start .Name });return nil ;};func (_gc *CT_SchemaLibrary )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {e .EncodeToken (start );if _gc .Schema !=nil {_bcb :=_a .StartElement {Name :_a .Name {Local :"\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}};for _ ,_ad :=range _gc .Schema {e .EncodeElement (_ad ,_bcb );};};e .EncodeToken (_a .EndElement {Name :start .Name });return nil ;};func NewCT_SchemaLibrary ()*CT_SchemaLibrary {_dg :=&CT_SchemaLibrary {};return _dg };func (_bac *SchemaLibrary )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079";return _bac .CT_SchemaLibrary .MarshalXML (e ,start );};

// Validate validates the SchemaLibrary and its children
func (_ef *SchemaLibrary )Validate ()error {return _ef .ValidateWithPath ("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_abe *SchemaLibrary )ValidateWithPath (path string )error {if _bb :=_abe .CT_SchemaLibrary .ValidateWithPath (path );_bb !=nil {return _bb ;};return nil ;};func NewSchemaLibrary ()*SchemaLibrary {_ec :=&SchemaLibrary {};_ec .CT_SchemaLibrary =*NewCT_SchemaLibrary ();return _ec ;};func NewCT_Schema ()*CT_Schema {_fc :=&CT_Schema {};return _fc };func init (){_fd .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043T\u005f\u0053\u0063\u0068\u0065\u006da",NewCT_Schema );_fd .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewCT_SchemaLibrary );_fd .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewSchemaLibrary );};