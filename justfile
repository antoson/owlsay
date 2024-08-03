default: test-all
alias b := build

# fortune := `fortune -s`
fortune := `cat wisdom`

build:
    go build -o owlsay owlsay.go
@fortune:
    echo fortune:
    echo '{{fortune}}'
    od -c wisdom
@test-args: build
    echo args:
    ./owlsay '{{fortune}}'
@test-stdin: build
    echo stdin:
    echo '{{fortune}}' | ./owlsay
@test-ab: build fortune
    echo cowsay:
    echo '{{fortune}}' | cowsay
    echo owlsay:
    echo '{{fortune}}' | ./owlsay
test-all: build && fortune test-args test-stdin

