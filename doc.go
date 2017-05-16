// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package cleancolumns generates table column names from unicode string, with
multiple languages support.

Example:

	package main

	import(
		"github.com/gophersgang/cleancolumns"
		"fmt"
	)

	func main () {
		text := cleancolumns.Make("Hellö Wörld хелло ворлд")
		fmt.Println(text) // Will print hello_world_khello_vorld

		someText := cleancolumns.Make("影師")
		fmt.Println(someText) // Will print: ying_shi

		enText := cleancolumns.MakeLang("This & that", "en")
		fmt.Println(enText) // Will print 'this_and_that'

		deText := cleancolumns.MakeLang("Diese & Dass", "de")
		fmt.Println(deText) // Will print 'diese_und_dass'

		cleancolumns.CustomSub = map[string]string{
			"water": "sand",
		}
		textSub := cleancolumns.Make("water is hot")
		fmt.Println(textSub) // Will print 'sand_is_hot'
	}

Requests or bugs?

https://github.com/gophersgang/cleancolumns/issues
*/
package cleancolumns
