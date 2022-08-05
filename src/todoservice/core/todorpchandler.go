package core

import (
	"github.com/todolist-project-rhttraining/src/todoservice/pkg/pb"
	"golang.org/x/net/context"
)

type Handler struct {
	svc TodoService
}

func NewHandler(todoSvc TodoService) Handler {
	return Handler{
		svc: todoSvc,
	}
}

func (h Handler) CreateTodo(ctx context.Context, request *pb.CreateTodoRequest, response *pb.ObjectIdResponse) error {
	id, err := h.svc.CreateTodo(ctx, request.UserId, request.Todo, request.Done)
	if err != nil {
		return err
	}

	response.Id = id
	return nil
}

func (h Handler) GetTodo(ctx context.Context, request *pb.GetTodoRequest, response *pb.GetTodoResponse) error {
	listColl, err := h.svc.GetTodo(ctx, request.Id)
	if err != nil {
		return err
	}

	var todoResp []*pb.Todo
	for _, v := range listColl {
		todoResp = append(todoResp, &pb.Todo{
			Id:   v.ID.String(),
			Todo: v.Todo,
			Done: v.Done,
		})
	}
	response.Todo = todoResp
	return nil
}
