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

package algo ;import _d "strconv";func _gg (_f byte )bool {return _f >='0'&&_f <='9'};func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_bb :=make ([]byte ,len (s )*cnt );_dad :=[]byte (s );for _bbb :=0;_bbb < cnt ;_bbb ++{copy (_bb [_bbb :],_dad );};return string (_bb );};

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_b ,_c :=0,0;for _b < len (lhs )&&_c < len (rhs ){_a :=lhs [_b ];_ad :=rhs [_c ];_e :=_gg (_a );_ed :=_gg (_ad );switch {case _e &&!_ed :return true ;case !_e &&_ed :return false ;case !_e &&!_ed :if _a !=_ad {return _a < _ad ;};_b ++;_c ++;default:_dd :=_b +1;_da :=_c +1;for _dd < len (lhs )&&_gg (lhs [_dd ]){_dd ++;};for _da < len (rhs )&&_gg (rhs [_da ]){_da ++;};_ce ,_ :=_d .ParseUint (lhs [_b :_dd ],10,64);_gf ,_ :=_d .ParseUint (rhs [_b :_da ],10,64);if _ce !=_gf {return _ce < _gf ;};_b =_dd ;_c =_da ;};};return len (lhs )< len (rhs );};