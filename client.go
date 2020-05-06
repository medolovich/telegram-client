package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/proxy"
)

const (
	tlgAPIURLBase = "https://api.telegram.org/"
)

func (t *TelegramClient) tlgAPIURL() string {
	return fmt.Sprintf("%sbot%s/", tlgAPIURLBase, t.botAPIToken)
}

func (t *TelegramClient) doRequest(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := t.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if v != nil {
		resp := response{
			Result: v,
		}
		if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
			return err
		}
		if !resp.Ok {
			return errors.New(resp.Description)
		}
	}
	return nil
}

func (t *TelegramClient) GetUpdates() (Updates, error) {
	res := Updates{}
	if err := t.doRequest(fmt.Sprintf("%sgetUpdates?offset=%d", t.tlgAPIURL(), t.updatesOffset), &res); err != nil {
		return nil, err
	}

	if len(res) > 0 {
		t.updatesOffset = res[len(res)-1].UpdateID + 1
	}

	return res, nil
}

func (t *TelegramClient) SendMessage(chatID int, text string, replyToMessage *Message) (*Message, error) {
	url := fmt.Sprintf("%ssendMessage?chat_id=%d&text=%s", t.tlgAPIURL(), chatID, text)
	if replyToMessage != nil {
		url += fmt.Sprintf("&reply_to_message_id=%d", replyToMessage.MessageID)
	}
	var msg Message
	if err := t.doRequest(url, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

func New(botAPIToken string, proxyURL string, proxyLogin string, proxyPass string) (*TelegramClient, error) {
	httpClient := new(http.Client)
	if proxyURL != "" {
		var dialer proxy.Dialer
		var err error
		if proxyLogin != "" && proxyPass != "" {
			auth := proxy.Auth{
				User:     proxyLogin,
				Password: proxyPass,
			}
			log.Println(proxyURL, auth)
			dialer, err = proxy.SOCKS5("tcp", proxyURL, &auth, proxy.Direct)
		} else {
			dialer, err = proxy.SOCKS5("tcp", proxyURL, nil, proxy.Direct)
		}
		if err != nil {
			return nil, err
		}
		httpClient.Transport = &http.Transport{
			Dial: dialer.Dial,
		}
	}
	return &TelegramClient{
		httpClient:    httpClient,
		updatesOffset: -1,
		botAPIToken:   botAPIToken,
	}, nil
}
