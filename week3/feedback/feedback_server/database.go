package main

import (
	"errors"
	pb "grab/week3/feedback/feedback"
)

type Feedbacks struct {
	ExistingBookingCode map[string]bool
	ExistingPassengerID map[int32]bool
	Feedbacks           []pb.PassengerFeedback
}

// CreateNewFeedbacks create new feedback database
func CreateNewFeedbacks() Feedbacks {
	return Feedbacks{
		ExistingBookingCode: make(map[string]bool),
		ExistingPassengerID: make(map[int32]bool),
		Feedbacks:           make([]pb.PassengerFeedback, 0),
	}
}

// AddFeedback add new feedback to database
func (fb *Feedbacks) AddFeedback(feedback *pb.PassengerFeedback) error {
	if fb.ExistingBookingCode[feedback.BookingCode] == true {
		return errors.New("Duplicated booking code")
	}
	fb.ExistingBookingCode[feedback.BookingCode] = true
	fb.ExistingPassengerID[feedback.PassengerID] = true
	fb.Feedbacks = append(fb.Feedbacks, *feedback)
	return nil
}

// DeleteFeedbacksByPassengerID delete all feedbacks by passengerID
func (fb *Feedbacks) DeleteFeedbacksByPassengerID(psgID int32) (int32, error) {
	if fb.ExistingPassengerID[psgID] == false {
		return 0, errors.New("Non-existing passengerId")
	}
	var newFeedbacks []pb.PassengerFeedback
	var cntDeleted int32
	for _, feedback := range fb.Feedbacks {
		if feedback.PassengerID == psgID {
			fb.ExistingBookingCode[feedback.BookingCode] = false
			cntDeleted++
		} else {
			newFeedbacks = append(newFeedbacks, feedback)
		}
	}
	fb.Feedbacks = newFeedbacks
	return cntDeleted, nil
}

// GetFeedbackByBookingCode get feedback by booking code
func (fb *Feedbacks) GetFeedbackByBookingCode(bookingCode string) (pb.PassengerFeedback, error) {
	if fb.ExistingBookingCode[bookingCode] == false {
		return pb.PassengerFeedback{}, errors.New("Non-existing booking code")
	}
	for _, feedback := range fb.Feedbacks {
		if feedback.BookingCode == bookingCode {
			return feedback, nil
		}
	}
	return pb.PassengerFeedback{}, nil
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
