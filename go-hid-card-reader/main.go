package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type Config struct {
	DataFileLoc string `json:"datafileloc"`
}

type WorkRecord struct { //define the structure for the work hour.
	MemberCardNumber string
	WorkType         string
	WorkTypeOther    string
	HoursWorked      string
	PictureLoc       string
	DateOfWork       string
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func open(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func thanks(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("thankyou.html")
	t.Execute(w, nil)
}

func addhours(w http.ResponseWriter, r *http.Request) {
	config, _ := LoadConfiguration("config.json")
	f, err := os.OpenFile(config.DataFileLoc, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()

	wr := new(WorkRecord)
	wr.MemberCardNumber = r.FormValue("membercardnumber")
	wr.WorkType = r.FormValue("worktype")
	wr.HoursWorked = r.FormValue("hoursworked")
	wr.DateOfWork = time.Now().Format(time.RFC850)
	b, err := json.Marshal(wr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	f.Write(b)
	f.Close()
	http.Redirect(w, r, "/thankyou.html", http.StatusFound)
}

func main() {
	fmt.Println("Starting the Application...")
	config, _ := LoadConfiguration("config.json")
	fmt.Println(config.DataFileLoc)
	http.HandleFunc("/addhours", addhours)
	http.HandleFunc("/", open)
	http.HandleFunc("/thankyou.html", thanks)
	http.ListenAndServe(":8080", nil)
}
