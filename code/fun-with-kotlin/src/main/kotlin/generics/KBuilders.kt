package generics

fun <K, V>buildMap(builderMap: MutableMap<K, V>.() -> Unit): MutableMap<K, V>{
    val result = mutableMapOf<K, V>()
    result.builderMap()
    return result
}

fun <T, A: MutableCollection<T>> Collection<T>.partitionTo(first: A, second: A, p: (T) -> Boolean): Pair<A, A> {
    for (e in this){
        if(p(e)){
            first.add(e)
        } else {
            second.add(e)
        }
    }

    return Pair(first, second)
}