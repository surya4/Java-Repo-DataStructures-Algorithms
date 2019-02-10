let adj_matrix = [
  [0,1,1,0,0,0,0,0],
  [0,0,1,1,0,0,0,0],
  [1,0,0,1,1,0,1,0],
  [0,1,1,0,1,0,0,0],
  [0,0,1,1,0,1,0,0],
  [0,0,0,0,1,0,0,1],
  [0,0,1,0,0,0,0,0],
  [0,0,0,0,0,1,1,1]
]

function printGraph() {
  for (var i = 0; i < adj_matrix.length; i++) {
    // console.log(adj_matrix[i][1]);
    process.stdout.write(String(i+1))
    process.stdout.write(" | ")
    // console.log("1234", adj_matrix[i])
    for (var j = 0; j < adj_matrix[i].length; j++) {
      if (adj_matrix[i][j] == 1) {
        process.stdout.write(String(j+1))
        process.stdout.write("--> ")
      }
    }
    process.stdout.write('null')
    console.log();
  }
}

printGraph();

// Ouput 

// 1 | 2--> 3--> null
// 2 | 3--> 4--> null
// 3 | 1--> 4--> 5--> 7--> null
// 4 | 2--> 3--> 5--> null
// 5 | 3--> 4--> 6--> null
// 6 | 5--> 8--> null
// 7 | 3--> null
// 8 | 6--> 7--> 8--> null