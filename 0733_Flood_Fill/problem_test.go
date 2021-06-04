package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type FloodFiller struct {
	image        [][]int
	sizeX, sizeY int
	oldColor     int
	newColor     int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	newFloodFiller(image, image[sr][sc], newColor).rfloodFill(sr, sc)
	return image
}

func newFloodFiller(image [][]int, oldColor, newColor int) *FloodFiller {
	return &FloodFiller{
		image: image, sizeX: len(image), sizeY: len(image[0]),
		oldColor: oldColor, newColor: newColor}
}

func (ff *FloodFiller) rfloodFill(x, y int) {
	if 0 <= x && x < ff.sizeX && 0 <= y && y < ff.sizeY && ff.image[x][y] == ff.oldColor {
		ff.image[x][y] = ff.newColor
		ff.rfloodFill(x-1, y)
		ff.rfloodFill(x+1, y)
		ff.rfloodFill(x, y-1)
		ff.rfloodFill(x, y+1)
	}
}

func TestFloodFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1}}

	expected := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1}}

	floodFill(image, 1, 1, 2)
	if !reflect.DeepEqual(image, expected) {
		t.Errorf("\noutput:\n%v\nexpected:\n%v", toString(image), toString(expected))
	}

	image = [][]int{
		{0, 0, 0},
		{0, 0, 0}}

	expected = [][]int{
		{2, 2, 2},
		{2, 2, 2}}

	floodFill(image, 0, 0, 2)
	if !reflect.DeepEqual(image, expected) {
		t.Errorf("\noutput:\n%v\nexpected:\n%v", toString(image), toString(expected))
	}
}

func toString(array [][]int) string {
	result := ""
	for _, r := range array {
		result += fmt.Sprintln(r)
	}
	return result
}
