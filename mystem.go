package mystem

/*
#cgo LDFLAGS: -lmystem_c_binding
#include "mystem.h"
*/
import "C"

import (
	"unsafe"
)

import "encoding/binary"

func StringToSymbols(word string) []C.ushort {
	var (
		symbols []C.ushort
	)
	for _, char := range word {
		symbols = append(symbols, C.ushort(char))
	}
	return symbols
}

func SymbolsToString(symbols *C.TSymbol, length C.int) string {
	var (
		runes []rune
	)
	bytes := C.GoBytes(unsafe.Pointer(symbols), length*2)
	for len(bytes) != 0 {
		letter := binary.LittleEndian.Uint16(bytes[:2])
		runes = append(runes, rune(letter))
		bytes = bytes[2:]
	}
	return string(runes)
}

type Analyses struct {
	handle unsafe.Pointer
}

type Lemma struct {
	handle unsafe.Pointer
}

func NewAnalyses(word string) *Analyses {
	cWord := StringToSymbols(word)
	cWordLength := len(cWord)
	analyses := new(Analyses)
	handle := C.MystemAnalyze((*C.TSymbol)(unsafe.Pointer(&cWord[0])), C.int(cWordLength))
	analyses.handle = handle
	return analyses
}

func (analyses *Analyses) Length() int {
	length := C.MystemAnalysesCount(analyses.handle)
	return (int)(length)
}

func (analyses *Analyses) Close() {
	C.MystemDeleteAnalyses(analyses.handle)
}

func (analyses *Analyses) GetLemma(i int) *Lemma {
	lemma := new(Lemma)
	handle := C.MystemLemma(unsafe.Pointer(analyses.handle), (C.int)(i))
	lemma.handle = handle
	return lemma
}

func (lemma *Lemma) TextLength() C.int {
	return C.MystemLemmaTextLen(lemma.handle)
}

func (lemma *Lemma) Text() string {
	return SymbolsToString(C.MystemLemmaText(lemma.handle), lemma.TextLength())
}
