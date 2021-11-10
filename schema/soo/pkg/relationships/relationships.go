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

package relationships ;import (_gg "encoding/xml";_c "fmt";_a "github.com/unidoc/unioffice";_cd "github.com/unidoc/unioffice/common/logger";);func (_b *CT_Relationships )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_ba :for {_cb ,_egc :=d .Token ();if _egc !=nil {return _egc ;};switch _bb :=_cb .(type ){case _gg .StartElement :switch _bb .Name {case _gg .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s",Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:_ecf :=NewRelationship ();if _ce :=d .DecodeElement (_ecf ,&_bb );_ce !=nil {return _ce ;};_b .Relationship =append (_b .Relationship ,_ecf );default:_cd .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u0020\u0025v",_bb .Name );if _cbe :=d .Skip ();_cbe !=nil {return _cbe ;};};case _gg .EndElement :break _ba ;case _gg .CharData :};};return nil ;};func NewCT_Relationships ()*CT_Relationships {_dg :=&CT_Relationships {};return _dg };type Relationships struct{CT_Relationships };const (ST_TargetModeUnset ST_TargetMode =0;ST_TargetModeExternal ST_TargetMode =1;ST_TargetModeInternal ST_TargetMode =2;);func (_egb *ST_TargetMode )UnmarshalXMLAttr (attr _gg .Attr )error {switch attr .Value {case "":*_egb =0;case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_egb =1;case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_egb =2;};return nil ;};type CT_Relationships struct{Relationship []*Relationship ;};func (_gga *CT_Relationships )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {e .EncodeToken (start );if _gga .Relationship !=nil {_cf :=_gg .StartElement {Name :_gg .Name {Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}};for _ ,_gbc :=range _gga .Relationship {e .EncodeElement (_gbc ,_cf );};};e .EncodeToken (_gg .EndElement {Name :start .Name });return nil ;};type ST_TargetMode byte ;func (_ca *Relationship )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {return _ca .CT_Relationship .MarshalXML (e ,start );};

// ValidateWithPath validates the Relationships and its children, prefixing error messages with path
func (_gcb *Relationships )ValidateWithPath (path string )error {if _df :=_gcb .CT_Relationships .ValidateWithPath (path );_df !=nil {return _df ;};return nil ;};func (_gf *Relationships )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s"});start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073";return _gf .CT_Relationships .MarshalXML (e ,start );};

// ValidateWithPath validates the Relationship and its children, prefixing error messages with path
func (_fee *Relationship )ValidateWithPath (path string )error {if _bde :=_fee .CT_Relationship .ValidateWithPath (path );_bde !=nil {return _bde ;};return nil ;};

// Validate validates the Relationships and its children
func (_aaa *Relationships )Validate ()error {return _aaa .ValidateWithPath ("\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073");};func NewRelationships ()*Relationships {_cec :=&Relationships {};_cec .CT_Relationships =*NewCT_Relationships ();return _cec ;};func (_acb ST_TargetMode )ValidateWithPath (path string )error {switch _acb {case 0,1,2:default:return _c .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_acb ));};return nil ;};func (_f *CT_Relationship )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {if _f .TargetModeAttr !=ST_TargetModeUnset {_af ,_fe :=_f .TargetModeAttr .MarshalXMLAttr (_gg .Name {Local :"\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"});if _fe !=nil {return _fe ;};start .Attr =append (start .Attr ,_af );};start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0054\u0061\u0072\u0067\u0065\u0074"},Value :_c .Sprintf ("\u0025\u0076",_f .TargetAttr )});start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0054\u0079\u0070\u0065"},Value :_c .Sprintf ("\u0025\u0076",_f .TypeAttr )});start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0049\u0064"},Value :_c .Sprintf ("\u0025\u0076",_f .IdAttr )});e .EncodeElement (_f .Content ,start );e .EncodeToken (_gg .EndElement {Name :start .Name });return nil ;};func (_ggd *Relationships )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_ggd .CT_Relationships =*NewCT_Relationships ();_acc :for {_fae ,_db :=d .Token ();if _db !=nil {return _db ;};switch _be :=_fae .(type ){case _gg .StartElement :switch _be .Name {case _gg .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s",Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:_eag :=NewRelationship ();if _afb :=d .DecodeElement (_eag ,&_be );_afb !=nil {return _afb ;};_ggd .Relationship =append (_ggd .Relationship ,_eag );default:_cd .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0052\u0065\u006c\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073 \u0025\u0076",_be .Name );if _cee :=d .Skip ();_cee !=nil {return _cee ;};};case _gg .EndElement :break _acc ;case _gg .CharData :};};return nil ;};

// Validate validates the Relationship and its children
func (_adb *Relationship )Validate ()error {return _adb .ValidateWithPath ("\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070");};type Relationship struct{CT_Relationship };func (_ecdf ST_TargetMode )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {return e .EncodeElement (_ecdf .String (),start );};

// Validate validates the CT_Relationship and its children
func (_eg *CT_Relationship )Validate ()error {return _eg .ValidateWithPath ("\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070");};func NewRelationship ()*Relationship {_feg :=&Relationship {};_feg .CT_Relationship =*NewCT_Relationship ();return _feg ;};

