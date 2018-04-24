package holidayjp

import (
	"log"
	"time"

	"github.com/tidwall/gjson"
)

//Holiday returns if today is a holiday/weekend or not
//returns boolean->is a holiday & string-> holiday description
func Holiday(t time.Time) (bool, string) {
	date := t.Format("Mon")

	//check first if weekeend
	if map[string]bool{
		"Sat": true,
		"Sun": true,
	}[date] {
		return true, "weekend"
	}

	//Check japanese holidays
	calendar := t.Format("2006-01-02")

	//get the japanese holiday
	json, err := Asset("asset/1970-2050.json")
	if err != nil {
		// Asset was not found.
		log.Fatal(err)
	}
	value := gjson.Get(string(json), calendar)
	if value.Exists() {
		desc := value.String()

		//get the english version
		json, err := Asset("asset/en.json")
		if err != nil {
			// Asset was not found.
			log.Fatal(err)
		}
		value = gjson.Get(string(json), desc)
		if value.Exists() {
			desc = desc + "- " + value.String()
		}

		return true, desc

	}

	return false, "not a holiday"
}
