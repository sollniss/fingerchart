# fingerchart
A tool to generate fingering charts.

This tool is mainly for not-so computer literate people. So no command line features.
You are expected to drop a text file onto the executable which then generate an SVG and a log file in the same directory as the text file.

# Example
Input:
```
eses1 xxx xxx O
es1 xxx xxx X
e1 xxo xxo O
eis1 xxx xxo X
eisis1 oxo xxo X

feses1 xxo xxx O
fes1 xox ooo O
f1 xxx ooo O
fis1 xxx xxo X
fisis1 ooo ooo X
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

Each line represents one note and it's fingering. A space in the fingering list creates a small gap in the SVG.
An empty line creates a new bar.
