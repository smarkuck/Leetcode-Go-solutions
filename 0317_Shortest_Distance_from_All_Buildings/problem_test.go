package leetcode

import "testing"

type Land struct {
	landGrid     [][]int
	distanceGrid [][]int
	sizeX, sizeY int
}

type DistancesCalculator struct {
	land            *Land
	queue           []Field
	visited         [][]bool
	currentDistance int
}

type Field struct {
	x, y int
}

var directions = []Field{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func shortestDistance(grid [][]int) int {
	return newLand(grid).getShortestDistance()
}

func newLand(grid [][]int) *Land {
	sizeX, sizeY := len(grid), len(grid[0])

	l := &Land{
		landGrid: grid, distanceGrid: nil,
		sizeX: sizeX, sizeY: sizeY}

	l.distanceGrid = make([][]int, sizeX)
	for x := 0; x < sizeX; x++ {
		l.distanceGrid[x] = make([]int, sizeY)
	}

	return l
}

func (l *Land) getShortestDistance() int {
	if l.tryComputeDistancesToBuildings() == false {
		return -1
	}

	return l.findShortestDistance()
}

func (l *Land) tryComputeDistancesToBuildings() bool {
	for x := 0; x < l.sizeX; x++ {
		for y := 0; y < l.sizeY; y++ {
			if l.landGrid[x][y] == 1 {
				dc := l.newDistancesCalculator(Field{x, y})
				if dc.tryComputeDistancesToBuilding() == false {
					return false
				}
			}
		}
	}

	return true
}

func (l *Land) newDistancesCalculator(building Field) *DistancesCalculator {
	dc := &DistancesCalculator{
		land: l, queue: []Field{building},
		visited: nil, currentDistance: 1}

	dc.visited = make([][]bool, l.sizeX)
	for x := 0; x < l.sizeX; x++ {
		dc.visited[x] = make([]bool, l.sizeY)
	}

	dc.visited[building.x][building.y] = true

	return dc
}

func (dc *DistancesCalculator) tryComputeDistancesToBuilding() bool {
	for len(dc.queue) > 0 {
		for fieldsToProcess := len(dc.queue); fieldsToProcess > 0; fieldsToProcess-- {
			target := dc.queue[0]
			dc.queue = dc.queue[1:]
			dc.computeDistancesToBuildingsForNeighbours(target)
		}
		dc.currentDistance++
	}

	return dc.setUnusedEmptyFieldsAndCheckIfCanGetToAllBuildings()
}

func (dc *DistancesCalculator) computeDistancesToBuildingsForNeighbours(target Field) {
	for _, direction := range directions {
		neighbour := Field{target.x + direction.x, target.y + direction.y}
		if dc.shouldVisit(neighbour) {
			dc.visited[neighbour.x][neighbour.y] = true
			if dc.land.landGrid[neighbour.x][neighbour.y] == 1 {
				continue
			}
			dc.land.distanceGrid[neighbour.x][neighbour.y] += dc.currentDistance
			dc.queue = append(dc.queue, neighbour)
		}
	}
}

func (dc *DistancesCalculator) shouldVisit(neighbour Field) bool {
	return 0 <= neighbour.x && neighbour.x < dc.land.sizeX &&
		0 <= neighbour.y && neighbour.y < dc.land.sizeY &&
		!dc.visited[neighbour.x][neighbour.y] &&
		dc.land.landGrid[neighbour.x][neighbour.y] < 2
}

func (dc *DistancesCalculator) setUnusedEmptyFieldsAndCheckIfCanGetToAllBuildings() bool {
	for x := 0; x < dc.land.sizeX; x++ {
		for y := 0; y < dc.land.sizeY; y++ {
			if !dc.visited[x][y] {
				if dc.land.landGrid[x][y] == 0 {
					dc.land.landGrid[x][y] = 3
				} else if dc.land.landGrid[x][y] == 1 {
					return false
				}
			}
		}
	}
	return true
}

func (l *Land) findShortestDistance() int {
	minDistance := 1000
	for x := 0; x < l.sizeX; x++ {
		for y := 0; y < l.sizeY; y++ {
			if l.landGrid[x][y] == 0 && l.distanceGrid[x][y] < minDistance {
				minDistance = l.distanceGrid[x][y]
			}
		}
	}

	if minDistance == 1000 {
		return -1
	}
	return minDistance
}

func TestShortestDistance(t *testing.T) {
	output := shortestDistance([][]int{
		{1, 0, 2, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0}})

	expected := 7
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = shortestDistance([][]int{{1, 0}})
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = shortestDistance([][]int{{1}})
	expected = -1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
