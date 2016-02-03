package incata

import (
	"testing"
	"time"
)

func TestJSONSerializer(t *testing.T) {

	expectedString := `{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00"}`

	location, err := time.LoadLocation("Europe/Athens")

	if err != nil {
		t.Fatalf("Error getting location!")
	}

	testData := TestData{
		Version:   1,
		Name:      "Joe",
		Balance:   12.99,
		BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
	}

	serializedString, err := NewJSONSerializer().Serialize(testData)

	if err != nil {
		t.Fatalf("Error in serialization! %s", err)
	}

	if expectedString != serializedString {
		t.Fatalf("Expected %s is different than serialized %s", expectedString, serializedString)
	}
}

func TestSerializerWithNull(t *testing.T)  {
    
    var m = make(map[int]int, 0)
    payload, err := NewJSONSerializer().Serialize(m)
    
    if err == nil {
        t.Fatalf("Should have failed! %s", payload)
    }
}

func BenchmarkJSONSerializer(b *testing.B) {

	var sert = NewJSONSerializer()

	location, err := time.LoadLocation("Europe/Athens")

	if err != nil {
		b.Fatalf("Error getting location!")
	}

	testData := TestData{
		Version:   1,
		Name:      "Joe",
		Balance:   12.99,
		BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
	}

	for n := 0; n < b.N; n++ {
		sert.Serialize(testData)
	}
}
