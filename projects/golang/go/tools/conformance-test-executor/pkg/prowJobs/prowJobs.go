package prowJobs

import "time"

const defaultTimeout = "6h"
const defaultMemoryRequest = "8Gi"
const defaultCpuRequest = "2"

type ProwJobOptions struct {
	RuntimeImage                string
	Timeout                     string
	CpuRequest                  string
	MemoryRequest               string
	PreExecuteCommands          []string
	PostExecuteCommands         []string
	JobIdentifier               string
}

func (g *ProwJobOptions) setDefaults() {
	if g.Timeout == "" {
		g.Timeout = defaultTimeout
	}

	if g.MemoryRequest == "" {
		g.MemoryRequest = defaultMemoryRequest
	}

	if g.CpuRequest == "" {
		g.CpuRequest = defaultCpuRequest
	}
}

func (g *ProwJobOptions) TemplateValues() map[string]interface{}{
	templateValues := make(map[string]interface{})
	templateValues["timeout"] = g.Timeout
	templateValues["memoryRequest"] = g.MemoryRequest
	templateValues["cpuRequest"] = g.CpuRequest
	templateValues["runtimeImage"] = g.RuntimeImage
	templateValues["jobId"] = g.JobIdentifier
	return templateValues
}

func ProwJobStartTime() string {
	return time.Now().Format("2006-01-02T15:04:05Z")
}