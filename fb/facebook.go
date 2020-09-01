package facebookgraph

import (
	fb "github.com/huandu/facebook/v2"
)

type Facebook struct {
	session *fb.Session
}

func NewFacebook(session *fb.Session) *Facebook {
	return &Facebook{session}
}

/*
type PagingCursor struct {
	Before string `mapstructure:"before"`
	After  string `mapstructure:"after"`
}

type Paging struct {
	Cursors  PagingCursor `mapstructure:"cursors"`
	Previous string       `mapstructure:"previous"`
	Next     string       `mapstructure:"next"`
}
*/
