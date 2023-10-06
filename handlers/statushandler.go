package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

func Status(w http.ResponseWriter, req *http.Request) {
	statusTemplate := template.Must(template.ParseFiles("handlers/status.html"))

	data := ReadData()

	err := statusTemplate.Execute(w, data)
	fmt.Fprint(w, err)
}

func Rerun(w http.ResponseWriter, req *http.Request) {

	rerunTemplate := template.Must(template.ParseFiles("handlers/rerun.html"))

	params, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		fmt.Fprint(w, err)
	}

	if value, hasParam := params["data"]; hasParam {
		if len(value) == 0 || value[0] == "" {
			err := rerunTemplate.Execute(w, "Requires paramater data please, cant be empty yo")
			fmt.Fprint(w, err)
			return
		}

		err := ReindexData(strings.ToLower(value[0]))

		if err != nil {
			err := rerunTemplate.Execute(w, err)
			fmt.Fprint(w, err)
			return
		}

		err = rerunTemplate.Execute(w, "Started indexing of data "+value[0])
		fmt.Fprint(w, err)
		return
	}

	err = rerunTemplate.Execute(w, "Param data cannot be empty, please provide parameter :(")
	fmt.Fprint(w, err)
}
