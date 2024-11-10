package builders

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
}