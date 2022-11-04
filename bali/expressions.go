/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
	exp "github.com/craterdog-bali/go-bali-document-notation/expressions"
)

// This method attempts to parse the arguments within a call. It returns a
// list of the arguments and whether or not the arguments were successfully
// parsed.
func (v *parser) parseArguments() (abs.ListLike[abs.ExpressionLike], *Token, bool) {
	var ok bool
	var token *Token
	var argument abs.ExpressionLike
	var arguments = col.List[abs.ExpressionLike]()
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an argument expression.
		return arguments, token, false
	}
	argument, token, ok = v.parseExpression()
	for ok {
		arguments.AddValue(argument)
		// Every subsequent argument must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more arguments.
			break
		}
		argument, token, ok = v.parseExpression()
		if !ok {
			var message = v.formatError("An unexpected token was received by the parser:", token)
			message += generateGrammar("$expression",
				"$arguments")
			panic(message)
		}
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar(")",
			"$intrinsic",
			"$function",
			"$arguments")
		panic(message)
	}
	return arguments, token, true
}

// This method adds the canonical format for the specified sequence of arguments
// to the state of the formatter.
func (v *formatter) formatArguments(arguments abs.Sequential[abs.ExpressionLike]) {
	v.state.AppendString("(")
	var iterator = age.Iterator(arguments)
	if iterator.HasNext() {
		var argument = iterator.GetNext()
		v.formatExpression(argument)
	}
	for iterator.HasNext() {
		v.state.AppendString(", ")
		var argument = iterator.GetNext()
		v.formatExpression(argument)
	}
	v.state.AppendString(")")
}

// This method attempts to parse a arithmetic expression. It returns the
// arithmetic expression and whether or not the arithmetic expression was
// successfully parsed.
func (v *parser) parseArithmetic(first abs.ExpressionLike) (abs.ArithmeticLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.ExpressionLike
	var expression abs.ArithmeticLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an arithmetic expression.
		return expression, token, false
	}
	if operator < abs.PLUS || operator > abs.MODULO {
		// This is not an arithmetic expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	second, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$arithmetic")
		panic(message)
	}
	expression = exp.Arithmetic(first, operator, second)
	return expression, token, true
}

// This method adds the canonical format for the specified arithmetic expression
// to the state of the formatter.
func (v *formatter) formatArithmetic(arithmetic abs.ArithmeticLike) {
	var first = arithmetic.GetFirst()
	v.formatExpression(first)
	var operator = arithmetic.GetOperator()
	v.formatOperator(operator)
	var second = arithmetic.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a chain expression. It returns the
// chain expression and whether or not the chain expression was
// successfully parsed.
func (v *parser) parseChaining(first abs.ExpressionLike) (abs.ChainingLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.ExpressionLike
	var expression abs.ChainingLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a chaining expression.
		return expression, token, false
	}
	if operator != abs.AMPERSAND {
		// This is not a chaining expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	second, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$chaining")
		panic(message)
	}
	expression = exp.Chaining(first, second)
	return expression, token, true
}

// This method adds the canonical format for the specified chaining expression
// to the state of the formatter.
func (v *formatter) formatChaining(chaining abs.ChainingLike) {
	var first = chaining.GetFirst()
	v.formatExpression(first)
	v.state.AppendString(" & ")
	var second = chaining.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a comparison expression. It returns the
// comparison expression and whether or not the comparison expression was
// successfully parsed.
func (v *parser) parseComparison(first abs.ExpressionLike) (abs.ComparisonLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.ExpressionLike
	var expression abs.ComparisonLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a comparison expression.
		return expression, token, false
	}
	if operator < abs.LESS || operator > abs.MATCHES {
		// This is not a comparison expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	second, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$comparison")
		panic(message)
	}
	expression = exp.Comparison(first, operator, second)
	return expression, token, true
}

