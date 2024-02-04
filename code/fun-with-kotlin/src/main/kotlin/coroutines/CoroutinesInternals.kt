package coroutines

import kotlin.coroutines.resume
import kotlin.coroutines.suspendCoroutine

/**
 * Coroutines used Continuation Passing Style (CPS)
 */
class CoroutinesInternals {
    suspend fun printNumber(){
        suspendCoroutine {
            println("Started Suspend coroutines")
            yetAnotherAsync()
            it.resume(Unit)
        }
    }

    private fun yetAnotherAsync() {
        for (i in 1..100) println(i)
    }

}

suspend fun main() {
    println("Start")
    val internals = CoroutinesInternals()
    internals.printNumber()
    println("End")
}
