package json_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SearchResult struct {
	Source            string `json:"source,omitempty"`
	Sha1              string `json:"sha1,omitempty"`
	AdaptationId      string `json:"adaptationId,omitempty"`
	AdaptationVersion string `json:"adaptationVersion,omitempty"`
}

type SearchResultV2 struct {
	NeSwId            string `json:"neSwId,omitempty"`
	Source            string `json:"source,omitempty"`
	Sha1              string `json:"sha1,omitempty"`
	AdaptationId      string `json:"adaptationId,omitempty"`
	AdaptationVersion string `json:"adaptationVersion,omitempty"`
}

func TestMarshal(t *testing.T) {
	fmt.Println("start test marshal")
	// t.Skip()

	// mashal to json text
	r := SearchResult{
		Source:            "mysrc",
		Sha1:              "asdfsadf7xl2q0x",
		AdaptationId:      "NOKLTE",
		AdaptationVersion: "LTE19A",
	}
	byteJson, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("marshaled to json:\t%s\n", string(byteJson))
	expectJsonText := `{"source":"mysrc","sha1":"asdfsadf7xl2q0x","adaptationId":"NOKLTE","adaptationVersion":"LTE19A"}`
	assert.Equal(t, expectJsonText, string(byteJson))

	// unmashal from json
	// jsonText := "{\"neSwId\":\"5G19ACLA_4\",\"source\":\"mysrc00\",\"sha1\":\"asdfsadf7xl2q0x888\",\"adaptationId\":\"NOKLTE\",\"adaptationVersion\":\"LTE19A\"}"
	jsonText :=
		`{
			"neSwId": "5G19ACLA_4",
			"source": "mysrc00",
			"sha1": "asdfsadf7xl2q0x888",
			"adaptationId": "NOKLTE",
			"adaptationVersion": "LTE19A"
		}`

	var sr SearchResult
	err = json.Unmarshal([]byte(jsonText), &sr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("object from json:\t%#v\n", sr)
	assert.NotNil(t, sr)
	assert.Equal(t, "mysrc00", sr.Source)
	assert.Equal(t, "asdfsadf7xl2q0x888", sr.Sha1)
	assert.Equal(t, "NOKLTE", sr.AdaptationId)

	var v2 SearchResultV2
	err = json.Unmarshal([]byte(jsonText), &v2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("object from json:\t%#v\n", v2)
	assert.NotNil(t, v2)
	assert.Equal(t, "5G19ACLA_4", v2.NeSwId)
	assert.Equal(t, "mysrc00", v2.Source)
}
