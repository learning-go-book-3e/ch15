package cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  37,
	}
	result := CreatePerson("Dennis", 37)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestCreatePersonIgnoreDate(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  37,
	}
	result := CreatePerson("Dennis", 37)
	if diff := cmp.Diff(expected, result, cmpopts.IgnoreFields(Person{}, "DateAdded")); diff != "" {
		t.Error(diff)
	}
	if result.DateAdded.IsZero() {
		t.Error("DateAdded wasn't assigned")
	}
}
