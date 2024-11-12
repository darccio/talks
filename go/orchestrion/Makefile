.PHONY: clean run run-uninstrumented setup

clean:
	rm orchestrion.tool.go golab2024
	go mod tidy

setup:
	orchestrion pin

run:
	orchestrion go build .
	DD_APPSEC_ENABLED=true DD_PROFILING_ENABLED=true DD_TRACE_AGENT_PORT=9126 ./golab2024 > /tmp/golab2024.log

run-uninstrumented:
	go run .

slides:
	marp --server --html .