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

package content_types ;import (_d "encoding/xml";_dc "fmt";_dcd "github.com/unidoc/unioffice";_e "github.com/unidoc/unioffice/common/logger";_da "regexp";);func NewOverride ()*Override {_cac :=&Override {};_cac .CT_Override =*NewCT_Override ();return _cac };func (_caf *Override )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_caf .CT_Override =*NewCT_Override ();for _ ,_bdaa :=range start .Attr {if _bdaa .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_agf ,_ddd :=_bdaa .Value ,error (nil );if _ddd !=nil {return _ddd ;};_caf .ContentTypeAttr =_agf ;continue ;};if _bdaa .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_bdab ,_fg :=_bdaa .Value ,error (nil );if _fg !=nil {return _fg ;};_caf .PartNameAttr =_bdab ;continue ;};};for {_ga ,_gd :=d .Token ();if _gd !=nil {return _dc .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073",_gd );};if _ded ,_ec :=_ga .(_d .EndElement );_ec &&_ded .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_fd *CT_Default )ValidateWithPath (path string )error {if !ST_ExtensionPatternRe .MatchString (_fd .ExtensionAttr ){return _dc .Errorf ("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029",path ,ST_ExtensionPatternRe ,_fd .ExtensionAttr );};if !ST_ContentTypePatternRe .MatchString (_fd .ContentTypeAttr ){return _dc .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_fd .ContentTypeAttr );};return nil ;};

// Validate validates the Default and its children
func (_caa *Default )Validate ()error {return _caa .ValidateWithPath ("\u0044e\u0066\u0061\u0075\u006c\u0074");};var ST_ContentTypePatternRe =_da .MustCompile (ST_ContentTypePattern );

// Validate validates the CT_Types and its children
func (_agb *CT_Types )Validate ()error {return _agb .ValidateWithPath ("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073");};var ST_ExtensionPatternRe =_da .MustCompile (ST_ExtensionPattern );type CT_Override struct{ContentTypeAttr string ;PartNameAttr string ;};

// Validate validates the CT_Default and its children
func (_cd *CT_Default )Validate ()error {return _cd .ValidateWithPath ("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074");};

// Validate validates the CT_Override and its children
func (_fa *CT_Override )Validate ()error {return _fa .ValidateWithPath ("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};type Override struct{CT_Override };func NewTypes ()*Types {_efa :=&Types {};_efa .CT_Types =*NewCT_Types ();return _efa };func NewCT_Override ()*CT_Override {_ee :=&CT_Override {};_ee .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _ee ;};func (_a *CT_Default )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_a .ExtensionAttr ="\u0078\u006d\u006c";_a .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";for _ ,_eb :=range start .Attr {if _eb .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_g ,_ge :=_eb .Value ,error (nil );if _ge !=nil {return _ge ;};_a .ExtensionAttr =_g ;continue ;};if _eb .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_ca ,_de :=_eb .Value ,error (nil );if _de !=nil {return _de ;};_a .ContentTypeAttr =_ca ;continue ;};};for {_ag ,_ad :=d .Token ();if _ad !=nil {return _dc .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073",_ad );};if _ac ,_gc :=_ag .(_d .EndElement );_gc &&_ac .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_ecc *Override )ValidateWithPath (path string )error {if _cga :=_ecc .CT_Override .ValidateWithPath (path );_cga !=nil {return _cga ;};return nil ;};func (_fe *Types )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_fe .CT_Types =*NewCT_Types ();_ae :for {_gdc ,_aa :=d .Token ();if _aa !=nil {return _aa ;};switch _aeg :=_gdc .(type ){case _d .StartElement :switch _aeg .Name {case _d .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_gbf :=NewDefault ();if _dccd :=d .DecodeElement (_gbf ,&_aeg );_dccd !=nil {return _dccd ;};_fe .Default =append (_fe .Default ,_gbf );case _d .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_aegc :=NewOverride ();if _eed :=d .DecodeElement (_aegc ,&_aeg );_eed !=nil {return _eed ;};_fe .Override =append (_fe .Override ,_aegc );default:_e .Log .Debug ("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076",_aeg .Name );if _gdb :=d .Skip ();_gdb !=nil {return _gdb ;};};case _d .EndElement :break _ae ;case _d .CharData :};};return nil ;};func (_adf *CT_Types )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {e .EncodeToken (start );if _adf .Default !=nil {_cb :=_d .StartElement {Name :_d .Name {Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}};for _ ,_ef :=range _adf .Default {e .EncodeElement (_ef ,_cb );};};if _adf .Override !=nil {_bg :=_d .StartElement {Name :_d .Name {Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}};for _ ,_eef :=range _adf .Override {e .EncodeElement (_eef ,_bg );};};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_abd *Default )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {return _abd .CT_Default .MarshalXML (e ,start );};func (_faa *Override )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {return _faa .CT_Override .MarshalXML (e ,start );};

// Validate validates the Override and its children
func (_ed *Override )Validate ()error {return _ed .ValidateWithPath ("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};const ST_ExtensionPattern ="\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b";type Default struct{CT_Default };func NewDefault ()*Default {_eea :=&Default {};_eea .CT_Default =*NewCT_Default ();return _eea };type CT_Default struct{ExtensionAttr string ;ContentTypeAttr string ;};type Types struct{CT_Types };func (_be *Default )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_be .CT_Default =*NewCT_Default ();for _ ,_age :=range start .Attr {if _age .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_efb ,_cgf :=_age .Value ,error (nil );if _cgf !=nil {return _cgf ;};_be .ExtensionAttr =_efb ;continue ;};if _age .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_fdc ,_fff :=_age .Value ,error (nil );if _fff !=nil {return _fff ;};_be .ContentTypeAttr =_fdc ;continue ;};};for {_ce ,_ffa :=d .Token ();if _ffa !=nil {return _dc .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073",_ffa );};if _fdf ,_fdg :=_ce .(_d .EndElement );_fdg &&_fdf .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_ea *CT_Override )ValidateWithPath (path string )error {if !ST_ContentTypePatternRe .MatchString (_ea .ContentTypeAttr ){return _dc .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_ea .ContentTypeAttr );};return nil ;};func (_gf *Types )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0054\u0079\u0070e\u0073";return _gf .CT_Types .MarshalXML (e ,start );};

// Validate validates the Types and its children
func (_dcf *Types )Validate ()error {return _dcf .ValidateWithPath ("\u0054\u0079\u0070e\u0073")};func NewCT_Default ()*CT_Default {_daf :=&CT_Default {};_daf .ExtensionAttr ="\u0078\u006d\u006c";_daf .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _daf ;};func (_b *CT_Override )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_dc .Sprintf ("\u0025\u0076",_b .ContentTypeAttr )});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"},Value :_dc .Sprintf ("\u0025\u0076",_b .PartNameAttr )});e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_dga *CT_Types )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_bd :for {_ba ,_afg :=d .Token ();if _afg !=nil {return _afg ;};switch _ggb :=_ba .(type ){case _d .StartElement :switch _ggb .Name {case _d .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_ggbd :=NewDefault ();if _fc :=d .DecodeElement (_ggbd ,&_ggb );_fc !=nil {return _fc ;};_dga .Default =append (_dga .Default ,_ggbd );case _d .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_eac :=NewOverride ();if _ff :=d .DecodeElement (_eac ,&_ggb );_ff !=nil {return _ff ;};_dga .Override =append (_dga .Override ,_eac );default:_e .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076",_ggb .Name );if _ebd :=d .Skip ();_ebd !=nil {return _ebd ;};};case _d .EndElement :break _bd ;case _d .CharData :};};return nil ;};func NewCT_Types ()*CT_Types {_cda :=&CT_Types {};return _cda };func (_dcg *CT_Default )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"},Value :_dc .Sprintf ("\u0025\u0076",_dcg .ExtensionAttr )});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_dc .Sprintf ("\u0025\u0076",_dcg .ContentTypeAttr )});e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_aca *CT_Types )ValidateWithPath (path string )error {for _dcc ,_bda :=range _aca .Default {if _bf :=_bda .ValidateWithPath (_dc .Sprintf ("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d",path ,_dcc ));_bf !=nil {return _bf ;};};for _fac ,_ddg :=range _aca .Override {if _gga :=_ddg .ValidateWithPath (_dc .Sprintf ("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d",path ,_fac ));_gga !=nil {return _gga ;};};return nil ;};const ST_ContentTypePattern ="\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024";

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_fbf *Types )ValidateWithPath (path string )error {if _afc :=_fbf .CT_Types .ValidateWithPath (path );_afc !=nil {return _afc ;};return nil ;};

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_ffd *Default )ValidateWithPath (path string )error {if _agbc :=_ffd .CT_Default .ValidateWithPath (path );_agbc !=nil {return _agbc ;};return nil ;};type CT_Types struct{Default []*Default ;Override []*Override ;};func (_af *CT_Override )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_af .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";for _ ,_gb :=range start .Attr {if _gb .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_cf ,_dd :=_gb .Value ,error (nil );if _dd !=nil {return _dd ;};_af .ContentTypeAttr =_cf ;continue ;};if _gb .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_gcg ,_dg :=_gb .Value ,error (nil );if _dg !=nil {return _dg ;};_af .PartNameAttr =_gcg ;continue ;};};for {_ab ,_bb :=d .Token ();if _bb !=nil {return _dc .Errorf ("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073",_bb );};if _db ,_fb :=_ab .(_d .EndElement );_fb &&_db .Name ==start .Name {break ;};};return nil ;};func init (){_dcd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073",NewCT_Types );_dcd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074",NewCT_Default );_dcd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewCT_Override );_dcd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0054\u0079\u0070e\u0073",NewTypes );_dcd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0044e\u0066\u0061\u0075\u006c\u0074",NewDefault );_dcd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewOverride );};