package RemoveCyrillic

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Description *string
	Age uint16
	Account *Account
	Address Address
}

type Account struct {
	Id uint64
	Username string
	Password string
}

type Address struct {
	Country string
	City string
	Street string
	HouseNumber int
}

func (p Person) String() string {
	return fmt.Sprintf("Person: %v, %v, account: %v", p.Name, p.Age, p.Account)
}

func TestFilter(t *testing.T) {
	// I just took test from Dias
	inputText := "this.is.указатель.for.test"
	inputPtr := &inputText
	expectedText := "this.is..for.test"
	expectedPtr := &expectedText

	testTable := []struct {
		data string
		input Person
		expected Person
	}{
		{
			input:Person{"Mahaмбет",inputPtr,15,
				&Account{1, "whiteбала", "мойwonderfullпароль"},
				Address{"Kazakстан", "Петропavlovsk", "Nazarбаева",456}},
			expected:Person{"Maha",expectedPtr, 15,
				&Account{1, "white", "wonderfull"},
				Address{"Kazak","avlovsk","Nazar",456}},
		},
		{
			input:Person{},
			expected:Person{},
		},
		{
			input:Person{"Олег", nil,30,
				&Account{2,"bigboss", "оскемен"},
				Address{"CanadaКанада","ТокиоKyzylorda","Kaskelenфоревер",7}},
			expected: Person{"", nil,30,
				&Account{2,"bigboss", ""},
				Address{"Canada","Kyzylorda","Kaskelen",7}},
		},
	}

	for _, testCase := range testTable {
		err := RemoveCyrillicFromStruct(&testCase.input)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(testCase.input, testCase.expected) {
			t.Errorf("incorrect result, expected %v, got %v", testCase.expected, testCase.input)
		}
	}
}