package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func Status(w http.ResponseWriter, req *http.Request) {
	statusTemplate := template.Must(template.ParseFiles("handlers/status.html"))

	data := map[string]*Data{
		"yeetmachine": {
			Version:       13,
			Status:        "FAILED",
			StatusMessage: "failed due to something wrong",
			LastSeen:      time.Now().Unix()},
		"SuperData": {
			Version:       13,
			Status:        "FAILED",
			StatusMessage: "failed due to something wrong",
			LastSeen:      time.Now().Unix()},
		"CoolData": {
			Version:       13,
			Status:        "FAILED",
			StatusMessage: "failed due to something wrong",
			LastSeen:      time.Now().Unix()},
	}

	err := statusTemplate.Execute(w, data)
	fmt.Fprint(w, err)
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
