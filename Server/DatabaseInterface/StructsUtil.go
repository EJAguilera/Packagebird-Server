package main

type Package struct {
	Name         string
	Description  string
	UUID         string
	Authors      []string
	Version      int64
	Source       string    // Path to source code, binaries on disk for server
	Dependencies []Package // Recursive reference to other packages
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
