package code_snippets

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt)

	sc.Scan()
	t := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(t[0])
	m, _ := strconv.Atoi(t[1])

	edges := make([][]int, n+1)
	for i := 0; i < m; i++ {
		sc.Scan()
		t := strings.Split(sc.Text(), " ")

		a, _ := strconv.Atoi(t[0])
		b, _ := strconv.Atoi(t[1])
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}

	visited := make([]bool, n+1)
	visited[0] = true

	var dfs func(int)
	dfs = func(v int) {
		visited[v] = true

		for _, val := range edges[v] {
			if !visited[val] {
				dfs(val)
			}
		}
	}

	dfs(1)
}
