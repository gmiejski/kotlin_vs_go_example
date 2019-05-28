package main

import java.time.Duration
import kotlin.system.measureTimeMillis

fun run(times: Int): List<Long> {
    return (1..times).map { measureTimeMillis { test() } }
}

fun printResults(results: List<Long>) {
    println("**********************************\n" +
            "Results:")
    results.forEach {
        val t = Duration.ofMillis(it)

        println("Total took: ${t.toMillis()}ms")
    }
}


fun main() {
    val executionsCount = 10
    printResults(run(executionsCount))
}