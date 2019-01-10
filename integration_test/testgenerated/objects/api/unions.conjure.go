// This file was generated by Conjure and should not be manually edited.

package api

import (
	"fmt"

	"github.com/palantir/pkg/safejson"
)

type ExampleUnion struct {
	typ         string
	str         *string
	strOptional **string
	other       *int
}

type exampleUnionDeserializer struct {
	Type        string   `json:"type" yaml:"type"`
	Str         *string  `json:"str" yaml:"str"`
	StrOptional **string `json:"strOptional" yaml:"strOptional"`
	Other       *int     `json:"other" yaml:"other"`
}

func (u *exampleUnionDeserializer) toStruct() ExampleUnion {
	return ExampleUnion{typ: u.Type, str: u.Str, strOptional: u.StrOptional, other: u.Other}
}

func (u *ExampleUnion) toSerializer() (interface{}, error) {
	switch u.typ {
	default:
		return nil, fmt.Errorf("unknown type %s", u.typ)
	case "str":
		return struct {
			Type string `json:"type" yaml:"type"`
			Str  string `json:"str" yaml:"str"`
		}{Type: "str", Str: *u.str}, nil
	case "strOptional":
		var strOptional *string
		if u.strOptional != nil {
			strOptional = *u.strOptional
		}
		return struct {
			Type        string  `json:"type" yaml:"type"`
			StrOptional *string `json:"strOptional" yaml:"strOptional"`
		}{Type: "strOptional", StrOptional: strOptional}, nil
	case "other":
		return struct {
			Type  string `json:"type" yaml:"type"`
			Other int    `json:"other" yaml:"other"`
		}{Type: "other", Other: *u.other}, nil
	}
}

func (u ExampleUnion) MarshalJSON() ([]byte, error) {
	ser, err := u.toSerializer()
	if err != nil {
		return nil, err
	}
	return safejson.Marshal(ser)
}

func (u *ExampleUnion) UnmarshalJSON(data []byte) error {
	var deser exampleUnionDeserializer
	if err := safejson.Unmarshal(data, &deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u ExampleUnion) MarshalYAML() (interface{}, error) {
	return u.toSerializer()
}

func (u *ExampleUnion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var deser exampleUnionDeserializer
	if err := unmarshal(&deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u *ExampleUnion) Accept(v ExampleUnionVisitor) error {
	switch u.typ {
	default:
		if u.typ == "" {
			return fmt.Errorf("invalid value in union type")
		}
		return v.VisitUnknown(u.typ)
	case "str":
		return v.VisitStr(*u.str)
	case "strOptional":
		var strOptional *string
		if u.strOptional != nil {
			strOptional = *u.strOptional
		}
		return v.VisitStrOptional(strOptional)
	case "other":
		return v.VisitOther(*u.other)
	}
}

type ExampleUnionVisitor interface {
	VisitStr(v string) error
	VisitStrOptional(v *string) error
	VisitOther(v int) error
	VisitUnknown(typeName string) error
}

func NewExampleUnionFromStr(v string) ExampleUnion {
	return ExampleUnion{typ: "str", str: &v}
}

func NewExampleUnionFromStrOptional(v *string) ExampleUnion {
	return ExampleUnion{typ: "strOptional", strOptional: &v}
}

func NewExampleUnionFromOther(v int) ExampleUnion {
	return ExampleUnion{typ: "other", other: &v}
}