// This method attempts to parse a complement expression. It returns the
// complement expression and whether or not the complement expression was
// successfully parsed.
func (v *parser) parseComplement() (abs.ComplementLike, *Token, bool) {
	var ok bool
	var token *Token
	var logical abs.ExpressionLike
	var expression abs.ComplementLike
	_, token, ok = v.parseKeyword("NOT")
	if !ok {
		// This is not an complement expression.
		return expression, token, false
	}
	logical, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$complement")
		panic(message)
	}
	expression = exp.Complement(logical)
	return expression, token, true
}

// This method attempts to parse a dereference expression. It returns the
// dereference expression and whether or not the dereference expression was
// successfully parsed.
func (v *parser) parseDereference() (abs.DereferenceLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var reference abs.ExpressionLike
	var expression abs.DereferenceLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a dereference expression.
		return expression, token, false
	}
	switch operator {
	case abs.AT:
		// Found a valid operator.
	default:
		// This is not a dereference expression.
		return expression, token, false
	}
	reference, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$dereference")
		panic(message)
	}
	expression = exp.Dereference(reference)
	return expression, token, true
}

// This method attempts to parse a power expression. It returns the
// power expression and whether or not the power expression was
// successfully parsed.
func (v *parser) parseExponential(base abs.ExpressionLike) (abs.ExponentialLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var exponent abs.ExpressionLike
	var expression abs.ExponentialLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an exponential expression.
		return expression, token, false
	}
	if operator  != abs.CARET {
		// This is not an exponential expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	exponent, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$exponential")
		panic(message)
	}
	expression = exp.Exponential(base, exponent)
	return expression, token, true
}

// This method attempts to parse an expression. It returns the expression and
// whether or not the expression was successfully parsed. The expressions are
// are checked in precedence order from highest to lowest precedence.
func (v *parser) parseExpression() (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var expression abs.ExpressionLike
	expression, token, ok = v.parseComponent()
	if !ok {
		// This must come before the parseIdentifier() for a variable.
		expression, token, ok = v.parseIntrinsic()
	}
	if !ok {
		expression, token, ok = v.parseVariable()
	}
	if !ok {
		expression, token, ok = v.parsePrecedence()
	}
	if !ok {
		expression, token, ok = v.parseDereference()
	}
	if !ok {
		expression, token, ok = v.parseRecursive()
	}
	if !ok {
		expression, token, ok = v.parseInversion()
	}
	if !ok {
		expression, token, ok = v.parseMagnitude()
	}
	if !ok {
		expression, token, ok = v.parseComplement()
	}
	return expression, token, ok
}

// This method attempts to parse a function expression. It returns the
// function expression and whether or not the function expression was
// successfully parsed.
func (v *parser) parseIntrinsic() (abs.IntrinsicLike, *Token, bool) {
	var ok bool
	var token *Token
	var function string
	var arguments abs.ListLike[abs.ExpressionLike]
	var expression abs.IntrinsicLike
	function, token, ok = v.parseIdentifier()
	if !ok {
		// This is not an function expression.
		return expression, token, false
	}
	arguments, token, ok = v.parseArguments()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$intrinsic",
			"$function",
			"$arguments")
		panic(message)
	}
	expression = exp.Intrinsic(function, arguments)
	return expression, token, true
}

// This method attempts to parse a inversion expression. It returns the
// inversion expression and whether or not the inversion expression was
// successfully parsed.
func (v *parser) parseInversion() (abs.InversionLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var numeric abs.ExpressionLike
	var expression abs.InversionLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an inversion expression.
		return expression, token, false
	}
	if operator < abs.PLUS || operator > abs.STAR {
		// This is not an inversion expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	numeric, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$inversion")
		panic(message)
	}
	expression = exp.Inversion(operator, numeric)
	return expression, token, true
}

// This method attempts to parse a message expression. It returns the
// message expression and whether or not the message expression was
// successfully parsed.
func (v *parser) parseInvocation(target abs.ExpressionLike) (abs.InvocationLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var message string
	var arguments abs.ListLike[abs.ExpressionLike]
	var expression abs.InvocationLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an invocation expression.
		return expression, token, false
	}
	if operator < abs.DOT || operator > abs.ARROW {
		// This is not an invocation expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	message, token, ok = v.parseIdentifier()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$method",
			"$invocation",
			"$method",
			"$arguments")
		panic(message)
	}
	arguments, token, ok = v.parseArguments()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$invocation",
			"$method",
			"$arguments")
		panic(message)
	}
	expression = exp.Invocation(target, operator, message, arguments)
	return expression, token, true
}

