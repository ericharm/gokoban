package main

import "sync"
import rl "github.com/gen2brain/raylib-go/raylib"

type InputEvent int32

const (
	LeftPressed InputEvent = iota
	RightPressed

	EnterPressed
	SpacePressed

	LeftHeld
	RightHeld
)

type InputEventQueue struct {
	events []InputEvent
}

var (
	inputEventQueueInstance *InputEventQueue
	inputEventQueueOnce     sync.Once
)

func GetInputEventQueue() *InputEventQueue {
	inputEventQueueOnce.Do(func() {
		inputEventQueueInstance = &InputEventQueue{
			events: make([]InputEvent, 0),
		}
	})
	return inputEventQueueInstance
}

func enqueueInputEvent(event InputEvent) {
	queue := GetInputEventQueue()
	queue.events = append(queue.events, event)
}

func DequeueInputEvent() InputEvent {
	queue := GetInputEventQueue()
	if len(queue.events) == 0 {
		return -1
	}
	event := queue.events[0]
	queue.events = queue.events[1:]
	return event
}

func ProcessInputEvents() {
	if rl.IsKeyPressed(rl.KeyLeft) {
		enqueueInputEvent(LeftPressed)
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		enqueueInputEvent(RightPressed)
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		enqueueInputEvent(EnterPressed)
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		enqueueInputEvent(SpacePressed)
	}
}

func ProcessRealTimeInput() {
	if rl.IsKeyDown(rl.KeyLeft) {
		enqueueInputEvent(LeftHeld)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		enqueueInputEvent(RightHeld)
	}
}
