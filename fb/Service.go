package facebookgraph

import (
	fb "github.com/huandu/facebook/v2"
)

type Service struct {
	session *fb.Session
}

func NewService(session *fb.Session) *Service {
	return &Service{session}
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
