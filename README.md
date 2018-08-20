# hlt

`hlt` is tool to highlight texts or lines matched a given pattern in files.

## Demo

![demo](resouces/demo002.gif)

## Usage

### highlight a line

```bash
# color charactors in lines including 'text' in file.txt
$ hlt line text file.txt
$ cat file.txt | hlt line text

# color charactors in lines red
$ cat file.txt | hlt line -c red text

# color background of lines blue
$ cat file.txt | hlt line -b blue text
```

### highlight a text

```bash
# color 'text' in file.txt
$ hlt word text file.txt
$ cat file.txt | hlt word text

# color 'text' red
$ cat file.txt | hlt word -c red text

# color background of `text` blue
$ cat file.txt | hlt word -b blue text
```

### Options

Settable color to options(`-b`, `-c`) of highlight commands are

- none
- blue
- green
- orange
- pink
- purple
- red
- yellow
- 0 ~ 255

## Install

```bash
# Instatll dep command
$ go get -u github.com/golang/dep/cmd/dep
# Install this repository
$ go get -u -d github.com/x-color/hlt
$ cd $GOPATH/src/github.com/x-color/hlt
# Install depended packages
$ dep ensure
# Build hlt command
$ go build ./cmd/hlt/
```

## License

MIT License
