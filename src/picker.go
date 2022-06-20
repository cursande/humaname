package humaname

import (
	"math/rand"
	"strings"
	"time"
)

type NameOption uint8

const (
	Both NameOption = iota
	First
	Last
)

type Gender uint8

const (
	All Gender = iota
	Male
	Female
)

func firstNames(gender Gender) string {
	switch gender {
	case Male:
		return maleFirstNames[rand.Intn(mfLength)]
	case Female:
		return femaleFirstNames[rand.Intn(ffLength)]
	default:
		if rand.Intn(2) == 1 {
			return maleFirstNames[rand.Intn(mfLength)]
		} else {
			return femaleFirstNames[rand.Intn(ffLength)]
		}
	}
}

func PickRandomName(option NameOption, gender Gender) string {
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	var first, last string

	switch option {
	case Both:
		first = firstNames(gender)
		last = lastNames[rand.Intn(lnLength)]
		sb.WriteString(first)
		sb.WriteString(" ")
		sb.WriteString(last)
	case First:
		first = firstNames(gender)
		sb.WriteString(first)
	case Last:
		last = lastNames[rand.Intn(lnLength)]
		sb.WriteString(last)
	}

	return sb.String()
}
