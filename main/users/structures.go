package users

import "time"

type UserAnswer struct {
	userID  int
	subject string
	rating  int
	date    time.Time
}

type UserDay struct {
	date    string
	answers []UserAnswer
}

type UserActivity struct {
	userID int
	days   []UserDay
}

func (activity UserActivity) answersCount() int {
	sum := 0
	for _, x := range activity.days {
		sum += len(x.answers)
	}
	return sum
}

func (activity UserActivity) averageRating() float64 {
	sum := 0.0
	count := 0.0
	for _, day := range activity.days {
		for _, answer := range day.answers {
			sum += float64(answer.rating)
			count += 1
		}
	}
	f := sum / count
	return f
}
