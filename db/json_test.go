package db

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/ake-persson/mapslice-json"
)

func TestMain(t *testing.T) {
	ms := mapslice.MapSlice{
		mapslice.MapItem{Key: "abc", Value: 123},
		mapslice.MapItem{Key: "def", Value: 456},
		mapslice.MapItem{Key: "ghi", Value: 789},
	}

	b, err := json.Marshal(ms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	ms = mapslice.MapSlice{}
	if err := json.Unmarshal(b, &ms); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ms)
}
