package mollymage

/*-
 * #%L
 * Codenjoy - it's a dojo-like platform from developers to developers.
 * %%
 * Copyright (C) 2021 Codenjoy
 * %%
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this program.  If not, see
 * <http://www.gnu.org/licenses/gpl-3.0.html>.
 * #L%
 */

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/codenjoyme/codenjoy-go-client/direction"
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

type Solver struct {
	actTimer uint
	isSpawn  bool
	// Decrese value of staying on the same cell ....
	visitedTimes map[engine.Point]int
	visited      []*engine.Point
	oldPotions   []*engine.Point
}

func NewSolver() *Solver {
	return &Solver{
		actTimer:     0,
		isSpawn:      true,
		visitedTimes: make(map[engine.Point]int),
		visited:      nil,
		oldPotions:   nil,
	}
}

func (solver *Solver) PushVisitedList(point *engine.Point) {
	if _, ok := solver.visitedTimes[*point]; !ok {
		solver.visitedTimes[*point] = 1
	} else {
		solver.visitedTimes[*point] += 1
		solver.visited = append(solver.visited, point)
	}
}

func (solver *Solver) PopVisitedList() {
	if len(solver.visited) > 13 {
		deletePoint := solver.visited[0]
		solver.visited = solver.visited[1:]
		delete(solver.visitedTimes, *deletePoint)
	}
}

func (solver *Solver) Answer(message string) engine.Action {
	board := newBoard(message)
	fmt.Println("Board \n" + board.String())
	action := nextAction(solver, board)
	fmt.Println("\nAnswer: " + action.String())
	fmt.Println("-------------------------------------------------------------")
	return action
}

func nextAction(solver *Solver, b *board) engine.Action {
	// Algorithm is based on graph A* algorithm.
	//
	// We don't care about other players,
	// We just try to not get on BOMBs and
	// place our bombs as often as possible and as close
	// as possible to chests
	//
	// we could actually estimate how much people in the area so
	// if there's a lot of chests but there's a lot of people there
	// we would chose a nearest chest as a target
	//
	// don't forget that chests by which side BOMBs are spawned has no value any more

	// estimationFunction(b)

	justSpawn := len(b.board.Find(elements["DEAD_HERO"])) > 0
	if justSpawn {
		solver = NewSolver()
	}

	place := b.findHero()

	solver.PopVisitedList()
	solver.PushVisitedList(place)

	fmt.Println("VISITED LIST: ", solver.visitedTimes)

	bestMove := findBestRoute(solver, b, place)

	solver.oldPotions = b.findPotions()

	if !justSpawn {
		if solver.actTimer == 0 {
			nextPlace := bestMove.Move(place)
			action := bombPlace(b, place, nextPlace, bestMove)

			if strings.Contains(action.String(), "ACT") {
				solver.actTimer = 4
			}

			return action
		} else {
			solver.actTimer -= 1
		}
	}

	return sideToAction(bestMove)
}

func bombPlace(b *board, curent, next *engine.Point, bestMove direction.Side) engine.Action {
	blastsAtCurent := b.predictBlastsAtPoint(curent)
	blastsAtNext := b.predictBlastsAtPoint(next)

	perks := b.findPerks()

	notPlace := false
	for _, blast := range blastsAtCurent {
		if isInSet(blast, perks...) {
			notPlace = true
			break
		}
	}

	notNextPlace := false
	for _, blast := range blastsAtNext {
		if isInSet(blast, perks...) {
			notNextPlace = true
			break
		}
	}

	notPlace = notPlace || !isThereEscape(b, curent, blastsAtCurent)
	notNextPlace = notNextPlace || !isThereEscape(b, next, blastsAtNext)

	if len(availableMoves(b, curent)) == 0 {
		return act
	}

	// todo: sometimes we don't have a way out do wee need to ACT on this????
	if notPlace && notNextPlace {
		return sideToAction(bestMove)
	}

	if notPlace {
		return sideToAction(bestMove) + ",ACT"
	}

	if notNextPlace {
		return "ACT," + sideToAction(bestMove)
	}

	inPlace := numberOfTreasuresInBlastRange(b, curent)
	inNextPlace := numberOfTreasuresInBlastRange(b, next)

	if inPlace > inNextPlace {
		return "ACT," + sideToAction(bestMove)
	} else {
		return sideToAction(bestMove) + ",ACT"
	}
}

func findBestRoute(solver *Solver, b *board, position *engine.Point) direction.Side {
	moves := []direction.Side{
		direction.Left,
		direction.Right,
		direction.Up,
		direction.Down,
		direction.Stop,
	}

	bestMove := direction.Stop
	bestValue := math.Inf(-1)
	for _, move := range moves {
		heroPos := move.Move(position)
		value := estimate(solver, b, heroPos)

		fmt.Println("move", move.String(), "value", value)

		if value > bestValue {
			bestValue = value
			bestMove = move
		}
	}

	fmt.Printf("Best move %s\n", bestMove)

	return bestMove
}

