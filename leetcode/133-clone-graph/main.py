# Definition for a undirected graph node
class UndirectedGraphNode:
    def __init__(self, x):
        self.label = x
        self.neighbors = []


# BFS
class Solution:
    # @param node, a undirected graph node
    # @return a undirected graph node
    def cloneGraph(self, node):
        if not node:
            return node
        root = UndirectedGraphNode(node.label)
        visited = {}
        visited[node] = root
        queue = []
        queue.append(node)
        while len(queue) > 0:
            head = queue.pop(0)
            for nb in head.neighbors:
                # 1. create new
                # 2. add new node to queue to ensure no revisit
                if nb not in visited:
                    temp = UndirectedGraphNode(nb.label)
                    visited[nb] = temp
                    visited[head].neighbors.append(temp)
                    queue.append(nb)
                # connect old node
                else:
                    visited[head].neighbors.append(visited[nb])
        return root
