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

	serializedString, err := NewJSONMarshaller().Serialize(testData)

	if err != nil {
		t.Fatalf("Error in serialization! %s", err)
	}

	if expectedString != serializedString {
		t.Fatalf("Expected %s is different than serialized %s", expectedString, serializedString)
	}
}

func TestJSONSerializerWithNull(t *testing.T) {

	var m = make(map[int]int, 0)
	payload, err := NewJSONMarshaller().Serialize(m)

	if err == nil {
		t.Fatalf("Should have failed! %s", payload)
	}
}

func TestJsonDeserializer(t *testing.T) {

	location, err := time.LoadLocation("Europe/Athens")

	if err != nil {
		t.Fatalf("Error getting location!")
	}

	expected := TestData{
		Version:   1,
		Name:      "Joe",
		Balance:   12.99,
		BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
	}

	actualData := `{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00"}`
	var actual TestData

	err = NewJSONMarshaller().Deserialize(actualData, &actual)

	if expected.Version != actual.Version {
		t.Fatalf("Version Expected: %s Actual: %s", expected.Version, actual.Version)
	}

	if expected.Name != actual.Name {
		t.Fatalf("Name Expected: %s Actual: %s", expected.Name, actual.Name)
	}

	if expected.Balance != actual.Balance {
		t.Fatalf("Balance Expected: %d Actual: %d", expected.Balance, actual.Balance)
	}

	if !expected.BirthDate.Equal(actual.BirthDate) {
		t.Fatalf("BirthDate Expected: %s Actual: %s", expected.BirthDate, actual.BirthDate)
	}
}

func TestJsonDeserializerError(t *testing.T) {

	actualData := `{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00------"}`

	var actual TestData

	err := NewJSONMarshaller().Deserialize(actualData, &actual)

	if err == nil {
		t.Fatal("Should have raised a error")
	}
}

func TestJsonDeserializerWrongTypeError(t *testing.T) {

	var actual TestData

	err := NewJSONMarshaller().Deserialize(123, &actual)

	if err == nil {
		t.Fatal("Should have raised a error")
	}
}

func BenchmarkJSONSerializer(b *testing.B) {

	var sert = NewJSONMarshaller()

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
