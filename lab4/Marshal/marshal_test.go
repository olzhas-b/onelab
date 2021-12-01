package Marshal

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
)

var rawJson = []byte(`[
  {
    "id": "1",
    "address": {
      "city_id": "5",
      "street": "Satbayev"
    },
    "Age": 20
  },
  {
    "id": 2,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "Age": "32"
  },
  {
    "id": "2",
    "address": {
      "city_id": 6,
      "street": "Al-Farabi"
    },
    "Age": "32"
  }
]`)
var rawXml = []byte(`
<User>
	<id>1</id>
	<address>
		<city_id>"22"</city_id>
		<street>tole by</street>
	</address>
	<age>18</age>
</User>
`)
type testStruct struct {
	Input       []byte
	Output      []byte
	Expected    []byte
}

func TestUnmarshalXML(t *testing.T) {
	var data = dataXML
	var users UsersTable
	err := xml.Unmarshal(data.Input, &users)
	if err != nil {
		t.Errorf("%+v", err)
	}
	data.Output, _ = xml.MarshalIndent(&users,  "", "\t")
	if !reflect.DeepEqual(data.Output, data.Expected) {
		t.Errorf("incorrect result, expected %v, got %v", data.Expected, data.Output)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	var data = dataJson
	var users UsersTable
	if err := json.Unmarshal(data.Input, &users); err != nil {
		t.Errorf("can't unmarshal data")
	}
	data.Output, _ = json.MarshalIndent(&users, "", "        ")
	if !reflect.DeepEqual(data.Output, data.Expected) {
		t.Errorf("we expected %v, got %v", data.Expected, string(data.Output))
	}
}


var dataJson = testStruct{
Input:[]byte(
`{
        "Users": [
                {
                        "id": 0,
                        "address": {
                                "city_id": "1",
                                "street": "tole by"
                        },
                        "age": 0
                },
				{
                        "id": "123",
                        "address": {
                                "city_id": "133",
                                "street": "tole by"
                        },
                        "age": "222"
                }
        ]
}`),
Expected:[]byte(
`{
        "Users": [
                {
                        "id": 0,
                        "address": {
                                "city_id": 1,
                                "street": "tole by"
                        },
                        "age": 0
                },
                {
                        "id": 123,
                        "address": {
                                "city_id": 133,
                                "street": "tole by"
                        },
                        "age": 222
                }
        ]
}`),
}

var dataXML = testStruct{
Input:[]byte(
`<UsersTable>
	<Users>
		<id>"0"</id>
		<address>
			<city_id>"12"</city_id>
			<street>tole by</street>
		</address>
		<age>123</age>
	</Users>
	<Users>
		<id>11</id>
		<address>
			<city_id>"122"</city_id>
			<street>kazbek by</street>
		</address>
		<age>"133"</age>
	</Users>
</UsersTable>
`),
Expected:[]byte(
`<UsersTable>
	<Users>
		<id>0</id>
		<address>
			<city_id>12</city_id>
			<street>tole by</street>
		</address>
		<age>123</age>
	</Users>
	<Users>
		<id>11</id>
		<address>
			<city_id>122</city_id>
			<street>kazbek by</street>
		</address>
		<age>133</age>
	</Users>
</UsersTable>`),
}