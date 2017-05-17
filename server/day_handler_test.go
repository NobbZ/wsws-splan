package server

import (
	"net/http"
	"sort"
	"strings"
	"sync"
	"testing"
)

func init() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go Start("localhost:9191", wg)
	wg.Wait()
}

func TestDayOptions(t *testing.T) {
	testCases := []struct {
		url     string
		methods []string
	}{
		{"http://localhost:9191/day", []string{"GET", "OPTIONS"}},
		{"http://localhost:9191/day/1", []string{"GET", "OPTIONS"}},
		{"http://localhost:9191/coffee", []string{"GET", "PUT", "OPTIONS"}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.url, func(t *testing.T) {
			req, _ := http.NewRequest("OPTIONS", testCase.url, nil)
			c := new(http.Client)
			resp, err := c.Do(req)

			if err != nil {
				t.Errorf("There was an error while doing OPTIONS for %s:\n%v", testCase.url, err)
			}

			parsedMethods := strings.Split(resp.Header["Allow"][0], ",")
			for i := range parsedMethods {
				parsedMethods[i] = strings.Fields(parsedMethods[i])[0]
			}

			if act, exp, eq := checkOptions(parsedMethods, testCase.methods); !eq {
				t.Errorf("Expected %s to allow %v, but it allowed %v.", testCase.url, exp, act)
			}
		})
	}
}

func checkOptions(got, want []string) ([]string, []string, bool) {
	act := make([]string, len(got))
	exp := make([]string, len(want))

	copy(act, got)
	copy(exp, want)

	sort.Strings(act)
	sort.Strings(exp)

	eq := len(act) == len(exp)
	if eq {
	loop:
		for i := range act {
			if act[i] != exp[i] {
				eq = false
				break loop
			}
		}
	}

	return act, exp, eq
}
