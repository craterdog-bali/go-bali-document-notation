/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package ranges_test

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	ele "github.com/bali-nebula/go-component-framework/elements"
	ran "github.com/bali-nebula/go-component-framework/ranges"
	col "github.com/craterdog/go-collection-framework/v2"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestIntervalsWithImplicitDurations(t *tes.T) {
	var s = ran.Interval[ele.Duration](3, abs.LEFT, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.False(t, s.ContainsValue(2))
	ass.True(t, s.ContainsValue(3))
	ass.True(t, s.ContainsValue(5))
	ass.False(t, s.ContainsValue(7))
	ass.False(t, s.ContainsValue(8))
	ass.Equal(t, ele.Duration(3), s.GetFirst())
	ass.Equal(t, abs.LEFT, s.GetExtent())
	ass.Equal(t, ele.Duration(7), s.GetLast())
	ass.Equal(t, ele.Duration(5), s.GetValue(3))
	ass.Equal(t, 0, s.GetIndex(2))
	ass.Equal(t, 1, s.GetIndex(3))
	ass.Equal(t, 3, s.GetIndex(5))
	ass.Equal(t, 0, s.GetIndex(7))
	ass.Equal(t, 0, s.GetIndex(8))
	ass.Equal(t, []ele.Duration{3, 4, 5, 6}, s.AsArray())
	var iterator = col.Iterator[ele.Duration](s)
	ass.Equal(t, ele.Duration(3), iterator.GetNext())
	iterator.ToEnd()
	ass.Equal(t, ele.Duration(6), iterator.GetPrevious())
}
func TestIntervalsWithExplicitDurations(t *tes.T) {
	var two = ele.DurationFromInt(2)
	var three = ele.DurationFromInt(3)
	var four = ele.DurationFromInt(4)
	var five = ele.DurationFromInt(5)
	var six = ele.DurationFromInt(6)
	var seven = ele.DurationFromInt(7)
	var eight = ele.DurationFromInt(8)
	var s = ran.Interval(three, abs.INCLUSIVE, seven)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.False(t, s.ContainsValue(two))
	ass.True(t, s.ContainsValue(three))
	ass.True(t, s.ContainsValue(five))
	ass.True(t, s.ContainsValue(seven))
	ass.False(t, s.ContainsValue(eight))
	ass.Equal(t, three, s.GetFirst())
	ass.Equal(t, abs.INCLUSIVE, s.GetExtent())
	ass.Equal(t, seven, s.GetLast())
	ass.Equal(t, five, s.GetValue(3))
	ass.Equal(t, 0, s.GetIndex(two))
	ass.Equal(t, 1, s.GetIndex(three))
	ass.Equal(t, 3, s.GetIndex(five))
	ass.Equal(t, 5, s.GetIndex(seven))
	ass.Equal(t, 0, s.GetIndex(eight))
	ass.Equal(t, []abs.DurationLike{
		three,
		four,
		five,
		six,
		seven,
	}, s.AsArray())
	var iterator = col.Iterator[abs.DurationLike](s)
	ass.Equal(t, three, iterator.GetNext())
	iterator.ToEnd()
	ass.Equal(t, seven, iterator.GetPrevious())
}

func TestIntervalsWithExplicitMoments(t *tes.T) {
	var two = ele.MomentFromInt(2)
	var three = ele.MomentFromInt(3)
	var four = ele.MomentFromInt(4)
	var five = ele.MomentFromInt(5)
	var six = ele.MomentFromInt(6)
	var seven = ele.MomentFromInt(7)
	var eight = ele.MomentFromInt(8)
	var s = ran.Interval(three, abs.RIGHT, seven)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.False(t, s.ContainsValue(two))
	ass.False(t, s.ContainsValue(three))
	ass.True(t, s.ContainsValue(five))
	ass.True(t, s.ContainsValue(seven))
	ass.False(t, s.ContainsValue(eight))
	ass.Equal(t, three, s.GetFirst())
	ass.Equal(t, abs.RIGHT, s.GetExtent())
	ass.Equal(t, seven, s.GetLast())
	ass.Equal(t, four, s.GetValue(1))
	ass.Equal(t, 0, s.GetIndex(three))
	ass.Equal(t, 1, s.GetIndex(four))
	ass.Equal(t, 2, s.GetIndex(five))
	ass.Equal(t, 4, s.GetIndex(seven))
	ass.Equal(t, 0, s.GetIndex(eight))
	ass.Equal(t, []abs.MomentLike{
		four,
		five,
		six,
		seven,
	}, s.AsArray())
}

func TestIntervalsWithImplicitMoments(t *tes.T) {
	var s = ran.Interval[ele.Moment](3, abs.EXCLUSIVE, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 3, s.GetSize())
	ass.False(t, s.ContainsValue(2))
	ass.False(t, s.ContainsValue(3))
	ass.True(t, s.ContainsValue(5))
	ass.False(t, s.ContainsValue(7))
	ass.False(t, s.ContainsValue(8))
	ass.Equal(t, ele.Moment(3), s.GetFirst())
	ass.Equal(t, abs.EXCLUSIVE, s.GetExtent())
	ass.Equal(t, ele.Moment(7), s.GetLast())
	ass.Equal(t, ele.Moment(4), s.GetValue(1))
	ass.Equal(t, 0, s.GetIndex(3))
	ass.Equal(t, 1, s.GetIndex(4))
	ass.Equal(t, 3, s.GetIndex(6))
	ass.Equal(t, 0, s.GetIndex(7))
	ass.Equal(t, []ele.Moment{4, 5, 6}, s.AsArray())
}

func TestIntervalsWithEmojis(t *tes.T) {
	var r1 = ran.RuneFromInt('😀')
	var r2 = ran.RuneFromInt('😆')
	var r3 = ran.RuneFromInt('🤣')
	var s = ran.Interval[abs.RuneLike](r1, abs.INCLUSIVE, r3)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 804, s.GetSize())
	ass.True(t, s.ContainsValue(r2))
	ass.Equal(t, r1, s.GetFirst())
	ass.Equal(t, abs.INCLUSIVE, s.GetExtent())
	ass.Equal(t, r3, s.GetLast())
	ass.Equal(t, r2, s.GetValue(7))
	ass.True(t, s.ContainsValue(r2))
}

