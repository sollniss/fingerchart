package main

import (
	"fmt"
	"math"
	"strconv"
)

type Bar struct {
	// Horizontal space between nodes.
	HorizontalNoteSpace float64

	// Vertical space between notes (half of the line space).
	VerticalNoteSpace float64

	notes []Note
}

func (b Bar) Print() (float64, float64, string) {
	var notes string

	barlength := 0.0

	// height of the key
	minNoteY := -9.0 // TODO: make dynamic
	maxNoteY := 30.0 // TODO: make dynamic

	maxFingerY := 0.0
	extraPaddingForLowNotes := 0.0
	for _, note := range b.notes {
		topModifier := 0.0
		bottomModifier := 0.0
		if note.Modifier != nil {
			barlength += note.Modifier.PaddingLeft
			topModifier += note.Modifier.PaddingTop

			// bottom is always going to be same so lets not offset for now
			//bottomModifier += note.Modifier.PaddingBottom
		}

		// find the space we need above the bar
		ny := note.getY(b.VerticalNoteSpace)

		// y position is the center, so we need to subtract space for the upper half
		if ny-b.VerticalNoteSpace-topModifier < minNoteY {
			minNoteY = ny - topModifier - b.VerticalNoteSpace
		}
		if ny+b.VerticalNoteSpace+bottomModifier > maxNoteY {
			maxNoteY = ny + bottomModifier + b.VerticalNoteSpace
		}

		// if the note has more than two fingerings, shift it to the right
		fingeringCount := len(note.Fingering)
		noteOffsetX := 0.0
		if fingeringCount > 2 {
			noteOffsetX = float64(5 * (fingeringCount - 1) / 2) // integer division

			// add padding to the left side of the note
			barlength += noteOffsetX
		}

		nx := strconv.FormatFloat(barlength, 'f', -1, 64)
		// print the note
		notes += fmt.Sprintf(`<g transform="translate(%s)">`, nx)
		notes += note.Print(b.VerticalNoteSpace)

		barlength += b.HorizontalNoteSpace

		if maxNoteY > 30.0 {
			// if we have at least one note below the bar, add extra padding
			extraPaddingForLowNotes = 5.0
		}

		var fingeringSVG string
		for i, fingering := range note.Fingering {
			// padding for the fingering
			fingerX := 5 + (10 * i)
			fingerY := 10.0

			for _, finger := range fingering {
				if finger.ID != "" {
					fingerYStr := strconv.FormatFloat(fingerY, 'f', -1, 64)
					fingeringSVG += fmt.Sprintf(`<g transform="translate(%d,%s)"><use href="#%s"/></g>`, fingerX, fingerYStr, finger.ID)
					fingerY += 10.0
				} else {
					// empty finger means start of a new group, so reduce the spacing a bit
					fingerY += 5.0
				}
			}

			if fingerY > maxFingerY {
				maxFingerY = fingerY
			}
		}

		fingeringOffsetX := -5 * (fingeringCount - 1)
		fingeringOffsetXStr := strconv.Itoa(fingeringOffsetX)
		notes += `<g transform="translate(` + fingeringOffsetXStr + `,%[1]s)">` // fingering group
		notes += fingeringSVG
		notes += `</g></g>` // fingering group, note group

		// add padding to the right side of the note
		barlength += noteOffsetX
	}

	// move all fingerings to the same offset
	maxNoteYStr := strconv.FormatFloat(maxNoteY+extraPaddingForLowNotes, 'f', -1, 64)
	notes = fmt.Sprintf(notes, maxNoteYStr)

	//fmt.Println(maxFingerY)
	maxNoteY += maxFingerY

	// space after the last note
	barlength += 0.75 * b.HorizontalNoteSpace

	lengthStr := strconv.FormatFloat(barlength, 'f', -1, 64)

	// negative because we move the bar in the opposite direction
	yOffsetStr := strconv.FormatFloat(-minNoteY, 'f', -1, 64)

	middleLine1 := strconv.FormatFloat(b.VerticalNoteSpace*2, 'f', -1, 64)
	middleLine2 := strconv.FormatFloat(b.VerticalNoteSpace*4, 'f', -1, 64)
	middleLine3 := strconv.FormatFloat(b.VerticalNoteSpace*6, 'f', -1, 64)
	barHeight := strconv.FormatFloat(b.VerticalNoteSpace*8, 'f', -1, 64)

	return barlength, math.Abs(maxNoteY - minNoteY), fmt.Sprintf(`<g transform="translate(0, %[1]s)">
		<path stroke="currentColor" stroke-linecap="square"
			d="M 0,0 V %[6]s
			M %[2]s,0 V %[6]s
			M 0,0 H %[2]s
			M 0,%[3]s H %[2]s
			M 0,%[4]s H %[2]s
			M 0,%[5]s H %[2]s
			M 0,%[6]s H %[2]s"/>
		<use href="#key" transform="translate(15, -3.5)"/>%[7]s</g>`, yOffsetStr, lengthStr, middleLine1, middleLine2, middleLine3, barHeight, notes)
}

func (b *Bar) AddNote(n Note) {
	b.notes = append(b.notes, n)
}
