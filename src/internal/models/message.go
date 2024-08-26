package models

type LinePayload struct {
	To       string    `json:"to"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Type     string  `json:"type"`
	AltText  string  `json:"altText"`
	Contents Content `json:"contents"`
}

type Body struct {
	Type       string    `json:"type"`
	Layout     string    `json:"layout"`
	Contents   []Content `json:"contents"`
	Spacing    string    `json:"spacing"`
	PaddingAll string    `json:"paddingAll"`
}

type Content struct {
	Type       string    `json:"type"`
	Body       *Body     `json:"body,omitempty"`
	Layout     string    `json:"layout,omitempty"`
	Contents   []Content `json:"contents,omitempty"`
	URL        string    `json:"url,omitempty"`
	Size       string    `json:"size,omitempty"`
	Flex       int64     `json:"flex,omitempty"`
	AspectMode string    `json:"aspectMode,omitempty"`
	Text       string    `json:"text,omitempty"`
	Weight     string    `json:"weight,omitempty"`
	Align      string    `json:"align,omitempty"`
	Wrap       bool      `json:"wrap,omitempty"`
	Gravity    string    `json:"gravity,omitempty"`
	Spacing    string    `json:"spacing,omitempty"`
	PaddingAll string    `json:"paddingAll,omitempty"`
	Margin     string    `json:"margin,omitempty"`
}
