let arr = [];

function insert(data) {
  arr.push(data);
  let pos = arr.length-1;
  heapifyUp(pos)
}

function heapifyUp(pos) {
  if (pos > -1) {
  let parent = Math.floor((pos + 1) / 2) - 1;
  if (arr[pos] < arr[parent]) {
    arr[parent] =  arr[parent] +  arr[pos];
    arr[pos] = arr[parent] -  arr[pos];
    arr[parent] = arr[parent] -  arr[pos];
  }
  heapifyUp(parent);
}
return;
}

function heapifyDown(pos) {
  let left = 2* pos + 1;
  let right = 2* pos + 2;
  let parent = pos;
  if (arr[left] && arr[left] < arr[parent]) {
    arr[parent] =  arr[parent] +  arr[left];
    arr[left] = arr[parent] -  arr[left];
    arr[parent] = arr[parent] -  arr[left];
  }
  if (arr[right]  && arr[right] < arr[parent]) {
    arr[parent] =  arr[parent] +  arr[right];
    arr[right] = arr[parent] -  arr[right];
    arr[parent] = arr[parent] -  arr[right];
  }
  if (left > arr.length-1 || right > arr.length-1) {
    return;
  } else {
    heapifyDown(left);
    heapifyDown(right);
  }
}

function remove(el) {
  for (var i = 0; i < arr.length; i++) {
   if (el == arr[i]) {
     let end = arr.pop();
     arr[i]  = end;
     let parent = Math.floor((i + 1) / 2) - 1;
     if (arr[i] < arr[parent]) {
      heapifyUp(i);
     } else {
      heapifyDown(i);
     }  
     break;
   }
}}

function isEmpty() {
  return arr.length == -1;
}

insert(5);
insert(4);
insert(8);
insert(6);
insert(1);
insert(14);
insert(2);
insert(7);
console.log("final min heap", arr);
remove(6);
console.log("final min heap after remove 6", arr);
console.log("is Heap Empty", isEmpty());