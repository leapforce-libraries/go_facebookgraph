package facebookgraph

import (
	fb "github.com/huandu/facebook/v2"
)

type Instagram struct {
	session *fb.Session
}

func NewInstagram(session *fb.Session) *Instagram {
	return &Instagram{session}
}

type PagingCursor struct {
	Before string `mapstructure:"before"`
	After  string `mapstructure:"after"`
}

type Paging struct {
	Cursors  PagingCursor `mapstructure:"cursors"`
	Previous string       `mapstructure:"previous"`
	Next     string       `mapstructure:"next"`
}
