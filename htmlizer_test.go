package htmlizer

import (
	"fmt"
	"testing"
)

func TestSimpleDOM(t *testing.T) {
	html := `<html>
  	<body>
    	<h1>Heading H1</h1>
      <p>This is the first text</p>
      <h2>heading h2</h2>
      <p>This is the second text</p>
    </body>
    <script>console.log("scripts are discarded")</script>
  </html>`

	ignore := []rune{}
	hr, _ := New(ignore)
	hr.Load(html)

	h1, _ := hr.GetValues("<h1>")
	h2, _ := hr.GetValues("<h2>")
	p, _ := hr.GetValues("<p>")
	readable := hr.HumanReadable()

	expectedH1 := "Heading H1"
	expectedH2 := "heading h2"
	if len(h1) != 1 {
		t.Fatal(fmt.Sprintf("There is one h1 tag, found %v", len(h1)))
	}
	if h1Content := h1[0].Value; h1Content != expectedH1 {
		t.Error(fmt.Sprintf("H1 content should be '%v', found %v", expectedH1, h1Content))
	}

	if len(h2) != 1 {
		t.Fatal(fmt.Sprintf("There is one h2 tag, found %v", len(h2)))
	}
	if h2Content := h2[0].Value; h2Content != expectedH2 {
		t.Error(fmt.Sprintf("H2 content should be '%v', found %v", expectedH2, h2Content))
	}

	expectedFirstP := "This is the first text"
	expectedSecondP := "This is the second text"
	if len(p) != 2 {
		t.Fatal(fmt.Sprintf("There are two p tags, found '%v'", len(p)))
	}
	if firstP := p[0].Value; firstP != expectedFirstP {
		t.Error(fmt.Sprintf("First P content should be '%v'', found %v", expectedFirstP, firstP))
	}
	if secondP := p[1].Value; secondP != expectedSecondP {
		t.Error(fmt.Sprintf("Second P content should be %v', found %v", expectedSecondP, secondP))
	}

	expectedHR := `Heading H1
This is the first text
heading h2
This is the second text`

	if readable != expectedHR {
		t.Error(fmt.Sprintf("Human Readable output should be %v, found", expectedHR, readable))
	}
}

/*
func TestSpecialChars(t *testing.T) {
	html := `<html>
  	<body>
      <p>testing the ignores now...
followed by this			<-- this are a tabs
How about this?</p>
    </body>
  </html>`

	ignore := []rune{'\t', '\n'}
	hr, _ := New(ignore)
	hr.Load(html)
	expectedHR := `testing the ignores now...
followed by this <-- this are tabs
How about this?
`
	p, _ := hr.GetValues("<p>")
	if actual := p[0].Value; actual != expectedHR {
		t.Error(fmt.Sprintf("Human readable should be: \n%v \nfound:\n %v", expectedHR, actual))
	}
}
*/
func testDropRunes(t *testing.T) {
	str := "This \n is \t a big mess \n \n"

	nan := []rune{}
	expectedNan := str
	actualNan := dropRunes(str, nan)

	if actualNan != expectedNan {
		t.Error(fmt.Sprintf("Final string should be %v, found %v", expectedNan, actualNan))
	}

	newLines := []rune{'\n'}
	expectedDropNewLine := "This is \t a big mess"
	actualDropNewLine := dropRunes(str, newLines)

	if actualDropNewLine != expectedDropNewLine {
		t.Error(fmt.Sprintf("Final string should be %v, found %v", expectedDropNewLine, actualDropNewLine))
	}

	all := []rune{'\n', '\t'}
	expectedDropAll := "This is a big mess"
	actualDropAll := dropRunes(str, all)

	if actualDropAll != expectedDropAll {
		t.Error(fmt.Sprintf("Final string should be %v, found %v", expectedDropAll, actualDropAll))
	}

}
