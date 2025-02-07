package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"stationkeep/calendar"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// some styling for cli use.
var invalidstyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("246"))
var validstyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("28"))

func main() {
	// Change this value to the date you want to check.
	tocheck := time.Now()

	schedule := calendar.ExtractSchedule()
	rooms_string := ""
	var today = make(map[string]string, 30)
	for room, list := range schedule {
		occupied := false
		for _, span := range list {
			if span.Overlap(tocheck) {
				occupied = true
				break
			}
		}
		h, m, _ := tocheck.Clock()

		if occupied {
			rooms_string += invalidstyle.Render(fmt.Sprintf("\n✖ %s : occupée à %dh%d.", room, h, m))
		} else {
			rooms_string = validstyle.Render(fmt.Sprintf("\n✔ %s : libre à %dh%d.", room, h, m)) + rooms_string
		}
		s := ""
		s += fmt.Sprintf("%s : \n", room)
		for _, span := range list {
			sy, sm, sd := span.Start.Date()
			ty, tm, td := tocheck.Date()
			if sy == ty && sm == tm && sd == td {
				s += fmt.Sprintf("- %02dh%02d - %02dh%02d\n", span.Start.Hour()+1, span.Start.Minute(), span.End.Hour()+1, span.End.Minute())
			}
		}
		today[room] = s

	}
	fmt.Println(rooms_string)
	keys := make([]string, 0, len(today))
	for k := range today {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	f, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	for _, room := range keys {
		fmt.Fprintf(w, "%s\n", today[room])
	}
	w.Flush()
}
