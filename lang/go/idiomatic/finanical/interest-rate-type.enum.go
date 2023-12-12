//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package finanical

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

type InterestRateType string

// /////////////////////////////////////////////////////////////////
//  InterestRateType Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t InterestRateType) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  InterestRateType JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t InterestRateType) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *InterestRateType) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, InterestRateTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  InterestRateType YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t InterestRateType) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *InterestRateType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, InterestRateTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  InterestRateType XML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t InterestRateType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return enumer.MarshalXML(t, e, start)
}

func (t *InterestRateType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return enumer.UnmarshalXML(t, InterestRateTypes.Parse, d, start)
}

// /////////////////////////////////////////////////////////////////
//  InterestRateType SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t InterestRateType) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *InterestRateType) Scan(value interface{}) error {
	return enumer.Scan(value, t, InterestRateTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type interestRateTypes struct {
	Simple   InterestRateType
	Compound InterestRateType
	Values   []InterestRateType
	Err      error
}

var InterestRateTypes = interestRateTypes{
	Simple:   InterestRateType("Simple"),
	Compound: InterestRateType("Compound"),
	Err:      fmt.Errorf("invalid InterestRateType"),
}

func init() {
	InterestRateTypes.Values = []InterestRateType{
		InterestRateTypes.Simple,
		InterestRateTypes.Compound,
	}
}

func (t interestRateTypes) newErr(a any, values ...InterestRateType) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		InterestRateTypes.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t interestRateTypes) ParseFrom(v string, values ...InterestRateType) (InterestRateType, error) {
	var found InterestRateType
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, InterestRateType](v)(value) {
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

func (t interestRateTypes) Parse(v string) (InterestRateType, error) {
	return t.ParseFrom(v, InterestRateTypes.Values...)
}

func (t interestRateTypes) IsFrom(v string, values ...InterestRateType) bool {
	for _, value := range values {
		if enumer.IsEq[string, InterestRateType](v)(value) {
			return true
		}
	}
	return false
}

func (t interestRateTypes) Is(v string) bool {
	return t.IsFrom(v, InterestRateTypes.Values...)
}
