package main

import "time"

type T struct {
	Data []struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			Start     time.Time  `json:"start"`
			Stop      *time.Time `json:"stop"`
			FirstTime bool       `json:"firstTime"`
			Name      string     `json:"name"`
			Private   bool       `json:"private"`
		} `json:"attributes"`
		Relationships struct {
			Server struct {
				Data struct {
					Type string `json:"type"`
					Id   string `json:"id"`
				} `json:"data"`
			} `json:"server"`
			Player struct {
				Data struct {
					Type string `json:"type"`
					Id   string `json:"id"`
				} `json:"data"`
			} `json:"player"`
			Identifiers struct {
				Data []struct {
					Type string `json:"type"`
					Id   string `json:"id"`
				} `json:"data"`
			} `json:"identifiers"`
		} `json:"relationships"`
	} `json:"data"`
	Included []interface{} `json:"included"`
}
