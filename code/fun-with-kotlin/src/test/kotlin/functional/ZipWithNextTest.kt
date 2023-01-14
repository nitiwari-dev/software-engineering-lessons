package functional

import org.assertj.core.api.Assertions
import org.assertj.core.api.Assertions.*
import org.junit.jupiter.api.BeforeAll
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.TestInstance
import org.mockito.Mockito.verify

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class ZipWithNextTest {
    lateinit var zipWithNext: ZipWithNext

    @BeforeAll
    fun setUp(){
        zipWithNext = ZipWithNext()
    }

    @Test
    fun `given list of numbers return the adjacent largest product using imperative style`(){
        val input = mutableListOf(1, 2, 3)
        val maxProduct = 6
        assertThat(zipWithNext.legacyLargestProduct(input)).isEqualTo(maxProduct)
    }

    @Test
    fun `given list of numbers return the adjacent largest product using functional style`(){
        val input = mutableListOf(1, 2, 3)
        val maxProduct = 5
        assertThat(zipWithNext.functionalLargestProduct(input)).isEqualTo(maxProduct)
    }
}