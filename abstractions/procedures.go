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

// ASSIGNMENT CONSTANTS

type Assignment int

const (
	_ Assignment = iota
	REGULAR
	DEFAULT
	TIMES
	DIVIDE
	PLUS
	MINUS
)

// STATEMENT INTERFACES

// This interface defines the methods supported by all accept-clause-like types.
type AcceptClauseLike interface {
	GetMessage() ExpressionLike
	SetMessage(message ExpressionLike)
}

// This interface defines the methods supported by all attribute-like types.
type AttributeLike interface {
	GetVariable() string
	SetVariable(variable string)
	GetIndex(index int) ExpressionLike
	SetIndex(index int, expression ExpressionLike)
	GetIndices() ListLike[ExpressionLike]
	SetIndices(indices ListLike[ExpressionLike])
}

// This interface defines the methods supported by all block-like types.
type BlockLike interface {
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetProcedure() ProcedureLike
	SetProcedure(procedure ProcedureLike)
}

// This interface defines the methods supported by all break-clause-like types.
type BreakClauseLike interface {
}

// This interface defines the methods supported by all checkout-clause-like types.
type CheckoutClauseLike interface {
	GetRecipient() any
	SetRecipient(recipient any)
	GetLevel() ExpressionLike
	SetLevel(level ExpressionLike)
	GetMoniker() ExpressionLike
	SetMoniker(moniker ExpressionLike)
}

// This interface defines the methods supported by all continue-clause-like types.
type ContinueClauseLike interface {
}

// This interface defines the methods supported by all discard-clause-like types.
type DiscardClauseLike interface {
	GetCitation() ExpressionLike
	SetCitation(citation ExpressionLike)
}

// This interface defines the methods supported by all evaluate-clause-like types.
type EvaluateClauseLike interface {
	GetRecipient() (recipient any, assignment Assignment)
	SetRecipient(recipient any, assignment Assignment)
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
}

// This interface defines the methods supported by all if-clause-like types.
type IfClauseLike interface {
	GetCondition() ExpressionLike
	SetCondition(condition ExpressionLike)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetStatements() ProcedureLike
	SetStatements(statements ProcedureLike)
}

// This interface defines the methods supported by all notarize-clause-like types.
type NotarizeClauseLike interface {
	GetDraft() ExpressionLike
	SetDraft(draft ExpressionLike)
	GetMoniker() ExpressionLike
	SetMoniker(moniker ExpressionLike)
}

// This interface defines the methods supported by all on-clause-like types.
type OnClauseLike interface {
	GetException() Symbolic
	SetException(exception Symbolic)
	GetHandler(index int) BlockLike
	SetHandler(index int, handler BlockLike)
	GetHandlers() ListLike[BlockLike]
	SetHandlers(blocks ListLike[BlockLike])
}

// This interface defines the methods supported by all post-clause-like types.
type PostClauseLike interface {
	GetMessage() ExpressionLike
	SetMessage(message ExpressionLike)
	GetBag() ExpressionLike
	SetBag(bag ExpressionLike)
}

// This interface consolidates all the interfaces supported by procedure-like
// types.
type ProcedureLike interface {
	Sequential[StatementLike]
	Indexed[StatementLike]
	Updatable[StatementLike]
	Malleable[StatementLike]
}

// This interface defines the methods supported by all publish-clause-like types.
type PublishClauseLike interface {
	GetEvent() ExpressionLike
	SetEvent(event ExpressionLike)
}

// This interface defines the methods supported by all reject-clause-like types.
type RejectClauseLike interface {
	GetMessage() ExpressionLike
	SetMessage(message ExpressionLike)
}

// This interface defines the methods supported by all retrieve-clause-like types.
type RetrieveClauseLike interface {
	GetRecipient() any
	SetRecipient(recipient any)
	GetBag() ExpressionLike
	SetBag(bag ExpressionLike)
}

// This interface defines the methods supported by all return-clause-like types.
type ReturnClauseLike interface {
	GetResult() ExpressionLike
	SetResult(result ExpressionLike)
}

// This interface defines the methods supported by all save-clause-like types.
type SaveClauseLike interface {
	GetDraft() ExpressionLike
	SetDraft(draft ExpressionLike)
	GetRecipient() any
	SetRecipient(recipient any)
}

// This interface defines the methods supported by all select-clause-like types.
type SelectClauseLike interface {
	GetControl() ExpressionLike
	SetControl(control ExpressionLike)
	GetOption(index int) BlockLike
	SetOption(index int, option BlockLike)
	GetOptions() ListLike[BlockLike]
	SetOptions(blocks ListLike[BlockLike])
}

// This interface defines the methods supported by all statement-like types.
type StatementLike interface {
	GetAnnotation() string
	SetAnnotation(annotation string)
	GetMainClause() any
	SetMainClause(mainClause any)
	GetNote() string
	SetNote(note string)
	GetOnClause() OnClauseLike
	SetOnClause(onClause OnClauseLike)
}

// This interface defines the methods supported by all throw-clause-like types.
type ThrowClauseLike interface {
	GetException() ExpressionLike
	SetException(exception ExpressionLike)
}

// This interface defines the methods supported by all while-clause-like types.
type WhileClauseLike interface {
	GetCondition() ExpressionLike
	SetCondition(condition ExpressionLike)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetStatements() ProcedureLike
	SetStatements(statements ProcedureLike)
}

// This interface defines the methods supported by all with-clause-like types.
type WithClauseLike interface {
	GetItem() Symbolic
	SetItem(exception Symbolic)
	GetSequence() ExpressionLike
	SetSequence(sequence ExpressionLike)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetStatements() ProcedureLike
	SetStatements(statements ProcedureLike)
}
