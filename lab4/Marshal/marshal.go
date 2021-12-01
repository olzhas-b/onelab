package Marshal

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
)

func (intOrString *IntOrString) UnmarshalJSON(data []byte) (err error) {
	s := string(data)
	s = strings.Replace(s, `"`, "", 2)
	num, err := strconv.Atoi(s)
	intOrString.Number = num
	return
}

func (intOrString IntOrString) MarshalJSON() ([]byte, error) {
	return json.Marshal(intOrString.Number)
}


func (intOrString *IntOrString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var data string
	err = d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	if string(data[0]) == `"` {
		data = data[1 : len(data)-1]
	}
	number, _ := strconv.Atoi(data)
	intOrString.Number = number
	return
}

func (intOrString *IntOrString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(intOrString.Number, start)
}


type IntOrString struct {
	Number int
}
type UsersTable struct {
	Users []User
}
type User struct {
	ID      IntOrString     `json:"id"      xml:"id"`
	Address Address         `json:"address" xml:"address"`
	Age     IntOrString     `json:"age"     xml:"age"`
}

type Address struct {
	CityID IntOrString      `json:"city_id" xml:"city_id"`
	Street string           `json:"street"  xml:"street"`
}
