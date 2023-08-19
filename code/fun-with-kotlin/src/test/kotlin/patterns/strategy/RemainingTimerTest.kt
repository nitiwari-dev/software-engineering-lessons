package patterns.strategy

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.BeforeAll
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.TestInstance

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class RemainingTimerTest{

    private lateinit var minutesSecondsTimer: RemainingTimer
    private lateinit var hrsMinutesTimer: RemainingTimer

    @BeforeAll
    fun setUp(){
        minutesSecondsTimer = RemainingTimer(RemainingTimer.MinutesSeconds())
        hrsMinutesTimer = RemainingTimer(RemainingTimer.HoursMinutes())
    }

    @Test
    fun `given time less than 59 seconds return 59 sec applying minutes seconds strategy`(){

        val inputTimeMillis = 59_000L
        val actual = minutesSecondsTimer.remainingTimer(inputTimeMillis)
        val expectedOutput = "59 sec"
        assertThat(actual).isEqualTo(expectedOutput)
    }

    @Test
    fun `given time less than 0 seconds return 0 sec applying minutes seconds strategy`(){

        val inputTime = 0L
        val actual = minutesSecondsTimer.remainingTimer(inputTime)
        val expectedOutput = "0 sec"
        assertThat(actual).isEqualTo(expectedOutput)
    }

    @Test
    fun `given time equal to 60 seconds return 1 minutes applying minutes seconds strategy`(){
        val inputTime = 60_000L
        val actual = minutesSecondsTimer.remainingTimer(inputTime)
        val expectedOutput = "1 min"
        assertThat(actual).isEqualTo(expectedOutput)
    }

    @Test
    fun `given time equal to 60 minutes return 1 hr applying hrs minutes strategy`(){
        val inputTime = 36_00_000L
        val actual = hrsMinutesTimer.remainingTimer(inputTime)
        val expectedOutput = "1 hr"
        assertThat(actual).isEqualTo(expectedOutput)
    }

    @Test
    fun `given time equal to 59 minutes return 59 min applying hrs minutes strategy`(){

        val inputTime = 35_40_000L
        val actual = hrsMinutesTimer.remainingTimer(inputTime)
        val expectedOutput = "59 min"
        assertThat(actual).isEqualTo(expectedOutput)
    }
}