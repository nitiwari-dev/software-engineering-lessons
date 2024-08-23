"use strict";
//Types
/**
 * 1. number
 * 2. string
 * 3. boolean
 * 4. null
 * 5. undefined
 * 6. void
 * 7. object
 * 8. array
 * 9. tuples
 * 10. Any
 * 11. Never
 * 12. Unknown
 *
 * Note: all lower case
 */
Object.defineProperty(exports, "__esModule", { value: true });
// number - both integer and float donatest he same type as number
var planetId = 1001;
var planetCordinates = 3.1415; // type infereance
// string
var username = "Mandolarian";
// boolean
var isUserFromEarth = false;
/*
any - if you want to not have type-checked
        use noImplicitAny within the tsconfig.json
*/
//BAD practice. As a developer its difficult to intefer the response
var result;
function networkRequest() {
    return "success";
}
result = networkRequest();
console.log(result);
//Good Practice
var resultObject;
resultObject = networkRequest();
console.log(resultObject);
// array
var gamesId = [1, 2, 3, 5];
gamesId.sort;
console.log(gamesId);
//union type either or in the type
function showCaseGameId(id) {
    console.log(id);
}
function showCaseGameNumber(id) {
    if (Array.isArray(id)) {
        id.sort();
        console.log(id);
    }
    else {
        console.log(id);
    }
}
showCaseGameNumber([9, 0, 0, 1]);
showCaseGameNumber(1);
