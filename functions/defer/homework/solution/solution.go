package solution

import "errors"

// MergeDictsJob is a job to merge dictionaries into a single dictionary.
type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	job.Merged = make(map[string]string)
	defer func() {
		job.IsFinished = true
	}()
	for _, dict := range job.Dicts {
		if dict == nil {
			return job, errNilDict
		}
		for key, value := range dict {
			job.Merged[key] = value
		}
	}

	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}

	job.Dicts = []map[string]string{job.Merged}
	return job, nil
}
