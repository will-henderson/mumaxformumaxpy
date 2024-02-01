package engine

import (
	"github.com/mumax/3/cuda"
)

func init() {
	DeclFunc("Reset", Reset, "Resets the solver and the geometry to mumax initialisation")
	DeclFunc("ResetGeometry", ResetGeometry, "Forgets all the defined regions")
	DeclFunc("ResetSolver", ResetSolver, "Sets solver parameters to their initial values")
}

func Reset() {
	ResetSolver()
	ResetGeometry()
}

func ResetSolver() {
	MaxErr = 1e-5
	Headroom = 0.8
	Dt_si = 1e-15
	FixDt = 0
	Time = 0
}

func ResetGeometry() {
	if !(globalmesh_.Size() == [3]int{0, 0, 0}) {
		ResetRegions()
		ResetGeom()
		ResetCentering()
	}
}

func ResetRegions() {
	regions.hist = nil
	regions.gpuCache.Free()
	mesh := regions.Mesh()
	regions.gpuCache = cuda.NewBytes(mesh.NCell())
	DefRegion(0, universe)
}

func ResetCentering() {
	postStep = nil
	lastShift = 0
	lastT = 0
	lastV = 0
	TotalShift = 0
	TotalYShift = 0
	Time = 0
}

func ResetGeom() {
	SetGeom(universe)
}
