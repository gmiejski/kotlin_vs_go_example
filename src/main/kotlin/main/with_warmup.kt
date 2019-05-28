package main


fun main() {
    val warmupCount = 10
    run(warmupCount)
    val executionsCount = 50
    printResults(run(executionsCount))
}