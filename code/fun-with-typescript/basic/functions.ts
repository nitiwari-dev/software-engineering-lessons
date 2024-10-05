//generic functionwith type 
function prepareOnboardingRequest(name: string, email: string, guestUser: boolean): object {
    let request = { "name": name, "email": email, "isGuest": guestUser }
    return request
}
console.log(prepareOnboardingRequest("Mando", "mando@madolarian.com", true));


// arrow function with type
let viewPagerEvent = (pageId: number, pageName: string, timestamp = Date.now()) => {
    return { "pageId": pageId, "pageName": pageName, "timestamp": timestamp }
}
let viewPagerEventDefault = (pageId: number, pageName: string) => {
    return viewPagerEvent(pageId, pageName)
}
console.log(viewPagerEventDefault(1001, "Onboarding"))

//using map and reduce as functions

const sessionsEpisodes = [
    {
        "session": "s1",
        "episodes": ["e2", "e1"]
    },

    {
        "session": "s2",
        "episodes": ["e1"]
    }
]


const groupEpisodes = sessionsEpisodes.reduce((groupeEpisodes: string[], sessionEpisode) => {
    return groupeEpisodes.concat(sessionEpisode.episodes)
}, [])
console.log(groupEpisodes);

const groupSessions = sessionsEpisodes.reduce((groupSessions: string[], sessionEpisode) => {
    return groupSessions.concat(sessionEpisode.session)
}, [])
console.log(groupSessions);


//function with void or never as good 

function fail(message: string): never {
    throw new Error(message)
}

function logFailure(message: string): void {
    console.log(message);
}

//function as type expression. similar to lamda function in java or kotlin world

function networkRequest(userId: string, request: (userId: string) => string){
    return request(userId)
}

let userId = "1001"
const request = networkRequest(userId, (): string => {
    return `onboarding completed for ${userId}`
})

console.log(request);

// type as function param

type InputRequest = (code: number) => void

function processInput(input: InputRequest){
    input.apply(null, [100])
}
const processReq: InputRequest = (code: number) => {
    console.log(code)
}
processInput(processReq)


// generics with type inteference, similar to reified type in kotlin solving the problem of type erasure during runtime
function myList<T>(t: T): T[]{
    return [t]
}

console.log(myList(10));


//extension function on the property of the type
function maxOf<Type extends {length: number}>(a: Type, b: Type): Type {
    if (a.length >= b.length) return a 
    else return b
}

console.log(maxOf([1, 3], [2,3,4]))

//extension function to mark the character a cap. This is just a example not recommeneded as we can simply write a fucntion

function capFirstCharacter<s extends string>(s: string): string {
    if (s.length == 0) return s
    return s[0].toUpperCase() + s.substring(1).toLowerCase()
}

console.log(capFirstCharacter("hello"));


export { };
