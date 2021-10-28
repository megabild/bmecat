package bmecat12

import (
	"encoding/xml"
)

type FeatureSystem struct {
	XMLName xml.Name `xml:"FEATURE_SYSTEM"`

	Name        string `xml:"FEATURE_SYSTEM_NAME"`
	Description string `xml:"FEATURE_SYSTEM_DESCR,omitempty"`
}

type FeatureGroup struct {
	XMLName     xml.Name          `xml:"FEATURE_GROUP"`
	ID          string            `xml:"FEATURE_GROUP_ID"`
	Name        string            `xml:"FEATURE_GROUP_NAME"`
	Description string            `xml:"FEATURE_GROUP_DESCR"`
	Templates   []FeatureTemplate `xml:"FEATURE_TEMPLATE"`
}

type FeatureTemplate struct {
	XMLName xml.Name `xml:"FEATURE_TEMPLATE"`
	Name    string   `xml:"FT_NAME"`
	Unit    string   `xml:"FT_UNIT"`
	Order   int      `xml:"FT_ORDER"`
}
