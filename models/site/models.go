package site

import "time"

//Setting setting
type Setting struct {
	ID        uint
	Key       string
	Val       []byte
	Flag      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Locale locale
type Locale struct {
	ID        uint
	Code      string
	Lang      string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Notice notice
type Notice struct {
	ID        uint
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//LeaveWord leave word
type LeaveWord struct {
	ID        uint
	Body      string
	CreatedAt time.Time
}

//Attachment attachment
type Attachment struct {
	ID        uint
	Name      string
	Title     string
	MediaType string
	CreatedAt time.Time
	UpdatedAt time.Time
}
