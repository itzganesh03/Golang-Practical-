package main

import "gorm.io/gorm"

// pascal case use User for Globally Access.
type User struct {
	gorm.Model
	// Cid          string `json:"id"`
	Cname        string `json:"name"`
	Clocation    string `json:"location"`
	Cpan         string `json:"pan"`
	Caddress     string `json:"address"`
	CmoNumber    string `json:"number"`
	Csex         string `json:"sex"`
	Cnationality string `json:"nationality"`
}
