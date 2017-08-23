# Concurrent-Maximum-Finder

### Overview

This application creates random integers and works out the maximum value concurrently. Users are asked to specify
* the seed to be used to generate pseudo random integers
* the number of psuedo random integers.

Each pseudo random number is assigned a process. Let us call it a node. Our aim is to have all nodes hold the maximum value. 

Four approaches are provided:
1. central.go: We create a single coordinator process. Each node sends the pseudo random number it holds to the coordinator via a channel. The coordinator works out the maximum by repeatedly reading in values from the nodes. Finally, the maximum is broadcast to all nodes. 
2. passiveRing.go: The nodes are configured in a ring. Every node reads in a value from the preceding node. Suppose that the value from the preceding node is x and the value stored in the current node is y. The largest of x and y is then passed on to the following node. This continues until one lap is completed. In the second round, the maximum is passed from one node to another. 
3. activeRing.go: Nodes are again configured in a ring. They pass their values to the following nodes concurrently. Each node keeps track of the maximum of all values that it has read in. This continues until all pseudo random numbers come back to the respective origin nodes. 
4. tree.go: The nodes are configured in a binary tree. All nodes, apart from leaves, read in values from their children, calculating the largest of the inputs and the values stored in the current nodes and passing the result to the parents if possible. At this end of this stage, the root holds the overall maximum. In the second stage, this maximum value is cascaded down the tree. 