type Estimation func(b *board, pos *engine.Point) float64

func estimate(solver *Solver, b *board, pos *engine.Point) float64 {
	funcs := []Estimation{
		debugFunction("path/enemy", enemyPathFunction),
		debugFunction("path/ghost", ghostPathFunction),
		debugFunction("path/blast", blastFunction(solver)),
		debugFunction("path/treasure", treasurePathFunction(solver)),
		debugFunction("path/perks", perkPathFunction(solver)),
		debugFunction("path/close/ghost", closeGhostFunction),
		debugFunction("path/close/dead-ghost", closeDeadGhostFunction),
		debugFunction("path/close/enemy", closeEnemyFunction),
		debugFunction("path/escape/blasts", escapeBlastFunction(solver)),
		debugFunction("path/escape/boom", escapeBoomBlastFunction),
		debugFunction("path/repeat", repeatPathFunction(solver)),
		debugFunction("target/enemy", enemyFunction),
		debugFunction("target/treasure", treasureFunction(solver)),
		debugFunction("target/perks", perksFunction(solver)),
		debugFunction("target/ghost", avoidGhostsFunction),
		debugFunction("target/barrier", avoidBarrierFunction),
	}

	w := 0.0
	for _, f := range funcs {
		w += f(b, pos)
	}

	return w
}

const (
	WAY_ENEMY_WEIGHT = 10
	// WAY_ENEMY_WEIGHT    = 15
	WAY_GHOST_WEIGHT    = 1
	WAY_PERK_WEIGHT     = 5
	WAY_TREASURE_WEIGHT = 5

	CLOSE_ENEMY_WEIGHT = -1
	CLOSE_GHOST_WEIGHT = -80
	// CLOSE_GHOST_WEIGHT      = -50
	CLOSE_DEAD_GHOST_WEIGHT = -100

	ENEMY_WEIGHT    = 0
	TREASURE_WEIGHT = 100
	PERK_WEIGHT     = 200

	BLAST_DANGER_MAX_WEIGHT     = -50
	BLAST_DANGER_UNKNOWN_WEIGHT = -100
	BLAST_DANGER_PATH_WEIGHT    = -50

	// BLAST_DANGER_WEIGHT = 10

	ESCAPE_BLAST_WEIGHT = -100
	ESCAPE_BOOM_WEIGHT  = -200

	REPEATED_PATH_WEIGHT = -2
)

func perkPathFunction(solver *Solver) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		blastMap := predictBlasts(b, solver)
		points := b.board.Find(PERKS...)
		points = skipBlasts(blastMap, points)

		return pathFunction(points, WAY_PERK_WEIGHT)(b, pos)
	}
}

func enemyPathFunction(b *board, pos *engine.Point) float64 {
	points := b.board.Find(elements["ENEMY_HERO"])
	return pathFunction(points, WAY_ENEMY_WEIGHT)(b, pos)
}

func ghostPathFunction(b *board, pos *engine.Point) float64 {
	points := b.board.Find(elements["GHOST"])
	return pathFunction(points, WAY_GHOST_WEIGHT)(b, pos)
}

func treasurePathFunction(solver *Solver) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		points := b.findTreasureBoxes()
		blastMap := predictBlasts(b, solver)
		points = skipBlasts(blastMap, points)

		return pathFunction(points, WAY_TREASURE_WEIGHT)(b, pos)
	}
}

func pathFunction(points []*engine.Point, weight float64) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		value := 0.0
		for _, enemy := range points {
			distance := distance(pos, enemy)
			value += float64(distance)
		}

		if value == 0.0 {
			return 0.0
		}

		return (1.0 / value) * weight * 10
	}
}

func avoidGhostsFunction(b *board, pos *engine.Point) float64 {
	return goFunction(b.board.Find(elements["DEAD_GHOST"], elements["GHOST"], elements["OPENING_TREASURE_BOX"]), math.Inf(-1))(b, pos)
}

func avoidBarrierFunction(b *board, pos *engine.Point) float64 {
	return goFunction(b.findBarriers(), math.Inf(-1))(b, pos)
}

func enemyFunction(b *board, pos *engine.Point) float64 {
	return goFunction(b.board.Find(PERKS...), ENEMY_WEIGHT)(b, pos)
}

func perksFunction(solver *Solver) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		blastMap := predictBlasts(b, solver)
		points := b.board.Find(PERKS...)
		points = skipBlasts(blastMap, points)

		return goFunction(points, PERK_WEIGHT)(b, pos)
	}
}

func treasureFunction(solver *Solver) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		blastMap := predictBlasts(b, solver)
		points := b.findTreasureBoxes()
		points = skipBlasts(blastMap, points)

		return goFunction(points, TREASURE_WEIGHT)(b, pos)
	}
}

