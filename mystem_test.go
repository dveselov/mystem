package mystem

import (
	"reflect"
	"testing"
)

func TestGetAnalysesCount(t *testing.T) {
	analyses := NewAnalyses("маша")
	defer analyses.Close()
	if analyses.Count() != 2 {
		t.Fail()
	}
}

func TestGetLemmaText(t *testing.T) {
	analyses := NewAnalyses("маша")
	defer analyses.Close()
	lemma := analyses.GetLemma(0)
	if lemma.TextLength() != 4 {
		t.Fail()
	}
	if lemma.Text() != "маша" {
		t.Fail()
	}
	lemma = analyses.GetLemma(1)
	if lemma.TextLength() != 6 {
		t.Fail()
	}
	if lemma.Text() != "махать" {
		t.Fail()
	}
}

func TestGetLemmaForm(t *testing.T) {
	analyses := NewAnalyses("маша")
	defer analyses.Close()
	lemma := analyses.GetLemma(0)
	if lemma.FormLength() != 4 {
		t.Fail()
	}
	if lemma.Form() != "маша" {
		t.Fail()
	}
}

func TestGetLemmaQuality(t *testing.T) {
	analyses := NewAnalyses("маша")
	lemma := analyses.GetLemma(0)
	if lemma.Quality() != Dictionary {
		t.Fail()
	}
	analyses.Close()
	analyses = NewAnalyses("майстем")
	defer analyses.Close()
	lemma = analyses.GetLemma(0)
	if lemma.Quality() != Bastard {
		t.Fail()
	}
}

func TestGetLemmaStemGrammemes(t *testing.T) {
	analyses := NewAnalyses("маша")
	defer analyses.Close()
	lemma := analyses.GetLemma(0)
	grammemes := lemma.StemGram()
	if !reflect.DeepEqual(grammemes, []int{Substantive, FirstName, Feminine, Animated}) {
		t.Fail()
	}
}

func TestGetLemmaFlexGrammemes(t *testing.T) {
	analyses := NewAnalyses("маша")
	defer analyses.Close()
	lemma := analyses.GetLemma(0)
	grammemes := lemma.FlexGram()
	if !reflect.DeepEqual(grammemes, [][]int{[]int{Nominative, Singular}}) {
		t.Fail()
	}
}

func TestGetLemmaForms(t *testing.T) {
	analyses := NewAnalyses("маша")
	defer analyses.Close()
	lemma := analyses.GetLemma(0)
	forms := lemma.GenerateForms()
	if forms.Count() != 10 {
		t.Fail()
	}
}

func TestGetForm(t *testing.T) {
	analyses := NewAnalyses("маша")
	defer analyses.Close()
	lemma := analyses.GetLemma(0)
	forms := lemma.GenerateForms()
	form := forms.Get(0)
	if form.TextLength() != 3 {
		t.Fail()
	}
	if form.Text() != "маш" {
		t.Fail()
	}
	if !reflect.DeepEqual(form.StemGram(), []int{Substantive, FirstName, Feminine, Animated}) {
		t.Fail()
	}
	if !reflect.DeepEqual(form.FlexGram()[0], []int{Accusative, Plural}) {
		t.Fail()
	}
}

func BenchmarkMystem(b *testing.B) {
	mystem := NewAnalyses("маша")
	defer mystem.Close()
	for i := 0; i < b.N; i++ {
		for i := 0; i < mystem.Count(); i++ {
			lemma := mystem.GetLemma(i)
			lemma.Text()
		}
	}
}
