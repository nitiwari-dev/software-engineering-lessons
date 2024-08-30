"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
//generic functionwith type 
function prepareOnboardingRequest(name, email, guestUser) {
    var request = { "name": name, "email": email, "isGuest": guestUser };
    return request;
}
console.log(prepareOnboardingRequest("Mando", "mando@madolarian.com", true));
// arrow function with type
var viewPagerEvent = function (pageId, pageName, timestamp) {
    if (timestamp === void 0) { timestamp = Date.now(); }
    return { "pageId": pageId, "pageName": pageName, "timestamp": timestamp };
};
var viewPagerEventDefault = function (pageId, pageName) {
    return viewPagerEvent(pageId, pageName);
};
console.log(viewPagerEventDefault(1001, "Onboarding"));
//using map and reduce as functions
var sessionsEpisodes = [
    {
        "session": "s1",
        "episodes": ["e2", "e1"]
    },
    {
        "session": "s2",
        "episodes": ["e1"]
    }
];
var groupEpisodes = sessionsEpisodes.reduce(function (groupeEpisodes, sessionEpisode) {
    return groupeEpisodes.concat(sessionEpisode.episodes);
}, []);
console.log(groupEpisodes);
var groupSessions = sessionsEpisodes.reduce(function (groupSessions, sessionEpisode) {
    return groupSessions.concat(sessionEpisode.session);
}, []);
console.log(groupSessions);
//function with void or never as good 
function fail(message) {
    throw new Error(message);
}
function logFailure(message) {
    console.log(message);
}
//function as type expression. similar to lamda function in java or kotlin world
function networkRequest(userId, request) {
    return request(userId);
}
var userId = "1001";
var request = networkRequest(userId, function () {
    return "onboarding completed for ".concat(userId);
});
console.log(request);
function processInput(input) {
    input.apply(null, [100]);
}
var processReq = function (code) {
    console.log(code);
};
processInput(processReq);
// generics with type inteference, similar to reified type in kotlin solving the problem of type erasure during runtime
function myList(t) {
    return [t];
}
console.log(myList(10));
//extension function on the property of the type
function maxOf(a, b) {
    if (a.length >= b.length)
        return a;
    else
        return b;
}
console.log(maxOf([1, 3], [2, 3, 4]));
//extension function to mark the character a cap. This is just a example not recommeneded as we can simply write a fucntion
function capFirstCharacter(s) {
    if (s.length == 0)
        return s;
    return s[0].toUpperCase() + s.substring(1).toLowerCase();
}
console.log(capFirstCharacter("hello"));
