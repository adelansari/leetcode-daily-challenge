package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	start := 0
	end := len(rooms) - 1
	dfs(rooms, visited, start, end)
	output := (len(rooms) == len(visited))

	return output
}

func dfs(rooms [][]int, visited map[int]bool, start, end int) {
	if _, ok := visited[start]; ok {
		return
	}
	visited[start] = true
	for _, next := range rooms[start] {
		dfs(rooms, visited, next, end)
	}
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {}, {3}}))
	fmt.Println(canVisitAllRooms([][]int{{2}, {}, {1}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {1, 4}, {2, 3, 4, 1}, {}, {4, 3, 2}}))
}
