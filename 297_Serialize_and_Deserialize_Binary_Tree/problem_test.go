package leetcode

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	deserializationInput []string
}

func Constructor() Codec {
	return Codec{deserializationInput: nil}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "n"
	}

	result := strconv.Itoa(root.Val)
	result += "," + c.serialize(root.Left)
	result += "," + c.serialize(root.Right)
	return result
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	c.deserializationInput = strings.Split(data, ",")
	return c.rdeserialize()
}

func (c *Codec) rdeserialize() *TreeNode {
	input := c.deserializationInput[0]
	c.deserializationInput = c.deserializationInput[1:]

	if input == "n" {
		return nil
	}

	value, _ := strconv.Atoi(input)
	return &TreeNode{Val: value, Left: c.rdeserialize(), Right: c.rdeserialize()}
}

func TestSerializeDeserialize(t *testing.T) {
	c := Constructor()
	tree := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
			Right: &TreeNode{Val: 5, Left: nil, Right: nil}}}

	serializationOutput := c.serialize(tree)
	expectedSerialization := "1,2,n,n,3,4,n,n,5,n,n"
	if serializationOutput != expectedSerialization {
		t.Errorf("\noutput:   %v\nexpected: %v", serializationOutput, expectedSerialization)
	}

	deserializationOutput := c.deserialize(serializationOutput)
	expectedDeserialization := tree
	if !reflect.DeepEqual(deserializationOutput, expectedDeserialization) {
		t.Errorf("\noutput:   %v\nexpected: %v",
			c.serialize(deserializationOutput), c.serialize(expectedDeserialization))
	}

	tree = nil
	serializationOutput = c.serialize(tree)
	expectedSerialization = "n"
	if serializationOutput != expectedSerialization {
		t.Errorf("\noutput:   %v\nexpected: %v", serializationOutput, expectedSerialization)
	}

	deserializationOutput = c.deserialize(serializationOutput)
	expectedDeserialization = tree
	if !reflect.DeepEqual(deserializationOutput, expectedDeserialization) {
		t.Errorf("\noutput:   %v\nexpected: %v",
			c.serialize(deserializationOutput), c.serialize(expectedDeserialization))
	}

	tree = &TreeNode{Val: 1, Left: nil, Right: nil}
	serializationOutput = c.serialize(tree)
	expectedSerialization = "1,n,n"
	if serializationOutput != expectedSerialization {
		t.Errorf("\noutput:   %v\nexpected: %v", serializationOutput, expectedSerialization)
	}

	deserializationOutput = c.deserialize(serializationOutput)
	expectedDeserialization = tree
	if !reflect.DeepEqual(deserializationOutput, expectedDeserialization) {
		t.Errorf("\noutput:   %v\nexpected: %v",
			c.serialize(deserializationOutput), c.serialize(expectedDeserialization))
	}

	tree = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}
	serializationOutput = c.serialize(tree)
	expectedSerialization = "1,2,n,n,n"
	if serializationOutput != expectedSerialization {
		t.Errorf("\noutput:   %v\nexpected: %v", serializationOutput, expectedSerialization)
	}

	deserializationOutput = c.deserialize(serializationOutput)
	expectedDeserialization = tree
	if !reflect.DeepEqual(deserializationOutput, expectedDeserialization) {
		t.Errorf("\noutput:   %v\nexpected: %v",
			c.serialize(deserializationOutput), c.serialize(expectedDeserialization))
	}
}
