package learningGame

type solution func(int) int

type result struct {
	command       int
	success       bool
	solution      solution
	solutionIndex int
}

type Machine struct {
	commandResults map[int][]result
	lastResult     *result
	solutions      []solution
}

func increment(x int) int {
	return x + 1
}

func multiplyBy0(x int) int {
	return x * 0
}

func divideBy2(x int) int {
	return x / 2
}

func multiplyBy100(x int) int {
	return x * 100
}

func modulo2(x int) int {
	return x % 2
}

func NewMachine() Machine {
	return Machine{
		commandResults: make(map[int][]result),
		lastResult:     nil,
		solutions: []solution{
			increment,
			multiplyBy0,
			divideBy2,
			multiplyBy100,
			modulo2,
		},
	}
}

func (m *Machine) Command(cmd, num int) int {
	commandResult, ok := m.commandResults[cmd]
	if ok {
		lastResult := commandResult[len(commandResult)-1]
		if lastResult.success {
			m.lastResult = &result{
				command:       lastResult.command,
				success:       false,
				solution:      lastResult.solution,
				solutionIndex: lastResult.solutionIndex,
			}

			return m.lastResult.solution(num)
		}

		sol := m.solutions[lastResult.solutionIndex+1]
		m.lastResult = &result{
			command:       cmd,
			success:       false,
			solution:      sol,
			solutionIndex: lastResult.solutionIndex + 1,
		}

		return m.lastResult.solution(num)
	}

	m.commandResults[cmd] = make([]result, 0)
	sol := m.solutions[0]
	m.lastResult = &result{
		command:       cmd,
		success:       false,
		solution:      sol,
		solutionIndex: 0,
	}

	return m.lastResult.solution(num)
}

func (m *Machine) Response(res bool) {
	m.lastResult.success = res
	m.commandResults[m.lastResult.command] = append(m.commandResults[m.lastResult.command], *m.lastResult)
}
