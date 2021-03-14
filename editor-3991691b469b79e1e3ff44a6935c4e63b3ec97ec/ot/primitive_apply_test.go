package ot_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/learningspoons-go/editor/ot"
)

func TestInsertApply(t *testing.T) {
	i := ot.Insert{Text: "hello"}
	w := new(strings.Builder)
	i.Apply(nil, w)
	if w.String() != "hello" {
		t.Errorf("expected 'hello', got '%s'", w.String())
	}
}

/*func TestInsertApplyUnicode(t *testing.T) {
	i := ot.Insert{Text: "안녕하세요"}
	w := new(strings.Builder)
	i.Apply(nil, w)
	if w.String() != "안녕하세요" {
		t.Errorf("expected '안녕하세요', got '%s'", w.String())
	}
}*/

func TestDeleteApply(t *testing.T) {
	d := ot.Delete{Count: 2}
	r := bytes.NewReader([]byte("hello"))
	d.Apply(r, nil)
	if all, err := ioutil.ReadAll(r); err != nil {
		t.Error(err)
		return
	} else if string(all) != "llo" {
		t.Errorf("expected text left in reader to be 'llo', got '%s'",
			string(all))
	}
}

/*func TestDeleteApplyUnicode(t *testing.T) {
	d := ot.Delete{Count: 2}
	r := bytes.NewReader([]byte("안녕하세요"))
	d.Apply(r, nil)
	if all, err := ioutil.ReadAll(r); err != nil {
		t.Error(err)
		return
	} else if string(all) != "하세요" {
		t.Errorf("expected text left in reader to be '하세요', got '%s'",
			string(all))
	}
}*/

func TestRetainApply(t *testing.T) {
	retain := ot.Retain{Count: 2}
	r := bytes.NewReader([]byte("hello"))
	w := new(strings.Builder)
	retain.Apply(r, w)
	if w.String() != "he" {
		t.Errorf("expected 'he', got '%s'", w.String())
	}

	if all, err := ioutil.ReadAll(r); err != nil {
		t.Error(err)
		return
	} else if string(all) != "llo" {
		t.Errorf("expected text left in reader to be 'llo' got '%s'",
			string(all))
	}
}

/*func TestRetainApplyUnicode(t *testing.T) {
	retain := ot.Retain{Count: 2}
	r := bytes.NewReader([]byte("안녕하세요"))
	w := new(strings.Builder)
	retain.Apply(r, w)
	if w.String() != "안녕" {
		t.Errorf("expected '안녕', got '%s'", w.String())
	}

	if all, err := ioutil.ReadAll(r); err != nil {
		t.Error(err)
		return
	} else if string(all) != "하세요" {
		t.Errorf("expected text left in reader to be '하세요' got '%s'",
			string(all))
	}
}*/
