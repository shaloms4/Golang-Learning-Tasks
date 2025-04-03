package usecases

import (
	"context"
	"time"

	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
)

type taskUsecase struct {
	taskRepo       Domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(repo Domain.TaskRepository, timeout time.Duration) Domain.TaskUsecase {
	return &taskUsecase{
		taskRepo:       repo,
		contextTimeout: timeout,
	}
}

func (u *taskUsecase) Create(ctx context.Context, task *Domain.Task) error {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.taskRepo.Create(c, task)
}

func (u *taskUsecase) FetchAll(ctx context.Context) ([]Domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.taskRepo.FetchAll(c)
}

func (u *taskUsecase) FetchByID(ctx context.Context, id string) (*Domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.taskRepo.FetchByID(c, id)
}

func (u *taskUsecase) Update(ctx context.Context, id string, task *Domain.Task) error {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.taskRepo.Update(c, id, task)
}

func (u *taskUsecase) Delete(ctx context.Context, id string) error {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.taskRepo.Delete(c, id)
}