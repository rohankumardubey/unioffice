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

package core_properties ;import (_a "encoding/xml";_d "fmt";_c "github.com/unidoc/unioffice";_e "github.com/unidoc/unioffice/common/logger";_gg "time";);

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (_aed *CT_Keyword )ValidateWithPath (path string )error {return nil };

// Validate validates the CT_CoreProperties and its children
func (_ae *CT_CoreProperties )Validate ()error {return _ae .ValidateWithPath ("\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};func (_fgb *CT_CoreProperties )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_bb :for {_cg ,_cfa :=d .Token ();if _cfa !=nil {return _cfa ;};switch _ea :=_cg .(type ){case _a .StartElement :switch _ea .Name {case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_fgb .Category =new (string );if _de :=d .DecodeElement (_fgb .Category ,&_ea );_de !=nil {return _de ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_fgb .ContentStatus =new (string );if _eab :=d .DecodeElement (_fgb .ContentStatus ,&_ea );_eab !=nil {return _eab ;};case _a .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_fgb .Created =new (_c .XSDAny );if _ga :=d .DecodeElement (_fgb .Created ,&_ea );_ga !=nil {return _ga ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_fgb .Creator =new (_c .XSDAny );if _bag :=d .DecodeElement (_fgb .Creator ,&_ea );_bag !=nil {return _bag ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_fgb .Description =new (_c .XSDAny );if _gba :=d .DecodeElement (_fgb .Description ,&_ea );_gba !=nil {return _gba ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_fgb .Identifier =new (_c .XSDAny );if _ec :=d .DecodeElement (_fgb .Identifier ,&_ea );_ec !=nil {return _ec ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_fgb .Keywords =NewCT_Keywords ();if _bab :=d .DecodeElement (_fgb .Keywords ,&_ea );_bab !=nil {return _bab ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_fgb .Language =new (_c .XSDAny );if _cfe :=d .DecodeElement (_fgb .Language ,&_ea );_cfe !=nil {return _cfe ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_fgb .LastModifiedBy =new (string );if _ee :=d .DecodeElement (_fgb .LastModifiedBy ,&_ea );_ee !=nil {return _ee ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_fgb .LastPrinted =new (_gg .Time );if _bd :=d .DecodeElement (_fgb .LastPrinted ,&_ea );_bd !=nil {return _bd ;};case _a .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_fgb .Modified =new (_c .XSDAny );if _db :=d .DecodeElement (_fgb .Modified ,&_ea );_db !=nil {return _db ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_fgb .Revision =new (string );if _eed :=d .DecodeElement (_fgb .Revision ,&_ea );_eed !=nil {return _eed ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_fgb .Subject =new (_c .XSDAny );if _ge :=d .DecodeElement (_fgb .Subject ,&_ea );_ge !=nil {return _ge ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_fgb .Title =new (_c .XSDAny );if _da :=d .DecodeElement (_fgb .Title ,&_ea );_da !=nil {return _da ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_fgb .Version =new (string );if _cbg :=d .DecodeElement (_fgb .Version ,&_ea );_cbg !=nil {return _cbg ;};default:_e .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u0020\u0025\u0076",_ea .Name );if _fgd :=d .Skip ();_fgd !=nil {return _fgd ;};};case _a .EndElement :break _bb ;case _a .CharData :};};return nil ;};func (_bdf *CT_Keywords )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {for _ ,_cfg :=range start .Attr {if _cfg .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_cfg .Name .Local =="\u006c\u0061\u006e\u0067"{_fag ,_dga :=_cfg .Value ,error (nil );if _dga !=nil {return _dga ;};_bdf .LangAttr =&_fag ;continue ;};};_ca :for {_gbe ,_bgg :=d .Token ();if _bgg !=nil {return _bgg ;};switch _bdg :=_gbe .(type ){case _a .StartElement :switch _bdg .Name {case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076\u0061\u006cu\u0065"}:_dfe :=NewCT_Keyword ();if _bba :=d .DecodeElement (_dfe ,&_bdg );_bba !=nil {return _bba ;};_bdf .Value =append (_bdf .Value ,_dfe );default:_e .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073\u0020\u0025\u0076",_bdg .Name );if _abg :=d .Skip ();_abg !=nil {return _abg ;};};case _a .EndElement :break _ca ;case _a .CharData :};};return nil ;};func (_ged *CoreProperties )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0063\u0070"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f"});start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"},Value :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"});start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0063\u0070\u003a\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073";return _ged .CT_CoreProperties .MarshalXML (e ,start );};func (_gfa *CT_Keywords )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {if _gfa .LangAttr !=nil {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_d .Sprintf ("\u0025\u0076",*_gfa .LangAttr )});};e .EncodeToken (start );if _gfa .Value !=nil {_dg :=_a .StartElement {Name :_a .Name {Local :"\u0063\u0070\u003a\u0076\u0061\u006c\u0075\u0065"}};for _ ,_eb :=range _gfa .Value {e .EncodeElement (_eb ,_dg );};};e .EncodeToken (_a .EndElement {Name :start .Name });return nil ;};func NewCoreProperties ()*CoreProperties {_gabb :=&CoreProperties {};_gabb .CT_CoreProperties =*NewCT_CoreProperties ();return _gabb ;};func (_fa *CT_CoreProperties )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {e .EncodeToken (start );if _fa .Category !=nil {_fg :=_a .StartElement {Name :_a .Name {Local :"c\u0070\u003a\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}};_c .AddPreserveSpaceAttr (&_fg ,*_fa .Category );e .EncodeElement (_fa .Category ,_fg );};if _fa .ContentStatus !=nil {_gb :=_a .StartElement {Name :_a .Name {Local :"\u0063\u0070:\u0063\u006f\u006et\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}};_c .AddPreserveSpaceAttr (&_gb ,*_fa .ContentStatus );e .EncodeElement (_fa .ContentStatus ,_gb );};if _fa .Created !=nil {_cb :=_a .StartElement {Name :_a .Name {Local :"\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064"}};e .EncodeElement (_fa .Created ,_cb );};if _fa .Creator !=nil {_b :=_a .StartElement {Name :_a .Name {Local :"\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}};e .EncodeElement (_fa .Creator ,_b );};if _fa .Description !=nil {_cf :=_a .StartElement {Name :_a .Name {Local :"\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}};e .EncodeElement (_fa .Description ,_cf );};if _fa .Identifier !=nil {_ff :=_a .StartElement {Name :_a .Name {Local :"\u0064\u0063\u003a\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}};e .EncodeElement (_fa .Identifier ,_ff );};if _fa .Keywords !=nil {_ab :=_a .StartElement {Name :_a .Name {Local :"c\u0070\u003a\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}};e .EncodeElement (_fa .Keywords ,_ab );};if _fa .Language !=nil {_dc :=_a .StartElement {Name :_a .Name {Local :"d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}};e .EncodeElement (_fa .Language ,_dc );};if _fa .LastModifiedBy !=nil {_ef :=_a .StartElement {Name :_a .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}};_c .AddPreserveSpaceAttr (&_ef ,*_fa .LastModifiedBy );e .EncodeElement (_fa .LastModifiedBy ,_ef );};if _fa .LastPrinted !=nil {_bc :=_a .StartElement {Name :_a .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u0050\u0072i\u006e\u0074\u0065\u0064"}};e .EncodeElement (_fa .LastPrinted ,_bc );};if _fa .Modified !=nil {_ba :=_a .StartElement {Name :_a .Name {Local :"\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}};e .EncodeElement (_fa .Modified ,_ba );};if _fa .Revision !=nil {_ed :=_a .StartElement {Name :_a .Name {Local :"c\u0070\u003a\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}};_c .AddPreserveSpaceAttr (&_ed ,*_fa .Revision );e .EncodeElement (_fa .Revision ,_ed );};if _fa .Subject !=nil {_gf :=_a .StartElement {Name :_a .Name {Local :"\u0064\u0063\u003a\u0073\u0075\u0062\u006a\u0065\u0063\u0074"}};e .EncodeElement (_fa .Subject ,_gf );};if _fa .Title !=nil {_df :=_a .StartElement {Name :_a .Name {Local :"\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}};e .EncodeElement (_fa .Title ,_df );};if _fa .Version !=nil {_bf :=_a .StartElement {Name :_a .Name {Local :"\u0063\u0070\u003a\u0076\u0065\u0072\u0073\u0069\u006f\u006e"}};_c .AddPreserveSpaceAttr (&_bf ,*_fa .Version );e .EncodeElement (_fa .Version ,_bf );};e .EncodeToken (_a .EndElement {Name :start .Name });return nil ;};func NewCT_CoreProperties ()*CT_CoreProperties {_f :=&CT_CoreProperties {};return _f };type CT_CoreProperties struct{Category *string ;ContentStatus *string ;Created *_c .XSDAny ;Creator *_c .XSDAny ;Description *_c .XSDAny ;Identifier *_c .XSDAny ;Keywords *CT_Keywords ;Language *_c .XSDAny ;LastModifiedBy *string ;LastPrinted *_gg .Time ;Modified *_c .XSDAny ;Revision *string ;Subject *_c .XSDAny ;Title *_c .XSDAny ;Version *string ;};

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (_gab *CT_CoreProperties )ValidateWithPath (path string )error {if _gab .Keywords !=nil {if _af :=_gab .Keywords .ValidateWithPath (path +"\u002fK\u0065\u0079\u0077\u006f\u0072\u0064s");_af !=nil {return _af ;};};return nil ;};

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (_fb *CoreProperties )ValidateWithPath (path string )error {if _efc :=_fb .CT_CoreProperties .ValidateWithPath (path );_efc !=nil {return _efc ;};return nil ;};func NewCT_Keyword ()*CT_Keyword {_eede :=&CT_Keyword {};return _eede };type CT_Keywords struct{LangAttr *string ;Value []*CT_Keyword ;};func NewCT_Keywords ()*CT_Keywords {_fd :=&CT_Keywords {};return _fd };type CoreProperties struct{CT_CoreProperties };

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (_bfg *CT_Keywords )ValidateWithPath (path string )error {for _gad ,_adg :=range _bfg .Value {if _gga :=_adg .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002fV\u0061\u006c\u0075\u0065\u005b\u0025\u0064\u005d",path ,_gad ));_gga !=nil {return _gga ;};};return nil ;};

// Validate validates the CoreProperties and its children
func (_ffc *CoreProperties )Validate ()error {return _ffc .ValidateWithPath ("\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};func (_afd *CT_Keyword )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {for _ ,_eca :=range start .Attr {if _eca .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_eca .Name .Local =="\u006c\u0061\u006e\u0067"{_gc ,_be :=_eca .Value ,error (nil );if _be !=nil {return _be ;};_afd .LangAttr =&_gc ;continue ;};};for {_gfc ,_ad :=d .Token ();if _ad !=nil {return _d .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u003a\u0020%\u0073",_ad );};if _ecf ,_ffb :=_gfc .(_a .CharData );_ffb {_afd .Content =string (_ecf );};if _dbg ,_ffa :=_gfc .(_a .EndElement );_ffa &&_dbg .Name ==start .Name {break ;};};return nil ;};func (_afdd *CoreProperties )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_afdd .CT_CoreProperties =*NewCT_CoreProperties ();_cc :for {_abgg ,_ecc :=d .Token ();if _ecc !=nil {return _ecc ;};switch _fad :=_abgg .(type ){case _a .StartElement :switch _fad .Name {case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_afdd .Category =new (string );if _daf :=d .DecodeElement (_afdd .Category ,&_fad );_daf !=nil {return _daf ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_afdd .ContentStatus =new (string );if _dd :=d .DecodeElement (_afdd .ContentStatus ,&_fad );_dd !=nil {return _dd ;};case _a .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_afdd .Created =new (_c .XSDAny );if _adb :=d .DecodeElement (_afdd .Created ,&_fad );_adb !=nil {return _adb ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_afdd .Creator =new (_c .XSDAny );if _bgc :=d .DecodeElement (_afdd .Creator ,&_fad );_bgc !=nil {return _bgc ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_afdd .Description =new (_c .XSDAny );if _efg :=d .DecodeElement (_afdd .Description ,&_fad );_efg !=nil {return _efg ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_afdd .Identifier =new (_c .XSDAny );if _ecaf :=d .DecodeElement (_afdd .Identifier ,&_fad );_ecaf !=nil {return _ecaf ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_afdd .Keywords =NewCT_Keywords ();if _fff :=d .DecodeElement (_afdd .Keywords ,&_fad );_fff !=nil {return _fff ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_afdd .Language =new (_c .XSDAny );if _gedg :=d .DecodeElement (_afdd .Language ,&_fad );_gedg !=nil {return _gedg ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_afdd .LastModifiedBy =new (string );if _faf :=d .DecodeElement (_afdd .LastModifiedBy ,&_fad );_faf !=nil {return _faf ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_afdd .LastPrinted =new (_gg .Time );if _fgdc :=d .DecodeElement (_afdd .LastPrinted ,&_fad );_fgdc !=nil {return _fgdc ;};case _a .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_afdd .Modified =new (_c .XSDAny );if _abd :=d .DecodeElement (_afdd .Modified ,&_fad );_abd !=nil {return _abd ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_afdd .Revision =new (string );if _feg :=d .DecodeElement (_afdd .Revision ,&_fad );_feg !=nil {return _feg ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_afdd .Subject =new (_c .XSDAny );if _bee :=d .DecodeElement (_afdd .Subject ,&_fad );_bee !=nil {return _bee ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_afdd .Title =new (_c .XSDAny );if _cfgd :=d .DecodeElement (_afdd .Title ,&_fad );_cfgd !=nil {return _cfgd ;};case _a .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_afdd .Version =new (string );if _feb :=d .DecodeElement (_afdd .Version ,&_fad );_feb !=nil {return _feb ;};default:_e .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072t\u0069e\u0073\u0020\u0025\u0076",_fad .Name );if _gd :=d .Skip ();_gd !=nil {return _gd ;};};case _a .EndElement :break _cc ;case _a .CharData :};};return nil ;};

// Validate validates the CT_Keyword and its children
func (_deb *CT_Keyword )Validate ()error {return _deb .ValidateWithPath ("\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064");};func (_bg *CT_Keyword )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {if _bg .LangAttr !=nil {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_d .Sprintf ("\u0025\u0076",*_bg .LangAttr )});};e .EncodeElement (_bg .Content ,start );e .EncodeToken (_a .EndElement {Name :start .Name });return nil ;};type CT_Keyword struct{LangAttr *string ;Content string ;};

// Validate validates the CT_Keywords and its children
func (_bae *CT_Keywords )Validate ()error {return _bae .ValidateWithPath ("C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073");};func init (){_c .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCT_CoreProperties );_c .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",NewCT_Keywords );_c .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064",NewCT_Keyword );_c .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCoreProperties );};