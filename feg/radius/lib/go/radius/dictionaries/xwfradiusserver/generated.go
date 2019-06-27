/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Code generated by radius-dict-gen. DO NOT EDIT.

package xwfradiusserver

import (
	"strconv"

	"fbc/lib/go/radius"
	"fbc/lib/go/radius/rfc2865"
)

const (
	_FacebookExpressWiFiRADIUSServer_VendorID = 99999
)

func _FacebookExpressWiFiRADIUSServer_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_FacebookExpressWiFiRADIUSServer_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return nil
}

func _FacebookExpressWiFiRADIUSServer_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _FacebookExpressWiFiRADIUSServer_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				values = append(values, vsa[2:int(vsaLen)])
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _FacebookExpressWiFiRADIUSServer_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _FacebookExpressWiFiRADIUSServer_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				return vsa[2:int(vsaLen)], true
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return nil, false
}

func _FacebookExpressWiFiRADIUSServer_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _FacebookExpressWiFiRADIUSServer_VendorID {
			i++
			continue
		}
		for j := 0; len(vsa[j:]) >= 3; {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa[j:]) || vsaLen < 3 {
				i++
				break
			}
			if vsaTyp == typ {
				vsa = append(vsa[:j], vsa[j+int(vsaLen):]...)
			}
			j += int(vsaLen)
		}
		if len(vsa) > 0 {
			copy(p.Attributes[rfc2865.VendorSpecific_Type][i][4:], vsa)
			i++
		} else {
			p.Attributes[rfc2865.VendorSpecific_Type] = append(p.Attributes[rfc2865.VendorSpecific_Type][:i], p.Attributes[rfc2865.VendorSpecific_Type][i+i:]...)
		}
	}
	return _FacebookExpressWiFiRADIUSServer_AddVendor(p, typ, attr)
}

type XWFRADIUSServerID uint32

var XWFRADIUSServerID_Strings = map[XWFRADIUSServerID]string{}

func (a XWFRADIUSServerID) String() string {
	if str, ok := XWFRADIUSServerID_Strings[a]; ok {
		return str
	}
	return "XWFRADIUSServerID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func XWFRADIUSServerID_Add(p *radius.Packet, value XWFRADIUSServerID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _FacebookExpressWiFiRADIUSServer_AddVendor(p, 1, a)
}

func XWFRADIUSServerID_Get(p *radius.Packet) (value XWFRADIUSServerID) {
	value, _ = XWFRADIUSServerID_Lookup(p)
	return
}

func XWFRADIUSServerID_Gets(p *radius.Packet) (values []XWFRADIUSServerID, err error) {
	var i uint32
	for _, attr := range _FacebookExpressWiFiRADIUSServer_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, XWFRADIUSServerID(i))
	}
	return
}

func XWFRADIUSServerID_Lookup(p *radius.Packet) (value XWFRADIUSServerID, err error) {
	a, ok := _FacebookExpressWiFiRADIUSServer_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = XWFRADIUSServerID(i)
	return
}

func XWFRADIUSServerID_Set(p *radius.Packet, value XWFRADIUSServerID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _FacebookExpressWiFiRADIUSServer_SetVendor(p, 1, a)
}

func XWFNormalizedMACAddress_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_AddVendor(p, 2, a)
}

func XWFNormalizedMACAddress_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_AddVendor(p, 2, a)
}

func XWFNormalizedMACAddress_Get(p *radius.Packet) (value []byte) {
	value, _ = XWFNormalizedMACAddress_Lookup(p)
	return
}

func XWFNormalizedMACAddress_GetString(p *radius.Packet) (value string) {
	return string(XWFNormalizedMACAddress_Get(p))
}

func XWFNormalizedMACAddress_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _FacebookExpressWiFiRADIUSServer_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XWFNormalizedMACAddress_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _FacebookExpressWiFiRADIUSServer_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XWFNormalizedMACAddress_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _FacebookExpressWiFiRADIUSServer_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func XWFNormalizedMACAddress_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _FacebookExpressWiFiRADIUSServer_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func XWFNormalizedMACAddress_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_SetVendor(p, 2, a)
}

func XWFNormalizedMACAddress_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_SetVendor(p, 2, a)
}

func XWFRADIUSSessionID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_AddVendor(p, 3, a)
}

func XWFRADIUSSessionID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_AddVendor(p, 3, a)
}

func XWFRADIUSSessionID_Get(p *radius.Packet) (value []byte) {
	value, _ = XWFRADIUSSessionID_Lookup(p)
	return
}

func XWFRADIUSSessionID_GetString(p *radius.Packet) (value string) {
	return string(XWFRADIUSSessionID_Get(p))
}

func XWFRADIUSSessionID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _FacebookExpressWiFiRADIUSServer_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XWFRADIUSSessionID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _FacebookExpressWiFiRADIUSServer_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XWFRADIUSSessionID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _FacebookExpressWiFiRADIUSServer_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func XWFRADIUSSessionID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _FacebookExpressWiFiRADIUSServer_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func XWFRADIUSSessionID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_SetVendor(p, 3, a)
}

func XWFRADIUSSessionID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_SetVendor(p, 3, a)
}

func XWFCVersion_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_AddVendor(p, 4, a)
}

func XWFCVersion_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_AddVendor(p, 4, a)
}

func XWFCVersion_Get(p *radius.Packet) (value []byte) {
	value, _ = XWFCVersion_Lookup(p)
	return
}

func XWFCVersion_GetString(p *radius.Packet) (value string) {
	return string(XWFCVersion_Get(p))
}

func XWFCVersion_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _FacebookExpressWiFiRADIUSServer_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XWFCVersion_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _FacebookExpressWiFiRADIUSServer_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XWFCVersion_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _FacebookExpressWiFiRADIUSServer_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func XWFCVersion_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _FacebookExpressWiFiRADIUSServer_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func XWFCVersion_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_SetVendor(p, 4, a)
}

func XWFCVersion_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _FacebookExpressWiFiRADIUSServer_SetVendor(p, 4, a)
}