// Case1: O(n2) approach | space - O(1)
function twoNumberSum1 (array, targetSum) {
  if (!array || !array.length || targetSum === null) return []
  for (let i = 0; i < array.length - 1; i++) {
    for (let j = i + 1; j < array.length; j++) {
      if (array[i] + array[j] === targetSum) {
        return [array[i], array[j]]
      }
    }
  }
  return []
}

// Case21: O(n) time | extra hash space - O(n)
// Problem here is object don't allow duplicate keys so it will fail
// for case like [0,4,3,0], 0
function twoNumberSum21 (array, targetSum) {
  if (!array || !array.length || targetSum === null) return []
  const hash = {}
  for (let i = 0; i < array.length; i++) {
    if (hash[targetSum - array[i]] !== undefined) {
      return [hash[targetSum - array[i]], array[i]]
    }
    hash[array[i]] = array[i]
  }
  return []
}

// Case21: using map - O(n) time | extra hash space - O(n)
function twoNumberSum22 (array, targetSum) {
  if (!array || !array.length || targetSum === null) return []
  const map = new Map()
  for (let i = 0; i < array.length; i++) {
    if (map.has(targetSum - array[i])) {
      return [map.get(targetSum - array[i]), i]
    }
    map.set(array[i], i)
  }
  return []
}

// Case3 :O(nlogn) time , space - O(n)
// Problem - if return expected in i,j then output comes wrong
function twoNumberSum31 (array, targetSum) {
  if (!array || !array.length || targetSum === null) return []
  array = array.sort((a, b) => a - b)
  for (let i = 0; i < array.length; i++) {
    for (let j = array.length - 1; j > i;) {
      if (array[i] + array[j] > targetSum) j--
      else if (array[i] + array[j] < targetSum) break
      else if (array[i] + array[j] === targetSum) return [array[i], array[j]]
    }
  }
  return []
}

// Case 3.2 - using while loop :O(nlogn) time , space - O(n)
// Problem - if return expected in i,j then output comes wrong
function twoNumberSum32 (array, targetSum) {
  if (!array || !array.length || targetSum === null) return []
  array = array.sort((a, b) => a - b)
  let left = 0
  let right = array.length - 1
  while (left < right) {
    if (array[left] + array[right] > targetSum) right--
    else if (array[left] + array[right] < targetSum) left++
    else return [array[left], array[right]]
  }
  return []
}

console.log(twoNumberSum1([3, 5, -4, 8, 11, 1, -1, 6], 10))
console.log(twoNumberSum21([3, 3], 6))
console.log(twoNumberSum22([3, 3], 6))
console.log(twoNumberSum31([3, 2, 4], 6))
console.log(twoNumberSum32([3, 2, 4], 6))
