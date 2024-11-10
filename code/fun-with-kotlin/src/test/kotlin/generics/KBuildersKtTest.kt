package generics

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class KBuildersKtTest {

    @Test
    fun `given entries into mutable map as lamda then return builder with mutable map and match`() {
        val actual = mutableMapOf(1 to "one")
        val expected = buildMap {
            put(1, "one")
        }

        assertThat(actual).isEqualTo(expected)
    }

    @Test
    fun `give empty mutable map should not add any entries `() {
        val actual = emptyMap<Int, String>()
        val expected = buildMap<Int, String> {}
        assertThat(actual).isEqualTo(expected)
    }

    @Test
    fun `given first and second choice return the partitioned output`(){
        val input = listOf(1, 2, 3, 4, 5, 6)
        val (even, odd) = input.partitionTo(mutableListOf(), mutableListOf()){ i -> i % 2 == 0 }

        assertThat(even).isEqualTo(listOf(2, 4, 6))
        assertThat(odd).isEqualTo(listOf(1, 3, 5))
    }
}