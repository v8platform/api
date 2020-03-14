package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(`(?smU)\A(?:\[\n|\r\n).*?(?:\n|\r\n)\]`)
	var str = `[
{
"type": "success",
"message": "",
"body": []
}
][
{
"type": "success",
"message": "",
"body": []
}
]`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
	b := re.MatchString(str)
	fmt.Println("found ", b)

}
