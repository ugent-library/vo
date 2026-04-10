package vo

import "regexp"

var (
	RuleFuzzyDate = "fuzzy_date"

	MessageFuzzyDate = "must be a date (YYYY, YYYY-MM, or YYYY-MM-DD)"

	reFuzzyDate = regexp.MustCompile(`^\d{4}(-\d{2}(-\d{2})?)?$`)
)

// FuzzyDate checks if the given string is a valid fuzzy date.
// Accepted formats: "2024", "2024-03", "2024-03-15".
//
//	err := vo.FuzzyDate("date", "2024-03")
func FuzzyDate(key, val string) *Error {
	if !reFuzzyDate.MatchString(val) {
		return NewError(key, RuleFuzzyDate).WithMessage(MessageFuzzyDate)
	}
	return nil
}
