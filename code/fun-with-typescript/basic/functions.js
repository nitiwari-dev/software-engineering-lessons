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
