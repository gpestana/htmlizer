## htmlizer

Parses only human readable content from HTML DOM.


### Example

```
import (
	"fmt"
	"github.com/gpestana/htmlizer"
)

func main() {
	html := `
		<html>
			<body>
				<h1>Heading H1</h1>
				<p>This is the first text</p>
				<h2>heading h2</h2>
				<p>This is the second text</p>
			</body>
			<script>console.log("scripts are discarded")</script>
		</html>`

	hizer := htmlizer.New()
	hizer.Load(html)

	fmt.Println(">> Struct:")
	fmt.Println(hizer)

	fmt.Println(">> Human readable content:")
	fmt.Println(hizer.HumanReadable())
}
```

Output:
```
	// >> Struct:
	// {[Heading H1 heading h2], [this is the first text this is the seconf text]}
	// >> Human readable content:
	// Heading H1
	// This is the first text
	// heading h2
	// This is the second text
```
