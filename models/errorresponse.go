package facebookgraph

type ErrorResponse struct {
	Error Error `mapstructure:"error"`
}

type Error struct {
	Message   string `mapstructure:"message"`
	Type      string `mapstructure:"type"`
	Code      int64  `mapstructure:"code"`
	FBTraceID string `mapstructure:"fbtrace_id"`
}
