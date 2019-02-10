var str = '2(3(b(2(ab)c3(d))';
let num_stack = [];
let char_stack = [];
var arr = str.split('');

function splitString(inp) {
  for (var i = 0; i < arr.length; i++) {
   if (isNaN(arr[i]) && arr[i] != '(' && arr[i] != ')') {
      char_stack.push(getString(i));
    } else {
      console.log('arr[i]', arr[i])
      num_stack.push(parseInt(arr[i]));
    }
  }
  decodeString();
}

function getString(index){
  var ret_str = '';
  for (let i = index; i < arr.length; i++) {
    // console.log(typeof(parseInt(arr[i])) , arr[i]);
    if (isNaN(arr[i]) && arr[i] != '(' && arr[i] != ')') {
      // console.log('abra', arr[i])
      ret_str += arr[i];
    } else {
      return ret_str;
    }
  }
}

function decodeString() {
  let final_string = '';
  console.log('i am abra',num_stack, char_stack)
  for (var i = 0; i < char_stack.length; i++) {
    // console.log('i am abra',char_stack[i])
    let start_index=0, end_index=0;
    for (var j = 0; j < char_stack.length; j++) {
      if (char_stack[j] === '(') {
        start_index = j;
      }
      if (char_stack[j] === ')') {
        end_index = j;
        if (start_index != end_index) {
          final_string = str.substring(start_index, end_index+1);
          final_string = final_string.repeat(start_index-1)
        }
        // console.log('dabra', final_string, start_index, end_index)
      }
     // console.log('i am abra',start_index, end_index)    
    }
  }
}

splitString(str);