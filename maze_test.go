package maze

import (
	"testing"
)

const n = 100

func TestMaze_GetDirections(t *testing.T) {

	var myMaze = New(10,3)
	myMaze.SetObstacleArea(1,1, 8,1)

	dir := myMaze.GetDirections(3,0, 3,2)
	expect := []byte{'W', 'W', 'W', 'N', 'N', 'E', 'E', 'E'}
	if len(dir) != len(expect) {
		t.Errorf("GetDirections(~)=%c want %c", dir, expect)
	} else {
		for i := 0; i < len(dir); i++ {
			if dir[i] != expect[i] {
				t.Errorf("GetDirections(~)=%c want %c", dir, expect)
			}
		}
	}

	myMaze = New(5,1)
	myMaze.SetObstacle(2, 0)
	dir = myMaze.GetDirections(0,0, 4,0)
	expect = []byte{}
	if len(dir) != len(expect) {
		t.Errorf("GetDirections(~)=%c want %c", dir, expect)
	} else {
		for i := 0; i < len(dir); i++ {
			if dir[i] != expect[i] {
				t.Errorf("GetDirections(~)=%c want %c", dir, expect)
			}
		}
	}

	//dir = myMaze.GetDirections(-1, -1, 10, 10) // panic: point out of bounds

	//dir = myMaze.GetDirections(1, 1, 2, 2) // panic: point out of bounds or on obstacle
}

func TestMaze_SetObstacle(t *testing.T) {
	var myMaze = New(5, 4)
	myMaze.SetObstacle(2, 2)

	if p := myMaze.field[2][2].obstacle; !p {
		t.Errorf("(2,2) is obstacle=%v want %v", p, true)
	}
}

func TestMaze_SetObstacleArea(t *testing.T) {
	var myMaze = New(5, 4)
	myMaze.SetObstacleArea(2, 2, 4, 3)

	for x := 2; x <= 4; x++ {
		for y := 2; y <= 3; y++ {
			if p := myMaze.field[x][y].obstacle; !p {
				t.Errorf("(%v, %v) is obstacle=%v want %v", x, y, p, true)
			}
		}
	}
}

func TestMaze_SetRoad(t *testing.T) {
	var myMaze = New(5, 4)
	myMaze.SetRoad(2, 2)

	if p := myMaze.field[2][2].obstacle; p {
		t.Errorf("(2,2) is road=%v want %v", !p, true)
	}
}

func TestMaze_SetRoadArea(t *testing.T) {
	var myMaze = New(5, 4)
	myMaze.SetRoadArea(2, 2, 4, 3)

	for x := 2; x <= 4; x++ {
		for y := 2; y <= 3; y++ {
			if p := myMaze.field[x][y].obstacle; p {
				t.Errorf("(%v, %v) is obstacle=%v want %v", x, y, !p, true)
			}
		}
	}
}
