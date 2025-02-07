package calendar

import (
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

// Represent a span of time. included beetween [Start, End]
type Span struct {
	Start time.Time
	End   time.Time
}

func (s Span) overlapS(s2 Span) bool {
	return (s.Start.Compare(s2.Start) >= 0 && s.Start.Compare(s2.End) <= 0) || (s.End.Compare(s2.Start) >= 0 && s.Start.Compare(s2.End) <= 0)
}

func (s Span) Overlap(t time.Time) bool {
	return (t.After(s.Start) && t.Before(s.End) || t.Equal(s.Start) || t.Equal(s.End))
}

func (s Span) fuse(s2 Span) Span {
	out := Span{}
	if s.Start.Compare(s2.Start) <= 0 {
		out.Start = s.Start
	} else {
		out.Start = s2.Start
	}
	if s.End.Compare(s2.End) >= 0 {
		out.End = s.End
	} else {
		out.End = s2.End
	}
	return out
}

// store the occupancy spans by room name.
var schedule = make(map[string][]Span)

// for each event, check if the event is in an 3IA room, if it is,
// add the span to the schedule.
func ExtractSchedule() map[string][]Span {
	events := (ParseData((GetDataFromUNI())))
	for _, event := range events {
		if strings.HasPrefix(event.Location, "E") || strings.Contains(event.Location, "AMPHI HERBRAND") {
			event_Span := Span{event.Span.Start, event.Span.End}
			rooms := strings.Split(event.Location, "\\,")
			for _, room := range rooms {
				room = strings.TrimSpace(room)
				if utf8.RuneCountInString(room) == 3 || strings.HasPrefix(room, "E06") || strings.HasPrefix(room, "AMPHI") {
					schedule[room] = append(schedule[room], event_Span)
				}
			}
		}
	}
	// Sort by room name.
	for _, list := range schedule {
		sort.Slice(list, func(i, j int) bool {
			return list[i].Start.Compare(list[j].Start) <= 0
		})
	}

	return schedule
}
