package compiler

import (
	"go/scanner"
	"go/token"
)

// A Token is a lexical unit returned by Scan.
//
type Token struct {
	Kind    uint
	Pos     token.Pos
	Literal string
}

// An ScanErrorHandler may be provided to Scanner.Init. If a syntax error is
// encountered and a handler was installed, the handler is called with a
// position and an error message. The position points to the beginning of
// the offending token.
//
type ScanErrorHandler func(pos token.Position, msg string)

// A Scanner holds the scanner's internal state while processing
// a given text.  It can be allocated as part of another data
// structure but must be initialized via Init before use.
//
type Scanner struct {
	// immutable state
	file   *token.File      // source file handle
	dir    string           // directory portion of file.Name()
	src    []byte           // source
	errors ErrorList        // errors
	err    ScanErrorHandler // error reporting; or nil
	mode   ScanMode         // scanning mode

	// scanning state
	ch         rune // current character
	offset     int  // character offset
	rdOffset   int  // reading offset (position after current character)
	lineOffset int  // current line offset
	insertSemi bool // insert a semicolon before next newline

	// public state - ok to modify
	ErrorCount int // number of errors encountered
}

const bom = 0xFEFF // byte order mark, only permitted as very first character
