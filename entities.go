package main

type Webhook struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}

type Result struct {
	UpdateId int     `json:"update_id"`
	Msg      Message `json:"message"`
}
type TelegramUsers struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

type Message struct {
	MessageId  int64  `json:"message_id"`
	FromDetail From   `json:"from"`
	ChatDetail Chat   `json:"chat"`
	Date       string `json:"date"`
	Text       string `json:"text"`
	Entities   Entity `json:"entities"`
}

type From struct {
	TelegramUsers
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	TelegramUsers
	Type string `json:"type"`
}

type Entity struct {
	Offset int16  `json:"offset"`
	Length int16  `json:"length"`
	Type   string `json:"type"`
}

type TelegramReplyMessage struct {
	ChatId   int64  `json:"chatId"`
	UserName string `json:"username"`
	Text     string `json:"text"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
