package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationCounterAddOneNormalAddOne(t *testing.T) {
	myAssert := assert.New(t)
	locationCounter := []int{0, 0, 0}
	baseString := "abcde"
	overflowFlag := locationCounterAddOne(&locationCounter, baseString)
	myAssert.False(overflowFlag)
	myAssert.Equal([]int{0, 0, 1}, locationCounter)
}

func TestLocationCounterAddOneWithOneCarry(t *testing.T) {
	myAssert := assert.New(t)
	locationCounter := []int{0, 0, 4}
	baseString := "abcde"
	overflowFlag := locationCounterAddOne(&locationCounter, baseString)
	myAssert.False(overflowFlag)
	myAssert.Equal([]int{0, 1, 0}, locationCounter)
}

func TestLocationCounterAddOneWithSomeCarry(t *testing.T) {
	myAssert := assert.New(t)
	locationCounter := []int{0, 4, 4, 4, 4}
	baseString := "abcde"
	overflowFlag := locationCounterAddOne(&locationCounter, baseString)
	myAssert.False(overflowFlag)
	myAssert.Equal([]int{1, 0, 0, 0, 0}, locationCounter)
}

func TestLocationCounterAddOneWithOverflow(t *testing.T) {
	myAssert := assert.New(t)
	locationCounter := []int{4, 4, 4, 4, 4}
	baseString := "abcde"
	overflowFlag := locationCounterAddOne(&locationCounter, baseString)
	myAssert.True(overflowFlag)
	myAssert.Equal([]int{0, 0, 0, 0, 0}, locationCounter)
}
