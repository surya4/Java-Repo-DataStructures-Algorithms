let adjList = {};

function Vertex(data) {
  this.data = data;
  this.tail = null;
}

function AdjacencyList(data) {
  this.head = null;
  this.tail = null;
}

function addEdge (e1, e2) {
  let node1 = new Vertex(e1);
  let node2 = new Vertex(e2);

  adjList[e1] = adjList[e1] || new AdjacencyList();
  adjList[e2] = adjList[e2] || new AdjacencyList();

  if (adjList[e1].head == null) {
    adjList[e1].head = node2;
    adjList[e1].tail = node2;
  } else {
    node2.tail = adjList[e1].tail;
    adjList[e1].tail = node2;
  }

  if (adjList[e2].head == null) {
    adjList[e2].head = node1;
    adjList[e2].tail = node1;
  } else {
    node1.tail = adjList[e2].tail;
    adjList[e2].tail = node1;
  }
 }

 function removeEdge(e1, e2) {
  if (!adjList) {
    console.log("empty adjList");
    return;
  }
  if (!adjList[e1]) {
    console.log("empty vertex");
    return;
  }
  let current = adjList[e1];
  while (current.tail.data != e2) {
    current = current.tail;
  }
  current.tail = current.tail.tail;
 }

function printGraph() {
   if (!adjList) {
     console.log("empty adjList")
   } else {
     for (var i in adjList) {
      current1 = adjList[i];
      process.stdout.write(String(Object.keys(adjList)[i-1]))
      process.stdout.write(" | ")
       if (current1) {
        while(current1.tail) {
          current1 = current1.tail;
          process.stdout.write(String(current1.data))
          process.stdout.write("--> ")
        }    
        process.stdout.write('null')
       }
       console.log(); 
     }
   }
 }

 addEdge(1,2)
 addEdge(1,3)
 addEdge(2,3)
 addEdge(2,4)
 addEdge(3,4)
 addEdge(3,5)
 addEdge(3,7)
 addEdge(4,5)
 addEdge(5,6)
 addEdge(6,8)
 addEdge(7,8)
 addEdge(8,8)

printGraph();
removeEdge(3,2);
console.log("-------------")
printGraph();
// Output : 

// 1--> 3--> 2--> null
// 2--> 4--> 3--> 1--> null
// 3--> 7--> 5--> 4--> 2--> 1--> null
// 4--> 5--> 3--> 2--> null
// 5--> 6--> 4--> 3--> null
// 6--> 8--> 5--> null
// 7--> 8--> 3--> null
// 8--> 8--> 8--> 7--> 6--> null