func goFunction(points []*engine.Point, weight float64) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		if isInSet(pos, points...) {
			return weight
		}

		return 0.0
	}
}

func closeEnemyFunction(b *board, pos *engine.Point) float64 {
	return goCloseFunction(b.findEnemyHeroes(), CLOSE_ENEMY_WEIGHT)(b, pos)
}

func closeGhostFunction(b *board, pos *engine.Point) float64 {
	return goCloseFunction(b.board.Find(elements["GHOST"]), CLOSE_GHOST_WEIGHT)(b, pos)
}

func closeDeadGhostFunction(b *board, pos *engine.Point) float64 {
	return goCloseFunction(b.board.Find(elements["DEAD_GHOST"], elements["OPENING_TREASURE_BOX"]), CLOSE_DEAD_GHOST_WEIGHT)(b, pos)
}

func goCloseFunction(points []*engine.Point, weight float64) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		for _, point := range points {
			distance := distance(pos, point)
			if distance == 1 {
				return weight
			}
		}

		return 0.0
	}
}

func blastFunction(solver *Solver) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		blasts := predictBlasts(b, solver)

		var closetPotion *engine.Point = nil
		minDistance := 0
		for _, blast := range blasts.blasts {
			if isInSet(pos, blast.points...) {
				distance := distance(pos, blast.potion)
				if distance < minDistance || closetPotion == nil {
					minDistance = distance
					closetPotion = blast.potion
				}
			}
		}

		if minDistance == 0 {
			return 0.0
		}

		potionTypes := []rune{
			elements["POTION_TIMER_1"],
			elements["POTION_TIMER_2"],
			elements["POTION_TIMER_3"],
			elements["POTION_TIMER_4"],
			elements["POTION_TIMER_5"],
		}

		potionType := b.getAt(closetPotion)

		danger := float64(BLAST_DANGER_UNKNOWN_WEIGHT)
		for i, pt := range potionTypes {
			if potionType == pt {
				danger = BLAST_DANGER_MAX_WEIGHT / float64(i+1)
				break
			}
		}

		return danger + (1.0/float64(minDistance))*BLAST_DANGER_PATH_WEIGHT
	}
}

func escapeBlastFunction(solver *Solver) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		blasts := predictBlasts(b, solver)
		for _, blasts := range blasts.blasts {
			if isInSet(pos, blasts.points...) {
				escape := isThereEscape(b, pos, blasts.points)
				if !escape {
					return ESCAPE_BLAST_WEIGHT
				}
			}
		}

		return 0.0
	}
}

func escapeBoomBlastFunction(b *board, pos *engine.Point) float64 {
	blasts := b.predictFutureBlasts()
	if isInSet(pos, blasts...) {
		return ESCAPE_BOOM_WEIGHT
	}

	return 0.0
}

var PERKS = []rune{
	elements["POTION_COUNT_INCREASE"],
	elements["POTION_REMOTE_CONTROL"],
	elements["POTION_IMMUNE"],
	elements["POTION_BLAST_RADIUS_INCREASE"],
	elements["POISON_THROWER"],
	elements["POTION_EXPLODER"],
}

func skipBlasts(m *BlastMap, points []*engine.Point) []*engine.Point {
	var out []*engine.Point
	for _, point := range points {
		if m.IsIn(point) {
			continue
		}

		out = append(out, point)
	}

	return out
}

// func bombPlateFunction(b *board, pos *engine.Point) float64 {
// 	return float64(numberOfTreasuresInBlastRange(b, pos)) * TREAUSURE_PLANT_WEIGHT
// }

func debugFunction(name string, f Estimation) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		r := f(b, pos)
		fmt.Printf("function name=%s value=%v\n", name, r)
		return r
	}
}

func isThereEscape(b *board, pos *engine.Point, blasts []*engine.Point) bool {
	checked := make(map[engine.Point]struct{})
	points := availableMoves(b, pos)

	for len(points) > 0 {
		point := points[0]
		if _, ok := checked[*point]; ok {
			points = points[1:]
			continue
		}

		if !isInSet(point, blasts...) {
			return true
		}

		checked[*point] = struct{}{}

		points = append(points, availableMoves(b, point)...)
	}

	return false
}

func repeatPathFunction(solver *Solver) Estimation {
	return func(b *board, pos *engine.Point) float64 {
		if count, ok := solver.visitedTimes[*pos]; ok {
			if count > 1 {
				return float64(count-1) * REPEATED_PATH_WEIGHT
			}
		}

		return 0.0
	}
}

func hiddenPotions(solver *Solver, b *board) []*engine.Point {
	ghosts := b.board.Find(elements["DEAD_GHOST"], elements["GHOST"])
	var hidden []*engine.Point
	for _, potion := range solver.oldPotions {
		if isInSet(potion, ghosts...) {
			hidden = append(hidden, potion)
		}
	}

	return hidden
}

