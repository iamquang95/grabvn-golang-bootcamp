package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/repository/mocks"
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

	req := &pb.ListTodoRequest{Limit: int32(1000000), NotCompleted: false}

	mockToDoRep.On("List", req.Limit, req.NotCompleted).Return(nil, errors.New("Too much items"))

	service := ToDo{ToDoRepo: mockToDoRep}

	_, err := service.ListTodo(nil, req)

	assert.NotNil(t, err)

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

	req := &pb.CreateTodoRequest{Item: todo}

	mockToDoRep.On("Insert", req.Item).Return(errors.New("Failed to insert empty title todo"))

	service := ToDo{ToDoRepo: mockToDoRep}

	_, err := service.CreateTodo(nil, req)

	assert.NotNil(t, err)

	mockToDoRep.AssertExpectations(t)

}
