package thundering_herd

import kotlinx.coroutines.*
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import java.lang.Thread.sleep
import java.util.concurrent.ConcurrentHashMap

/**
 * Solve 1: Using thread safe operations
 *
 * As per [[ThunderingHerdProblem.kt]] we found there are cache misses due to concurrent operations
 * Lets take example to how to solve them, there the computeIfAbsent is thread safe and locks the first fetch operations and post the cache updates
 * all the next subsequent request are fetched from cache to get rid of cache miss
 *
 * Output:
 *
 * fetching from DB for key=IN // Fetch from DB once and all the next concurrent operations return  cache result
 * India
 * result key=IN, result=India
 * result key=IN, result=India
 * result key=IN, result=India
 * result key=IN, result=India
 * result key=IN, result=India
 *
 * Solve 2: using Mutex
 * - Mutex (mutual exclusion) is a sync primitive used to control access over the shared  resource. It allow only one thread to execeute the piece of code
 * In Coroutine world it `suspend` the thread and don't block the thread completely making it efficient use of shared resources
 * It works on 3 steps
 * - Lock acquisition - acquire the lock (in Java world it is used using ReentrantLock)
 * - Execute the code
 * - Release the lock
 *
 * output:
 * solve 2: mutex solve
 * fetching from DB for key=IN
 * result key=IN, result=India
 * result key=IN, result=India
 * result key=IN, result=India
 * result key=IN, result=India
 * result key=IN, result=India
 *
 *
 *
 * Issue and solutions
 *
 * If ```fetchFromDb``` takes too long, the mutex will block all waiting threads or coroutines for that key.
 * Optimizations like double-checked locking, asynchronous fetching, per-key locks, and timeouts can mitigate delays and improve responsiveness.
 * The choice of approach depends on the trade-offs you're willing to make between complexity, performance, and consistency.
 */

private val cacheData = ConcurrentHashMap<String, String>()

@DelicateCoroutinesApi
fun main() {

    threadSafeSolve()
    clearCache()
    mutexSolve()

}

val mutex = Mutex()
fun mutexSolve() {
    println("solve 2: mutex solve")
    runBlocking {
        val keyToBeSearched = "IN"
        val jobCount = 5
        val jobsList = List(jobCount) {
            GlobalScope.launch() {
                val result = fetchValueUsingKeyUsingMutex(keyToBeSearched)
                println("result key=$keyToBeSearched, result=$result")
            }
        }

        //join all the jobs
        for (job in jobsList) {
            job.join()
        }
    }

}

suspend fun fetchValueUsingKeyUsingMutex(keyToBeSearched: String): String {
    return cacheData[keyToBeSearched] ?: mutex.withLock {
        cacheData.computeIfAbsent(keyToBeSearched) {
            fetchFromDb(keyToBeSearched)
        }
    }
}

//Double-check the cache before acquiring the lock to avoid unnecessary waiting.
suspend fun fetchValueUsingKeyUsingDoubleLocking(keyToBeSearched: String): String {
    cacheData[keyToBeSearched]?.let { return it } // early exit
    mutex.lock() // lock
    try {
        return cacheData[keyToBeSearched] ?: fetchFromDb(keyToBeSearched).also { cacheData[keyToBeSearched] = it }
    } finally {
        mutex.unlock() //unlock
    }
}

//Use coroutines to allow other threads or coroutines to proceed with unrelated keys while waiting for the database fetch.
suspend fun fetchValueUsingKeyAsync(keyToBeSearched: String): String {
        return cacheData[keyToBeSearched] ?: run {
            mutex.withLock {
                return cacheData[keyToBeSearched] ?:  fetchFromDb(keyToBeSearched).also { cacheData[keyToBeSearched] = it }
            }
        }
}

val locks = ConcurrentHashMap<String, Mutex>()
//per key locking
suspend fun fetchValueUsingPerKeyLocking(keyToBeSearched: String): String {
    val locks = locks.computeIfAbsent(keyToBeSearched){
        Mutex() // lock per key
    }
    return locks.withLock {
        cacheData[keyToBeSearched]?: fetchFromDb(keyToBeSearched).also { cacheData[keyToBeSearched] = it }
    }
}






fun clearCache() {
    cacheData.clear()
}

@DelicateCoroutinesApi
 fun threadSafeSolve() {
    runBlocking {
        val keyToBeSearched = "IN"
        val jobCount = 5
        val jobsList = List(jobCount) {
            GlobalScope.launch() {
                val result = fetchValueUsingKey(keyToBeSearched)
                println("result key=$keyToBeSearched, result=$result")
            }
        }

        //join all the jobs
        for (job in jobsList) {
            job.join()
        }
    }
}



private fun fetchValueUsingKey(key: String): String {
   return cacheData.computeIfAbsent(key){
       val result = fetchFromDb(key)
       result
   }
}

private fun fetchFromDb(key: String): String {
    val result = "India"
    sleep(100) // Simulate Db operations delay
    println("fetching from DB for key=$key")
    return result
}

private suspend fun fetchFromDbWithTimeout(key: String): String {
    return withTimeout(1000){
        fetchFromDb(key)
    }
}