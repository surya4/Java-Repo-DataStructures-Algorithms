// O(n) time and O(1) space
function isValidSubsequenc1 (array, sequence) {
  if (!array || !array.length || !sequence || !sequence.length) return false
  if (array.length < sequence.length) return false
  let j = 0
  for (let i = 0; i < array.length; i++) {
    if (j === sequence.length) break
    if (array[i] === sequence[j]) j++
  }
  return j === sequence.length
}

function isValidSubsequenc2 (array, sequence) {
  if (!array || !array.length || !sequence || !sequence.length) return false
  if (array.length < sequence.length) return false
  let i = 0
  let j = 0
  while (i < array.length && j < sequence.length) {
    if (array[i] === sequence[j]) j++
    i++
  }
  return j === sequence.length
}

console.log(isValidSubsequenc1([5, 1, 22, 25, 6, -1, 8, 10], [1, 6, -1, 10]))
console.log(isValidSubsequenc1([5, 1, 22, 25, 6, -1, 8, 10], [1, 6, -1]))
console.log(isValidSubsequenc2([5, 1, 1, 22, 25, 6, -1, 8, 10], [1, 6, -1]))
console.log(isValidSubsequenc2([5, 1, 1, 22, 25, 6, -1, 8, 10], [1, 6, -1, 5]))
