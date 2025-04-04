package bot

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func IsUserTrusted(userID int) bool {
	// Read the file containing trusted user IDs
	file, err := os.Open("./notes/users.txt")
	if err != nil {
		log.Printf("Error opening trusted users file: %v", err)
		return false
	}
	defer file.Close()

	// Scan through each line in the file
	var trustedUserIDs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trustedUserIDs = append(trustedUserIDs, scanner.Text())
	}

	// Check if the user ID is in the list of trusted users
	for _, id := range trustedUserIDs {
		if fmt.Sprintf("%d", userID) == id {
			return true
		}
	}

	// If the user is not found in the list, they are not trusted
	return false
}
