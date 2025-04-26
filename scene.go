package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SceneNode interface {
	Update()
	Draw()
	AddChild(node SceneNode)
}

type BaseNode struct {
	children []SceneNode
}

func NewBaseNode() *BaseNode {
	return &BaseNode{
		children: make([]SceneNode, 0),
	}
}

func (root *BaseNode) AddChild(node SceneNode) {
	root.children = append(root.children, node)
}

func (root *BaseNode) Update() {
	for _, child := range root.children {
		child.Update()
	}
}

func (root *BaseNode) Draw() {

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, child := range root.children {
		child.Draw()
	}

	rl.EndDrawing()
}
