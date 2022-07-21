package query

import (
	"strings"
	"testing"
)

func testKeyOrder(t *testing.T, f Order, keys []string, expect []string) {
	t.Helper()

	e := make([]Entry, len(keys))
	for i, k := range keys {
		e[i] = Entry{Key: k}
	}

	res := ResultsWithEntries(Query{}, e)
	res = NaiveOrder(res, f)
	actualE, err := res.Rest()
	if err != nil {
		t.Fatal(err)
	}

	actual := make([]string, len(actualE))
	for i, e := range actualE {
		actual[i] = e.Key
	}

	if len(actual) != len(expect) {
		t.Error("expect != actual.", expect, actual)
	}

	if strings.Join(actual, "") != strings.Join(expect, "") {
		t.Error("expect != actual.", expect, actual)
	}
}

func TestOrderByKey(t *testing.T) {

	testKeyOrder(t, OrderByKey{}, sampleKeys, []string{
		"/a",
		"/ab",
		"/ab/c",
		"/ab/cd",
		"/ab/ef",
		"/ab/fg",
		"/abce",
		"/abcf",
	})
	testKeyOrder(t, OrderByKeyDescending{}, sampleKeys, []string{
		"/abcf",
		"/abce",
		"/ab/fg",
		"/ab/ef",
		"/ab/cd",
		"/ab/c",
		"/ab",
		"/a",
	})
}
