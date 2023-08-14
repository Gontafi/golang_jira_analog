package models

type Workflows struct {
	Id          int
	Name        string
	Description string
}

type WorkflowSteps struct {
	Id          int
	WorkflowsId int
	Name        string
	Description string
}

type WorkflowTransition struct {
	Id          int
	WorkflowsId int
	FromStepId  int
	ToStepId    int
	Name        string
	Description string
}
