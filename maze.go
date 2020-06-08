package maze

import "fmt"

// Package maze lets you build mazes and navigate between points.
//
// Maze uses Breadth-first search for path-finding, which guarantees finding
// the shortest path between two points. The package also has tools for easy
// manipulation of the maze's terrain.
//
// The pathfinding algorithm only allows non-diagonal movements. Each point
// is marked as either road (walkable) or obstacle (not walkable). Due to
// utilization of Breadth-first search, pathfinding in maze will be costly
// in open environments with long distances between points.

type point struct {
	// Represents a point in the maze
	x int
	y int
	obstacle bool
	marked bool
	path bool
	lastPoint *point
}

func newPoint(x int, y int) *point {
	// Point constructor
	var point point
	point.x = x
	point.y = y
	return &point
}

type Maze struct {
	// Represents a maze
	width int
	height int
	field [][]point

}

func New(width int, height int) *Maze {
	// Maze constructor
	var maze Maze
	maze.width = width
	maze.height = height

	// Fill field with points and assign their coordinates
	field := make([][]point, width)
	for i := range field {
		field[i] = make([]point, height)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			field[x][y].x = x
			field[x][y].y = y
		}
	}
	maze.field = field
	return &maze
}

func (m *Maze) pointInBounds(x int, y int) bool {
	return x >= 0 && x < m.width && y >= 0 && y < m.height
}

func (m *Maze) SetObstacle(x int, y int) {
	if m.pointInBounds(x, y) {
		m.field[x][y].obstacle = true
	} else {
		panic("Point out of bounds")
	}
}

func (m *Maze) SetObstacleArea(fromX int, fromY int, toX int, toY int) {
	if m.pointInBounds(fromX, fromY) && m.pointInBounds(toX, toY) {
		if fromX > toX {
			temp := toX
			toX = fromX
			fromX = temp
		}
		if fromY > toY {
			temp := toY
			toY = fromY
			fromY = temp
		}
		for i := fromX; i <= toX; i++ {
			for j := fromY; j <= toY; j++ {
				m.field[i][j].obstacle = true
			}
		}
	} else {
		panic("Point out of bounds")
	}
}

func (m *Maze) SetRoad(x int, y int) {
	if m.pointInBounds(x, y) {
		m.field[x][y].obstacle = false
	} else {
		panic("Point out of bounds")
	}
}

func (m *Maze) SetRoadArea(fromX int, fromY int, toX int, toY int) {
	if m.pointInBounds(fromX, fromY) && m.pointInBounds(toX, toY) {
		if fromX > toX {
			temp := toX
			toX = fromX
			fromX = temp
		}
		if fromY > toY {
			temp := toY
			toY = fromY
			fromY = temp
		}
		for i := fromX; i <= toX; i++ {
			for j := fromY; j <= toY; j++ {
				m.field[i][j].obstacle = false
			}
		}
	} else {
		panic("Point out of bounds")
	}
}

func (m *Maze) GetDirections(fromX int, fromY int, toX int, toY int) []byte {
	// Returns a slice with directions in the form of cardial marks
	if m.pointInBounds(fromX, fromY) && m.pointInBounds(toX, toY) {
		m.resetMazePath()

		var list []byte
		directionsList := m.directionsHelper(fromX, fromY, toX, toY)
		for len(directionsList) > 1 {
			n := len(directionsList) - 1
			p1 := directionsList[n]
			p2 := directionsList[n-1]
			if p2.x-p1.x == 0 {
				if p2.y-p1.y == 1 {
					list = append(list, 'N')
				} else {
					list = append(list, 'S')
				}
			} else if p2.x-p1.x == 1 {
				list = append(list, 'E')
			} else {
				list = append(list, 'W')
			}
			directionsList = directionsList[:n]
		}
		return list
	} else {
		panic("Point out of bounds")
	}
}

