package facebookgraph

type ErrorResponse struct {
	Error Error `mapstructure:"error"`
}

type Error struct {
	Message      string `mapstructure:"message"`
	Type         string `mapstructure:"type"`
	Code         int64  `mapstructure:"code"`
	ErrorSubcode int64  `mapstructure:"error_subcode"`
	FBTraceID    string `mapstructure:"fbtrace_id"`
}
