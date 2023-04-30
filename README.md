# fingerchart
A tool to generate fingering charts.

This tool is mainly for not-so computer literate people. So no command line features.
You are expected to drop a text file onto the executable which then generate an SVG in the same directory as the text file. 
A log file is also generated in the same directory as the executable, to allow some degree of debugging.

# Example
Input:
```
#background-color=#FFCDBC
#color=#130303
#node-color=#2D080A
#hole-color=#7C3626
#scale=1.3
eses1 xxx xxx O
es1 xxx xxx X
e1 xxx xxo X
eis1 xxx xoo X,xxx oox O
eisis1 xxx oox X

feses-1 ooo ooo X
fes0 ooo oox X
f1 ooo oxx X
fis2 ooo xxx X
fisis3 oox xxx X
```

Output:

![SVG output](https://raw.githubusercontent.com/sollniss/fingerchart/main/examples/example.svg)

# Usage

All note names are in German. The number behind the note's name specifies it's height. Negative numbers are valid.
Here is a list of all supported note names:
```
c, cis, cisis, ces, ceses,
d, dis, disis, des, deses,
e, eis, eisis, es, eses,
f, fis, fisis, fes, feses,
g, gis, gisis, ges, geses,
a, ais, aisis, as, ases,
h, his, hisis, b, heses
```

`x` and `o` represend a pressed, and unpressed hole. `X` and `O` represent pressed and unpressed keys.

Each line represents one note and it's fingering(s). A space in the fingering list creates a small gap in the SVG. You can add multiple fingerings to a note by separating them with a `,`.
An empty line creates a new bar.

## Config

Config values are all optional and start with a `#`. The following default values are used when the parameter is unset:
```
#background-color=transparent
#color=black
#node-color=black
#hole-color=black
#scale=1
```
(Under the hood, the values are simply omitted from the SVG if not set.)