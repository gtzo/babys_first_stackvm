package main

import "fmt"

/*
A comically simple stack VM implementation.
The VM only supports basic arithmetic.
*/

func (s *StackMachine) push(num int64) {
  s.stack = append(s.stack, num)
  s.size++
}

func (s *StackMachine) pop() int64 {
  ret := s.stack[len(s.stack)-1]
  s.stack = s.stack[:len(s.stack)-1]
  s.size--
  return ret
}

func (s StackMachine) show() {
  fmt.Println(s.stack)
}

func (s *StackMachine) handleInstructionFrame(instr InstructionFrame) {
  switch instr.instruction {
  case LOAD:
    s.push(instr.value)
  case ADD:
    a := s.pop()
    b := s.pop()
    s.push(a + b)
  }
}

func (s *StackMachine) runProgram(program Program) {
  for _, frame := range program.frames {
    s.handleInstructionFrame(frame)
    s.show()
  }
}

// Do some cutting-edge addition
func main() {
  s := StackMachine{}
  a := InstructionFrame{LOAD,10}
  b := InstructionFrame{LOAD,10}
  c := InstructionFrame{ADD,0}
  frames := []InstructionFrame{a,b,c}
  code := Program{frames}
  s.runProgram(code)
}
