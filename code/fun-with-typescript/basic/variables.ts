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


export {}
