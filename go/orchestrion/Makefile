.PHONY: clean run run-uninstrumented setup

clean:
	rm orchestrion.tool.go
	go mod tidy

setup:
	orchestrion pin

run:
	DD_PROFILING_ENABLED=true DD_TRACE_AGENT_PORT=9126 orchestrion go run .

run-uninstrumented:
	go run .
