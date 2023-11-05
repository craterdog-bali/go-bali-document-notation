/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	mat "math"
	tim "time"
)

// CONSTANT DEFINITIONS

const Epoch = Moment(0)

// MOMENT IMPLEMENTATION

// This constructor creates a new moment in time element for the current time
// in the UTC timezone.
func Now() abs.MomentLike {
	var now = tim.Now().UTC().UnixMilli()
	return Moment(now)
}

// This constructor creates a new moment in time element from the specified
// integer number of milliseconds since the UNIX Epoch in the UTC timezone.
func MomentFromInt(milliseconds int) abs.MomentLike {
	return Moment(milliseconds)
}

// This constructor returns the earliest value for a moment.
func MinimumMoment() abs.MomentLike {
	return Moment(mat.MinInt)
}

// This constructor returns the latest value for a moment.
func MaximumMoment() abs.MomentLike {
	return Moment(mat.MaxInt)
}

// This type defines the methods associated with moment in time elements. It
// extends the native Go int type and represents the number of milliseconds
// after the UNIX Epoch (Midnight, January 1, 1970 UTC) for a moment of tim.
// All moments are based on UTC.
type Moment int

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v Moment) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this discrete element.
func (v Moment) AsInteger() int {
	return int(v)
}

// TEMPORAL INTERFACE

// This method returns the total number of milliseconds since the UNIX Epoch
// in this moment.
func (v Moment) AsMilliseconds() float64 {
	return float64(v)
}

// This method returns the total number of seconds since the UNIX Epoch
// in this moment.
func (v Moment) AsSeconds() float64 {
	return float64(v.AsMilliseconds()) / 1000.0 // milliseconds per second
}

// This method returns the total number of minutes since the UNIX Epoch
// in this moment.
func (v Moment) AsMinutes() float64 {
	return v.AsSeconds() / 60.0 // seconds per minute
}

// This method returns the total number of hours since the UNIX Epoch
// in this moment.
func (v Moment) AsHours() float64 {
	return v.AsMinutes() / 60.0 // minutes per hour
}

// This method returns the total number of days since the UNIX Epoch
// in this moment.
func (v Moment) AsDays() float64 {
	return v.AsHours() / 24.0 // hours per day
}

// This method returns the total number of weeks since the UNIX Epoch
// in this moment.
func (v Moment) AsWeeks() float64 {
	return v.AsDays() / 7.0 // days per week
}

// This method returns the total number of months since the UNIX Epoch
// in this moment.
func (v Moment) AsMonths() float64 {
	return v.AsDays() / DaysPerMonth
}

// This method returns the total number of years since the UNIX Epoch
// in this moment.
func (v Moment) AsYears() float64 {
	return v.AsDays() / DaysPerYear
}

// This method returns the millisecond part of this moment.
func (v Moment) GetMilliseconds() int {
	var t = v.asTime()
	return t.Nanosecond() / 1e6
}

// This method returns the second part of this moment.
func (v Moment) GetSeconds() int {
	var t = v.asTime()
	return t.Second()
}

// This method returns the minute part of this moment.
func (v Moment) GetMinutes() int {
	var t = v.asTime()
	return t.Minute()
}

// This method returns the hour part of this moment.
func (v Moment) GetHours() int {
	var t = v.asTime()
	return t.Hour()
}

// This method returns the day part of this moment.
func (v Moment) GetDays() int {
	var t = v.asTime()
	return t.Day()
}

// This method returns the week part of this moment.
func (v Moment) GetWeeks() int {
	var t = v.asTime()
	var _, week = t.ISOWeek()
	return week
}

// This method returns the month part of this moment.
func (v Moment) GetMonths() int {
	var t = v.asTime()
	return int(t.Month())
}

// This method returns the year part of this moment.
func (v Moment) GetYears() int {
	var t = v.asTime()
	return t.Year()
}

// PRIVATE INTERFACE

// This function returns the go Time value for the specified UNIX milliseconds.
func (v Moment) asTime() tim.Time {
	var milliseconds int64 = int64(v)
	return tim.UnixMilli(milliseconds).UTC()
}

// MOMENTS LIBRARY

// This singleton creates a unique name space for the library functions for
// moment elements.
var Moments = &moments{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for moment elements.
type moments struct{}

// This library function returns the duration of time between the two specified
// moments in tim.
func (l *moments) Duration(first, second abs.MomentLike) abs.DurationLike {
	return Duration(second.AsInteger() - first.AsInteger())
}

// This library function returns the moment in time that is earlier than the
// specified moment in time by the specified duration of tim.
func (l *moments) Earlier(moment abs.MomentLike, duration abs.DurationLike) abs.MomentLike {
	return Moment(moment.AsInteger() - duration.AsInteger())
}

// This library function returns the moment in time that is later than the
// specified moment in time by the specified duration of tim.
func (l *moments) Later(moment abs.MomentLike, duration abs.DurationLike) abs.MomentLike {
	return Moment(moment.AsInteger() + duration.AsInteger())
}