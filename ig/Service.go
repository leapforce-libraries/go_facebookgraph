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
