/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language_test

import (
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

/*
const m = "<-1009-08-07T06:05:04.321>"

const p = `"ca+t"?`

const r = "<https://google.com>"

const s = "$getItem"

const v = "v1.2.3"
*/

func TestParserWithAngles(t *tes.T) {
	var angleStrings = []string{
		`~0`,
		`~π`,
		`~1.23456789E-10`,
	}

	var angleValues = []ele.Angle{
		0,
		mat.Pi,
		1.23456789e-10,
	}

	for index, s := range angleStrings {
		var angle = lan.ParseSource(s).GetEntity().(ele.Angle)
		ass.Equal(t, angleValues[index], angle)
	}

	for index, angle := range angleValues {
		var s = lan.FormatValue(angle)
		ass.Equal(t, angleStrings[index], s)
	}
}

func TestParserWithBooleans(t *tes.T) {
	var booleanStrings = []string{
		`false`,
		`true`,
	}

	var booleanValues = []ele.Boolean{
		false,
		true,
	}

	for index, s := range booleanStrings {
		var boolean = lan.ParseSource(s).GetEntity().(ele.Boolean)
		ass.Equal(t, booleanValues[index], boolean)
	}

	for index, boolean := range booleanValues {
		var s = lan.FormatValue(boolean)
		ass.Equal(t, booleanStrings[index], s)
	}
}

func TestParserWithDurations(t *tes.T) {
	var durationStrings = []string{
		`~P0W`,
		`~-P13W`,
		`~P12Y3M4DT5H6M7.890S`,
		`~-P12Y3M4D`,
		`~P3M4DT5H6M`,
		`~-PT6M7.890S`,
		`~P12YT6M7.890S`,
		`~-P12Y4DT5H6M7S`,
	}

	var durationValues = []ele.Duration{
		0,
		-7862400000,
		386936629890,
		-386918262000,
		8253198000,
		-367890,
		378683791890,
		-379047391000,
	}

	for index, s := range durationStrings {
		var duration = lan.ParseSource(s).GetEntity().(ele.Duration)
		ass.Equal(t, durationValues[index], duration)
	}

	for index, duration := range durationValues {
		var s = lan.FormatValue(duration)
		ass.Equal(t, durationStrings[index], s)
	}
}

func TestParserWithMoments(t *tes.T) {
	var momentStrings = []string{
		`<1776>`,
		`<1970>`,
		`<2009-04>`,
		`<1962-04-25>`,
		`<-1-02-03T04:05:06.789>`,
	}

	var momentValues = []ele.Moment{
		-6122044800000,
		0,
		1238544000000,
		-242611200000,
		-62195889293211,
	}

	for index, s := range momentStrings {
		var moment = lan.ParseSource(s).GetEntity().(ele.Moment)
		ass.Equal(t, momentValues[index], moment)
	}

	for index, moment := range momentValues {
		var s = lan.FormatValue(moment)
		ass.Equal(t, momentStrings[index], s)
	}
}