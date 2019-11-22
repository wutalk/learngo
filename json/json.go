package json

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type operationStatus struct {
	OperationID string    `json:"operationId"`
	Timestamp   time.Time `json:"timestamp"`
	Status      string    `json:"status"`
}

type OpStatus struct {
	ID          int64     `db:"id,primarykey" json:"-"`
	OperationID string    `db:"operation_id" json:"operationId"`
	Status      string    `db:"status" json:"status"`
	CreatedAt   time.Time `db:"created_at" json:"timestamp"`
}

func (s *operationStatus) String() string {
	return fmt.Sprintf("OperationID:\"%s\", Timestamp:%s, Status:\"%s\"", s.OperationID, s.Timestamp.UTC().Format(time.RFC3339), s.Status)
}

func MarshalAndUnmarshal() {
	st := operationStatus{
		OperationID: "7465fc42f22a4a87b047b1bb2430e56f",
		Timestamp:   time.Now().UTC(), /* .Format(time.RFC3339) */
		Status:      "ongoing",
	}
	fmt.Printf("original object:\t%s\n", st)
	byteJson, err := json.Marshal(st)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("marshaled to json:\t%s\n", string(byteJson))

	// unmashal from json
	jsonText := "{\"operationId\":\"7465fc42f22a4a87b047b1bb2430e56f\",\"timestamp\":\"2019-11-12T12:04:18Z\",\"status\":\"ongoing\"}"
	var opStatus operationStatus
	err = json.Unmarshal([]byte(jsonText), &opStatus)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("object from json:\t%#v\n", opStatus)

	timeStr := "2019-11-12T12:12:03Z"
	t, _ := time.Parse(time.RFC3339, timeStr)
	fmt.Println(t)
}

func MarshalAndUnmarshalDB() {
	st := OpStatus{
		OperationID: "7465fc42f22a4a87b047b1bb2430e56f",
		CreatedAt:   time.Now().UTC(), /* .Format(time.RFC3339) */
		Status:      "ongoing",
	}
	fmt.Printf("original object:\t%s\n", st)
	byteJson, err := json.Marshal(st)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("marshaled to json:\t%s\n", string(byteJson))

	// unmashal from json
	jsonText := `{"operationId":"7465fc42f22a4a87b047b1bb2430e56f","status":"ongoing","timestamp":"2019-11-22T12:09:50.088135927Z"}`
	var opStatus OpStatus
	err = json.Unmarshal([]byte(jsonText), &opStatus)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("object from json:\t%#v\n", opStatus)

	timeStr := "2019-11-12T12:12:03Z"
	t, _ := time.Parse(time.RFC3339, timeStr)
	fmt.Println(t)
}
