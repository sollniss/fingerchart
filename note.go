package main

import (
	"fmt"
	"strconv"
)

const EPSILON float64 = 0.001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

type NoteModifier struct {
	// ID of the SVG element.
	ID string

	// Extra space needed on the left side of the note.
	PaddingLeft float64

	// Extra space needed above the note.
	PaddingTop float64

	// Extra space needed below the note.
	PaddingBottom float64
}

type BaseNote struct {
	// ID of the SVG element.
	ID string

	// X-position of the d note.
	BaseX float64

	// X-position of the d note.
	BaseY float64
}

type Note struct {
	BaseNote

	// Position relative to the d note.
	Position int

	// Current octave (0 is first octave).
	Octave int

	Modifier *NoteModifier

	Fingering [][]Finger
}

// getY calculates the y position of the note.
func (n Note) getX() float64 {
	return n.BaseX
}

// getY calculates the y position of the note.
func (n Note) getY(verticalNoteSpace float64) float64 {
	return n.BaseY - float64(n.Position)*verticalNoteSpace - float64(n.Octave)*verticalNoteSpace*7.0
}

// AtOctave(0) is actually the first octave
func (n Note) AtOctave(octave int) Note {
	n.Octave = octave
	return n
}

func (n Note) WithFingering(fingering [][]Finger) Note {
	n.Fingering = fingering
	return n
}

func (n *Note) AddFinger(f []Finger) {
	n.Fingering = append(n.Fingering, f)
}

func (n Note) Print(verticalNoteSpace float64) string {
	y := n.getY(verticalNoteSpace)
	lineX := strconv.FormatFloat(n.getX()-12.5, 'f', -1, 64) // TODO: automatic offset (note width + 1/2 of the length that extends from the note)

	var lines string

	// need extra line(s) above if note is lower than or equal to c (d + 1 vertical unit)
	offset := n.BaseY + verticalNoteSpace
	// lower means y is greater
	for y > offset || floatEquals(y, offset) {
		lineY := strconv.FormatFloat(offset, 'f', -1, 64)
		lines += fmt.Sprintf(`<path stroke="black" d="M %s,%s h 14"/>`, lineX, lineY) // TODO: automatic length

		// move up
		offset += verticalNoteSpace * 2
	}

	if lines == "" {
		// need extra lines below if higher than or equal to a' (d + 11 vertical units)
		offset = n.BaseY - (11 * verticalNoteSpace)
		// higher means y is smaller
		for y < offset || floatEquals(y, offset) {
			lineY := strconv.FormatFloat(offset, 'f', -1, 64)
			lines += fmt.Sprintf(`<path stroke="black" d="M %s,%s h 14"/>`, lineX, lineY) // TODO: automatic length

			// move down
			offset -= verticalNoteSpace * 2
		}
	}

	var modifier string
	if n.Modifier != nil {
		modifier = fmt.Sprintf(`<use href="#%s"/>`, n.Modifier.ID)
		//x += n.Modifier.PaddingLeft
	}

	strX := strconv.FormatFloat(n.getX(), 'f', -1, 64)
	strY := strconv.FormatFloat(y, 'f', -1, 64)

	return fmt.Sprintf(`%s<g transform="translate(%s, %s)">%s<use href="#%s"/></g>`, lines, strX, strY, modifier, n.ID)
}
