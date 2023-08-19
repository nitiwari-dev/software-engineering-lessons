package patterns.strategy

import java.util.concurrent.TimeUnit

open class RemainingTimer(private val strategy: RemainingTimeStrategy) {

    fun remainingTimer(inputMills: Long): String =  strategy.execute(inputMills)

    fun interface RemainingTimeStrategy {
        fun execute(timeInMillis: Long): String
    }


    class HoursMinutes: RemainingTimeStrategy {
        override fun execute(timeInMillis: Long): String {
            val toHours = TimeUnit.MILLISECONDS.toHours(timeInMillis)
            val toMinutes = convertMillsToMinutes(timeInMillis, toHours)
            return if (toHours == 0L) "$toMinutes min"
            else if(toMinutes == 0L)"$toHours hr" else "$toHours hr $toMinutes min"
        }

        private fun convertMillsToMinutes(timeInMillis: Long, toHours: Long) =
            TimeUnit.MILLISECONDS.toMinutes(timeInMillis - TimeUnit.HOURS.toMillis(toHours))

    }

    class MinutesSeconds: RemainingTimeStrategy {
        override fun execute(timeInMillis: Long): String {
            val toMinutes = TimeUnit.MILLISECONDS.toMinutes(timeInMillis)

            return if(toMinutes == 0L){
                val toSeconds = TimeUnit.MILLISECONDS.toSeconds(timeInMillis)
                "$toSeconds sec"
            }else {
                val toSeconds = TimeUnit.MILLISECONDS.toSeconds(timeInMillis - TimeUnit.MINUTES.toMillis(toMinutes))
                if (toSeconds > 0) "$toMinutes min $toSeconds sec" else "$toMinutes min"
            }
        }

    }
<<<<<<< HEAD
}

