// using summation of numbers
const missingNumber1 = (input) => {
  if (!input.length) return
  const last = input[input.length - 1]
  let expectedSum = (last * (last + 1)) / 2
  for (let i = 0; i < input.length; i++) {
    expectedSum -= input[i]
  }
  return expectedSum
}

// using xor
// as xor gives - 1 when both are different hence can be used
// 0,0 - 0
// 0,1 - 1
// 1,0 - 1
// 0,0 - 0
const missingNumber2 = (input) => {
  if (!input.length) return
  let xor1 = 1
  let xor2 = input[0]

  for (let i = 2; i <= input[input.length - 1]; i++) {
    xor1 = xor1 ^ i
  }

  for (let i = 1; i < input.length; i++) {
    xor2 = xor2 ^ input[i]
  }

  return xor1 ^ xor2
}

const input = [1, 2, 3, 4, 5, 7, 8]
console.log(input, ' missing output -->', missingNumber1(input))
console.log(input, ' missing output -->', missingNumber2(input))