func TestIntervalsWithRunes(t *tes.T) {
	var rA = ran.RuneFromInt('A')
	var ra = ran.RuneFromInt('a')
	var rb = ran.RuneFromInt('b')
	var rc = ran.RuneFromInt('c')
	var rd = ran.RuneFromInt('d')
	var re = ran.RuneFromInt('e')
	var rf = ran.RuneFromInt('f')
	var rg = ran.RuneFromInt('g')
	var s = ran.Interval[abs.RuneLike](ra, abs.LEFT, rf)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.False(t, s.ContainsValue(rA))
	ass.True(t, s.ContainsValue(ra))
	ass.True(t, s.ContainsValue(rc))
	ass.False(t, s.ContainsValue(rf))
	ass.False(t, s.ContainsValue(rg))
	ass.Equal(t, ra, s.GetFirst())
	ass.Equal(t, abs.LEFT, s.GetExtent())
	ass.Equal(t, rf, s.GetLast())
	ass.Equal(t, ra, s.GetValue(1))
	ass.Equal(t, 1, s.GetIndex(ra))
	ass.Equal(t, 4, s.GetIndex(rd))
	ass.Equal(t, 0, s.GetIndex(rf))
	ass.Equal(t, []abs.RuneLike{ra, rb, rc, rd, re}, s.AsArray())
}

func TestIntervalsWithIntegers(t *tes.T) {
	var i1 = ran.IntegerFromInt(1)
	var i2 = ran.IntegerFromInt(2)
	var i3 = ran.IntegerFromInt(3)
	var i4 = ran.IntegerFromInt(4)
	var i5 = ran.IntegerFromInt(5)
	var i6 = ran.IntegerFromInt(6)
	var s = ran.Interval[abs.IntegerLike](i1, abs.RIGHT, i5)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.False(t, s.ContainsValue(i1))
	ass.True(t, s.ContainsValue(i2))
	ass.True(t, s.ContainsValue(i5))
	ass.False(t, s.ContainsValue(i6))
	ass.Equal(t, i1, s.GetFirst())
	ass.Equal(t, abs.RIGHT, s.GetExtent())
	ass.Equal(t, i5, s.GetLast())
	ass.Equal(t, i2, s.GetValue(1))
	ass.Equal(t, 0, s.GetIndex(i1))
	ass.Equal(t, 1, s.GetIndex(i2))
	ass.Equal(t, 4, s.GetIndex(i5))
	ass.Equal(t, 0, s.GetIndex(i6))
	ass.Equal(t, []abs.IntegerLike{i2, i3, i4, i5}, s.AsArray())
}
