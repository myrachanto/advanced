package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const customDateFormat = `"2006-01-02"`

type CustomDate struct {
	time.Time
}

func (c *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	parsed, err := time.Parse(customDateFormat, str)
	if err != nil {
		return err
	}
	c.Time = parsed
	return nil
}
func (c *CustomDate) Marshal(b []byte) ([]byte, error) {
	return []byte(c.Format(customDateFormat)), nil
}

type User struct {
	Name    string     `json:"name"`
	Age     int        `json:"age"`
	Address string     `json:"address"`
	Dob     CustomDate `json:"dob"`
}

func main() {
	jsonData := `{"name": "Tony Stark","age":50, "address":"stark industries", "dob": "2024-08-12"}`
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("Parsed User: %+v\n", user)

	out, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("Marshalled JSON:", string(out))
}
