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

func (s *server) AddCustomerFeedback(ctx context.Context, in *pb.AddCustomerFeedbackRequest) (*pb.AddCustomerFeedbackResponse, error) {
	log.Println("add new customer feedback:", in.Feedback)
	return &pb.AddCustomerFeedbackResponse{Msg: "Success"}, nil
}

func (s *server) GetFeedbacksByPassengerID(ctx context.Context, in *pb.GetFeedbacksByPassengerIdRequest) (*pb.GetFeedbacksByPassengerIdResponse, error) {
	log.Println("get feedback by passengerId =", in.PassengerID)
	return &pb.GetFeedbacksByPassengerIdResponse{Feedbacks: make([]*pb.PassengerFeedback, 0)}, nil
}

func (s *server) GetFeedbackByBookingCode(ctx context.Context, in *pb.GetFeedbackByBookingCodeRequest) (*pb.GetFeedbackByBookingCodeResponse, error) {
	log.Println("get feedback by bookingCode =", in.BookingCode)
	return &pb.GetFeedbackByBookingCodeResponse{Feedback: &pb.PassengerFeedback{}}, nil
}

func (s *server) DeleteFeedbacksByPassengerID(ctx context.Context, in *pb.DeleteFeedbacksByPassengerIdRequest) (*pb.DeleteFeedbacksByPassengerIdResponse, error) {
	log.Println("delete feedback by passengerId =", in.PassengerID)
	return &pb.DeleteFeedbacksByPassengerIdResponse{DeletedFeedbacks: 0}, nil
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
