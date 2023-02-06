package apns

import (
	"github.com/sideshow/apns2"
	"log"
	"notification/config"
	. "notification/models"
	"notification/push/base"
	"strings"
)

type Sender struct {
	base.Sender
}

func (s *Sender) Send() {
	for _, token := range s.Tokens {
		res, err := client.Push(&apns2.Notification{
			DeviceToken: token,
			Topic:       config.Config.IOSPackageName,
			Payload:     constructPayload(s.Message),
		})
		if err != nil || res == nil {
			log.Printf("APNS push error: %s\n", err)
			return
		}
		if res.StatusCode != 200 {
			log.Printf("APNS push failed: %d %s\n",
				res.StatusCode, res.Reason)
			if strings.Contains(res.Reason, "DeviceToken") {
				// device token is expired, remove it from database
				s.ExpiredTokens = append(s.ExpiredTokens, token)
			}
		}
	}
}

func constructPayload(message *Message) any {
	return Map{"aps": Map{"alert": Map{
		"title":    message.Title,
		"subtitle": message.Description,
		"body":     message.Data,
	}}}
}
