package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/repository/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func TestGetToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.GetTodoRequest{Id: "123"}
	mockToDoRep.On("Get", req.Id).Return(toDo, nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.GetTodo(nil, req)

	expectedRes := &pb.GetTodoResponse{Item: toDo}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestGetToDoNonExistElem(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}

	mockErr := errors.New("Non-exist item")
	req := &pb.GetTodoRequest{Id: "td100"}

	mockToDoRep.On("Get", req.Id).Return(nil, mockErr)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.GetTodo(nil, req)
	expectedErr := grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", mockErr)

	assert.Nil(t, res)
	assert.Equal(t, expectedErr, err)
	mockToDoRep.AssertExpectations(t)
}

func TestListToDo(t *testing.T) {

	mockToDoRep := &mocks.ToDo{}

	var toDoList []*pb.Todo

	toDoList = append(toDoList, &pb.Todo{Title: "ToDo1"})

	toDoList = append(toDoList, &pb.Todo{Title: "ToDo2"})

	req := &pb.ListTodoRequest{Limit: int32(2), NotCompleted: true}

	mockToDoRep.On("List", req.Limit, req.NotCompleted).Return(toDoList, nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.ListTodo(nil, req)

	expectedRes := &pb.ListTodoResponse{Items: toDoList}

	assert.Nil(t, err)

	assert.Equal(t, expectedRes, res)

	mockToDoRep.AssertExpectations(t)

}

func TestListToDoTooLargeLimit(t *testing.T) {

	mockToDoRep := &mocks.ToDo{}

	mockErr := errors.New("Too much items")
	req := &pb.ListTodoRequest{Limit: int32(1000000), NotCompleted: false}

	mockToDoRep.On("List", req.Limit, req.NotCompleted).Return(nil, mockErr)

	service := ToDo{ToDoRepo: mockToDoRep}

	_, err := service.ListTodo(nil, req)
	expectedErr := grpc.Errorf(codes.NotFound, "Could not list items from the database: %s", mockErr)

	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)

	mockToDoRep.AssertExpectations(t)

}

func TestCreateToDo(t *testing.T) {

	mockToDoRep := &mocks.ToDo{}

	todo := &pb.Todo{Title: "Hello"}

	req := &pb.CreateTodoRequest{Item: todo}

	mockToDoRep.On("Insert", req.Item).Return(nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.CreateTodo(nil, req)

	assert.Nil(t, err)
	assert.NotNil(t, res.Id)

	mockToDoRep.AssertExpectations(t)

}

func TestCreateToDoEmpty(t *testing.T) {

	mockToDoRep := &mocks.ToDo{}

	todo := &pb.Todo{}

	mockErr := errors.New("Failed to insert empty title todo")
	req := &pb.CreateTodoRequest{Item: todo}

	mockToDoRep.On("Insert", req.Item).Return(mockErr)

	service := ToDo{ToDoRepo: mockToDoRep}

	_, err := service.CreateTodo(nil, req)

	expectedErr := grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", mockErr)

	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)

	mockToDoRep.AssertExpectations(t)
}

func TestDeleteToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}

	req := &pb.DeleteTodoRequest{Id: "td1"}
	mockToDoRep.On("Delete", req.Id).Return(nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.DeleteTodo(nil, req)

	expectedRes := &pb.DeleteTodoResponse{}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestDeleteToDoNonExistItem(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}

	req := &pb.DeleteTodoRequest{Id: "td100"}
	mockErr := errors.New("non-exist item")
	mockToDoRep.On("Delete", req.Id).Return(mockErr)

	service := ToDo{ToDoRepo: mockToDoRep}
	res, err := service.DeleteTodo(nil, req)

	expectedErr := grpc.Errorf(codes.Internal, "Could not delete item from the database: %s", mockErr)

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.Equal(t, expectedErr, err)
	mockToDoRep.AssertExpectations(t)
}
