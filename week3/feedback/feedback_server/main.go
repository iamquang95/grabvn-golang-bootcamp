package main

import (
	"context"
	pb "grab/week3/feedback/feedback"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

var feedbackDB = CreateNewFeedbacks()

func (s *server) AddPassengerFeedback(ctx context.Context, in *pb.AddPassengerFeedbackRequest) (*pb.AddPassengerFeedbackResponse, error) {
	log.Println("add new customer feedback:", in.Feedback)
	err := feedbackDB.AddFeedback(in.Feedback)
	if err != nil {
		return &pb.AddPassengerFeedbackResponse{Msg: "Error"}, err
	}
	return &pb.AddPassengerFeedbackResponse{Msg: "Success"}, nil
}

func (s *server) GetFeedbacksByPassengerID(ctx context.Context, in *pb.GetFeedbacksByPassengerIdRequest) (*pb.GetFeedbacksByPassengerIdResponse, error) {
	log.Println("get feedback by passengerId =", in.PassengerID)
	feedbacks, err := feedbackDB.GetFeedbackByPassengerID(in.PassengerID)
	return &pb.GetFeedbacksByPassengerIdResponse{Feedbacks: feedbacks}, err
}

func (s *server) GetFeedbackByBookingCode(ctx context.Context, in *pb.GetFeedbackByBookingCodeRequest) (*pb.GetFeedbackByBookingCodeResponse, error) {
	log.Println("get feedback by bookingCode =", in.BookingCode)
	feedback, err := feedbackDB.GetFeedbackByBookingCode(in.BookingCode)
	return &pb.GetFeedbackByBookingCodeResponse{Feedback: &feedback}, err
}

func (s *server) DeleteFeedbacksByPassengerID(ctx context.Context, in *pb.DeleteFeedbacksByPassengerIdRequest) (*pb.DeleteFeedbacksByPassengerIdResponse, error) {
	log.Println("delete feedback by passengerId =", in.PassengerID)
	deletedFeedbacks, err := feedbackDB.DeleteFeedbacksByPassengerID(in.PassengerID)
	return &pb.DeleteFeedbacksByPassengerIdResponse{DeletedFeedbacks: deletedFeedbacks}, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomerFeedbackServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
