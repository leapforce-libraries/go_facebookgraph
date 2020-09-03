package facebookgraph

type Attachments struct {
	Data []Attachment `mapstructure:"data"`
}

type Attachment struct {
	Description string               `mapstructure:"description"`
	Media       AttachmentMediaImage `mapstructure:"media"`
	Target      AttachmentTarget     `mapstructure:"target"`
	Title       string               `mapstructure:"title"`
	Type        string               `mapstructure:"type"`
	URL         string               `mapstructure:"url"`
}

type AttachmentMedia struct {
	Image AttachmentMediaImage `mapstructure:"image"`
}

type AttachmentMediaImage struct {
	Height int64  `mapstructure:"height"`
	Source string `mapstructure:"src"`
	Width  string `mapstructure:"width"`
}

type AttachmentTarget struct {
	URL string `mapstructure:"url"`
}
