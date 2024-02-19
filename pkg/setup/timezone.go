package setup

import (
	"log"
	"time"
)

// example : "America/Sao_Paulo" -> -03
func TimeZone(expectedTimeZone string) {

	// OS timezone
	//location, _ := time.LoadLocation("Local")
	current, _ := time.Now().Zone()
	if current != expectedTimeZone {
		log.Default().Fatalf("Current Timezone %s is not expected %s . Verify OS env 'TZ' or change 'TZ_BOT' value", current, expectedTimeZone)
	}
	log.Default().Println("Current timezone is ok.")
}
