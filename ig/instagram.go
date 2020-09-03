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
