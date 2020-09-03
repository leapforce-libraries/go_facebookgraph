package models

type PagingCursor struct {
	Before string `mapstructure:"before"`
	After  string `mapstructure:"after"`
}

type Paging struct {
	Cursors  PagingCursor `mapstructure:"cursors"`
	Previous string       `mapstructure:"previous"`
	Next     string       `mapstructure:"next"`
}
