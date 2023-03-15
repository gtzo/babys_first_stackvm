package main

type Stack interface {
  push(int64)
  pop() int64
}

type StackMachine struct {
  size int64
  stack []int64
}
type Instruction string

type InstructionFrame struct {
  instruction Instruction
  value int64
}

type Program struct {
  frames []InstructionFrame
}

const (
  // value field is ignored in all instructions except LOAD
  ADD = "ADD"
  SUB = "SUB"
  MUL = "MUL"
  DIV = "DIV"
  LOAD = "LOAD"
)
