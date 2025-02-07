package calendar

import (
	"io"
	"net/http"
	"strings"
	"time"
)

// Array of links for every calendar of the UFR Science, for the 2nd semester of 2024-2025
var links = [...]string{
	"https://aderead.univ-orleans.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?data=823645bb1e00905432a7c3c4434de87d96005fcb2911a7e5aa39ea9f9d208e417a99563a9aeb0376b7abd9810b6512156b52ef39550f73b83da7d3c05cc18657b20901fdebde31b831fa42185d6ea6ea3d0222416033db799d896f8feba7d694,1",
	"https://aderead.univ-orleans.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?data=97ac4bca4504437facf44b5a224533bc515d81d140b1cee98fdd5b8fa5a97792fe70a3681cc863b03bf4b6693e7b74c211327cffd07a0ec30cccb7a8f8ca5de7f29616d40e5c26f5d508daf7a330426fb26d8bdb6e18cff95495071c8c388aa070fe89da822070bc9293ef6202b8fb2c,1",
	"https://aderead.univ-orleans.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?data=f0f9a8b8b255476f593e88f7644107294f69f29f39ec7eb92a13f98c8408a65f314e0a5811d856d22a864e20e0cdd81befac0955bd76f5f19e8a913832d0acb434c26ee04bf3a1cb9fabb2f2f64f3113555914ef806a9681e4d4dc7eec478d5cb27ee4db933f95149e7fbe9191670ba25fde7ed30f05dc2dd5c3c476c452580a,1",
	"https://aderead.univ-orleans.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?data=80ded1ba6dafb36ee5ac3837cff68d332cac53ed2d0b20b6fdc92c4259f549f49d4a05f9d15acf91ab80937163ea92214592592d3b282f749c9a606b710f264f6250ba3fea2e12caebbcd166cfe884766459ed9756bb6c8968c42ed82dc0671ed7cf1a57d0131369f66861c6913cc45a166c54e36382c1aa3eb0ff5cb8980cdb,1",
}

// Download every event for the whole semester
func GetDataFromUNI() []byte {
	var data []byte

	for _, link := range links {

		// get the files on aderead
		resp, err := http.Get(link)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// read the data
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		data = append(data, bytes...)
	}
	return data
}

// Parse the raw data (.ics file) to usable slice of events.
func ParseData(data_bytes []byte) []Event {
	data_string := string(data_bytes)
	blocks := strings.Split(data_string, "BEGIN:VEVENT")

	var events []Event

	for _, block := range blocks {
		if strings.HasPrefix(strings.TrimSpace(block), "DTSTAMP") {
			lines := strings.Split(block, "\n")
			// extract location
			location := strings.TrimPrefix(lines[5], "LOCATION:")

			// extract timestamps
			start, err := time.Parse("DTSTART:20060102T150405Z\x0d", lines[2])
			if err != nil {
				panic(err)
			}
			end, err := time.Parse("DTEND:20060102T150405Z\x0d", lines[3])
			if err != nil {
				panic(err)
			}
			events = append(events, Event{location, Span{start, end}})
		}
	}
	return events
}
