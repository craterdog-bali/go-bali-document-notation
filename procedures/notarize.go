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

// NOTARIZE CLAUSE IMPLEMENTATION

// This constructor creates a new notarize clause.
func NotarizeClause(document, moniker abs.ExpressionLike) abs.NotarizeClauseLike {
	var v = &notarizeClause{}
	// Perform argument validation.
	v.SetDocument(document)
	v.SetMoniker(moniker)
	return v
}

// This type defines the structure and methods associated with a notarize
// clause.
type notarizeClause struct {
	document abs.ExpressionLike
	moniker  abs.ExpressionLike
}

// This method returns the document expression for this notarize clause.
func (v *notarizeClause) GetDocument() abs.ExpressionLike {
	return v.document
}

// This method sets the document expression for this notarize clause.
func (v *notarizeClause) SetDocument(document abs.ExpressionLike) {
	if document == nil {
		panic("A notarize clause requires a document expression.")
	}
	v.document = document
}

// This method returns the citation moniker for this notarize clause.
func (v *notarizeClause) GetMoniker() abs.ExpressionLike {
	return v.moniker
}

// This method sets the citation moniker for this notarize clause.
func (v *notarizeClause) SetMoniker(moniker abs.ExpressionLike) {
	if moniker == nil {
		panic("A notarize clause requires a citation moniker.")
	}
	v.moniker = moniker
}
