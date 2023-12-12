//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package location

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

type StateName string

// /////////////////////////////////////////////////////////////////
//  StateName Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t StateName) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  StateName JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t StateName) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *StateName) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, StateNames.Parse)
}

// /////////////////////////////////////////////////////////////////
//  StateName YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t StateName) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *StateName) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, StateNames.Parse)
}

// /////////////////////////////////////////////////////////////////
//  StateName XML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t StateName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return enumer.MarshalXML(t, e, start)
}

func (t *StateName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return enumer.UnmarshalXML(t, StateNames.Parse, d, start)
}

// /////////////////////////////////////////////////////////////////
//  StateName SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t StateName) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *StateName) Scan(value interface{}) error {
	return enumer.Scan(value, t, StateNames.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type stateNames struct {
	Alabama                     StateName
	Alaska                      StateName
	AmericanSamoa               StateName
	Arizona                     StateName
	Arkansas                    StateName
	California                  StateName
	Colorado                    StateName
	Connecticut                 StateName
	Delaware                    StateName
	DistrictOfColumbia          StateName
	FederatedStatesOfMicronesia StateName
	Florida                     StateName
	Georgia                     StateName
	Guam                        StateName
	Hawaii                      StateName
	Idaho                       StateName
	Illinois                    StateName
	Indiana                     StateName
	Iowa                        StateName
	Kansas                      StateName
	Kentucky                    StateName
	Louisiana                   StateName
	Maine                       StateName
	MarshallIslands             StateName
	Maryland                    StateName
	Massachusetts               StateName
	Michigan                    StateName
	Minnesota                   StateName
	Mississippi                 StateName
	Missouri                    StateName
	Montana                     StateName
	Nebraska                    StateName
	Nevada                      StateName
	NewHampshire                StateName
	NewJersey                   StateName
	NewMexico                   StateName
	NewYork                     StateName
	NorthCarolina               StateName
	NorthDakota                 StateName
	NorthernMarianaIslands      StateName
	Ohio                        StateName
	Oklahoma                    StateName
	Oregon                      StateName
	Palau                       StateName
	Pennsylvania                StateName
	PuertoRico                  StateName
	RhodeIsland                 StateName
	SouthCarolina               StateName
	SouthDakota                 StateName
	Tennessee                   StateName
	Texas                       StateName
	Utah                        StateName
	Vermont                     StateName
	VirginIslands               StateName
	Virginia                    StateName
	Washington                  StateName
	WestVirginia                StateName
	Wisconsin                   StateName
	Wyoming                     StateName
	Values                      []StateName
	Err                         error
}

var StateNames = stateNames{
	Alabama:                     StateName("Alabama"),
	Alaska:                      StateName("Alaska"),
	AmericanSamoa:               StateName("American Samoa"),
	Arizona:                     StateName("Arizona"),
	Arkansas:                    StateName("Arkansas"),
	California:                  StateName("California"),
	Colorado:                    StateName("Colorado"),
	Connecticut:                 StateName("Connecticut"),
	Delaware:                    StateName("Delaware"),
	DistrictOfColumbia:          StateName("District Of Columbia"),
	FederatedStatesOfMicronesia: StateName("Federated States Of Micronesia"),
	Florida:                     StateName("Florida"),
	Georgia:                     StateName("Georgia"),
	Guam:                        StateName("Guam"),
	Hawaii:                      StateName("Hawaii"),
	Idaho:                       StateName("Idaho"),
	Illinois:                    StateName("Illinois"),
	Indiana:                     StateName("Indiana"),
	Iowa:                        StateName("Iowa"),
	Kansas:                      StateName("Kansas"),
	Kentucky:                    StateName("Kentucky"),
	Louisiana:                   StateName("Louisiana"),
	Maine:                       StateName("Maine"),
	MarshallIslands:             StateName("Marshall Islands"),
	Maryland:                    StateName("Maryland"),
	Massachusetts:               StateName("Massachusetts"),
	Michigan:                    StateName("Michigan"),
	Minnesota:                   StateName("Minnesota"),
	Mississippi:                 StateName("Mississippi"),
	Missouri:                    StateName("Missouri"),
	Montana:                     StateName("Montana"),
	Nebraska:                    StateName("Nebraska"),
	Nevada:                      StateName("Nevada"),
	NewHampshire:                StateName("New Hampshire"),
	NewJersey:                   StateName("New Jersey"),
	NewMexico:                   StateName("New Mexico"),
	NewYork:                     StateName("New York"),
	NorthCarolina:               StateName("North Carolina"),
	NorthDakota:                 StateName("North Dakota"),
	NorthernMarianaIslands:      StateName("Northern Mariana Islands"),
	Ohio:                        StateName("Ohio"),
	Oklahoma:                    StateName("Oklahoma"),
	Oregon:                      StateName("Oregon"),
	Palau:                       StateName("Palau"),
	Pennsylvania:                StateName("Pennsylvania"),
	PuertoRico:                  StateName("Puerto Rico"),
	RhodeIsland:                 StateName("Rhode Island"),
	SouthCarolina:               StateName("South Carolina"),
	SouthDakota:                 StateName("South Dakota"),
	Tennessee:                   StateName("Tennessee"),
	Texas:                       StateName("Texas"),
	Utah:                        StateName("Utah"),
	Vermont:                     StateName("Vermont"),
	VirginIslands:               StateName("Virgin Islands"),
	Virginia:                    StateName("Virginia"),
	Washington:                  StateName("Washington"),
	WestVirginia:                StateName("West Virginia"),
	Wisconsin:                   StateName("Wisconsin"),
	Wyoming:                     StateName("Wyoming"),
	Err:                         fmt.Errorf("invalid StateName"),
}

func init() {
	StateNames.Values = []StateName{
		StateNames.Alabama,
		StateNames.Alaska,
		StateNames.AmericanSamoa,
		StateNames.Arizona,
		StateNames.Arkansas,
		StateNames.California,
		StateNames.Colorado,
		StateNames.Connecticut,
		StateNames.Delaware,
		StateNames.DistrictOfColumbia,
		StateNames.FederatedStatesOfMicronesia,
		StateNames.Florida,
		StateNames.Georgia,
		StateNames.Guam,
		StateNames.Hawaii,
		StateNames.Idaho,
		StateNames.Illinois,
		StateNames.Indiana,
		StateNames.Iowa,
		StateNames.Kansas,
		StateNames.Kentucky,
		StateNames.Louisiana,
		StateNames.Maine,
		StateNames.MarshallIslands,
		StateNames.Maryland,
		StateNames.Massachusetts,
		StateNames.Michigan,
		StateNames.Minnesota,
		StateNames.Mississippi,
		StateNames.Missouri,
		StateNames.Montana,
		StateNames.Nebraska,
		StateNames.Nevada,
		StateNames.NewHampshire,
		StateNames.NewJersey,
		StateNames.NewMexico,
		StateNames.NewYork,
		StateNames.NorthCarolina,
		StateNames.NorthDakota,
		StateNames.NorthernMarianaIslands,
		StateNames.Ohio,
		StateNames.Oklahoma,
		StateNames.Oregon,
		StateNames.Palau,
		StateNames.Pennsylvania,
		StateNames.PuertoRico,
		StateNames.RhodeIsland,
		StateNames.SouthCarolina,
		StateNames.SouthDakota,
		StateNames.Tennessee,
		StateNames.Texas,
		StateNames.Utah,
		StateNames.Vermont,
		StateNames.VirginIslands,
		StateNames.Virginia,
		StateNames.Washington,
		StateNames.WestVirginia,
		StateNames.Wisconsin,
		StateNames.Wyoming,
	}
}

func (t stateNames) newErr(a any, values ...StateName) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		StateNames.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t stateNames) ParseFrom(v string, values ...StateName) (StateName, error) {
	var found StateName
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, StateName](v)(value) {
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

func (t stateNames) Parse(v string) (StateName, error) {
	return t.ParseFrom(v, StateNames.Values...)
}

func (t stateNames) IsFrom(v string, values ...StateName) bool {
	for _, value := range values {
		if enumer.IsEq[string, StateName](v)(value) {
			return true
		}
	}
	return false
}

func (t stateNames) Is(v string) bool {
	return t.IsFrom(v, StateNames.Values...)
}
