package main

type TaskStatus int

const (
	TaskNotStarted TaskStatus = iota
	TaskFinished
	TaskActive
	TaskDisabled
)

type ProjectStage int

const (
	PreCon
)

type Task struct {
	status          TaskStatus
	subTasks        []Task
	prerequisites   []Task
	chat            Chat
	personOfContact Contact
}

// TODO: Placeholder types
type Message int
type Chat int

type Contact struct {
	firstName  string
	lastName   string
	email      string
	phone      string
	company    string
	department string
	team       string
}
