package lib

import (
	"15/stack"
	"math"
)

// https://en.wikipedia.org/wiki/Iterative_deepening_A*
// path              current search path (acts like a stack)
// node              current node (last node in current path)
// cost              the cost to reach current node
// f                 estimated cost of the cheapest path (root..node..goal)
// h(node)           estimated cost of the cheapest path (node..goal)
// cost(node, succ)  step cost function
// is_goal(node)     goal test
// successors(node)  node expanding function, expand nodes ordered by g + h(node)
// ida_star(root)    return either NOT_FOUND or a pair with the best path and its cost

// procedure ida_star(root)
//     bound := h(root)
//     path := [root]
//     loop
//         t := search(path, 0, bound)
//         if t = FOUND then return (path, bound)
//         if t = ∞ then return NOT_FOUND
//         bound := t
//     end loop
// end procedure

// function search(path, g, bound)
//     node := path.last
//     f := g + h(node)
//     if f > bound then return f
//     if is_goal(node) then return FOUND
//     min := ∞
//     for succ in successors(node) do
//         if succ not in path then
//             path.push(succ)
//             t := search(path, g + cost(node, succ), bound)
//             if t = FOUND then return FOUND
//             if t < min then min := t
//             path.pop()
//         end if
//     end for
//     return min
// end function

// Iterative deepening A*
func Solver(b *Board) (stack.Stack[*Board], int) {
	bound := b.NeededMoves()
	path := stack.NewStack[*Board]()
	path.Push(b)

	for {
		cost, path, found := search(path, 0, bound)

		if found {
			return path, bound
		}

		if cost == math.MaxInt {
			return path, -1
		}

		bound = cost

	}
}

func search(path stack.Stack[*Board], cost int, bound int) (int, stack.Stack[*Board], bool) {
	last := path.Peek()
	estimatedCost := cost + last.NeededMoves()

	if last.SolvedFast() {
		return cost, path, true
	}

	if estimatedCost > bound {
		return estimatedCost, path, false
	}

	minCost := math.MaxInt

	possibleDirections := last.PossibleDirections()

	for _, direction := range possibleDirections {
		step := last.Copy()
		step.Move(direction)

		// TODO: check not only last but every node in the path
		if *step == *last {
			continue
		}

		path.Push(step)

		cost, path, found := search(path, cost+1, bound)

		if found {
			return cost, path, true
		}

		if cost < minCost {
			minCost = cost
		}

		path.Pop()
	}

	return minCost, path, false
}
