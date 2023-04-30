package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Finger struct {
	ID         string
	PaddingTop float64
}

type SVGConfig struct {
	BackgroundColor string
	Color           string
	NodeColor       string
	HoleColor       string

	Scale float64
}

type colored struct {
	Color string
}

func (c *SVGConfig) Update(s string) error {
	split := strings.SplitN(s, "=", 2)
	if len(split) < 2 {
		return errors.New("invalid format: " + s)
	}

	key := strings.TrimLeft(split[0], "#")
	val := split[1]
	switch key {
	case "background-color":
		c.BackgroundColor = val
	case "color":
		c.Color = val
	case "node-color":
		c.NodeColor = val
	case "hole-color":
		c.HoleColor = val
	case "scale":
		c.Scale, _ = strconv.ParseFloat(val, 64)
	default:
		return errors.New("unknown key: " + key)
	}

	return nil
}

func main() {
	logger, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	defer logger.Sync()
	defer logger.Close()

	var inputFile, outputFile string
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
		outputFile = inputFile[:strings.LastIndex(inputFile, ".")] + ".svg"
	} else {
		return
	}

	chart := FingeringChart{
		BarSpace: 20,
	}
	bar := Bar{
		HorizontalNoteSpace: 30,
		VerticalNoteSpace:   3,
	}

	logger.WriteString("\r\nopening " + inputFile)
	readFile, err := os.Open(inputFile)
	if err != nil {
		logger.WriteString("\r\nerror opening file: " + err.Error())
		return
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	config := SVGConfig{}

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		// empty line = start of a new bar
		if line == "" {
			chart.AddBar(bar)
			bar = Bar{
				HorizontalNoteSpace: 30,
				VerticalNoteSpace:   3,
			}
			continue
		}

		// config values start with a '#'
		if line[0] == '#' {
			err := config.Update(line)
			if err != nil {
				logger.WriteString("\r\nerror parsing config: " + err.Error())
			}
			continue
		}

		// cis1 oxx ooo x,oxx oxx o
		split := strings.SplitN(line, " ", 2)
		noteStr := split[0]
		fingeringStr := split[1]

		note := parseNote(noteStr)
		note.Fingering = parseFingering(fingeringStr)

		bar.AddNote(note)
	}

	chart.AddBar(bar)

	logger.WriteString("\r\nwriting " + outputFile)
	f, err := os.Create(outputFile)
	if err != nil {
		logger.WriteString("\r\nerror creating output file: " + err.Error())
		return
	}

	_, err = f.WriteString(chart.Print(config))
	if err != nil {
		logger.WriteString("\r\nerror writing chart: " + err.Error())
		return
	}

	err = f.Sync()
	if err != nil {
		logger.WriteString("\r\nerror flushing file: " + err.Error())
		return
	}

	err = f.Close()
	if err != nil {
		logger.WriteString("\r\nerror closing file: " + err.Error())
		return
	}
}

