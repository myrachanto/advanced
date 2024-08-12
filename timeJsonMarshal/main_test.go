package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestUnmarshalJSON(t *testing.T) {
	jsonData := `{"name": "Tony Stark","age":50, "address":"stark industries", "dob": "2024-08-12"}`
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %v", err)
	}

	expectedDob := CustomDate{Time: time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC)}
	if !user.Dob.Equal(expectedDob.Time) {
		t.Errorf("Expected DOB %v, got %v", expectedDob.Time, user.Dob.Time)
	}

	expectedName := "Tony Stark"
	if user.Name != expectedName {
		t.Errorf("Expected Name %v, got %v", expectedName, user.Name)
	}

	expectedAge := 50
	if user.Age != expectedAge {
		t.Errorf("Expected Age %v, got %v", expectedAge, user.Age)
	}

	expectedAddress := "stark industries"
	if user.Address != expectedAddress {
		t.Errorf("Expected Address %v, got %v", expectedAddress, user.Address)
	}
}

func TestMarshalJSON(t *testing.T) {
	user := User{
		Name:    "Tony Stark",
		Age:     50,
		Address: "stark industries",
		Dob:     CustomDate{Time: time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC)},
	}

	data, err := json.Marshal(&user)
	if err != nil {
		t.Errorf("Error marshalling JSON: %v", err)
	}

	expectedJSON := `{"name":"Tony Stark","age":50,"address":"stark industries","dob":"2024-08-12T00:00:00Z"}`
	if string(data) != expectedJSON {
		t.Errorf("Expected JSON %v, got %v", expectedJSON, string(data))
	}
}

func TestInvalidDateFormat(t *testing.T) {
	jsonData := `{"name": "Tony Stark","age":50, "address":"stark industries", "dob": "12-08-2024"}`
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err == nil {
		t.Error("Expected error unmarshalling JSON due to invalid date format, but got none")
	}
}
