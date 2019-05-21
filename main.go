package main

import (
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/micmonay/keybd_event"
)

func main() {

	time.Sleep(2 * time.Second)
	println(" Starting...")

	hashtag := "#StardewValleyCommunityGamePlay"
	consumerKey := "Adp7EPVRLHlZA2SQwH0MxukzZ"
	consumerSecret := "BFXhXUXNnDsedHggXWDWah0pU6b1k4EqM59F2yLHlEKe0gcomj"
	tokenStr := "54723793-3AbtX0lTCzZ9H9X26jKLfpbFah9viNHqP32v7G6TM"
	tokenSecret := "XHMXDotstZndYA6eVUPUVq3mPsP5ecCLk4IsrqlUEDAoV"

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(tokenStr, tokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	params := &twitter.StreamFilterParams{
		Track:         []string{hashtag},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		panic(err)
	}

	for message := range stream.Messages {

		tweet := message.(*twitter.Tweet)
		messageText := strings.Replace(tweet.Text, hashtag, "", -1)

		println(messageText)
		whenKeyThenAction(messageText, "A", &kb, keybd_event.VK_A)
		whenKeyThenAction(messageText, "W", &kb, keybd_event.VK_W)
		whenKeyThenAction(messageText, "D", &kb, keybd_event.VK_D)
		whenKeyThenAction(messageText, "S", &kb, keybd_event.VK_S)
		whenKeyThenAction(messageText, "C", &kb, keybd_event.VK_C)
		whenKeyThenAction(messageText, "X", &kb, keybd_event.VK_X)
		whenKeyThenAction(messageText, "E", &kb, keybd_event.VK_E)
		whenKeyThenAction(messageText, "F", &kb, keybd_event.VK_F)
		whenKeyThenAction(messageText, "M", &kb, keybd_event.VK_M)
		whenKeyThenAction(messageText, "1", &kb, keybd_event.VK_1)
		whenKeyThenAction(messageText, "2", &kb, keybd_event.VK_2)
		whenKeyThenAction(messageText, "3", &kb, keybd_event.VK_3)
		whenKeyThenAction(messageText, "4", &kb, keybd_event.VK_4)
		whenKeyThenAction(messageText, "5", &kb, keybd_event.VK_5)
		whenKeyThenAction(messageText, "6", &kb, keybd_event.VK_6)
		whenKeyThenAction(messageText, "7", &kb, keybd_event.VK_7)
		whenKeyThenAction(messageText, "8", &kb, keybd_event.VK_8)
		whenKeyThenAction(messageText, "9", &kb, keybd_event.VK_9)
		whenKeyThenAction(messageText, "0", &kb, keybd_event.VK_0)
	}
}

func whenKeyThenAction(messageText string, char string, kb *keybd_event.KeyBonding, key int) {
	if checkKey(messageText, char) {
		println(" - Pressing ", char)
		press(kb, key)
	}
}

func checkKey(text string, char string) bool {
	return strings.Contains(text, char)
}

func press(kb *keybd_event.KeyBonding, key int) {
	kb.SetKeys(key)
	kb.Launching()
	println("")
}
