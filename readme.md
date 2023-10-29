# go-reloaded

The Text Processor is a Go-based project aimed at automating text editing, completion, and auto-correction tasks. The application reads an input text file and applies a series of predefined text transformations to produce an edited output text file.

## Modifications
The following modifications are supported by the tool:

- *(hex)*: Replace hexadecimal numbers with their decimal equivalents.
- *(bin)*: Replace binary numbers with their decimal equivalents.
- *(up)*: Convert words to uppercase.
- *(low*): Convert words to lowercase.
- *(cap)*: Capitalize words.
- *(modifier, number)*: Modifies the specified number of words before the modifier.


## Run Locally

* Clone the project

```bash
  git clone git@git.01.alem.school:yzhumyro/go-reloaded.git
```

* Go to the project directory

```bash
  cd go-reloaded
```

* Prepare example

* Start project

```bash
  go run . sample.txt result.txt
```

## Author

- [@yzhumyro](https://01.alem.school/git/yzhumyro)