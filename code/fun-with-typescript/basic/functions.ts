

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







export { };
