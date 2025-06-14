package data

import "encoding/json"

func DeserializeLLMOutput(output string) (*LogAnalysisOutput, error) {
	var logAnalysisOutput LogAnalysisOutput
	err := json.Unmarshal([]byte(output), &logAnalysisOutput)
	if err != nil {
		return nil, err
	}

	return &logAnalysisOutput, nil
}
