package main

import "sync"
import rl "github.com/gen2brain/raylib-go/raylib"

type SceneNode interface {
	Update()
	Draw()
	AddChild(node SceneNode)
}

type RootNode struct {
	children []SceneNode
}

func NewRootNode() *RootNode {
	return &RootNode{
		children: make([]SceneNode, 0),
	}
}

func (root *RootNode) AddChild(node SceneNode) {
	root.children = append(root.children, node)
}

func (root *RootNode) Update() {
	for _, child := range root.children {
		child.Update()
	}
}

func (root *RootNode) Draw() {

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, child := range root.children {
		child.Draw()
	}

	rl.EndDrawing()
}

type Application struct {
	RootNode SceneNode
}

var (
	applicationInstance *Application
	applicationOnce     sync.Once
)

func GetApplication() *Application {
	applicationOnce.Do(func() {
		applicationInstance = &Application{
			RootNode: NewRootNode(),
		}
	})
	return applicationInstance
}
