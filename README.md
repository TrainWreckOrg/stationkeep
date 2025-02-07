# 🚉 Stationkeep
Stationkeep is a simple go tool used to check the availability of each room in 3IA.

## Usage :
### In the terminal (like a hacker (▀̿Ĺ̯▀̿ ̿))
just run the .exe like this
```bash
./stationkeep.exe
```

or if you prefer to run from source with go installed
```bash
go run .
```
(if you want to re-build the .exe, you can run `go build`)

you then get a colored print of the availability of each room at the time you launched the command.

### Double clicking the .exe file
just double click, do not type anything in the terminal box, just wait a bit.

In both case you will get a out.txt file detailing the span of time in wich the room is **occupied**.

## Notes :
- **the out.txt times are the times during wich the room is not free**
- no data is stored on your computer, hence the long loading time each time you launch the programm
- yes i could probably make it faster, i will do it, at some point, or not
- to change the day for wich you check, modify the main.go file and run or rebuild.