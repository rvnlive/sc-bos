package statuspb

import (
	"google.golang.org/protobuf/proto"

	"github.com/vanti-dev/sc-bos/pkg/gen"
)

type ProblemMerger struct {
	BiggestProblem    *gen.StatusLog_Problem
	MostRecentNominal *gen.StatusLog_Problem

	problems []*gen.StatusLog_Problem
}

func (pm *ProblemMerger) Build() *gen.StatusLog {
	out := &gen.StatusLog{}
	if pm.BiggestProblem == nil {
		// no problems
		out.Level = gen.StatusLog_NOMINAL
		if pm.MostRecentNominal != nil {
			out.Level = pm.MostRecentNominal.Level
			out.Description = pm.MostRecentNominal.Description
			out.RecordTime = pm.MostRecentNominal.RecordTime
		}
	} else {
		out.Level = pm.BiggestProblem.Level
		out.Description = pm.BiggestProblem.Description
		out.RecordTime = pm.BiggestProblem.RecordTime

		if pm.BiggestProblem.Level == gen.StatusLog_OFFLINE && pm.MostRecentNominal != nil {
			out.Level = gen.StatusLog_REDUCED_FUNCTION
		}
	}
	out.Problems = pm.problems // this reuses a reference, but for our use I'm sure it's fine
	return out
}

func (pm *ProblemMerger) AddStatusLog(sl *gen.StatusLog) {
	pm.AddProblems(sl.Problems)
	if sl.Level == gen.StatusLog_NOMINAL {
		switch {
		case pm.MostRecentNominal == nil:
			pm.MostRecentNominal = &gen.StatusLog_Problem{
				Level:       gen.StatusLog_NOMINAL,
				Description: sl.Description,
				RecordTime:  sl.RecordTime,
			}
		case sl.RecordTime == nil: // do nothing
		case pm.MostRecentNominal.RecordTime == nil:
			pm.MostRecentNominal.RecordTime = sl.RecordTime
			pm.MostRecentNominal.Description = sl.Description
		case sl.RecordTime.AsTime().After(pm.MostRecentNominal.RecordTime.AsTime()):
			pm.MostRecentNominal.RecordTime = sl.RecordTime
			pm.MostRecentNominal.Description = sl.Description
		}
	}
}

func (pm *ProblemMerger) AddProblemMessages(problems []proto.Message) {
	for _, problem := range problems {
		pm.AddProblem(problem.(*gen.StatusLog_Problem))
	}
}

func (pm *ProblemMerger) AddProblems(problems []*gen.StatusLog_Problem) {
	for _, problem := range problems {
		pm.AddProblem(problem)
	}
}

func (pm *ProblemMerger) AddProblem(problem *gen.StatusLog_Problem) {
	if problem.Level == gen.StatusLog_NOMINAL {
		switch {
		case pm.MostRecentNominal == nil:
			pm.MostRecentNominal = problem
		case pm.MostRecentNominal.RecordTime == nil:
			pm.MostRecentNominal = problem
		case problem.RecordTime == nil:
		case problem.RecordTime.AsTime().After(pm.MostRecentNominal.RecordTime.AsTime()):
			pm.MostRecentNominal = problem
		}
		return
	}
	pm.problems = append(pm.problems, problem)
	if pm.BiggestProblem == nil || problem.Level > pm.BiggestProblem.Level {
		pm.BiggestProblem = problem
	}
	if pm.BiggestProblem.Level == problem.Level && problem.RecordTime != nil {
		// make sure we're reporting the earliest problem at this level
		switch {
		case pm.BiggestProblem.RecordTime == nil:
			pm.BiggestProblem = problem
		case problem.RecordTime.AsTime().Before(pm.BiggestProblem.RecordTime.AsTime()):
			pm.BiggestProblem = problem
		}
	}
}
