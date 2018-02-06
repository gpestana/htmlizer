## htmlizer

[![Build Status](https://travis-ci.org/gpestana/htmlizer.svg?branch=master)](https://travis-ci.org/gpestana/htmlizer)

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

### Just use it!

MIT License

Copyright (c) [2018] [Goncalo Pestana]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
