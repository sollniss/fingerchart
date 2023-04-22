package main

import (
	"fmt"
	"strconv"
)

type FingeringChart struct {
	// Space between bars.
	BarSpace float64
	bars     []Bar
}

func (c *FingeringChart) AddBar(b Bar) {
	c.bars = append(c.bars, b)
}

func (c FingeringChart) Print() string {
	svg := `<defs>
		<path id="key" d="m0 27 .9 4.9c.5 2.6-1.2 5.2-3.9 5.5-2.4.2-4.7-1.2-4.9-3.5-.1-1.4.9-3 2.2-3.1 1.2-.1 2.4.5 2.6 2.3.2 1.9-1.1 2.2-2.3 2.8.4.4 1.4.9 2.3.8 2.1-.2 3.7-2.4 3.2-4.6l-.9-4.8c-.4 0-.7.1-1.1.2-4.6.4-8.7-2.7-9.2-8.4-.4-4.4 4.3-9 6.7-11.2l-.4-2.1c-.7-3.4-.3-6.4 1.3-8.9.5-.8 1.5-2.3 2.1-2.3.5 0 2.5 2.6 2.8 6 .3 3.6-.8 7.6-4.2 10.5l.8 4.3c2.5-.2 5.8.8 6.2 5.2-.1 2.1-1.4 5.4-4.2 6.4zm-3.9-15.3c-.1.1-.2.1-.5.4-2 1.5-5.4 5.1-5.1 8.6.4 4.6 4.6 6.3 7 6.1.5 0 .9-.1 1.4-.2l-1.5-8c-1.5.2-2.7 1.4-2.6 3.1.1 1.5.7 1.8 1.6 2.6.4.4.1.6-.2.4-1.5-.6-2.7-2.3-2.9-3.9-.2-1.9 1.2-4.5 3.5-5.1zm3.9-11.8c0-.3-.2-1.6-1.3-1.5-2.4.2-3.3 4.9-2.6 8.2l.1.5c1.8-1.7 4-4.6 3.8-7.2zm-.2 26.2c1.8-.8 2.5-3.1 2.3-4.5-.1-1.6-1.6-3.4-3.8-3.3z"/>
		<path id="note" d="m0 0c.1 2-3.5 3.1-5.4 3.1-2 0-5.4-1.4-5.5-3.1-.1-2 3.2-3.1 5.3-3.1 2.5-.1 5.5 1.1 5.6 3.1zm-3.2.6c-.1-2-1-3.5-2.9-3.3-1.1.1-1.8.6-1.7 2.1.1 2 1 3.5 3 3.3 1.1-.2 1.7-.6 1.6-2.1z"/>
		<path id="sharp" d="m-13 1.5c.2-.1.3.5.3.9 0 .4 0 1-.3 1.1l-1 .4v3.4c0 .1-.2.1-.3.1-.1 0-.3 0-.3-.1v-3.1l-1.7.7v3.1c0 .1-.2.2-.3.2-.1 0-.3 0-.3-.2v-3l-1 .4c-.2.1-.2-.6-.2-1 0-.4 0-1 .2-1l1-.4v-4l-1 .4c-.1.1-.2-.6-.2-1 0-.4 0-1 .2-1l1-.4v-3.3c0-.1.2-.3.3-.3.2 0 .4.1.4.3v3l1.7-.7v-3c0-.1.2-.3.3-.3.2 0 .4.1.4.3v2.8l1-.4c.3-.1.3.5.3.9 0 .4 0 1-.3 1.1l-1 .4v3.9zm-3.3 1.2 1.7-.7v-3.9l-1.7.7z"/>
		<path id="double-sharp" d="m-15 0c.6-.6 1.4-.7 2.2-.7.1 0 .3-.2.3-.3l.2-2.2c0-.1-.1-.2-.2-.2l-2.2.2c-.1 0-.3.2-.3.3 0 .82-.1 1.6-.7 2.2-.6-.6-.7-1.4-.7-2.2 0-.1-.2-.3-.3-.3l-2.2-.2c-.1 0-.2.1-.2.2l.2 2.2c0 .1.2.3.3.3.8 0 1.6.1 2.2.7-.6.6-1.4.7-2.2.7-.1 0-.3.2-.3.3l-.2 2.2c0 .1.1.2.2.2l2.2-.2c.1 0 .3-.2.3-.3 0-.8.1-1.6.7-2.2.6.6.7 1.4.7 2.2 0 .1.2.3.3.3l2.2.2c.1 0 .2-.1.2-.2l-.2-2.2c0-.1-.2-.3-.3-.3-.8 0-1.6-.1-2.2-.7z"/>
		<path id="flat" d="m-17-1.7c.9-.6 1.8-1.2 2.6-1.1.7.1 1.6.7 1.6 1.9 0 2-1.3 2.7-2.8 3.9-.7.5-2 2-2 .9v-12.4c0-.2.1-.3.3-.3.2 0 .3.1.3.3zm0 4.7c1.1-.8 1.8-1 2.3-2.2.3-.9.1-1.9-.5-2.2-.6-.3-1.2.3-1.8.7z"/>
		<g id="double-flat">
			<use href="#flat" transform="translate(-4)"/>
			<use href="#flat"/>
		</g>
		
		<ellipse id="dot" ry="4" rx="4" cy="0" cx="29.5"/>
		<ellipse id="circle" ry="3.5" rx="3.5" cy="0" cx="29.5" style="fill:none;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:1;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"/>
		<rect id="rect-fill" width="7.5" height="7.5" x="25.75" y="-3.75"/>
		<rect id="rect-outline" width="7.5" height="7.5" x="25.75" y="-3.75" style="fill:none;stroke:#000000;stroke-width:1;stroke-linecap:round;stroke-linejoin:round;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"/>
	</defs><g transform="translate(1,1)">`

	maxWidth := 0.0
	heightOffset := 0.0
	for _, bar := range c.bars {
		w, h, s := bar.Print()

		y := strconv.FormatFloat(heightOffset, 'f', -1, 64)
		svg += fmt.Sprintf(`<g transform="translate(0,%s)">`, y)
		svg += s
		svg += `</g>`

		if w > maxWidth {
			maxWidth = w
		}

		heightOffset += h + c.BarSpace
	}

	// remove padding for last bar
	heightOffset -= c.BarSpace

	svg += `</g></svg>` // g is for the translate(1,1)

	// 1pt one the left + 1pt on the right (transform(1,1)) + 1pt for the vertical line in the bar (i think??)
	strW := strconv.FormatFloat(maxWidth+3, 'f', -1, 64)
	strH := strconv.FormatFloat(heightOffset+1, 'f', -1, 64)
	strW150per := strconv.FormatFloat((maxWidth+3)*1.5, 'f', -1, 64)
	strH150per := strconv.FormatFloat((heightOffset+1)*1.5, 'f', -1, 64)
	return fmt.Sprintf(`<svg viewBox="0 0 %s %s" width="%s" height="%s" xmlns="http://www.w3.org/2000/svg">`, strW, strH, strW150per, strH150per) + svg
}