// This method attempts to parse a logical expression. It returns the
// logical expression and whether or not the logical expression was
// successfully parsed.
func (v *parser) parseLogical(first abs.ExpressionLike) (abs.LogicalLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.ExpressionLike
	var expression abs.LogicalLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a logical expression.
		return expression, token, false
	}
	if operator < abs.NOT || operator > abs.XOR {
		// This is not a logical expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	second, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$logical")
		panic(message)
	}
	expression = exp.Logical(first, operator, second)
	return expression, token, true
}

// This method attempts to parse a magnitude expression. It returns the
// magnitude expression and whether or not the magnitude expression was
// successfully parsed.
func (v *parser) parseMagnitude() (abs.MagnitudeLike, *Token, bool) {
	var ok bool
	var token *Token
	var numeric abs.ExpressionLike
	var expression abs.MagnitudeLike
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		// This is not an magnitude expression.
		return expression, token, false
	}
	numeric, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$magnitude")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("|",
			"$magnitude")
		panic(message)
	}
	expression = exp.Magnitude(numeric)
	return expression, token, true
}

// This method attempts to parse a precedence expression. It returns the
// precedence expression and whether or not the precedence expression was
// successfully parsed.
func (v *parser) parsePrecedence() (abs.PrecedenceLike, *Token, bool) {
	var ok bool
	var token *Token
	var inner abs.ExpressionLike
	var expression abs.PrecedenceLike
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an precedence expression.
		return expression, token, false
	}
	inner, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$precedence")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$precedence")
		panic(message)
	}
	expression = exp.Precedence(inner)
	return expression, token, true
}

// This method attempts to parse a left recursive expression. It returns
// the left recursive expression and whether or not the left recursive
// expression was successfully parsed.
func (v *parser) parseRecursive() (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var expression abs.ExpressionLike
	expression, token, ok = v.parseExpression()
	if !ok {
		// This is not a left recursive expression.
		return expression, token, false
	}
	expression, token, ok = v.parseInvocation(expression)
	if !ok {
		expression, token, ok = v.parseItem(expression)
	}
	if !ok {
		expression, token, ok = v.parseChaining(expression)
	}
	if !ok {
		expression, token, ok = v.parseExponential(expression)
	}
	if !ok {
		expression, token, ok = v.parseArithmetic(expression)
	}
	if !ok {
		expression, token, ok = v.parseComparison(expression)
	}
	if !ok {
		expression, token, ok = v.parseLogical(expression)
	}
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("operator",
			"$invocation",
			"$item",
			"$chaining",
			"$exponential",
			"$arithmetic",
			"$comparison",
			"$logical")
		panic(message)
	}
	return expression, token, true
}

// This method attempts to parse an item expression. It returns the
// item expression and whether or not the item expression was successfully
// parsed.
func (v *parser) parseItem(composite abs.ExpressionLike) (abs.ItemLike, *Token, bool) {
	var ok bool
	var token *Token
	var indices abs.ListLike[abs.ExpressionLike]
	var expression abs.ItemLike
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		// This is not an item expression.
		return expression, token, false
	}
	indices, token, ok = v.parseIndices()
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$expression",
			"$item",
			"$composite",
			"$indices")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = v.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("]",
			"$item",
			"$composite",
			"$indices")
		panic(message)
	}
	expression = exp.Value(composite, indices)
	return expression, token, true
}

