package main

type runner interface {
	name() string
	run(distance int) (seconds int)
}

type tortoise struct {
	pace int
}

type hare struct {
	nap int
}

func (t tortoise) name() string {
	return "trevor"
}

func (t tortoise) run(distance int) (seconds int) {
	return distance * t.pace
}

func (h hare) name() string {
	return "harry"
}

func (h hare) run(distance int) (seconds int) {

	if distance <= 3 {
		return distance * 1
	}

	return distance*1 + h.nap
}

func race(distance int, runners ...runner) string {
	var fastestName string
	fastestTime := 100000
	for _, runner := range runners {
		time := runner.run(distance)
		// println(runner.name(), "ran", time)
		if time < fastestTime {
			fastestTime = time
			fastestName = runner.name()
		}
	}
	return fastestName
}

// func main() {
// 	distance := 50
// 	t := tortoise{pace: 2}
// 	h := hare{nap: 15}
// 	println("winner is", race(distance, []runner{t, h}...))
// }