func (m *Maze) resetMazeMarks() {
	for x := 0; x < m.width; x++ {
		for y := 0; y < m.height; y++ {
			m.field[x][y].marked = false
		}
	}
}

func (m *Maze) resetMazePath() {
	for x := 0; x < m.width; x++ {
		for y := 0; y < m.height; y++ {
			m.field[x][y].path = false
			m.field[x][y].lastPoint = nil
		}
	}
}

func (m *Maze) directionsHelper(fromX int, fromY int, toX int, toY int) []*point {
	if m.pointInBounds(fromX, fromY) && m.pointInBounds(toX, toY) &&
		!m.field[fromX][fromY].obstacle && !m.field[toX][toY].obstacle {
		// Returns a list containing references to the points visited in shortest path
		m.resetMazeMarks()

		var directionsList []*point

		var queue []*point
		queue = append(queue, &m.field[fromX][fromY])

		for len(queue) > 0 {
			// Search using Breadth-first search
			var tempPoint *point
			tempPoint = queue[0]
			queue[0] = nil
			queue = queue[1:]
			tempPoint.marked = true

			if tempPoint.x == toX && tempPoint.y == toY {
				// When path found, generate directionsList
				p := tempPoint
				for p != nil {
					directionsList = append(directionsList, p)
					p = p.lastPoint
				}
				break
			}

			// Add adjacent points to queue
			// Eastern point
			if tempPoint.x+1 < m.width {
				if p := &m.field[tempPoint.x+1][tempPoint.y]; !p.marked && !p.obstacle {
					p.lastPoint = tempPoint
					queue = append(queue, p)
				}
			}
			// Northern point
			if tempPoint.y+1 < m.height {
				if p := &m.field[tempPoint.x][tempPoint.y+1]; !p.marked && !p.obstacle {
					m.field[tempPoint.x][tempPoint.y+1].lastPoint = tempPoint
					queue = append(queue, p)
				}
			}
			// Western point
			if tempPoint.x-1 >= 0 {
				if p := &m.field[tempPoint.x-1][tempPoint.y]; !p.marked && !p.obstacle {
					m.field[tempPoint.x-1][tempPoint.y].lastPoint = tempPoint
					queue = append(queue, p)
				}
			}
			// Southern point
			if tempPoint.y-1 >= 0 {
				if p := &m.field[tempPoint.x][tempPoint.y-1]; !p.marked && !p.obstacle {
					m.field[tempPoint.x][tempPoint.y-1].lastPoint = tempPoint
					queue = append(queue, p)
				}
			}
		}
		return directionsList
	} else {
		panic("Point out of bounds or on obstacle")
	}
}


func (m *Maze) Print() {
	// Draws a map of the maze
	s := ""
	for i := 0; i < m.width; i++ {
		s += "-"
	}
	fmt.Printf("+" + s + "+\n")
	for y := m.height-1; y >= 0; y-- {
		fmt.Printf("|")
		for x := 0; x < m.width; x++ {
			if m.field[x][y].obstacle {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("+" + s + "+\n")
}

func (m *Maze) PrintDirections(fromX int, fromY int, toX int, toY int) {
// Draws a map with directions
	if m.pointInBounds(fromX, fromY) && m.pointInBounds(toX, toY) {
		m.resetMazePath()

		directionsList := m.directionsHelper(fromX, fromY, toX, toY)
		for i := 0; i < len(directionsList); i++ {
			p := directionsList[i]
			m.field[p.x][p.y].path = true
		}
		s := ""
		for i := 0; i < m.width; i++ {
			s += "-"
		}
		fmt.Printf("+" + s + "+\n")
		for y := m.height - 1; y >= 0; y-- {
			fmt.Printf("|")
			for x := 0; x < m.width; x++ {
				if m.field[x][y].path {
					fmt.Printf(".")
				} else if m.field[x][y].obstacle {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("|\n")
		}
		fmt.Printf("+" + s + "+\n")
	} else {
		panic("Point out of bounds")
	}
}
