// using ascii values
const reverseString1 = (input) => {
  if (!input.length) return
  let output
  for (let i = input.length - 1; i >= 0; i--) {
    if ((input[i].charCodeAt() >= 65 && input[i].charCodeAt() <= 90) || // A-Z
    (input[i].charCodeAt() >= 97 && input[i].charCodeAt() <= 122) || // a-z
    (input[i].charCodeAt() >= 48 && input[i].charCodeAt() <= 57)) { // 0-9
      output = output ? output + input[i] : input[i]
    }
  }
  return output
}

// using regex
const reverseString2 = (input) => {
  if (!input.length) return
  let output
  for (let i = input.length - 1; i >= 0; i--) {
    if (input[i].match(/^[0-9a-zA-Z]+$/)) {
      output = output ? output + input[i] : input[i]
    }
  }
  return output
}

const input = 'hello_w1o%rl)d'
console.log(input, ' reverse output -->', reverseString1(input))
console.log(input, ' reverse output -->', reverseString2(input))
