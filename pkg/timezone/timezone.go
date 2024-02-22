package timezone

import (
	"log"
	"os"
	"time"
)

// Load example : check(America/Sao_Paulo)
// Load Timezone configuration
func Load(expectedTimeZone string) {
	current := os.Getenv("TZ")
	if current != expectedTimeZone {
		location, err := time.LoadLocation(expectedTimeZone)
		time.Local = location
		if err != nil {
			log.Default().Fatalf("TZ: Current Timezone '%s' is not expected '%s'. Can't change it manually. Verify OS env 'TZ' or change 'TZ_APP' value", current, expectedTimeZone)
		}
	}
}
