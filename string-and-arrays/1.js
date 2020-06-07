// Simply iterate loop from end to beginning and concat into a string - O(n)
const reverseString1 = (input) => {
  if (!input.length) return
  let output
  for (let i = input.length - 1; i >= 0; i--) {
    output = output ? output + input[i] : input[i]
  }
  return output
}

// Iterate from both end, i from beginneing and j from end and store in reverse order
// store jth value in ith position and ith value in jth position of array, convert the array
// the array to string - O(n)
const reverseString2 = (input) => {
  if (!input.length) return
  const output = []
  for (let i = 0, j = input.length; i < j; i++, j--) {
    output[i] = input[j]
    output[j] = input[i]
  }
  return output.join('')
}

// using a stack - O(n)
const reverseString3 = (input) => {
  if (!input.length) return
  const stack = []; const output = []
  for (let i = 0; i < input.length; i++) {
    stack.push(input[i])
  }
  for (let i = 0; stack.length > 0; i++) {
    output.push(stack.pop())
  }
  return output.join('')
}

const input = 'hello_world'

console.log(input, ' reverse output -->', reverseString1(input))
console.log(input, ' reverse output -->', reverseString2(input))
console.log(input, ' reverse output -->', reverseString3(input))
