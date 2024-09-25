package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	var err error

	pattern, err := p.Compile(`^(.*)-(.*)$`)
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	replacement, err := template.New("substitution").Parse(`{{if eq (index $ 1) "No SKU"}}{{index $ 2}}{{else}}{{index $ 1}}{{end}}`)
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	for _, str := range strs {

		matches := pattern.FindAllStringSubmatchIndex(str, -1)

		lastMatchIndex := 0
		result := ""

		for _, match := range matches {
			// collect group strings
			groups := make([]string, len(match)/2)
			for i := 0; i < len(match); i += 2 {
				if match[i] >= 0 && match[i+1] >= 0 {
					groups[i/2] = str[match[i]:match[i+1]]
				}
			}

			// render template with groups
			var buf bytes.Buffer
			if err := replacement.Execute(&buf, groups); err != nil {
				fmt.Println("🚫", err)
			}

			// substitute!
			result = result + str[lastMatchIndex:match[0]] + buf.String()
			lastMatchIndex = match[1]
		}

		result = result + str[lastMatchIndex:]
		fmt.Println(result, " ⏩ ", str)
	}
}

var strs = []string{`727235402187-183`, `No SKU-HW55`}
