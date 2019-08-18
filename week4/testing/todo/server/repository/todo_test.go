// +build integration persistence

package repository

import (
	"testing"
	"time"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/suite"
	"github.com/xuanit/testing/todo/pb"
)

type ToDoRepositorySuite struct {
	db *pg.DB
	suite.Suite
	todoRep ToDoImpl
}

func (s *ToDoRepositorySuite) SetupSuite() {
	// Connect to PostgresQL
	s.db = pg.Connect(&pg.Options{
		User:                  "postgres",
		Password:              "example",
		Database:              "todo",
		Addr:                  "localhost" + ":" + "5433",
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
	})

	// Create Table
	s.db.CreateTable(&pb.Todo{}, nil)

	s.todoRep = ToDoImpl{DB: s.db}
}

func (s *ToDoRepositorySuite) TearDownSuite() {
	s.db.DropTable(&pb.Todo{}, nil)
	s.db.Close()
}

func (s *ToDoRepositorySuite) TestInsert() {
	item := &pb.Todo{Id: "new_item", Title: "meeting"}
	err := s.todoRep.Insert(item)

	s.Nil(err)

	newTodo, err := s.todoRep.Get(item.Id)
	s.Nil(err)
	s.Equal(item, newTodo)

	err = s.todoRep.Delete(item.Id)
	s.Nil(err)
}

func (s *ToDoRepositorySuite) TestDeleteNonExistItem() {
	err := s.todoRep.Delete("new")
	s.NotNil(err)
}

func (s *ToDoRepositorySuite) TestList() {
	item1 := &pb.Todo{Id: "td1", Title: "Todo 1", Completed: true}
	err1 := s.todoRep.Insert(item1)
	s.Nil(err1)

	item2 := &pb.Todo{Id: "td2", Title: "Todo 2", Completed: false}
	err2 := s.todoRep.Insert(item2)
	s.Nil(err2)

	res, err := s.todoRep.List(2, false)
	expectedRes := []*pb.Todo{item1, item2}

	s.Nil(err)
	s.Equal(expectedRes, res)

	res, err = s.todoRep.List(1, false)
	expectedRes = []*pb.Todo{item1}

	s.Nil(err)
	s.Equal(expectedRes, res)
}

func (s *ToDoRepositorySuite) TestGetNonExistingElem() {
	result, err := s.todoRep.Get("non-exist")
	s.NotNil(err)
	s.Nil(nil, result)
}

func TestToDoRepository(t *testing.T) {
	suite.Run(t, new(ToDoRepositorySuite))
}
