package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type programm struct {
	Data []Programm `json:"epg_data"`
}

var fullXML = ""

func main() {
	go func() {
		// ticker := time.NewTicker(24 * time.Hour)
		// for range ticker.C {
		buildXML()
		// }
	}()
	http.HandleFunc("/", HomeRouterHandler)  // установим роутер
	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// HomeRouterHandler comment
func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fullXML)
}
func buildXML() {
	var wg sync.WaitGroup
	var channelXML = ""
	var programmXML = ""
	syncProgrammXML := make(chan string)
	client := http.Client{}
	resp, err := client.Get("http://ott.watch/api/channel_now")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	channels := map[string]Channel{}
	json.NewDecoder(resp.Body).Decode(&channels)
	if err != nil {
		fmt.Println(err)
		return
	}
	wg.Add(len(channels))
	for _, element := range channels {
		element.Imgsrc = Img{"http://ott.watch/images/" + element.Img}
		marshalled, _ := xml.MarshalIndent(element, "", "   ")
		channelXML += string(marshalled) + "\n"
		go func(element Channel, syncProgrammXML chan string) {
			defer wg.Done()
			resp, err := client.Get("http://ott.watch/api/channel/" + element.ChID)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			programm := programm{}
			json.NewDecoder(resp.Body).Decode(&programm)
			if len(programm.Data) == 0 {
				fmt.Println(element.ChannelName)
				fmt.Println(element.ChID)
			}
			marshalled, _ := xml.MarshalIndent(programm.Data, "", "   ")
			syncProgrammXML <- string(marshalled)
		}(element, syncProgrammXML)
	}
	go func(text *string, syncProgrammXML chan string) {
		for value := range syncProgrammXML {
			programmXML += value + "\n"
		}
	}(&programmXML, syncProgrammXML)
	wg.Wait()
	close(syncProgrammXML)
	fullXML = channelXML + programmXML
}
