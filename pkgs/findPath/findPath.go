package findPath

import (
	"strconv"
	"strings"
	"wikipediaScraper/pkgs/linklist"
	"wikipediaScraper/pkgs/utils"
)

var data = utils.GetData()
var graph map[string] []string
var visited map[string] bool
type vertex struct {
	label string
	path string
}
func bfs(init string, final string) string {
	var Queue = linklist.MakeLinklist()
	f := graph[init]
	for _, v := range f {
		Queue.AddBack(vertex{label: v, path: init + "-"+ v})
		visited[v] = true
	}
	for !Queue.IsEmpty() {
		m := Queue.Data.(vertex)
		f := graph[m.label]
		if m.label == final {
			return m.path
		}
		for _, v := range f {
			if !visited[v] {
				Queue.AddBack(vertex{label: v, path: m.path +"-"+ v })
				visited[v] = true
			}
		}
		Queue.PopTop()
	}
	return "-1"
}
func Find(i string, f string) []string{
	visited = make(map[string] bool)
	utils.LoadMap(&graph)
	init := strconv.Itoa(utils.IndexOf(i,data))
	final := strconv.Itoa(utils.IndexOf(f,data))
	path := bfs(init, final)
	ev := strings.Split(path,"-")
	st := make([]string, len(ev))
	for ix, v := range ev {
		i, _ := strconv.Atoi(v)
		st[ix] = data[i]
	}
	return st
}
