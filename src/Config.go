package main

type DatabaseConfiguration struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Namedb   string `json:"namedb"`
}

type DataConfiguration struct {
	Database DatabaseConfiguration `json:"database"`
}
