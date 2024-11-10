package builders

fun <K, V>buildMap(builderMap: MutableMap<K, V>.() -> Unit): MutableMap<K, V>{
    val result = mutableMapOf<K, V>()
    result.builderMap()
    return result
}