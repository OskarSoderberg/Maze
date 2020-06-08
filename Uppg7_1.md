# package maze

Package maze lets you build mazes and navigate between points.

Maze uses Breadth-first search for path-finding, which guarantees finding
the shortest path between two points. The package also has tools for easy
manipulation of the maze's terrain.

The pathfinding algorithm only allows non-diagonal movements. Each point
is marked as either road (walkable) or obstacle (not walkable). Due to
utilization of Breadth-first search, pathfinding in maze will be costly
in open environments with long distances between points.


## Index
- type Maze
    - func New(width int, height int) \*Maze
    - func (m \*Maze) SetObstacle()
    - func (m \*Maze) SetObstacleArea(fromX int, fromY int, toX int, toY int)
    - func (m \*Maze) SetRoad()
    - func (m \*Maze) SetRoadArea(fromX int, fromY int, toX int, toY int)
    - func (m \*Maze) GetDirections(fromX int, fromY int, toX int, toY int) []byte
    - func (m \*Maze) Print()
    - func (m \*Maze) PrintDirections()

## type Maze
Maze represents a maze with fixed width and height. Each point is either an
obstacle or a road.
### func New(width int, height int) \*Maze
New creates a new maze with chosen width and height. By default all points
are roads.
### func (m \*Maze) SetObstacle(x int, y int)
SetObstacle sets the point (x, y) to an obstacle.
### func (m \*Maze) SetObstacleArea(fromX int, fromY int, toX int, toY int)
SetObstacleArea sets the rectangle with corners (fromX, fromY), (toX, toY)
to obstacles.
### func (m \*Maze) SetRoad()
SetRoad sets the point (x, y) to a road.
### func (m \*Maze) SetRoadArea(fromX int, fromY int, toX int, toY int)
SetRoadArea sets the rectangle with corners (fromX, fromY), (toX, toY)
to roads.
### func (m \*Maze) GetDirections(fromX int, fromY int, toX int, toY int) []byte
GetDirections returns an array of directions with each element being of the form
'N', 'E', 'W' or 'S', representing the four cardinal directions. Each element in
the array represents one step in the corresponding direction. If no path solves
the maze an empty list is returned.
### func (m \*Maze) Print()
Print draws the maze with '#' marking an obstacle and ' ' marking a road.
### func (m \*Maze) PrintDirections(fromX int, fromY int, toX int, toY int)
PrintDirections draws a map with directions on it. '#' marks an obstacle,
' ' marks a road and '.' marks the recommended path. If no path solves
the maze no path will be printed.
