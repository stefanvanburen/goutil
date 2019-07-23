workflow "build and test" {
  on = "pull_request"
  resolves = ["go test"]
}

action "go build" {
  uses = "docker://golang:1.12.7"
  runs = "go"
  args = "build ./..."
}

action "go test" {
  uses = "docker://golang:1.12.7"
  needs = ["go build"]
  runs = "go"
  args = "test -cover -race ./..."
}
