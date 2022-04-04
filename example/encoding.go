package example

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func marshalJson(member *Member) []byte {
	json, errMarshal := json.Marshal(member)
	if errMarshal != nil {
		panic(errMarshal)
	}

	return json
}

func unmarshalJson(jsonBytes []byte) *Member {
	var member Member
	errUnmarshal := json.Unmarshal(jsonBytes, &member)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	return &member
}

func marshalXml(member *Member) []byte {
	xml, errMarshal := xml.Marshal(member)
	if errMarshal != nil {
		panic(errMarshal)
	}

	return xml
}

func unmarshalXml(xmlBytes []byte) *Member {
	var member Member
	errUnmarshal := xml.Unmarshal(xmlBytes, &member)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	return &member
}

func ExampleEncoding() {
	member := &Member{"Hong Gildong", 26, true}
	jsonFromMember := marshalJson(member)
	memberFromJson := unmarshalJson(jsonFromMember)

	xmlFromMember := marshalXml(member)
	memberFromXml := unmarshalXml(xmlFromMember)

	fmt.Printf("JSON : %s\nSTRUCT : %#v\n", jsonFromMember, memberFromJson)
	fmt.Printf("XML  : %s\nSTRUCT : %#v\n", xmlFromMember, memberFromXml)

}
