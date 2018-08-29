# hlt

`hlt` is tool to highlight texts or lines matched a given pattern in files.

## Demo

### simple highlight

![demo02](resouces/demo002.gif)

### formatting

![demo03](resouces/demo003.gif)

### highlight lines N to M

![demo04](resouces/demo004.gif)

## Usage

### highlight

#### Basic usage

```bash
# highlight charactors in lines including 'text' in file.txt
$ hlt line text file.txt
$ cat file.txt | hlt line text

# highlight 'text' in file.txt
$ hlt word text file.txt
$ cat file.txt | hlt word text

# highlight line 2 to 10 and 20 to the end in file.txt
$ hlt linen 2~10,20~ file.txt
$ cat file.txt | hlt linen 2~10
```

#### Changing highlight color

```bash
# highlight charactors in lines red
$ cat file.txt | hlt line -c red text

# highlight background of lines blue
$ cat file.txt | hlt line -b blue text
```

#### Formatting

```bash
# bold format
$ cat file.txt | hlt line -B text

# italic format
$ cat file.txt | hlt line -I text
```

### Options

#### Color Control

| option | description |
|:-|:-|
| --background, -b | color background of matching lines |
| --charactor, -c | color charactors of matching lines |

Settable colors are

- none
- black
- blue
- cyan
- green
- magenta
- red
- yellow
- 0 ~ 255

#### Format Control

| option | description |
|:-|:-|
| --bold, -B | format matching texts bold |
| --hide, -H | hide matching texts |
| --italic, -I | format matching texts italic |
| --strikethrough, -S | strikethrough matching texts |
| --underline, -U | underline matching texts |

#### Context Control

This options can be set to `line` command only.

| option | description |
|:-|:-|
| --after | highlight num lines after matching lines |
| --before | highlight num lines before matching lines |

## Customize default highlight style

`hlt` consults `$HOME/.hlt/config.yaml`.

#### Colors

You can change highlight colors, `background` is none, `charactor` is red on by default.

```yaml
background: none
charactor: green
```

#### Formats

You can change highlight formats, these formats are `false` on by default. 

```yaml
bold: true
hide: false
italic: false
strikethrough: false
underline: false
```

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
