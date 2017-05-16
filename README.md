cleancolumns
====

Fork from [gosimple/slug](github.com/gosimple/slug) to cleanup table column names with strange characters.


-----

Package `slug` generate slug from unicode string, URL-friendly slugify with
multiple languages support.

[![GoDoc](https://godoc.org/github.com/gosimple/slug?status.png)](https://godoc.org/github.com/gosimple/slug)
[![Build Status](https://drone.io/github.com/gosimple/slug/status.png)](https://drone.io/github.com/gosimple/slug/latest)

[Documentation online](http://godoc.org/github.com/gosimple/slug)

## Example

	package main

	import(
		"github.com/gophersgang/cleancolumns"
	    "fmt"
	)

	func main () {
		text := cleancolumns.Make("Hellö Wörld хелло ворлд")
		fmt.Println(text) // Will print: "hello_world_khello_vorld"

		someText := cleancolumns.Make("影師")
		fmt.Println(someText) // Will print: "ying_shi"

		enText := cleancolumns.MakeLang("This & that", "en")
		fmt.Println(enText) // Will print: "this_and_that"

		deText := cleancolumns.MakeLang("Diese & Dass", "de")
		fmt.Println(deText) // Will print: "diese_und_dass"

		cleancolumns.CustomSub = map[string]string{
			"water": "sand",
		}
		textSub := cleancolumns.Make("water is hot")
		fmt.Println(textSub) // Will print: "sand_is_hot"
	}

### Requests or bugs?
<https://github.com/gophersgang/cleancolumns/issues>

## Installation

	go get -u github.com/gophersgang/cleancolumns

## License

The source files are distributed under the
[Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
