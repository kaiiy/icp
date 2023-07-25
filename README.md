# icp

Icp is a tool for copying files, incrementing a number within the filename.

## Installation

To install this tool directly from the GitHub repository, run the following command:

```sh
$ go install github.com/kaiiy/icp@latest
```

## Usage

```sh
$ icp [filename]
```

Replace [filename] with the name of your file, for example `1_report.txt`. The output will be a new file named with the next number, e.g. `2_report.txt`.