func parseNote(n string) Note {
	base := BaseNote{
		ID:    "note",
		BaseX: 40,
		BaseY: 27,
	}

	flat := &NoteModifier{
		ID: "flat",
		//PaddingLeft:   5, // force same bar length for same note count
		PaddingTop:    6,
		PaddingBottom: 0.5,
	}

	doubleFlat := &NoteModifier{
		ID:            "double-flat",
		PaddingLeft:   5, // force same bar length for same note count
		PaddingTop:    6,
		PaddingBottom: 0.5,
	}

	sharp := &NoteModifier{
		ID: "sharp",
		//PaddingLeft:   6, // force same bar length for same note count
		PaddingTop:    4,
		PaddingBottom: 2,
	}

	doubleSharp := &NoteModifier{
		ID: "double-sharp",
		//PaddingLeft:   6, // force same bar length for same note count
		PaddingTop:    4,
		PaddingBottom: 2,
	}

	octave := map[string]Note{
		"c":     {BaseNote: base, Position: -1},
		"cis":   {BaseNote: base, Modifier: sharp, Position: -1},
		"cisis": {BaseNote: base, Modifier: doubleSharp, Position: -1},
		"ces":   {BaseNote: base, Modifier: flat, Position: -1},
		"ceses": {BaseNote: base, Modifier: doubleFlat, Position: -1},

		"d":     {BaseNote: base, Position: 0},
		"dis":   {BaseNote: base, Modifier: sharp, Position: 0},
		"disis": {BaseNote: base, Modifier: doubleSharp, Position: 0},
		"des":   {BaseNote: base, Modifier: flat, Position: 0},
		"deses": {BaseNote: base, Modifier: doubleFlat, Position: 0},

		"e":     {BaseNote: base, Position: 1},
		"eis":   {BaseNote: base, Modifier: sharp, Position: 1},
		"eisis": {BaseNote: base, Modifier: doubleSharp, Position: 1},
		"es":    {BaseNote: base, Modifier: flat, Position: 1},
		"eses":  {BaseNote: base, Modifier: doubleFlat, Position: 1},

		"f":     {BaseNote: base, Position: 2},
		"fis":   {BaseNote: base, Modifier: sharp, Position: 2},
		"fisis": {BaseNote: base, Modifier: doubleSharp, Position: 2},
		"fes":   {BaseNote: base, Modifier: flat, Position: 2},
		"feses": {BaseNote: base, Modifier: doubleFlat, Position: 2},

		"g":     {BaseNote: base, Position: 3},
		"gis":   {BaseNote: base, Modifier: sharp, Position: 3},
		"gisis": {BaseNote: base, Modifier: doubleSharp, Position: 3},
		"ges":   {BaseNote: base, Modifier: flat, Position: 3},
		"geses": {BaseNote: base, Modifier: doubleFlat, Position: 3},

		"a":     {BaseNote: base, Position: 4},
		"ais":   {BaseNote: base, Modifier: sharp, Position: 4},
		"aisis": {BaseNote: base, Modifier: doubleSharp, Position: 4},
		"as":    {BaseNote: base, Modifier: flat, Position: 4},
		"ases":  {BaseNote: base, Modifier: doubleFlat, Position: 4},

		"h":     {BaseNote: base, Position: 5},
		"his":   {BaseNote: base, Modifier: sharp, Position: 5},
		"hisis": {BaseNote: base, Modifier: doubleSharp, Position: 5},
		"b":     {BaseNote: base, Modifier: flat, Position: 5},
		"heses": {BaseNote: base, Modifier: doubleFlat, Position: 5},
	}

	var i int
	for _, v := range n {
		if unicode.IsDigit(v) || v == '-' {
			break
		}
		i++
	}

	note, ok := octave[n[:i]]
	if !ok {
		return Note{}
	}

	oct, err := strconv.ParseInt(n[i:], 10, 32)
	if err != nil {
		// we got a note but no octave, so return the note with the 0th octave
		return note.AtOctave(-1)
	}

	// since we are using the first octave as the base, we need to subtract 1
	return note.AtOctave(int(oct - 1))
}

func parseFingering(f string) [][]Finger {
	pressed := Finger{
		ID: "dot",
	}

	unpressed := Finger{
		ID: "circle",
	}

	keyPressed := Finger{
		ID: "rect-fill",
	}

	keyUnpressed := Finger{
		ID: "rect-outline",
	}

	space := Finger{}

	fingering := make([][]Finger, 0)
	fingerPattern := make([]Finger, 0)

	for _, finger := range f {
		switch finger {
		case ',':
			// begin a new pattern
			fingering = append(fingering, fingerPattern)
			fingerPattern = make([]Finger, 0)
		case 'o':
			fingerPattern = append(fingerPattern, unpressed)
		case 'x':
			fingerPattern = append(fingerPattern, pressed)
		case 'O':
			fingerPattern = append(fingerPattern, keyUnpressed)
		case 'X':
			fingerPattern = append(fingerPattern, keyPressed)
		case ' ':
			fingerPattern = append(fingerPattern, space)
		}
	}
	fingering = append(fingering, fingerPattern)
	return fingering
}
