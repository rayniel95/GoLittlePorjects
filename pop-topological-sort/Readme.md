# About

Another way to perform topological sorting on a directed acyclic graph G=(V,E) is to repeatedly find a vertex of in-degree 0, output it, and remove it and
all of its outgoing edges from the graph. Explain how to implement this idea so
that it runs in time O(V+E). What happens to this algorithm if G has cycles?
