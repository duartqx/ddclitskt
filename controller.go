package main

import (
	"fmt"
	"time"
)

type TaskController struct {
	args       *ArgParser
	view       *TaskView
	repository *TaskRepository
}

func NewTaskController(args *ArgParser, view *TaskView, repository *TaskRepository) *TaskController {
	return &TaskController{args: args, view: view, repository: repository}
}

func (tc TaskController) Serve() error {
	switch {
	case *tc.args.Insert:
		return tc.insert()
	case *tc.args.Completed:
		return tc.completed()
	case *tc.args.Find:
		return tc.find()
	case *tc.args.Update:
		return tc.update()
	default:
		return tc.started()
	}
}

func (tc TaskController) insert() (err error) {
	if *tc.args.Tag == "" || *tc.args.Description == "" {
		return fmt.Errorf("Tag or Url can't be empty")
	}

	task := &Task{Tag: *tc.args.Tag, Description: *tc.args.Description}
	err = tc.repository.Create(task)
	if err != nil {
		return err
	}
	return tc.view.Render(&[]*Task{task})
}

func (tc TaskController) completed() (err error) {
	yesterday := time.Now().AddDate(0, 0, -1)
	tasks, err := tc.repository.FindByEndDate(&yesterday)
	if err != nil {
		return err
	}
	return tc.view.Render(tasks)
}

func (tc TaskController) find() (err error) {
	if *tc.args.Tag == "" {
		tasks, err := tc.repository.FindStarted()
		if err != nil {
			return err
		}
		return tc.view.Render(tasks)
	}

	task, err := tc.repository.FindByTag(*tc.args.Tag)
	if err != nil {
		return err
	}
	return tc.view.Render(&[]*Task{task})

}

func (tc TaskController) update() (err error) {

	if *tc.args.Tag == "" {
		return fmt.Errorf("Tag can't be empty to update")
	}
	task, err := tc.repository.FindByTag(*tc.args.Tag)
	if err != nil {
		return err
	}
	if err := tc.repository.UpdateEndAt(task); err != nil {
		return err
	}

	return tc.view.Render(&[]*Task{task})
}

func (tc TaskController) started() (err error) {
	tasks, err := tc.repository.FindStarted()
	if err != nil {
		return err
	}
	return tc.view.Render(tasks)
}
