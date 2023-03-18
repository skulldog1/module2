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

// BEGIN (write your solution here)
func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	if len(job.Dicts) < 2 {
		job.IsFinished = true
		return job, errNotEnoughDicts
	}

	for _, dict := range job.Dicts {
		if dict == nil {
			job.IsFinished = true
			return job, errNilDict
		}

		for key, value := range dict {
			job.Merged[key] = value
		}
	}

	job.Dicts = []map[string]string{job.Merged}
	job.IsFinished = true
	return job, nil
}

// END

// Пример работы
// ExecuteMergeDictsJob(&MergeDictsJob{}) // &MergeDictsJob{IsFinished: true}, "at least 2 dictionaries are required"
// ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"},nil}}) // &MergeDictsJob{IsFinished: true, Dicts: []map[string]string{{"a": "b"},nil}}, "nil dictionary"
// ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"},{"b": "c"}}}) // &MergeDictsJob{IsFinished: true, Dicts: []map[string]string{{"a": "b", "b": "c"}}}, nil
/*
Задание: реализуй функцию ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error), которая выполняет джобу MergeDictsJob и возвращает ее.

Алгоритм обработки джобы следующий:

перебрать по порядку все словари job.Dicts и записать каждое ключ-значение в результирующую мапу job.Merged;

если в структуре job.Dicts меньше двух словарей, возвращается ошибка errNotEnoughDicts = errors.New(at least 2 dictionaries are required);

если в структуре job.Dicts встречается словарь в виде нулевого значения nil, то возвращается ошибка errNilDict = errors.New(nil dictionary);

независимо от успешного выполнения или ошибки в возвращаемой структуре MergeDictsJob, поле IsFinished должно быть заполнено как true.
*/
