package d07nospaceleftondevice

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const moveCommand = "cd"
const listCommand = "ls"
const dirIdentifier = "dir"
const rootDir = "/"
const upDir = ".."
const largeDirectory = 100000
const totalSpace = 70000000
const requiredFreeSpace = 30000000

type Monitor struct {
	smallFileSizeCount int
	spaceToFree        int
	currentResidual    int
}

func (m *Monitor) inspectDirectorySize(node *DirNode) {
	for _, d := range node.directories {
		m.inspectDirectorySize(d)
	}
	if node.directorySize <= largeDirectory {
		m.smallFileSizeCount += node.directorySize
	}
	if node.directorySize >= m.spaceToFree && node.directorySize < m.currentResidual {
		m.currentResidual = node.directorySize
	}
}

func (m *Monitor) calculateSpaceToFree(root *DirNode) {
	m.spaceToFree = requiredFreeSpace - (totalSpace - root.directorySize)
	m.currentResidual = root.directorySize
}

type DirNode struct {
	name          string
	parent        *DirNode
	directories   map[string]*DirNode
	filesize      int
	directorySize int
}

func newDir(name string, parent *DirNode) *DirNode {
	return &DirNode{
		name:          name,
		filesize:      0,
		parent:        parent,
		directories:   map[string]*DirNode{},
		directorySize: -1,
	}
}

func (d *DirNode) AddDirectory(s string) {
	if _, ok := d.directories[s]; ok {
		return
	}
	d.directories[s] = newDir(s, d)
}

func (d *DirNode) AddFile(size, name string) {
	val, err := strconv.Atoi(size)
	if err != nil {
		log.Fatal(err)
	}
	d.filesize += val
}

func (d *DirNode) GetSize() int {
	val := d.filesize
	for _, v := range d.directories {
		val += v.GetSize()
	}
	d.directorySize = val
	return val
}

func BuildDirectories(input []string, root *DirNode) {
	i := 0
	currentNode := root
	for i < len(input)-1 {
		command, target := parseCommand(input[i])
		switch command {
		case listCommand:
			i = addListedObjects(input, i, currentNode)
		case moveCommand:
			currentNode = moveDirectory(target, currentNode)
			i += 1
		}
	}
}

func addListedObjects(input []string, index int, current *DirNode) int {
	for index < len(input)-1 {
		index += 1
		line := input[index]
		switch line[0] {
		case '$':
			return index
		case 'd':
			components := strings.Split(line, " ")
			current.AddDirectory(components[1])
		default:
			components := strings.Split(line, " ")
			current.AddFile(components[0], components[1])
		}
	}
	return index
}

func moveDirectory(target string, c *DirNode) *DirNode {
	switch target {
	case upDir:
		return c.parent
	}
	return c.directories[target]
}

func parseCommand(s string) (string, string) {
	components := strings.Split(s, " ")
	if len(components) == 2 {
		return listCommand, ""
	}
	return moveCommand, components[2]
}

func InspectFilesystem(lines []string) (int, int) {
	root := newDir(rootDir, nil)
	root.parent = root
	BuildDirectories(lines[1:], root)
	root.GetSize()

	monitor := &Monitor{}
	monitor.calculateSpaceToFree(root)
	monitor.inspectDirectorySize(root)
	return monitor.smallFileSizeCount, monitor.currentResidual
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := InspectFilesystem(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
