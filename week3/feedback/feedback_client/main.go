package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "grab/week3/feedback/feedback"
)

const (
	address = "localhost:50051"
)

func addPassengerFeedback(ctx context.Context, c pb.CustomerFeedbackClient, in *pb.PassengerFeedback) {
	resp, err := c.AddPassengerFeedback(ctx, &pb.AddPassengerFeedbackRequest{Feedback: in})
	if err != nil {
		log.Println("failed to add passenger feedback", err)
		return
	}
	log.Println("success to add passenger feedback with msg:", resp.Msg)
}

func getFeedbacksByPassengerID(ctx context.Context, c pb.CustomerFeedbackClient, psgID int32) {
	resp, err := c.GetFeedbacksByPassengerID(ctx, &pb.GetFeedbacksByPassengerIdRequest{PassengerID: psgID})
	if err != nil {
		log.Println("failed to get feedback by passengerID:", psgID, "with error:", err)
		return
	}
	log.Println("feedback by passengerID", psgID, ":")
	for _, feedback := range resp.Feedbacks {
		log.Println(feedback)
	}
}

func getFeedbackByBookingCode(ctx context.Context, c pb.CustomerFeedbackClient, bookingCode string) {
	resp, err := c.GetFeedbackByBookingCode(ctx, &pb.GetFeedbackByBookingCodeRequest{BookingCode: bookingCode})
	if err != nil {
		log.Println("failed to get feedback by bookingCode", bookingCode, "with error:", err)
		return
	}
	log.Println("feedback by bookingCode", bookingCode, ":")
	log.Println(resp.Feedback)
}

func deleteFeedbacksByPassengerID(ctx context.Context, c pb.CustomerFeedbackClient, psgID int32) {
	resp, err := c.DeleteFeedbacksByPassengerID(ctx, &pb.DeleteFeedbacksByPassengerIdRequest{PassengerID: psgID})
	if err != nil {
		log.Println("failed to delete feedbacks by passengerID", psgID, "with error:", err)
		return
	}
	log.Println("deleted", resp.DeletedFeedbacks, "feedbacks of passengerID", psgID)
}

func simulateData(ctx context.Context, c pb.CustomerFeedbackClient) {
	feedback1 := pb.PassengerFeedback{
		BookingCode: "abc123",
		PassengerID: 1,
		Feedback:    "good",
	}
	feedback2 := pb.PassengerFeedback{
		BookingCode: "abc124",
		PassengerID: 1,
		Feedback:    "ok",
	}
	feedback3 := pb.PassengerFeedback{
		BookingCode: "abc125",
		PassengerID: 2,
		Feedback:    "bad",
	}
	addPassengerFeedback(ctx, c, &feedback1)
	addPassengerFeedback(ctx, c, &feedback1)
	addPassengerFeedback(ctx, c, &feedback3)
	getFeedbackByBookingCode(ctx, c, feedback1.BookingCode)
	getFeedbackByBookingCode(ctx, c, feedback2.BookingCode)
	getFeedbacksByPassengerID(ctx, c, feedback1.PassengerID)
	addPassengerFeedback(ctx, c, &feedback2)
	getFeedbackByBookingCode(ctx, c, feedback2.BookingCode)
	getFeedbacksByPassengerID(ctx, c, feedback1.PassengerID)
	deleteFeedbacksByPassengerID(ctx, c, feedback3.PassengerID)
	getFeedbacksByPassengerID(ctx, c, feedback3.PassengerID)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect", err)
	}
	defer conn.Close()
	c := pb.NewCustomerFeedbackClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	simulateData(ctx, c)
}
