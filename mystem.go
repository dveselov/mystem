package mystem

/*
#cgo LDFLAGS: -lmystem_c_binding
#include "mystem.h"
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

// Lemma quality constants
const (
	Dictionary  = 0        // слово из словаря
	Bastard     = 1        // не словарное
	Sob         = 2        // из "быстрого словаря"
	Prefixoid   = 4        // словарное + стандартный префикс (авто- мото- кино- фото-) всегда в компании с Bastard или Sob
	Foundling   = 8        // непонятный набор букв, но проходящий в алфавит
	BadRequest  = 16       // доп. флаг.: "плохая лемма" при наличии "хорошей" альтернативы ("махать" по форме "маша")
	FromEnglish = 65536    // переведено с английского
	ToEnglish   = 131072   // переведено на английский
	Untranslit  = 262144   // "переведено" с транслита
	Overrode    = 1048576  // текст леммы был перезаписан
	Fix         = 16777216 // слово из фикс-листа
)

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
