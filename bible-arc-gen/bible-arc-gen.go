package main

import (
	"fmt"
	"math"
	"os"
)

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

// Generate SVG arc paths
func generateSVGArcs(books []struct {
	Name  string
	Words int
}, radius float64) string {
	totalWords := 0
	for _, book := range books {
		totalWords += book.Words
	}

	startAngle := 0.0
	var svgPaths string

	for bookIdx, book := range books {
		sweepAngle := (float64(book.Words) / float64(totalWords)) * 360.0
		endAngle := startAngle + sweepAngle

		x1, y1 := polarToCartesian(radius, startAngle)
		x2, y2 := polarToCartesian(radius, endAngle)

		largeArcFlag := 0
		if sweepAngle > 180 {
			largeArcFlag = 1
		}

		svgPaths += fmt.Sprintf("<path id=\"%s\" d=\"M %.2f %.2f A %.2f %.2f 0 %d 1 %.2f %.2f\" stroke=\"%s\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"butt\"/>",
			book.Name, x1, y1, radius, radius, largeArcFlag, x2, y2, idxToHSL(bookIdx))
		svgPaths += "\n"
		startAngle = endAngle
	}

	return svgPaths
}

func main() {
	bibleBooks := []struct {
		Name  string
		Words int
	}{
		{"Genesis", 38262}, {"Exodus", 32685}, {"Leviticus", 24541}, {"Numbers", 32896}, {"Deuteronomy", 28352},
		{"Joshua", 18854}, {"Judges", 18966}, {"Ruth", 2574}, {"1 Samuel", 25048}, {"2 Samuel", 20600},
		{"1 Kings", 24513}, {"2 Kings", 23517}, {"1 Chronicles", 20365}, {"2 Chronicles", 26069}, {"Ezra", 7440},
		{"Nehemiah", 10480}, {"Esther", 5633}, {"Job", 10702}, {"Psalms", 42704}, {"Proverbs", 15038},
		{"Ecclesiastes", 5579}, {"Song of Solomon", 2658}, {"Isaiah", 37036}, {"Jeremiah", 42654}, {"Lamentations", 3411},
		{"Ezekiel", 39401}, {"Daniel", 11602}, {"Hosea", 5174}, {"Joel", 2033}, {"Amos", 4216}, {"Obadiah", 669},
		{"Jonah", 1320}, {"Micah", 3152}, {"Nahum", 1284}, {"Habakkuk", 1475}, {"Zephaniah", 1616}, {"Haggai", 1131},
		{"Zechariah", 6443}, {"Malachi", 1781}, {"Matthew", 23143}, {"Mark", 14949}, {"Luke", 25640}, {"John", 18658},
		{"Acts", 24229}, {"Romans", 9467}, {"1 Corinthians", 9462}, {"2 Corinthians", 6046}, {"Galatians", 3084},
		{"Ephesians", 3022}, {"Philippians", 2183}, {"Colossians", 1979}, {"1 Thessalonians", 1837}, {"2 Thessalonians", 1022},
		{"1 Timothy", 2244}, {"2 Timothy", 1666}, {"Titus", 896}, {"Philemon", 430}, {"Hebrews", 4953},
		{"James", 2304}, {"1 Peter", 2483}, {"2 Peter", 1553}, {"1 John", 2517}, {"2 John", 298}, {"3 John", 294},
		{"Jude", 608}, {"Revelation", 11952},
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
