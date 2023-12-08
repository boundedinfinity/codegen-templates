//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package people

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

type SexType string

// /////////////////////////////////////////////////////////////////
//  SexType Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t SexType) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  SexType JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t SexType) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *SexType) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, SexTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  SexType YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t SexType) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *SexType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, SexTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  SexType XML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t SexType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return enumer.MarshalXML(t, e, start)
}

func (t *SexType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return enumer.UnmarshalXML(t, SexTypes.Parse, d, start)
}

// /////////////////////////////////////////////////////////////////
//  SexType SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t SexType) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *SexType) Scan(value interface{}) error {
	return enumer.Scan(value, t, SexTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type sexTypes struct {
	Male   SexType
	Female SexType
	Other  SexType
	Values []SexType
	Err    error
}

var SexTypes = sexTypes{
	Male:   SexType("male"),
	Female: SexType("female"),
	Other:  SexType("other"),
	Err:    fmt.Errorf("invalid SexType"),
}

func init() {
	SexTypes.Values = []SexType{
		SexTypes.Male,
		SexTypes.Female,
		SexTypes.Other,
	}
}

func (t sexTypes) newErr(a any, values ...SexType) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		SexTypes.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t sexTypes) ParseFrom(v string, values ...SexType) (SexType, error) {
	var found SexType
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, SexType](v)(value) {
			found = value
			ok = true
			break
		}
	}

	if !ok {
		return found, t.newErr(v, values...)
	}

	return found, nil
}

func (t sexTypes) Parse(v string) (SexType, error) {
	return t.ParseFrom(v, SexTypes.Values...)
}

func (t sexTypes) IsFrom(v string, values ...SexType) bool {
	for _, value := range values {
		if enumer.IsEq[string, SexType](v)(value) {
			return true
		}
	}
	return false
}

func (t sexTypes) Is(v string) bool {
	return t.IsFrom(v, SexTypes.Values...)
}
