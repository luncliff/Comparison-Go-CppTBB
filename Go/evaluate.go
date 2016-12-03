package main

import "obst"

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

func EvaluateSeq(cfg *Config, tree *obst.Tree) {
	N := tree.Size() // cfg.N

	for i := N; i >= 0; i-- {
		for j := i; j <= N; j++ {
			root, cost := tree.Calculate(i, j)
			tree.root[i][j] = root
			tree.cost[i][j] = cost
		}
	}
	return
}

// Dependency ...
//  	Data set for chunk calculation & synchronization
type Dependency struct {
	PreSet  [2]chan int
	PostSet [2]chan int
}

// Wait ...
func (rcv *Dependency) Wait() {
	if rcv.PreSet[0] != nil {
		<-rcv.PreSet[0]
	}
	if rcv.PreSet[1] != nil {
		<-rcv.PreSet[1]
	}
	return
}

// Notify ...
func (rcv *Dependency) Notify() {
	if rcv.PostSet[0] != nil {
		rcv.PostSet[0] <- 1
	}
	if rcv.PostSet[1] != nil {
		rcv.PostSet[1] <- 1
	}
}

func Chunk(tree *obst.Tree, i int, j int, width int, dep *Dependency) {
	dep.Wait()

	for row := i + width; row >= 0; row-- {
		for col := j; col <= j+width; col++ {
			root, cost := tree.Calculate(row, col)
			tree.root[row][col] = root
			tree.cost[row][col] = cost
		}
	}

	dep.Notify()
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// EvaluatePar ...
//  	Parallel Chunk processing
func EvaluatePar(cfg *Config, tree *obst.Tree, shd *Shared) {
	N := cfg.N
	VP := cfg.VP

	for i, x := 1, 0; x < VP; x++ {
		for j, y := 0, 0; y < VP; y++ {

			dep := new(Dependency)
			dep.PreSet[0] = nil
			dep.PreSet[1] = nil
			dep.PostSet[0] = nil
			dep.PostSet[1] = nil

			go Chunk(tree, i, j, N/VP, dep)

			j += VP
		}
		i += VP
	}
	<-shd.finish
	return
}
