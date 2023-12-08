//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package currency

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

type CurrencyCode string

// /////////////////////////////////////////////////////////////////
//  CurrencyCode Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t CurrencyCode) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  CurrencyCode JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t CurrencyCode) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *CurrencyCode) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, CurrencyCodes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  CurrencyCode YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t CurrencyCode) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *CurrencyCode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, CurrencyCodes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  CurrencyCode XML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t CurrencyCode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return enumer.MarshalXML(t, e, start)
}

func (t *CurrencyCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return enumer.UnmarshalXML(t, CurrencyCodes.Parse, d, start)
}

// /////////////////////////////////////////////////////////////////
//  CurrencyCode SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t CurrencyCode) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *CurrencyCode) Scan(value interface{}) error {
	return enumer.Scan(value, t, CurrencyCodes.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type currencyCodes struct {
	USD    CurrencyCode
	EUR    CurrencyCode
	AED    CurrencyCode
	AFN    CurrencyCode
	Values []CurrencyCode
	Err    error
}

var CurrencyCodes = currencyCodes{
	USD: CurrencyCode("USD"),
	EUR: CurrencyCode("EUR"),
	AED: CurrencyCode("AED"),
	AFN: CurrencyCode("AFN"),
	Err: fmt.Errorf("invalid CurrencyCode"),
}

func init() {
	CurrencyCodes.Values = []CurrencyCode{
		CurrencyCodes.USD,
		CurrencyCodes.EUR,
		CurrencyCodes.AED,
		CurrencyCodes.AFN,
	}
}

func (t currencyCodes) newErr(a any, values ...CurrencyCode) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		CurrencyCodes.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t currencyCodes) ParseFrom(v string, values ...CurrencyCode) (CurrencyCode, error) {
	var found CurrencyCode
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, CurrencyCode](v)(value) {
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

func (t currencyCodes) Parse(v string) (CurrencyCode, error) {
	return t.ParseFrom(v, CurrencyCodes.Values...)
}

func (t currencyCodes) IsFrom(v string, values ...CurrencyCode) bool {
	for _, value := range values {
		if enumer.IsEq[string, CurrencyCode](v)(value) {
			return true
		}
	}
	return false
}

func (t currencyCodes) Is(v string) bool {
	return t.IsFrom(v, CurrencyCodes.Values...)
}
