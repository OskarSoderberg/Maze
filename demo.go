package main

import (
	"fmt"
	"maze"
)

func main() {
	var myMaze = maze.New(11, 9)
	myMaze.SetObstacleArea(2, 1, 7, 1)
	myMaze.SetObstacleArea(7, 2, 8, 2)
	myMaze.SetObstacleArea(1, 3, 1, 4)
	myMaze.SetObstacleArea(3, 3, 5, 3)
	myMaze.SetObstacleArea(7, 4, 10, 4)
	myMaze.SetObstacle(4,4)
	myMaze.SetObstacle(2,5)
	myMaze.SetObstacle(6,5)
	myMaze.SetObstacle(0,6)
	myMaze.SetObstacle(4,6)
	myMaze.SetObstacle(7,8)
	myMaze.SetObstacle(9,7)
	myMaze.SetObstacleArea(2,7,4,7)
	myMaze.SetObstacleArea(8,6,9,6)

	fmt.Printf("Maze: \n")
	myMaze.Print()

	fmt.Printf("Path from (5,0) to (5,8): \n")
	fmt.Printf("%c\n", myMaze.GetDirections(5,0,5,8))

	myMaze.PrintDirections(5,0,5,8)

	fmt.Printf("Path from (0,3) to (10,5): \n")
	fmt.Printf("%c\n", myMaze.GetDirections(0,3,10,5))

	myMaze.PrintDirections(0,3,10,5)

}
