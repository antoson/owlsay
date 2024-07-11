default: test-all
alias b := build

fortune := `fortune -s`

build:
    go build -o owlsay owlsay.go
@test-args:
    echo args:
    ./owlsay '{{fortune}}'
@test-stdin:
    echo stdin:
    echo '{{fortune}}' | ./owlsay
test-all: build && test-args test-stdin

