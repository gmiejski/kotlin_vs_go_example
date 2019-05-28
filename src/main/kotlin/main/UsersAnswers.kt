package main

import java.time.Duration
import java.time.ZonedDateTime
import java.time.temporal.ChronoUnit
import java.util.*

val subjects = listOf("MATH", "BIOLOGY", "HISTORY", "PHYSICS")
val ratings = (1..5)

val end: ZonedDateTime = ZonedDateTime.now()
val maxTime: Duration = Duration.ofDays(21)
const val MINIMUM_ANSWERS = 100

fun generate(answers: Int, usersCount: Int): List<UserAnswer> {
    val random = Random()

    val seconds = maxTime.seconds
    if (seconds > Int.MAX_VALUE) throw RuntimeException("overvalue")

    return (1..answers).map {
        UserAnswer(
            random.nextInt(usersCount),
            Answer(
                subjects.random(),
                ratings.random(),
                end.minus((1..14).random().toLong(), ChronoUnit.DAYS)
            )
        )
    }
}

fun toUserActivity(user: Int, answers: List<Answer>): UserActivity {
    val days = answers.groupBy { it.date.toLocalDate() }
        .map { UserDay(it.key, it.value) }
        .sortedBy { it.date }

    return UserActivity(user, days)
}

fun test() {
    val generateLastTwoWeeks = (14 downTo 1).map { end.minusDays(it.toLong()) }

    val data = generate(10000, 50)

    val usersActivities = data.filter { it.answer.date.isAfter(end.minusDays(14)) }
        .groupBy { it.user }
        .mapValues { it.value.map { usersAnswer -> usersAnswer.answer } }
        .map { toUserActivity(it.key, it.value) }
        .filter { it.answersCount() > MINIMUM_ANSWERS }

    usersActivities.forEach {
        println("User: ${it.user}")
        generateLastTwoWeeks.zip(it.days).forEach { day ->
            println("${day.first.toLocalDate()} -> ${day.second.answers.size}")
        }
    }

    val goodAnswersPercentage = usersActivities.percentageMatching { it.avgRating() > 3.0 }
    println("Good answerers: ${"%.2f".format(goodAnswersPercentage)}%")
}

fun List<UserActivity>.percentageMatching(precondition: (UserActivity) -> Boolean): Double {
    return this.count(precondition) / this.count().toDouble() * 100
}