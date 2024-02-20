package setup

import (
	"log"
	"time"
)

// TimeZone example : America/Sao_Paulo -> TimeZone("-03")
// Check for misconfiguration with Timezone
func TimeZone(expectedTimeZone string) {
	current, _ := time.Now().Zone()
	if current != expectedTimeZone {
		log.Default().Fatalf("Current Timezone %s is not expected %s . Verify OS env 'TZ' or change 'TZ_BOT' value", current, expectedTimeZone)
	}
	log.Default().Println("Current timezone is ok.")
}
