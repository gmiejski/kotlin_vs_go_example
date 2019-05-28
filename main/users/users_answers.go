package users

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var subjects = []string{"MATH", "BIOLOGY", "HISTORY", "PHYSICS"}
var ratings = []int{1, 2, 3, 4, 5}
var now = time.Now()
var generatedDatesMaxDiff = 21

func timeInDays(daysDiff int) time.Duration {
	return time.Duration(daysDiff) * time.Hour * 24
}

func getDay(t time.Time) string {
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}

func generate(answers int, usersCount int) []UserAnswer {
	result := make([]UserAnswer, answers)
	for i := 0; i < answers; i++ {
		userID := rand.Intn(usersCount) + 1
		subject := subjects[rand.Intn(len(subjects))]
		rating := ratings[rand.Intn(len(ratings))]
		duration := -timeInDays(rand.Intn(generatedDatesMaxDiff))
		date := now.Add(duration)
		answer := UserAnswer{userID: userID, subject: subject, rating: rating, date: date}
		result[i] = answer
	}
	return result
}

func Test(answers int, users int) {
	data := generate(answers, users)

	lastTwoWeeks := filterLastDays(data, 14)
	activities := generateUsersActivity(lastTwoWeeks)
	activities = filterByActivitiesCount(activities, 100)
	printActivitiesDiagram(activities)
	fmt.Println(fmt.Sprintf("Good answerers: %.2f%%", percentageAnswersOver(activities, 3.0)))
}

func percentageAnswersOver(activities []UserActivity, over float64) float64 {
	trueFor := 0.0
	for _, act := range activities {
		if act.averageRating() > over {
			trueFor++
		}
	}
	return trueFor / float64(len(activities)) * 100
}

func printActivitiesDiagram(activities []UserActivity) {
	dateRanges := generateDates()

	for _, x := range activities {
		fmt.Printf("User: %d\n", x.userID)
		for i, _ := range dateRanges {
			fmt.Printf("%s -> %d\n", dateRanges[i], len(x.days[i].answers))
		}
	}
}

func generateDates() []string {
	var result []string
	for i := 14; i > 1; i-- {
		result = append(result, getDay(now.Add(-timeInDays(i))))
	}
	return result
}

func filterByActivitiesCount(activities []UserActivity, minimumAnswers int) []UserActivity {
	var result []UserActivity
	for _, x := range activities {
		if x.answersCount() > minimumAnswers {
			result = append(result, x)
		}
	}
	return result
}

func generateUsersActivity(answers []UserAnswer) []UserActivity {
	answersPerUser := make(map[int][]UserAnswer)
	for _, answer := range answers {
		answersPerUser[answer.userID] = append(answersPerUser[answer.userID], answer)
	}

	var activities []UserActivity
	for userID, userAnswers := range answersPerUser {
		groupedByDay := answersGroupedByDay(userAnswers)

		activities = append(activities, UserActivity{userID: userID, days: groupedByDay})
	}

	return activities
}

func answersGroupedByDay(answers []UserAnswer) []UserDay {
	var result []UserDay
	byDay := make(map[string][]UserAnswer)
	for _, answer := range answers {
		date := getDay(answer.date)
		byDay[date] = append(byDay[date], answer)
	}

	for day, answersDay := range byDay {
		result = append(result, UserDay{date: day, answers: answersDay})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].date < (result[j].date)
	})
	return result
}

func filterLastDays(answers []UserAnswer, maximumDaysDiff int) []UserAnswer {
	result := make([]UserAnswer, 0)
	from := now.Add(-timeInDays(maximumDaysDiff))

	for _, answer := range answers {
		if answer.date.After(from) {
			result = append(result, answer)
		}
	}
	return result
}
