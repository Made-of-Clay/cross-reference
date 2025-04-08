package main

import (
	"fmt"
	"math"
	"os"
)

type Book struct {
	Name     string
	Chapters int
}

// Convert polar coordinates to Cartesian (SVG coordinate space)
func polarToCartesian(radius, angleDegrees float64) (float64, float64) {
	angleRadians := angleDegrees * (math.Pi / 180.0)
	x := radius * math.Cos(angleRadians)
	y := radius * math.Sin(angleRadians)
	return x, y
}

// Convert int to CSS Hex color
func intToCSSColor(num int) string {
	// Ensure the value is within the range of 0-16777215 (0xFFFFFF)
	if num < 0 || num > 16777215 {
		num = 0 // default to black if out of range
	}

	// Format the integer to a hex value with the prefix "#"
	return fmt.Sprintf("#%06X", num)
}

func idxToHSL(idx int) string {
	return fmt.Sprintf("hsl(%d , 100%%, 50%%)", idx*90)
}

func idxToHex(idx int) string {
	colors := []string{
		"#D32F2F",
		"#F4511E",
		"#388E3C",
		"#00695C",
		"#1976D2",
		"#842cb9",
		"#53cdd3",
	}
	return colors[idx%len(colors)]
}

// Generate SVG arc paths
func generateSVGArcs(books []Book, radius float64) string {
	totalChapters := 0
	for _, book := range books {
		totalChapters += book.Chapters
	}

	startAngle := 0.0
	var svgPaths string

	for bookIdx, book := range books {
		sweepAngle := (float64(book.Chapters) / float64(totalChapters)) * 360.0
		endAngle := startAngle + sweepAngle

		x1, y1 := polarToCartesian(radius, startAngle)
		x2, y2 := polarToCartesian(radius, endAngle)

		largeArcFlag := 0
		if sweepAngle > 180 {
			largeArcFlag = 1
		}

		svgPaths += fmt.Sprintf("<path id=\"%s\" d=\"M %.2f %.2f A %.2f %.2f 0 %d 1 %.2f %.2f\" stroke=\"%s\" stroke-width=\"10\" data-chapters=\"%d\" fill=\"none\" stroke-linecap=\"butt\"/>",
			book.Name, // id
			x1,
			y1,
			radius,
			radius,
			largeArcFlag,
			x2,
			y2,
			idxToHex(bookIdx), // stroke
			book.Chapters,     // data-chapters
		)
		svgPaths += "\n"
		startAngle = endAngle
	}

	return svgPaths
}

func main() {
	bibleBooks := []Book{
		{"Genesis", 50},
		{"Exodus", 40},
		{"Leviticus", 27},
		{"Numbers", 36},
		{"Deuteronomy", 34},
		{"Joshua", 24},
		{"Judges", 21},
		{"Ruth", 4},
		{"1 Samuel", 31},
		{"2 Samuel", 24},
		{"1 Kings", 22},
		{"2 Kings", 25},
		{"1 Chronicles", 29},
		{"2 Chronicles", 36},
		{"Ezra", 10},
		{"Nehemiah", 13},
		{"Esther", 10},
		{"Job", 42},
		{"Psalms", 150},
		{"Proverbs", 31},
		{"Ecclesiastes", 12},
		{"Song of Solomon", 8},
		{"Isaiah", 66},
		{"Jeremiah", 52},
		{"Lamentations", 5},
		{"Ezekiel", 48},
		{"Daniel", 12},
		{"Hosea", 14},
		{"Joel", 3},
		{"Amos", 9},
		{"Obadiah", 1},
		{"Jonah", 4},
		{"Micah", 7},
		{"Nahum", 3},
		{"Habakkuk", 3},
		{"Zephaniah", 3},
		{"Haggai", 2},
		{"Zechariah", 14},
		{"Malachi", 4},
		{"Matthew", 28},
		{"Mark", 16},
		{"Luke", 24},
		{"John", 21},
		{"Acts", 28},
		{"Romans", 16},
		{"1 Corinthians", 16},
		{"2 Corinthians", 13},
		{"Galatians", 6},
		{"Ephesians", 6},
		{"Philippians", 4},
		{"Colossians", 4},
		{"1 Thessalonians", 5},
		{"2 Thessalonians", 3},
		{"1 Timothy", 6},
		{"2 Timothy", 4},
		{"Titus", 3},
		{"Philemon", 1},
		{"Hebrews", 13},
		{"James", 5},
		{"1 Peter", 5},
		{"2 Peter", 3},
		{"1 John", 5},
		{"2 John", 1},
		{"3 John", 1},
		{"Jude", 1},
		{"Revelation", 22},
	}

	svgArcs := generateSVGArcs(bibleBooks, 100)

	// Write SVG paths to a file
	file, err := os.Create("bible_arcs.svg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"-110 -110 220 220\">\n")
	file.WriteString(svgArcs)
	file.WriteString("\n</svg>")

	fmt.Println("SVG file 'bible_arcs.svg' has been generated.")
}