func distance(pos *engine.Point, to *engine.Point) int {
	return absSubInt(pos.X(), to.X()) + absSubInt(pos.Y(), to.Y())
}

func isInSet(pos *engine.Point, arr ...*engine.Point) bool {
	for _, point := range arr {
		if point.Equal(pos) {
			return true
		}
	}

	return false
}

func randomMove() direction.Side {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(4)
	switch n {
	case 0:
		return direction.Up
	case 1:
		return direction.Down
	case 2:
		return direction.Left
	default:
		return direction.Right
	}
}

func sideToAction(side direction.Side) engine.Action {
	switch side {
	case direction.Down:
		return down
	case direction.Up:
		return up
	case direction.Left:
		return left
	case direction.Right:
		return right
	case direction.Stop:
		return stop
	}

	panic("unreachable")
}

func absSubInt(lhs, rhs int) int {
	if lhs > rhs {
		return lhs - rhs
	} else {
		return rhs - lhs
	}
}

type BlastMap struct {
	blasts []BlastInformation
}

func (m *BlastMap) IsIn(pos *engine.Point) bool {
	for _, blast := range m.blasts {
		if isInSet(pos, blast.points...) {
			return true
		}
	}

	return false
}

type BlastInformation struct {
	points []*engine.Point
	potion *engine.Point
}

func predictBlasts(b *board, solver *Solver) *BlastMap {
	points := b.board.Find(
		elements["POTION_TIMER_1"],
		elements["POTION_TIMER_2"],
		elements["POTION_TIMER_3"],
		elements["POTION_TIMER_4"],
		elements["POTION_TIMER_5"],
		elements["POTION_HERO"],
		elements["OTHER_POTION_HERO"],
		elements["ENEMY_POTION_HERO"],
	)

	hidden := hiddenPotions(solver, b)
	points = append(points, hidden...)

	var info []BlastInformation
	for _, potion := range points {
		barriers := b.predictBlastsAtPoint(potion)
		info = append(info, BlastInformation{barriers, potion})
	}
	return &BlastMap{info}
}

func (b *board) predictBlastsAtPoint(pt *engine.Point) []*engine.Point {
	barriers := []*engine.Point{pt}
	barriers = append(barriers, b.predictBlastsAffectionForOneSide(pt, engine.StepLeft)...)
	barriers = append(barriers, b.predictBlastsAffectionForOneSide(pt, engine.StepRight)...)
	barriers = append(barriers, b.predictBlastsAffectionForOneSide(pt, engine.StepUp)...)
	barriers = append(barriers, b.predictBlastsAffectionForOneSide(pt, engine.StepDown)...)
	return barriers
}

func (b *board) predictBlastsAffectionForOneSide(pt *engine.Point, nextStep Move) []*engine.Point {
	var barriers []*engine.Point
	barriers = appendIfMissing(barriers, b.findWalls()...)
	barriers = appendIfMissing(barriers, b.findGhosts()...)
	barriers = appendIfMissing(barriers, b.findTreasureBoxes()...)
	barriers = appendIfMissing(barriers, b.findPotions()...)
	// enemies can move away
	// barriers = appendIfMissing(barriers, b.findOtherHeroes()...)
	// barriers = appendIfMissing(barriers, b.findEnemyHeroes()...)

	var points []*engine.Point
	for i := 0; i < BLAST_RANGE; i++ {
		pt = nextStep(pt)
		if !pt.IsValid(b.board.Size()) {
			break
		}
		points = append(points, pt)
		for _, barrier := range barriers {
			if barrier.Equal(pt) {
				return points
			}
		}
	}
	return points
}

func blastRange(b *board, pt *engine.Point) []*engine.Point {
	moves := []Move{
		engine.StepLeft,
		engine.StepRight,
		engine.StepUp,
		engine.StepDown,
	}

	var points []*engine.Point
	for _, move := range moves {
		points = append(points, b.predictBlastsForOneSide(pt, move)...)
	}

	return points
}

func numberOfTreasuresInBlastRange(b *board, pt *engine.Point) int {
	treasures := b.findTreasureBoxes()
	points := blastRange(b, pt)

	i := 0
	for _, treasure := range treasures {
		if isInSet(treasure, points...) {
			i += 1
		}
	}

	return i
}

func availableMoves(b *board, pos *engine.Point) []*engine.Point {
	barriers := b.findBarriers()

	moves := []direction.Side{
		direction.Left,
		direction.Right,
		direction.Up,
		direction.Down,
	}

	var points []*engine.Point
	for _, move := range moves {
		point := move.Move(pos)
		if isInSet(point, barriers...) {
			continue
		}

		points = append(points, point)
	}

	return points
}
