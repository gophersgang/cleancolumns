// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cleancolumns

func init() {
	// Merge language subs with the default one
	for _, sub := range []*map[rune]string{&deSub, &enSub, &plSub, &esSub} {
		for key, value := range defaultSub {
			(*sub)[key] = value
		}
	}
}

var defaultSub = map[rune]string{
	'"':  "",
	'\'': "",
	'’':  "",
	'‒':  "_", // figure dash
	'–':  "_", // en dash
	'—':  "_", // em dash
	'―':  "_", // horizontal bar
}

var deSub = map[rune]string{
	'&': "und",
	'@': "an",
}

var enSub = map[rune]string{
	'&': "and",
	'@': "at",
}

var plSub = map[rune]string{
	'&': "i",
	'@': "na",
}

var esSub = map[rune]string{
	'&': "y",
	'@': "en",
}
