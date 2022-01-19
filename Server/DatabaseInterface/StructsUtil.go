package main

type Package struct {
	Name         string
	Description  string
	UUID         string
	Authors      []string
	Version      int64
	Source       string    // Path to source code, binaries on disk for server
	Dependencies []Package // Recursive reference to other packages
	Graph        []DependencyGraph
}

type Project struct {
	Name          string
	Description   string
	UUID          string
	LatestVersion int64
	Packages      []Package
}

type Team struct {
	TeamName string
	Members  []Member
}

type DependencyGraph struct {
	Dependencies []Package
}

// Recursively iterates through list of dependencies, building on passed dependency-graph through parameter
func BuildGraph(item Package, graph DependencyGraph) {
	for _, dep := range item.Dependencies {

		// Special loop that checks that package hasn't already been added to dependency-graph to prevent duplicates
		var found bool = false
		for _, dups := range graph.Dependencies {
			if dep.UUID == dups.UUID {
				found = true
			}
		}

		// If item not already found, add to list 'graph', recursively call on package to see if more dependencies nested
		if !found {
			graph.Dependencies = append(graph.Dependencies, dep)
			BuildGraph(dep, graph)
		}
	}
}
