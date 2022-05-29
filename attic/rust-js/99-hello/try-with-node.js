const pkg = require('./hello/pkg/hello.js')
const {hello} = pkg

console.log(hello("Bob")) 
console.log(hello("Jane")) 
console.log(hello("John")) 

