package temp

import (
	"MTCTrueTechHack/internal/entities"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetEvents() []entities.Event {
	eventTypes := []string{
		"children",
		"concerts",
		"excursions",
		"exhibitions",
		"festivals",
		"musicals",
		"show",
		"sport",
		"standups",
		"theater",
	}

	var events []entities.Event
	for _, event := range eventTypes {
		path := fmt.Sprintf("temp/data/%v.json", event)

		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("File opening error: %v", err.Error())
			return []entities.Event{}
		}
		defer file.Close()

		var event = make(map[string]entities.Event)

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("Data read error: %v", err.Error())
			return []entities.Event{}
		}
		err = json.Unmarshal(data, &event)
		if err != nil {
			fmt.Printf("Unmarshaling error: %v", err.Error())
			return []entities.Event{}
		}

		for _, value := range event {
			events = append(events, value)
		}
	}

	return events
}
