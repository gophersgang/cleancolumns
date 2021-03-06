// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cleancolumns

import (
	"bytes"
	"regexp"
	"sort"
	"strings"

	"github.com/rainycape/unidecode"
)

// SubStruct represents a string substitution
type SubStruct struct {
	In  string
	Out string
}

var (
	// CustomSub stores custom substitution map
	CustomSub map[string]string

	// CustomSubOrdered stores custom substitution slice, to have some control for the order of replacement
	CustomSubOrdered []SubStruct

	// CustomRuneSub stores custom rune substitution map
	CustomRuneSub map[rune]string

	// MaxLength stores maximum slug length.
	// It's smart so it will cat slug after full word.
	// By default slugs aren't shortened.
	// If MaxLength is smaller than length of the first word, then returned
	// slug will contain only substring from the first word truncated
	// after MaxLength.
	MaxLength int

	regexpNonAuthorizedChars  = regexp.MustCompile("[^a-z0-9-_]")
	regexpMultipleDashes      = regexp.MustCompile("-+")
	regexpMultipleUnderscores = regexp.MustCompile("_+")
)

//=============================================================================

// Make returns slug generated from provided string. Will use "en" as language
// substitution.
func Make(s string) (slug string) {
	return MakeLang(s, "en")
}

// MakeLang returns slug generated from provided string and will use provided
// language for chars substitution.
func MakeLang(s string, lang string) (slug string) {
	slug = strings.TrimSpace(s)

	// Custom substitutions
	// Always substitute runes first
	slug = SubstituteRune(slug, CustomRuneSub)
	slug = SubstituteOrdered(slug, CustomSubOrdered)
	slug = Substitute(slug, CustomSub)

	// Process string with selected substitution language
	switch lang {
	case "de":
		slug = SubstituteRune(slug, deSub)
	case "en":
		slug = SubstituteRune(slug, enSub)
	case "pl":
		slug = SubstituteRune(slug, plSub)
	case "es":
		slug = SubstituteRune(slug, esSub)
	default: // fallback to "en" if lang not found
		slug = SubstituteRune(slug, enSub)
	}

	// Process all non ASCII symbols
	slug = unidecode.Unidecode(slug)

	slug = strings.ToLower(slug)

	// Process all remaining symbols
	slug = regexpNonAuthorizedChars.ReplaceAllString(slug, "_")
	slug = regexpMultipleDashes.ReplaceAllString(slug, "_")
	slug = regexpMultipleUnderscores.ReplaceAllString(slug, "_")
	slug = strings.Trim(slug, "_")

	if MaxLength > 0 {
		slug = smartTruncate(slug)
	}

	return slug
}

// Substitute returns string with superseded all substrings from
// provided substitution map. Substitution map will be applied in alphabetic
// order. Many passes, on one substitution another one could apply.
func Substitute(s string, sub map[string]string) (buf string) {
	buf = s
	var keys []string
	for k := range sub {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		buf = strings.Replace(buf, key, sub[key], -1)
	}
	return
}

// SubstituteOrdered has controlled order during replacements
func SubstituteOrdered(s string, sub []SubStruct) (buf string) {
	buf = s
	for _, subItem := range sub {
		buf = strings.Replace(buf, subItem.In, subItem.Out, -1)
	}
	return buf
}

// SubstituteRune substitutes string chars with provided rune
// substitution map. One pass.
func SubstituteRune(s string, sub map[rune]string) string {
	var buf bytes.Buffer
	for _, c := range s {
		if d, ok := sub[c]; ok {
			buf.WriteString(d)
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func smartTruncate(text string) string {
	if len(text) < MaxLength {
		return text
	}

	var truncated string
	words := strings.SplitAfter(text, "_")
	// If MaxLength is smaller than length of the first word return word
	// truncated after MaxLength.
	if len(words[0]) > MaxLength {
		return words[0][:MaxLength]
	}
	for _, word := range words {
		if len(truncated)+len(word)-1 <= MaxLength {
			truncated = truncated + word
		} else {
			break
		}
	}
	return strings.Trim(truncated, "_")
}
