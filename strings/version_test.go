/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings_test

import (
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestBadVersion(t *tes.T) {
	var _, ok = str.VersionFromString(`v0`)
	ass.False(t, ok)
}

func TestEmptyVersion(t *tes.T) {
	var _, ok = str.VersionFromString(`v`)
	ass.False(t, ok)
}

func TestVersion(t *tes.T) {
	var v1, ok = str.VersionFromString("v1.2.3")
	ass.True(t, ok)
	ass.Equal(t, "v1.2.3", lan.FormatValue(v1))
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 3, v1.GetSize())
	ass.Equal(t, 1, v1.GetItem(1))
	ass.Equal(t, 3, v1.GetItem(-1))
	var v2 = str.VersionFromOrdinals(v1.AsArray())
	ass.Equal(t, v1.String(), lan.FormatValue(v2))
	var v3 = str.VersionFromOrdinals(v1.GetItems(1, 2))
	ass.Equal(t, "v1.2", string(v3))
	ass.Equal(t, 3, v1.GetIndex(3))
}

func TestVersionsLibrary(t *tes.T) {
	var v1, _ = str.VersionFromString("v1")
	var v2, _ = str.VersionFromString("v2.3")
	var v3 = str.Versions.Concatenate(v1, v2)
	ass.Equal(t, "v1.2.3", lan.FormatValue(v3))

	ass.False(t, str.Versions.IsValidNextVersion(v1, v1))
	ass.Equal(t, "v2", lan.FormatValue(str.Versions.GetNextVersion(v1, 1)))
	ass.True(t, str.Versions.IsValidNextVersion(v1, str.Versions.GetNextVersion(v1, 1)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v1, 1), v1))
	ass.Equal(t, "v1.1", lan.FormatValue(str.Versions.GetNextVersion(v1, 2)))
	ass.True(t, str.Versions.IsValidNextVersion(v1, str.Versions.GetNextVersion(v1, 2)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v1, 2), v1))
	ass.Equal(t, "v1.1", lan.FormatValue(str.Versions.GetNextVersion(v1, 2)))
	ass.Equal(t, "v1.1", lan.FormatValue(str.Versions.GetNextVersion(v1, 3)))

	ass.False(t, str.Versions.IsValidNextVersion(v2, v2))
	ass.Equal(t, "v3", lan.FormatValue(str.Versions.GetNextVersion(v2, 1)))
	ass.True(t, str.Versions.IsValidNextVersion(v2, str.Versions.GetNextVersion(v2, 1)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v2, 1), v2))
	ass.Equal(t, "v2.4", lan.FormatValue(str.Versions.GetNextVersion(v2, 2)))
	ass.True(t, str.Versions.IsValidNextVersion(v2, str.Versions.GetNextVersion(v2, 2)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v2, 2), v2))
	ass.Equal(t, "v2.3.1", lan.FormatValue(str.Versions.GetNextVersion(v2, 3)))
	ass.True(t, str.Versions.IsValidNextVersion(v2, str.Versions.GetNextVersion(v2, 3)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v2, 3), v2))

	ass.False(t, str.Versions.IsValidNextVersion(v3, v3))
	ass.Equal(t, "v2", lan.FormatValue(str.Versions.GetNextVersion(v3, 1)))
	ass.True(t, str.Versions.IsValidNextVersion(v3, str.Versions.GetNextVersion(v3, 1)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v3, 1), v3))
	ass.Equal(t, "v1.3", lan.FormatValue(str.Versions.GetNextVersion(v3, 2)))
	ass.True(t, str.Versions.IsValidNextVersion(v3, str.Versions.GetNextVersion(v3, 2)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v3, 2), v3))
	ass.Equal(t, "v1.2.4", lan.FormatValue(str.Versions.GetNextVersion(v3, 3)))
	ass.True(t, str.Versions.IsValidNextVersion(v3, str.Versions.GetNextVersion(v3, 3)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v3, 3), v3))
	ass.Equal(t, "v1.2.3.1", lan.FormatValue(str.Versions.GetNextVersion(v3, 4)))
	ass.True(t, str.Versions.IsValidNextVersion(v3, str.Versions.GetNextVersion(v3, 4)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v3, 4), v3))
}