// This method attempts to parse the an operator. It returns the operator and
// whether or not the operator was successfully parsed.
func (v *parser) parseOperator() (abs.Operator, *Token, bool) {
	var token = v.nextToken()
	var operator abs.Operator
	if token.Type != TokenDelimiter {
		v.backupOne()
		return operator, token, false
	}
	switch token.Value {
	case "&":
		operator = abs.AMPERSAND
	case "<-":
		operator = abs.ARROW
	case ":=":
		operator = abs.ASSIGN
	case "@":
		operator = abs.AT
	case "|":
		operator = abs.BAR
	case "^":
		operator = abs.CARET
	case "?=":
		operator = abs.DEFAULT
	case "-=":
		operator = abs.DIFFERENCE
	case ".":
		operator = abs.DOT
	case "=":
		operator = abs.EQUAL
	case "IS":
		operator = abs.IS
	case "<":
		operator = abs.LESS
	case "MATCHES":
		operator = abs.MATCHES
	case "-":
		operator = abs.MINUS
	case "//":
		operator = abs.MODULO
	case ">":
		operator = abs.MORE
	case "NOT":
		operator = abs.NOT
	case "OR":
		operator = abs.OR
	case "+":
		operator = abs.PLUS
	case "*=":
		operator = abs.PRODUCT
	case "/=":
		operator = abs.QUOTIENT
	case "SANS":
		operator = abs.SANS
	case "/":
		operator = abs.SLASH
	case "*":
		operator = abs.STAR
	case "+=":
		operator = abs.SUM
	case "~":
		operator = abs.TILDA
	case "≠":
		operator = abs.UNEQUAL
	case "XOR":
		operator = abs.XOR
	default:
		var message = v.formatError("An unexpected operator token was received by the parser:", token)
		message += generateGrammar("operator",
			"$expression")
		panic(message)
	}
	return operator, token, true
}

// This method adds the canonical format for the specified operator to the
// state of the formatter.
func (v *formatter) formatOperator(operator abs.Operator) {
	switch operator {
	case abs.AMPERSAND:
		v.state.AppendString("&")
	case abs.ARROW:
		v.state.AppendString("<-")
	case abs.ASSIGN:
		v.state.AppendString(":=")
	case abs.AT:
		v.state.AppendString("@")
	case abs.BAR:
		v.state.AppendString("|")
	case abs.CARET:
		v.state.AppendString("^")
	case abs.DEFAULT:
		v.state.AppendString("?=")
	case abs.DIFFERENCE:
		v.state.AppendString("-=")
	case abs.DOT:
		v.state.AppendString(".")
	case abs.EQUAL:
		v.state.AppendString("=")
	case abs.IS:
		v.state.AppendString("IS")
	case abs.LESS:
		v.state.AppendString("<")
	case abs.MATCHES:
		v.state.AppendString("MATCHES")
	case abs.MINUS:
		v.state.AppendString("-")
	case abs.MODULO:
		v.state.AppendString("//")
	case abs.MORE:
		v.state.AppendString(">")
	case abs.NOT:
		v.state.AppendString("NOT")
	case abs.OR:
		v.state.AppendString("OR")
	case abs.PLUS:
		v.state.AppendString("+")
	case abs.PRODUCT:
		v.state.AppendString("*=")
	case abs.QUOTIENT:
		v.state.AppendString("/=")
	case abs.SANS:
		v.state.AppendString("SANS")
	case abs.SLASH:
		v.state.AppendString("/")
	case abs.STAR:
		v.state.AppendString("*")
	case abs.SUM:
		v.state.AppendString("+=")
	case abs.TILDA:
		v.state.AppendString("~")
	case abs.UNEQUAL:
		v.state.AppendString("≠")
	case abs.XOR:
		v.state.AppendString("XOR")
	default:
		var message = fmt.Sprintf("An unexpected operator token was received by the formatter: %v\n", operator)
		panic(message)
	}
}

// This method attempts to parse an identifier. It returns the identifier
// string and whether or not the identifier was successfully parsed.
func (v *parser) parseVariable() (abs.VariableLike, *Token, bool) {
	var variable abs.VariableLike
	var token = v.nextToken()
	if token.Type != TokenIdentifier {
		v.backupOne()
		return variable, token, false
	}
	variable = exp.Variable(token.Value)
	return variable, token, true
}

// This method adds the canonical format for the specified identifier to the
// state of the formatter.
func (v *formatter) formatVariable(variable abs.VariableLike) {
	var identifier = variable.GetIdentifier()
	v.state.AppendString(identifier)
}
