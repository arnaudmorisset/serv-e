package tests

import (
	"serv-e/internal"
	"testing"
)

type reverseTest struct {
	Name     string
	Input    []internal.Record
	Expected []internal.Record
}

type equalityTest struct {
	Name     string
	RecordA  internal.Record
	RecordB  internal.Record
	Expected bool
}

func recordsAreEquals(a, b []internal.Record) bool {
	if len(a) != len(b) {
		return false
	}
	for i, a_elt := range a {
		b_elt := b[i]
		if !a_elt.Equals(b_elt) {
			return false
		}
	}
	return true
}

func TestRecordEquality(t *testing.T) {
	tests := []equalityTest{
		equalityTest{
			Name:     "empty records must be equal",
			RecordA:  internal.Record{},
			RecordB:  internal.Record{},
			Expected: true,
		},
		equalityTest{
			Name:     "different IDs must not be equal",
			RecordA:  internal.Record{Id: "19:48"},
			RecordB:  internal.Record{Id: "19:49"},
			Expected: false,
		},
		equalityTest{
			Name:     "different headers must not be equal",
			RecordA:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}},
			RecordB:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"notbar"}}},
			Expected: false,
		},
		equalityTest{
			Name:     "different headers must not be equal (2)",
			RecordA:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}},
			RecordB:  internal.Record{Id: "19:48", Headers: map[string][]string{}},
			Expected: false,
		},
		equalityTest{
			Name:     "different body must not be equal",
			RecordA:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK"},
			RecordB:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: "NOK"},
			Expected: false,
		},
		equalityTest{
			Name:     "different body must not be equal (2)",
			RecordA:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: ""},
			RecordB:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: "NOK"},
			Expected: false,
		},
		equalityTest{
			Name:     "same records must be equal",
			RecordA:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"},
			RecordB:  internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"},
			Expected: true,
		},
	}

	for i, test := range tests {
		if test.RecordA.Equals(test.RecordB) != test.Expected {
			if test.Expected {
				t.Errorf("error running test %d (%s): expected equality", i, test.Name)
			} else {
				t.Errorf("error running test %d (%s): did not expected equality", i, test.Name)
			}
		}
	}
}

func TestReverseRecords(t *testing.T) {

	tests := []reverseTest{
		reverseTest{
			Name:     "empty records",
			Input:    []internal.Record{},
			Expected: []internal.Record{},
		},
		reverseTest{
			Name: "one record",
			Input: []internal.Record{
				internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"},
			},
			Expected: []internal.Record{
				internal.Record{Id: "19:48", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"},
			},
		},
		reverseTest{
			Name: "multiple records",
			Input: []internal.Record{
				internal.Record{Id: "19:47", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"},
				internal.Record{Id: "19:49", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"},
			},
			Expected: []internal.Record{
				internal.Record{Id: "19:49", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"},
				internal.Record{Id: "19:47", Headers: map[string][]string{"foo": {"bar"}}, Body: "OK!"},
			},
		},
	}

	for i, test := range tests {
		internal.ReverseRecords(&test.Input)
		if !recordsAreEquals(test.Input, test.Expected) {
			t.Errorf("error running test %d (%s): got %+v want %+v", i, test.Name, test.Input, test.Expected)
		}
	}

}
