package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
	"net/http"
	"time"
)

//go:embed karmakut.ico
var icon []byte

//go:embed redx.ico
var redx []byte

const urlQ = "https://api.battlemetrics.com/servers/11522563/player-count-history?start=%s&stop=%s&resolution=raw"

type CountData struct {
	Data []struct {
		Type       string `json:"type"`
		Attributes struct {
			Timestamp time.Time `json:"timestamp"`
			Value     int       `json:"value"`
		} `json:"attributes"`
	} `json:"data"`
}

var playerCount int

func GetPlayerCount() CountData {
	data := CountData{}

	current := time.Now().UTC().Format("2006-01-02T15:06:00Z")
	prev := time.Now().UTC().Add(-time.Hour).Format("2006-01-02T15:06:00Z")

	req, _ := http.NewRequest("GET", fmt.Sprintf(urlQ, prev, current), nil)

	req.Header.Add("accept", "application/vnd.api+json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")
	req.AddCookie(&http.Cookie{Name: "_pbjs_userid_consent_data", Value: "3524755945110770"})
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	buf := bytes.NewBuffer(nil)

	buf.ReadFrom(res.Body)

	json.Unmarshal(buf.Bytes(), &data)

	return data

}

func onReady() {
	systray.SetTitle("Karmakut Liberation Player Count")
	systray.SetTooltip("Current Player Count:" + fmt.Sprintf("%d", playerCount))
	systray.SetIcon(icon)

	mGetNow := systray.AddMenuItem("Get Now", "Get the current player count")
	go func() {
		for {
			<-mGetNow.ClickedCh
			data := GetPlayerCount()
			playerCount = data.Data[0].Attributes.Value
			systray.SetTooltip("Current Liberation Population:" + fmt.Sprintf("%d", playerCount))
			dialog.Message("Current Player Count: %d", playerCount).Title("Karmakut Liberation Player Count").Info()
		}
	}()
	mGetNow.SetIcon(icon)

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
	mQuit.SetIcon(redx)

}

func UpdatePlayerCount() {
	for {
		data := GetPlayerCount()
		playerCount = data.Data[0].Attributes.Value
		systray.SetTooltip("Current Liberation Population:" + fmt.Sprintf("%d", playerCount))
		if playerCount > 30 {
			dialog.Message("Current Player Count: %d", playerCount).Title("Karmakut Liberation Player Count").Info()
		}
		time.Sleep(time.Minute)
	}

}

func main() {
	go UpdatePlayerCount()
	systray.Run(onReady, nil)

}
