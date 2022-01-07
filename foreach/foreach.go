package foreach

import (
	"fmt"
	"log"
	"strings"
)

type Graph struct {
	Nodes []Node
	Edges []Edge
}

type Edge struct {
	From string // from node
	To   string // to node
}

type Node struct {
	Id string
}

func ForeachGraph(graphNodes []Node, graphEdges []Edge) ([]string, error) {

	res := make([]string, 0) // 最终的执行路径

	nodes := graphNodes
	edges := graphEdges

	for {
		if len(nodes) < 1 {
			break
		}
		var rootIds []string // 存储入度为 0 的节点 id
		query := make(map[string]int)
		for _, item := range nodes {
			query[item.Id] = 0
		}

		for _, item := range edges {
			v, ok := query[item.To]
			if !ok {
				continue
			}
			query[item.To] = v + 1
		}

		for k, v := range query {
			if v == 0 {
				rootIds = append(rootIds, k)
			}
		}

		if len(rootIds) < 1 {
			log.Printf("存在环，请检查输入")
			return nil, fmt.Errorf("存在环，请检查输入")
		}

		// A,B 单节点则 A
		res = append(res, strings.Join(rootIds, ","))

		// 移除节点
		var tmpNodes []Node
		for _, item := range nodes {
			if !contains(rootIds, item.Id) {
				tmpNodes = append(tmpNodes, item)
			}
		}

		// 移除关系
		var tmpEdges []Edge
		for _, item := range edges {
			if !contains(rootIds, item.From) {
				tmpEdges = append(tmpEdges, item)
			}
		}

		nodes = tmpNodes
		edges = tmpEdges
	}

	return res, nil
}

func contains(source []string, target string) bool {
	for _, a := range source {
		if a == target {
			return true
		}
	}
	return false
}
