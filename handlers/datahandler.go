package handlers

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var data DataContainer

func InitData() {
	data = DataContainer{
		data: map[string]*Data{"yeetmachine": {
			Version:       13,
			Status:        "FAILED",
			StatusMessage: "failed due to something wrong",
			LastSeen:      time.Now().Unix()},
			"superdata": {
				Version:       13,
				Status:        "FAILED",
				StatusMessage: "failed due to something wrong",
				LastSeen:      time.Now().Unix()},
			"cooldata": {
				Version:       13,
				Status:        "FAILED",
				StatusMessage: "failed due to something wrong",
				LastSeen:      time.Now().Unix()},
		}}
}

type DataContainer struct {
	mu   sync.Mutex
	data map[string]*Data
}

type Data struct {
	Version       int
	Status        string
	StatusMessage string
	LastSeen      int64
}

func (data *Data) HumanLastSeen() string {
	t := time.Unix(data.LastSeen, 0)
	return fmt.Sprint(t.UTC())
}

func ReadData() map[string]*Data {
	return data.data
}

func ReindexData(name string) error {

	data.mu.Lock()
	defer data.mu.Unlock()

	if value, hasData := data.data[name]; hasData {

		if value.Status != "NOTINDEXING" && value.Status != "FAILED" {
			return errors.New("Work already started on " + name)
		}

		value.Status = "INDEXING"
		data.data[name] = value
		return nil
	}

	return errors.New("no data named " + name)
}
