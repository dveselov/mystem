package mystem

/*
#cgo LDFLAGS: -lmystem_c_binding
#include "mystem.h"

char* get_flex_gram_by_id(char** grammemes, int id) {
	return grammemes[id];
};
*/
import "C"

import (
	"encoding/binary"
	"unsafe"
)

func stringToSymbols(word string) []C.ushort {
	var (
		symbols []C.ushort
	)
	for _, char := range word {
		symbols = append(symbols, C.ushort(char))
	}
	return symbols
}

func symbolsToString(symbols *C.TSymbol, length C.int) string {
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

func NewAnalyses(word string) *Analyses {
	cWord := stringToSymbols(word)
	cWordLength := len(cWord)
	analyses := new(Analyses)
	handle := C.MystemAnalyze((*C.TSymbol)(unsafe.Pointer(&cWord[0])), C.int(cWordLength))
	analyses.handle = handle
	return analyses
}

func (analyses *Analyses) Count() int {
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

type Lemma struct {
	handle unsafe.Pointer
}

func (lemma *Lemma) TextLength() C.int {
	return C.MystemLemmaTextLen(lemma.handle)
}

func (lemma *Lemma) Text() string {
	return symbolsToString(C.MystemLemmaText(lemma.handle), lemma.TextLength())
}

func (lemma *Lemma) FormLength() C.int {
	return C.MystemLemmaFormLen(lemma.handle)
}

func (lemma *Lemma) Form() string {
	return symbolsToString(C.MystemLemmaForm(lemma.handle), lemma.FormLength())
}

func (lemma *Lemma) Quality() int {
	return int(C.MystemLemmaQuality(lemma.handle))
}

func (lemma *Lemma) StemGram() []int {
	var (
		grammemes []int
	)
	rawGrammemes := []byte(C.GoString(C.MystemLemmaStemGram(lemma.handle)))
	for _, grammeme := range rawGrammemes {
		grammemes = append(grammemes, int(grammeme))
	}
	return grammemes
}

func (lemma *Lemma) FlexGramNum() int {
	return int(C.MystemLemmaFlexGramNum(lemma.handle))
}

func (lemma *Lemma) FlexGram() [][]int {
	var (
		grammemes [][]int
	)
	gramCount := lemma.FlexGramNum()
	rawGram := C.MystemLemmaFlexGram(lemma.handle)
	for i := 0; i < gramCount; i++ {
		var (
			currentGramSet []int
		)
		currentRawGramSet := []byte(C.GoString(C.get_flex_gram_by_id(rawGram, C.int(i))))
		for _, grammeme := range currentRawGramSet {
			currentGramSet = append(currentGramSet, int(grammeme))
		}
		grammemes = append(grammemes, currentGramSet)
	}
	return grammemes
}
