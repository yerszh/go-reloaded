#  go-reloaded


This tool is designed to perform various text modifications on a given input text file and save the modified text to an output file. It supports the following modifications:

+ Hexadecimal to Decimal Conversion: Replace every instance of (hex) with the decimal equivalent of the preceding hexadecimal number. For example, "1E (hex) files were added" becomes "30 files were added."

+ Binary to Decimal Conversion: Replace every instance of (bin) with the decimal equivalent of the preceding binary number. For example, "It has been 10 (bin) years" becomes "It has been 2 years."

+ Uppercase Conversion: Convert the word before (up) to uppercase. For example, "Ready, set, go (up)!" becomes "Ready, set, GO!"

+ Lowercase Conversion: Convert the word before (low) to lowercase. For example, "I should stop SHOUTING (low)" becomes "I should stop shouting."

+ Capitalization: Capitalize the word before (cap). For example, "Welcome to the Brooklyn bridge (cap)" becomes "Welcome to the Brooklyn Bridge."

+ Modifier with Number: For (low), (up), and (cap), if followed by a number, apply the modification to the specified number of words. For example, "This is so exciting (up, 2)" becomes "This is SO EXCITING."

+ Punctuation Formatting: Ensure proper spacing around punctuation marks (., ,, !, ?, :, and ;). For example, "I was sitting over there ,and then BAMM !!" becomes "I was sitting over there, and then BAMM!!". Handle exceptions for groups of punctuation like "...", "!?", formatting them as "I was thinking... You were right."

+ A/An Conversion: Replace "a" with "an" if the next word begins with a vowel (a, e, i, o, u) or "h". For example, "There it was. A amazing rock!" becomes "There it was. An amazing rock!"

+ Single Quotation Marks: Ensure proper placement of single quotation marks (' ') around words or phrases. For example, "I am exactly how they describe me: ' awesome '" becomes "I am exactly how they describe me: 'awesome'".

+ Multiple Words in Quotation Marks: If there are multiple words between single quotation marks, place the marks next to the corresponding words. For example, "As Elton John said: ' I am the most well-known homosexual in the world '" becomes "As Elton John said: 'I am the most well-known homosexual in the world'".

## Run Locally

Clone the project

```
git clone git@git.01.alem.school:yzhumyro/go-reloaded.git
```

Go to the project directory

```
cd go-reloaded/
```

**Examples how to run the program:**

_make sure you are in project directory_

```
go run sample.txt result.txt
```

Usage:

```
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
$ go run . sample.txt result.txt
$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
$ cat sample.txt
There is no greater agony than bearing a untold story inside you.
$ go run . sample.txt result.txt
$ cat result.txt
There is no greater agony than bearing an untold story inside you.
$ cat sample.txt
Punctuation tests are ... kinda boring ,don't you think !?
$ go run . sample.txt result.txt
$ cat result.txt
Punctuation tests are... kinda boring, don't you think!?
```
