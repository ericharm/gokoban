package main

import "sync"

type Application struct {
	BaseNode SceneNode
}

var (
	applicationInstance *Application
	applicationOnce     sync.Once
)

func GetApplication() *Application {
	applicationOnce.Do(func() {
		applicationInstance = &Application{
			BaseNode: NewBaseNode(),
		}
	})
	return applicationInstance
}
