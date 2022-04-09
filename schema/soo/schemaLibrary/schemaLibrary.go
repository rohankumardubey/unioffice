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

package schemaLibrary ;import (_e "encoding/xml";_d "fmt";_a "github.com/unidoc/unioffice";_eb "github.com/unidoc/unioffice/common/logger";);

// Validate validates the SchemaLibrary and its children
func (_ad *SchemaLibrary )Validate ()error {return _ad .ValidateWithPath ("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_bag *CT_SchemaLibrary )ValidateWithPath (path string )error {for _aff ,_cg :=range _bag .Schema {if _bc :=_cg .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d",path ,_aff ));_bc !=nil {return _bc ;};};return nil ;};func (_fde *SchemaLibrary )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079";return _fde .CT_SchemaLibrary .MarshalXML (e ,start );};func (_ea *CT_Schema )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for _ ,_g :=range start .Attr {if _g .Name .Local =="\u0075\u0072\u0069"{_gf ,_ec :=_g .Value ,error (nil );if _ec !=nil {return _ec ;};_ea .UriAttr =&_gf ;continue ;};if _g .Name .Local =="\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"{_eag ,_c :=_g .Value ,error (nil );if _c !=nil {return _c ;};_ea .ManifestLocationAttr =&_eag ;continue ;};if _g .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"{_gb ,_cd :=_g .Value ,error (nil );if _cd !=nil {return _cd ;};_ea .SchemaLocationAttr =&_gb ;continue ;};if _g .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"{_ae ,_af :=_g .Value ,error (nil );if _af !=nil {return _af ;};_ea .SchemaLanguageAttr =&_ae ;continue ;};};for {_ce ,_ba :=d .Token ();if _ba !=nil {return _d .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073",_ba );};if _ff ,_de :=_ce .(_e .EndElement );_de &&_ff .Name ==start .Name {break ;};};return nil ;};func NewCT_Schema ()*CT_Schema {_b :=&CT_Schema {};return _b };func (_fd *CT_Schema )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {if _fd .UriAttr !=nil {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u006d\u0061\u003a\u0075\u0072\u0069"},Value :_d .Sprintf ("\u0025\u0076",*_fd .UriAttr )});};if _fd .ManifestLocationAttr !=nil {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"},Value :_d .Sprintf ("\u0025\u0076",*_fd .ManifestLocationAttr )});};if _fd .SchemaLocationAttr !=nil {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"},Value :_d .Sprintf ("\u0025\u0076",*_fd .SchemaLocationAttr )});};if _fd .SchemaLanguageAttr !=nil {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"},Value :_d .Sprintf ("\u0025\u0076",*_fd .SchemaLanguageAttr )});};e .EncodeToken (start );e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};type SchemaLibrary struct{CT_SchemaLibrary };func NewSchemaLibrary ()*SchemaLibrary {_cgc :=&SchemaLibrary {};_cgc .CT_SchemaLibrary =*NewCT_SchemaLibrary ();return _cgc ;};type CT_SchemaLibrary struct{Schema []*CT_Schema ;};

// Validate validates the CT_SchemaLibrary and its children
func (_ece *CT_SchemaLibrary )Validate ()error {return _ece .ValidateWithPath ("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_afd *SchemaLibrary )ValidateWithPath (path string )error {if _fcg :=_afd .CT_SchemaLibrary .ValidateWithPath (path );_fcg !=nil {return _fcg ;};return nil ;};type CT_Schema struct{UriAttr *string ;ManifestLocationAttr *string ;SchemaLocationAttr *string ;SchemaLanguageAttr *string ;};

// Validate validates the CT_Schema and its children
func (_aed *CT_Schema )Validate ()error {return _aed .ValidateWithPath ("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da");};func (_ac *CT_SchemaLibrary )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_db :for {_da ,_ed :=d .Token ();if _ed !=nil {return _ed ;};switch _fe :=_da .(type ){case _e .StartElement :switch _fe .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_baa :=NewCT_Schema ();if _cb :=d .DecodeElement (_baa ,&_fe );_cb !=nil {return _cb ;};_ac .Schema =append (_ac .Schema ,_baa );default:_eb .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v",_fe .Name );if _gbg :=d .Skip ();_gbg !=nil {return _gbg ;};};case _e .EndElement :break _db ;case _e .CharData :};};return nil ;};func NewCT_SchemaLibrary ()*CT_SchemaLibrary {_ef :=&CT_SchemaLibrary {};return _ef };func (_ee *CT_SchemaLibrary )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {e .EncodeToken (start );if _ee .Schema !=nil {_fc :=_e .StartElement {Name :_e .Name {Local :"\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}};for _ ,_df :=range _ee .Schema {e .EncodeElement (_df ,_fc );};};e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_cea *CT_Schema )ValidateWithPath (path string )error {return nil };func (_eg *SchemaLibrary )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_eg .CT_SchemaLibrary =*NewCT_SchemaLibrary ();_ege :for {_dea ,_ga :=d .Token ();if _ga !=nil {return _ga ;};switch _ceg :=_dea .(type ){case _e .StartElement :switch _ceg .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_gfg :=NewCT_Schema ();if _dfb :=d .DecodeElement (_gfg ,&_ceg );_dfb !=nil {return _dfb ;};_eg .Schema =append (_eg .Schema ,_gfg );default:_eb .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076",_ceg .Name );if _cdg :=d .Skip ();_cdg !=nil {return _cdg ;};};case _e .EndElement :break _ege ;case _e .CharData :};};return nil ;};func init (){_a .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043T\u005f\u0053\u0063\u0068\u0065\u006da",NewCT_Schema );_a .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewCT_SchemaLibrary );_a .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewSchemaLibrary );};