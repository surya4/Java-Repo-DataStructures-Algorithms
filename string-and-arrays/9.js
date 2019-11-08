const stack = [];
const map = {
  '(': ')',
  '[': ']',
  '{': '}'
}


function parenCheck(inp) {
  for (let i = 0; i < inp.length; i++) {
    if (inp[i] == "{" || inp[i] == "(" || inp[i] == "[") {
      stack.push(inp[i]);
    } else {
      const p = stack.pop();
      if (inp[i] != map[p]) {
        return false;
      }
    }
  }
  if (stack.length !== 0) return false;
  return true; 
}

const input = "{([()])}";
console.log("parethesis check", parenCheck(input));