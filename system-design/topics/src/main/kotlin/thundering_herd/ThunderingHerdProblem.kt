package thundering_herd

import kotlinx.coroutines.*
import java.lang.Thread.sleep
import java.util.concurrent.ConcurrentHashMap

/**
 * Imagine you're racing in a track competition, but instead of a regular track, you have to run through a field filled with hurdles. These hurdles are like small fences you have to jump over while running. Now, let's say there are a bunch of people racing alongside you, and everyone wants to win. But here's the catch: only one person can win, and if you bump into someone else while jumping over a hurdle, both of you fall down and lose precious time.
 * Now, picture this happening not just with a few people, but with a whole bunch of racers, all trying to jump over hurdles at the same time. It gets chaotic, right? That's basically the thundering herd problem.
 * In computer terms, the "thundering herd" describes what happens when a lot of computer processes, like those in a server, all rush to do the same thing at the same time. Just like in the race, this can lead to inefficiency or even crashes, because too many processes are trying to access the same resource all at once. It's like everyone trying to jump over the same hurdle at the same time, causing a big pile-up.
 *
 * Kotlin:
 * In the example where we have span 5 coroutines jobs trying to get the keys and you will find that we miss the cache due to concurrent operation open
 * it can also lead to file locking of DB lockint
 *
 * Output:
 * fetching from DB for key=IN
 * fetching from DB for key=IN
 * fetching from DB for key=IN
 * result key=IN, result=India // result
 * result key=IN, result=India // result
 * result key=IN, result=India // result
 * fetching from DB for key=IN // result
 * result key=IN, result=India
 * fetching from DB for key=IN
 * result key=IN, result=India
 */

val cache = ConcurrentHashMap<String, String>()

@DelicateCoroutinesApi
fun main() {

    runBlocking {
        val keyToBeSearched = "IN"
        val jobCount = 5
        val jobsList = List(jobCount){
            GlobalScope.launch(){
                val result = fetchValueUsingKey(keyToBeSearched)
                println("result key=$keyToBeSearched, result=$result")
            }
        }

        //join all the jobs
        for (job in jobsList){
            job.join()
        }
    }

}


private fun fetchValueUsingKey(key: String): String {
    if (cache[key] == null){
        cache[key] = fetchFromDb(key)
    }
    return cache[key]!!
}

private fun fetchFromDb(key: String): String {
    val result = "India"
    sleep(100) // Simulate Db operations delay
    println("fetching from DB for key=$key")
    return result
}