// ValidateWithPath validates the CT_Relationship and its children, prefixing error messages with path
func (_fb *CT_Relationship )ValidateWithPath (path string )error {if _da :=_fb .TargetModeAttr .ValidateWithPath (path +"\u002fT\u0061r\u0067\u0065\u0074\u004d\u006f\u0064\u0065\u0041\u0074\u0074\u0072");_da !=nil {return _da ;};return nil ;};func (_dbf *ST_TargetMode )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_faf ,_fbd :=d .Token ();if _fbd !=nil {return _fbd ;};if _cgg ,_caa :=_faf .(_gg .EndElement );_caa &&_cgg .Name ==start .Name {*_dbf =1;return nil ;};if _ge ,_eba :=_faf .(_gg .CharData );!_eba {return _c .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_faf );}else {switch string (_ge ){case "":*_dbf =0;case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_dbf =1;case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_dbf =2;};};_faf ,_fbd =d .Token ();if _fbd !=nil {return _fbd ;};if _dae ,_dgb :=_faf .(_gg .EndElement );_dgb &&_dae .Name ==start .Name {return nil ;};return _c .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_faf );};

// ValidateWithPath validates the CT_Relationships and its children, prefixing error messages with path
func (_ecd *CT_Relationships )ValidateWithPath (path string )error {for _ea ,_bg :=range _ecd .Relationship {if _dd :=_bg .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0052el\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u005b\u0025\u0064\u005d",path ,_ea ));_dd !=nil {return _dd ;};};return nil ;};func (_ad *Relationship )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_ad .CT_Relationship =*NewCT_Relationship ();for _ ,_fef :=range start .Attr {if _fef .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"{_ad .TargetModeAttr .UnmarshalXMLAttr (_fef );continue ;};if _fef .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074"{_fa ,_bc :=_fef .Value ,error (nil );if _bc !=nil {return _bc ;};_ad .TargetAttr =_fa ;continue ;};if _fef .Name .Local =="\u0054\u0079\u0070\u0065"{_gbe ,_dc :=_fef .Value ,error (nil );if _dc !=nil {return _dc ;};_ad .TypeAttr =_gbe ;continue ;};if _fef .Name .Local =="\u0049\u0064"{_aa ,_bd :=_fef .Value ,error (nil );if _bd !=nil {return _bd ;};_ad .IdAttr =_aa ;continue ;};};for {_fca ,_cdba :=d .Token ();if _cdba !=nil {return _c .Errorf ("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0052\u0065\u006c\u0061\u0074i\u006f\u006e\u0073\u0068\u0069\u0070\u003a\u0020\u0025\u0073",_cdba );};if _de ,_fce :=_fca .(_gg .EndElement );_fce &&_de .Name ==start .Name {break ;};};return nil ;};func (_ee *CT_Relationship )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {for _ ,_ed :=range start .Attr {if _ed .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"{_ee .TargetModeAttr .UnmarshalXMLAttr (_ed );continue ;};if _ed .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074"{_ec ,_gb :=_ed .Value ,error (nil );if _gb !=nil {return _gb ;};_ee .TargetAttr =_ec ;continue ;};if _ed .Name .Local =="\u0054\u0079\u0070\u0065"{_cc ,_cg :=_ed .Value ,error (nil );if _cg !=nil {return _cg ;};_ee .TypeAttr =_cc ;continue ;};if _ed .Name .Local =="\u0049\u0064"{_cdb ,_eb :=_ed .Value ,error (nil );if _eb !=nil {return _eb ;};_ee .IdAttr =_cdb ;continue ;};};for {_fg ,_ga :=d .Token ();if _ga !=nil {return _c .Errorf ("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068i\u0070:\u0020\u0025\u0073",_ga );};if _ecc ,_d :=_fg .(_gg .CharData );_d {_ee .Content =string (_ecc );};if _fc ,_gc :=_fg .(_gg .EndElement );_gc &&_fc .Name ==start .Name {break ;};};return nil ;};func (_ddb ST_TargetMode )String ()string {switch _ddb {case 0:return "";case 1:return "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c";case 2:return "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c";};return "";};

// Validate validates the CT_Relationships and its children
func (_gge *CT_Relationships )Validate ()error {return _gge .ValidateWithPath ("\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073");};func (_fcea ST_TargetMode )MarshalXMLAttr (name _gg .Name )(_gg .Attr ,error ){_bcd :=_gg .Attr {};_bcd .Name =name ;switch _fcea {case ST_TargetModeUnset :_bcd .Value ="";case ST_TargetModeExternal :_bcd .Value ="\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c";case ST_TargetModeInternal :_bcd .Value ="\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c";};return _bcd ,nil ;};func (_afe ST_TargetMode )Validate ()error {return _afe .ValidateWithPath ("")};type CT_Relationship struct{TargetModeAttr ST_TargetMode ;TargetAttr string ;TypeAttr string ;IdAttr string ;Content string ;};func NewCT_Relationship ()*CT_Relationship {_e :=&CT_Relationship {};return _e };func init (){_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073",NewCT_Relationships );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070",NewCT_Relationship );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073",NewRelationships );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070",NewRelationship );};