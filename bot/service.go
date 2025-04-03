package bot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func SendMessagesInBatches(c tele.Context, message string, total int) error {
	const batchSize = 1
	const delayBetweenBatches = 5

	for i := 0; i < total; i += batchSize {
		remaining := total - i
		count := batchSize
		if remaining < batchSize {
			count = remaining
		}

		for j := 0; j < count; j++ {
			if err := c.Send(message); err != nil {
				log.Printf("Failed to send message #%d: %v", i+j+1, err)
				continue
			}
		}

		if i+batchSize < total {
			time.Sleep(time.Duration(delayBetweenBatches) * time.Second)
		}
	}

	return nil
}
