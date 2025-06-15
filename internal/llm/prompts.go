package llm

import "fmt"

var Prompt = func(msg string) string {
	return fmt.Sprintf(`
You are an expert in debugging Go logs. Analyze the following log:

%s

Return **only** a well-formatted JSON object with the following fields:

- analysis: A brief explanation of what the log indicates.
- cause: The most likely root cause of the error.
- severity: Level of severity (e.g., low, moderate, high, critical).
- time_of_occurrence: The timestamp of the error from the log.
- stacktrace_insight: Key insight from the stack trace that reveals where or why the error happened.
- file: Identify the source file that contains the code triggering the error, not just where it was logged in the stack trace, but the actual file responsible for the failure.
- line_number: The specific line number of origin if identifiable.
- summary: A short summary of the error.
- comprehensive_detail: A concise explanation (1–3 sentences) of how this error affects the system — both technically (code behavior) and from the user’s perspective (UX impact).
- suggested_way_to_fix: Recommended steps to resolve or mitigate the issue.

Only output valid JSON. Do not include any additional explanation or formatting outside the JSON.
`, msg)
}
