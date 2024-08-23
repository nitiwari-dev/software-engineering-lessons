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


// number - both integer and float donatest he same type as number
let planetId : number = 1001
let planetCordinates = 3.1415 // type infereance

// string
let username: string = "Mandolarian"

// boolean
let isUserFromEarth:boolean = false

/*  
any - if you want to not have type-checked
        use noImplicitAny within the tsconfig.json
*/  

//BAD practice. As a developer its difficult to intefer the response

let result: any
function networkRequest(){
    return "success"
}
result = networkRequest()
console.log(result);


//Good Practice
let resultObject: string
resultObject = networkRequest()
console.log(resultObject);


// array
let gamesId: Array<number> = [1, 2, 3, 5]
gamesId.sort
console.log(gamesId)

//union type either or in the type as input

function showCaseGameId(id: number | string){
    console.log(id);
}

function showCaseGameNumber(id: number[] | number){
    if (Array.isArray(id)){
        id.sort()
        console.log(id);
    } else {
        console.log(id)
    }
}

// union type either or in the type as out
function compare(a: number, b: number): -1 | 0 | 1 {
    return a == b ? 0 : a > b ? 1: -1
}
let a = 10
let b = 20
console.log("comparing ", compare(10, 20));


showCaseGameNumber([9, 0 ,0, 1])
showCaseGameNumber(1)

// type alias

type Growth = {
    rate: number,
    value: number
}


function findRateOfGrowth(growth: Growth): number {
    return growth.rate/100 * growth.value
}

let growth: Growth = {rate: 10, value:10}
console.log(findRateOfGrowth(growth))


// type alias - using interfaces

interface GrowthRate {
    ratePercentage: number
    value: number
}

function findGrowthRate2(growth: GrowthRate): number {
    return growth.ratePercentage *  growth.value
}

console.log(findGrowthRate2({ratePercentage: 10, value: 5}));


// type cast with example
interface Auth {
    token: string
}

interface HMAC extends Auth {
    alog: string
}

function verifyAuth(token: string): Auth {
    let auth: HMAC = {alog: "SHA256", token: token}
    return auth
}

let auth = <HMAC>verifyAuth("1234")
console.log(auth.alog);
console.log(auth.token);


// null and undefined using strictNullChecks into tsconfig.json
// using ! as not null assertion (for kotlin programmers its !!)

function validateToken(token: string | null): boolean {
    if (token != null){
        console.log(token);
        return true
    }

    return false
}


// Truthiness narrowing safe guarding and coerce 
// Typeguard find the type of variable e.g typeof id === "number" will return true if id is a number type
// This process of typeguarding and checking with more specific checks are called narrowing in ts
// we can have typeof for "string" "number" "object" "boolean" "undefined" "function"
// Operator ? and ! are same as kotlin behaviour
// handle null cases with

//using trutiness narrowing
// function findAllCustomers(ids: string[] | string | null){
    //if (idstypeof ids == "object"){


function findAllCustomers(ids: string[] | string | null){
    if (typeof ids == "object"){
        ids?.forEach(element => {
            console.log(element)
        });
    }
}

findAllCustomers([])


export {}



