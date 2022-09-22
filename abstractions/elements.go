/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// ELEMENT INTERFACES

// This interface defines the methods supported by all discrete elements.
type Discrete interface {
	AsBoolean() bool
	AsInteger() int
}

// This interface defines the methods supported by all continuous elements.
type Continuous interface {
	AsReal() float64
}

// This interface defines the methods supported by all polar elements.
type Complex interface {
	IsUndefined() bool
	IsInfinite() bool
	GetReal() float64
	GetImaginary() float64
	GetMagnitude() float64
	GetPhase() float64
}

// This interface defines the methods supported by all polarized elements.
type Polarized interface {
	IsNegative() bool
}

// This interface defines the methods supported by all matchable pattern
// elements.
type Matchable interface {
	MatchesText(text string) bool
}

// This interface defines the methods supported by all segmented resource
// elements.
type Segmented interface {
	GetScheme() string
	GetAuthority() string
	GetPath() string
	GetQuery() string
	GetFragment() string
}

// This interface defines the methods supported by all temporal elements.
type Temporal interface {
	// Return the entire time in specific units.
	AsMilliseconds() float64
	AsSeconds() float64
	AsMinutes() float64
	AsHours() float64
	AsDays() float64
	AsWeeks() float64
	AsMonths() float64
	AsYears() float64

	// Return a specific part of the entire time.
	GetMilliseconds() int
	GetSeconds() int
	GetMinutes() int
	GetHours() int
	GetDays() int
	GetWeeks() int
	GetMonths() int
	GetYears() int
}

// These are locked to the Earth's daily revolutions.

const MillisecondsPerSecond int = 1000
const MillisecondsPerMinute int = MillisecondsPerSecond * 60
const MillisecondsPerHour int = MillisecondsPerMinute * 60
const MillisecondsPerDay int = MillisecondsPerHour * 24
const MillisecondsPerWeek int = MillisecondsPerDay * 7

// These are locked to the Earth's yearly orbit around the sun.

const MillisecondsPerYear int = 31556952000
const MillisecondsPerMonth int = MillisecondsPerYear / 12 // An average but exact value.

// Tying the two together is where things get messy.

const DaysPerMonth float64 = float64(MillisecondsPerMonth) / float64(MillisecondsPerDay)   // ~30.436875 days/month
const WeeksPerMonth float64 = float64(MillisecondsPerMonth) / float64(MillisecondsPerWeek) // ~4.348125 weeks/month
const DaysPerYear float64 = float64(MillisecondsPerYear) / float64(MillisecondsPerDay)     // ~365.2425 days/year

// LIBRARY INTERFACES

// This library interface defines the functions supported by all libraries of
// scalable numerical elements.
type Scalable[T ~int | ~float64] interface {
	Inverse(number T) T
	Sum(first, second T) T
	Difference(first, second T) T
	Scaled(number T, factor float64) T
}

// This library interface defines the functions supported by all libraries of
// fully numerical elements.
type Numerical[T ~int | ~float64 | ~complex128] interface {
	Reciprocal(number T) T
	Conjugate(number T) T
	Product(first, second T) T
	Quotient(first, second T) T
	Remainder(first, second T) T
	Exponential(base, power T) T
	Logarithm(base, value T) T
}

// This library interface defines the functions supported by all libraries of
// anglular elements.
type Angular[T ~float64] interface {
	Complement(angle T) T
	Supplement(angle T) T
	Conjugate(angle T) T
	Cosine(angle T) float64
	ArcCosine(x float64) T
	Sine(angle T) float64
	ArcSine(y float64) T
	Tangent(angle T) float64
	ArcTangent(x, y float64) T
}

// This library interface defines the functions supported by all libraries of
// elements that support boolean logic.
type Logical[T any] interface {
	Not(boolean T) T
	And(first, second T) T
	Sans(first, second T) T
	Or(first T, second T) T
	Xor(first T, second T) T
}

// This library interface defines the functions supported by all libraries of
// time-relative elements.
type Relative[M ~int, D ~int] interface {
	Duration(first, second M) D
	Earlier(moment M, duration D) M
	Later(moment M, duration D) M
}