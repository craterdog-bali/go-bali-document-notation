/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// PUBLISH CLAUSE IMPLEMENTATION

// This constructor creates a new publish clause.
func PublishClause(event abs.ExpressionLike) abs.PublishClauseLike {
	var v = &publishClause{}
	// Perform argument validation.
	v.SetEvent(event)
	return v
}

// This type defines the structure and methods associated with a publish
// clause.
type publishClause struct {
	event abs.ExpressionLike
}

// This method returns the event expression for this publish clause.
func (v *publishClause) GetEvent() abs.ExpressionLike {
	return v.event
}

// This method sets the event expression for this publish clause.
func (v *publishClause) SetEvent(event abs.ExpressionLike) {
	if event == nil {
		panic("A publish clause requires an event.")
	}
	v.event = event
}