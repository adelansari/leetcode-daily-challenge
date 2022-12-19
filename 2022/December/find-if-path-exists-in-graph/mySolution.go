package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {
	adjList := make([][]int, n)
	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		adjList[node1] = append(adjList[node1], node2)
		adjList[node2] = append(adjList[node2], node1)
	}

	visited := make([]bool, n)
	for i := range visited {
		visited[i] = false
	}

	dfs(source, adjList, visited)
	output := visited[destination]

	return output
}

func dfs(node int, adjList [][]int, visited []bool) {
	if visited[node] {
		return
	}
	visited[node] = true

	for _, nextNode := range adjList[node] {
		dfs(nextNode, adjList, visited)
	}
}

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
}
