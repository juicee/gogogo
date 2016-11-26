package main

import (
	_ "bytes"
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	mode    []float64
	stddev  float64
}

type solution struct {
	solve1 complex128
	solve2 complex128
}

type quadraticParam struct {
	a int64
	b int64
	c int64
}

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom       = `</body></html>`
	anError          = `<p class="error">%s</p>`
	pageTopQuadratic = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Quadratic</h3>
<p>Quadratic Equation Solver</p>`
	formQuadratic = `<form action="/" method="POST">
<label for="numbers">Solves equations of the form ax² + bx + c</label><br />
<pre>
<input type="text" name="factA" size="5">x² + <input type="text" name="factB" size="5">x + <input type="text" name="factC" size="5"> -> <input type="submit" value="Calculate">
</pre>
</form>`
)

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) (median float64) {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func findModes(numbers []float64) (modes []float64) {
	mapMode := map[float64]int{} // 存储数字对应出现次数
	maxNum := 0                  // 存储当前出现次数的最大值
	for _, num := range numbers {
		if _, ok := mapMode[num]; ok {
			mapMode[num] += 1
			if mapMode[num] == maxNum {
				modes = append(modes, num)
			} else if mapMode[num] > maxNum {
				modes = nil
				modes = append(modes, num)
				maxNum = mapMode[num]
			} else {
				continue
			}
		} else {
			mapMode[num] = 1
		}
	}
	if len(modes) == len(mapMode) {
		modes = []float64{}
	}
	return modes
}

func calcStddev(numbers []float64) float64 {
	lnum := float64(len(numbers))
	if lnum <= 1 {
		return 0
	}
	var total float64
	for _, num := range numbers {
		total += num
	}
	avg := total / lnum
	var sumPow float64
	for _, num := range numbers {
		sumPow += math.Pow(num-avg, 2)
	}
	stddev := math.Sqrt(sumPow / (lnum - 1))
	return stddev
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	stats.mode = findModes(numbers)
	stats.stddev = calcStddev(numbers)
	return stats
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))

		} else {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func homePageQuadratic(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTopQuadratic, formQuadratic)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if params, message, ok := processQuadraticRequest(request); ok {
			solutions := solve(float64(params.a), float64(params.b), float64(params.c))
			fmt.Fprint(writer, formatSolutions(solutions, params))

		} else {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Mode</td><td>%v</td></tr>
<tr><td>Std. Dev.</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.mode, stats.stddev)
}
func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

func processQuadraticRequest(request *http.Request) (quadraticParam, string, bool) {
	var params = quadraticParam{}
	paramA, foundA := request.Form["factA"]
	paramB, foundB := request.Form["factB"]
	paramC, foundC := request.Form["factC"]
	if foundA && len(paramA) > 0 {
		if x, err := strconv.ParseInt(paramA[0], 10, 64); err != nil {
			return params, "'factA' is invalid", false
		} else {
			params.a = x
		}
	}
	if foundB && len(paramB) > 0 {
		if x, err := strconv.ParseInt(paramB[0], 10, 64); err != nil {
			return params, "'factB' is invalid", false
		} else {
			params.b = x
		}
	}
	if foundC && len(paramC) > 0 {
		if x, err := strconv.ParseInt(paramC[0], 10, 64); err != nil {
			return params, "'factC' is invalid", false
		} else {
			params.c = x
		}
	}

	if !foundA || !foundB || !foundC {
		return params, "", false
	}
	return params, "", true
}

func solve(a float64, b float64, c float64) solution {
	judgeSign := float64(math.Pow(b, 2) - 4*a*c)
	var solutions solution
	if judgeSign >= 0 {
		solutions.solve1 = complex((-b+math.Sqrt(judgeSign))/(2*a), 0)
		solutions.solve2 = complex((-b-math.Sqrt(judgeSign))/(2*a), 0)
	} else {
		cmplxJudge := cmplx.Sqrt(complex(judgeSign, 0))
		solutions.solve1 = (complex(-b, 0) + cmplxJudge) / complex(2*a, 0)
		solutions.solve2 = (complex(-b, 0) - cmplxJudge) / complex(2*a, 0)
	}
	return solutions
}

func formatQuestion(params quadraticParam) string {
	var output string
	if params.a > 0 {
		if params.a == 1 {
			output += "x²"
		} else {
			output += strconv.FormatInt(params.a, 10) + "x²"
		}
	} else if params.a < 0 {
		if params.a == -1 {
			output += "-x²"
		} else {
			output += strconv.FormatInt(params.a, 10) + "x²"
		}
	}

	if params.a != 0 && params.b > 0 {
		output += " + "
	}

	if params.b > 0 {
		if params.b == 1 {
			output += "x"
		} else {
			output += strconv.FormatInt(params.b, 10) + "x"
		}
	} else if params.b < 0 {
		if params.b == -1 {
			output += " - x"
		} else {
			output += " - " + strconv.FormatInt(-params.b, 10) + "x"
		}
	}

	if params.b != 0 && params.c > 0 {
		output += " + "
	}

	if params.c > 0 {
		output += strconv.FormatInt(params.c, 10)
	} else if params.c < 0 {
		output += " - " + strconv.FormatInt(-params.c, 10)
	}
	return output
}

func formatSolution(solutions solution) string {
	return fmt.Sprintf(`x=%v or x=%v`, solutions.solve1, solutions.solve2)
}

func formatSolutions(solutions solution, params quadraticParam) string {
	return fmt.Sprintf(`%v -> %v`, formatQuestion(params), formatSolution(solutions))
}

func main() {
	//http.HandleFunc("/", homePage)
	http.HandleFunc("/", homePageQuadratic)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}
