package main

import "fmt"

/*
Есть дерево категорий товаров, размещаемых на X. Нужно найти заданную подкатегорию и вывести путь до неё
от корневой категории, если она присутствует в дереве
### in
class Category {
  name: string;
  children: Array;
}
root = {
  name: "root",
  children: [
    {
      name: "Бытовая техника",
      children: [
        {
          name: "Телевизоры",
          children: [
            { name: "ЭЛТ", children: [] },
            { name: "LED", children: [] },
            { name: "OLED", children: [] }
          ]
        },
        {
          name: "Холодильники",
          children: [
            { name: "Двухкамерные", children: [] },
            { name: "Однокамерные", children: [] }
          ]
        },
        {
          name: "Утюги",
          children: []
        }
      ]
    },
    {
      name: "Растения",
      children: [
        { name: "Комнатные", children: [] },
        { name: "Садовые", children: [] }
      ]
    }
  ]
}
search_name = "OLED"
### out
Бытовая техника > Телевизоры > OLED
*/

type Node struct {
	Name     string
	Children []Node
}

var root = Node{
	Name: "root",
	Children: []Node{
		{
			Name: "Бытовая техника",
			Children: []Node{
				{
					Name: "Телевизоры",
					Children: []Node{
						{Name: "ЭЛТ", Children: []Node{}},
						{Name: "LED", Children: []Node{}},
						{Name: "OLED", Children: []Node{}},
					},
				},
				{
					Name: "Холодильники",
					Children: []Node{
						{Name: "Двухкамерные", Children: []Node{}},
						{Name: "Однокамерные", Children: []Node{}},
					},
				},
				{
					Name:     "Растения",
					Children: []Node{},
				},
			},
		},
		{
			Name: "Растения",
			Children: []Node{
				{Name: "Комнатные", Children: []Node{}},
				{Name: "Садовые", Children: []Node{}},
			},
		},
	},
}

func recur(root Node, name string) string {
	if root.Name == name {
		return name
	}

	for _, child := range root.Children {
		deepName := recur(child, name)
		if deepName == "" {
			continue
		}

		return fmt.Sprintf("%s > %s", root.Name, deepName)
	}

	return ""
}

func catStack(root *Node, name string) string {
	stack := []*Node{root}
	parent := make(map[*Node]*Node)

	var found *Node

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if n.Name == name {
			found = n
			break
		}

		for i := range n.Children {
			child := &n.Children[i]
			parent[&n.Children[i]] = n
			stack = append(stack, child)
		}
	}

	if found == nil {
		return ""
	}

	res := found.Name
	currParent := parent[found]

	for currParent != nil {
		if currParent.Name != "root" {
			res = currParent.Name + " > " + res
		}
		currParent = parent[currParent]
	}
	return res
}

func main() {
	fmt.Println(catStack(&root, "Двухкамерные"))
	fmt.Println(recur(root, "Растения"))
}
