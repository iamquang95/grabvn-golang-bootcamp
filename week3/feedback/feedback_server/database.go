package main

import (
	"errors"
	pb "grab/week3/feedback/feedback"
)

type Feedbacks struct {
	Feedbacks           map[string]pb.PassengerFeedback
	ExistingPassengerID map[int32]bool
}

// CreateNewFeedbacks create new feedback database
func CreateNewFeedbacks() Feedbacks {
	return Feedbacks{
		ExistingPassengerID: make(map[int32]bool),
		Feedbacks:           make(map[string]pb.PassengerFeedback),
	}
}

// AddFeedback add new feedback to database
func (fb *Feedbacks) AddFeedback(feedback *pb.PassengerFeedback) error {
	if _, ok := fb.Feedbacks[feedback.BookingCode]; ok {
		return errors.New("Duplicated booking code")
	}
	fb.Feedbacks[feedback.BookingCode] = *feedback
	fb.ExistingPassengerID[feedback.PassengerID] = true
	return nil
}

// DeleteFeedbacksByPassengerID delete all feedbacks by passengerID
func (fb *Feedbacks) DeleteFeedbacksByPassengerID(psgID int32) (int32, error) {
	if fb.ExistingPassengerID[psgID] == false {
		return 0, errors.New("Non-existing passengerId")
	}
	var cntDeleted int32
	for _, feedback := range fb.Feedbacks {
		if feedback.PassengerID == psgID {
			delete(fb.Feedbacks, feedback.BookingCode)
			cntDeleted++
		}
	}
	return cntDeleted, nil
}

// GetFeedbackByBookingCode get feedback by booking code
func (fb *Feedbacks) GetFeedbackByBookingCode(bookingCode string) (pb.PassengerFeedback, error) {
	feedback, ok := fb.Feedbacks[bookingCode]
	if !ok {
		return pb.PassengerFeedback{}, errors.New("Non-existing booking code")
	}
	return feedback, nil
}

// GetFeedbackByPassengerID get feedbacks by passengerID
func (fb *Feedbacks) GetFeedbackByPassengerID(psgID int32) (feedbacks []*pb.PassengerFeedback, err error) {
	if fb.ExistingPassengerID[psgID] == false {
		err = errors.New("Non-exisiting passenger")
		return
	}
	for _, feedback := range fb.Feedbacks {
		if feedback.PassengerID == psgID {
			feedbacks = append(feedbacks, &feedback)
		}
	}
	return feedbacks, nil
}
