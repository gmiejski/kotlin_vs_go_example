package main

import java.time.LocalDate
import java.time.ZonedDateTime

data class UserDay(val date: LocalDate, val answers: List<Answer>)

data class Answer(val subject: String, val rating: Int, val date: ZonedDateTime)
data class UserAnswer(val user: Int, val answer: Answer)


data class UserActivity(val user: Int, val days: List<UserDay>) {
    fun answersCount(): Int {
        return days.map { it.answers.size }.sum()
    }

    fun avgRating(): Double {
        return days.flatMap { it.answers }
            .map { it.rating }
            .average()
    }
}
