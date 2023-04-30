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
//BAD practice
var result;
function networkRequest() {
    return { "success": true };
}
result = networkRequest();
console.log(result);
var resultObject;
resultObject = networkRequest();
console.log(resultObject);
