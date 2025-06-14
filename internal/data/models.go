package data

import (
	"context"
	"database/sql"
	"time"
)

type LogAnalysisOutput struct {
	Analysis            string    `json:"analysis"`
	Cause               string    `json:"cause"`
	Severity            string    `json:"severity"`
	TimeOfOccurrence    time.Time `json:"time_of_occurrence"`
	StacktraceInsight   string    `json:"stacktrace_insight"`
	File                string    `json:"file"`
	LineNumber          string    `json:"line_number"`
	Summary             string    `json:"summary"`
	ComprehensiveDetail string    `json:"comprehensive_detail"`
	SuggestedWayToFix   string    `json:"suggested_way_to_fix"`
}

type LogAnalysisDB struct {
	DB *sql.DB
}

func (m LogAnalysisDB) ExtractAndInsertLogAnalysis(output string) error {
	deserializeLLMOutput, err := DeserializeLLMOutput(output)
	if err != nil {
		return err
	}

	err = m.Insert(deserializeLLMOutput)
	if err != nil {
		return err
	}

	return nil
}

func (m LogAnalysisDB) Insert(output *LogAnalysisOutput) error {
	query := `INSERT INTO log_analysis 
    			(analysis, cause, severity, time_of_occurrence, stack_trace, file, line_number, summary, comprehensive_details, suggested_way_to_fix, created_at)
    			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{
		output.Analysis,
		output.Cause,
		output.Severity,
		output.TimeOfOccurrence,
		output.StacktraceInsight,
		output.File,
		output.LineNumber,
		output.Summary,
		output.ComprehensiveDetail,
		output.SuggestedWayToFix,
		time.Now(),
	}

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
