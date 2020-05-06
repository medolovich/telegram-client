package client

import "net/http"

type TelegramClient struct {
	httpClient    *http.Client
	updatesOffset int
	botAPIToken   string
}

type response struct {
	Ok          bool        `json:"ok"`
	Description string      `json:"description"`
	Result      interface{} `json:"result"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Chat struct {
	ID                          int    `json:"id"`
	Type                        string `json:"type"`
	Username                    string `json:"username"`
	FirstName                   string `json:"first_name"`
	LastName                    string `json:"last_name"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

type Message struct {
	MessageID      int      `json:"message_id"`
	From           User     `json:"from"`
	Date           int      `json:"date"`
	Chat           Chat     `json:"chat"`
	ForwardFrom    User     `json:"forward_from"`
	ForwardDate    int      `json:"forward_date"`
	ReplyToMessage *Message `json:"reply_to_message"`
	Text           string   `json:"text"`
}

type Update struct {
	UpdateID    int     `json:"update_id"`
	Message     Message `json:"message"`
	ChannelPost Message `json:"channel_post"`
}

type Updates []Update
