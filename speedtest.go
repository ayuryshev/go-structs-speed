package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

type SpeedTest struct {
	TestFunc            func()
	ItemsDefinition     string
	OperationDefinition string
	DurationsNs         []int64
}

//MeasureDurations   runs  function f, returns time.Duration of execution or error
func (test *SpeedTest) MeasureDurations(times int) (err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error: %v\n %v\n", err, debug.Stack())
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < times; i++ {
		wg.Add(1)
		startT := time.Now()

		go func() {
			defer wg.Done()
			test.TestFunc()
		}()

		wg.Wait()
		stopT := time.Now()
		test.DurationsNs = append(test.DurationsNs, stopT.Sub(startT).Nanoseconds())
	}
	return
}

func (test SpeedTest) MinNs() (minNs int64) {
	for _, durNs := range test.DurationsNs {
		if durNs < minNs {
			minNs = durNs
		}
	}
	return
}

func (test SpeedTest) MaxNs() (maxNs int64) {
	for _, durNs := range test.DurationsNs {
		if durNs > maxNs {
			maxNs = durNs
		}
	}
	return
}

func (test SpeedTest) SumNs() (sumNs int64) {
	for _, durNs := range test.DurationsNs {
		sumNs += durNs
	}
	return
}

func (test SpeedTest) OpCostNs() (ns float64) {
	if lenDur := len(test.DurationsNs); lenDur > 0 {
		ns = math.Floor(float64(test.SumNs())/float64(TestQty*lenDur)*100) / 100
	}
	return
}

func (test SpeedTest) String() string {
	return fmt.Sprintf(`Structure: %v . Operation: %v
Durations = %v
Sum duration = %v. 
Operation cost: ns/op = %v
`, test.ItemsDefinition, test.OperationDefinition,
		test.DurationsNs, test.SumNs(), test.OpCostNs())
}

type SpeedTests []SpeedTest

func (tests SpeedTests) Len() int           { return len(tests) }
func (tests SpeedTests) Swap(i, j int)      { tests[i], tests[j] = tests[j], tests[i] }
func (tests SpeedTests) Less(i, j int) bool { return tests[i].SumNs() < tests[j].SumNs() }

const htmlTmplt = `<table>
<tr><th>Structure</th><th>Operation</th><th>Summary Duration(Ns)</th><th><th>Op Cost(ns)</th></tr>
{{with .Tests}}{{range .}}<tr>
        <td><pre>{{.ItemsDefinition}}</pre></td><td><pre>{{.OperationDefinition}}</pre></td><td>{{.SumNs}}</td>
        <td>{{.OpCostNs}}ns</td>
    </tr>{{end}}{{end}}
</table>
`

func (tests SpeedTests) HTML() {

	t := template.New("HtmlReport")
	if t, err := t.Parse(htmlTmplt); err != nil {
		log.Printf("Error %v", err)
	} else {
		if err := t.Execute(os.Stdout, struct{ Tests SpeedTests }{Tests: tests}); err != nil {
			log.Printf("Error %v", err)
		}
	}

	return
}
