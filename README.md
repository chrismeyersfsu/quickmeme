## Quickstart

```
cd src && go build -o quickmeme
GIFROOT=/home/<user>/Downloads/gifs/ ./src/quickmeme
```

OR

Download the latest release binary https://github.com/chrismeyersfsu/quickmeme/releases `GIFROOT=/home/<user>/Downloads/gifs/ ./quickmeme`

## Features

* Put `.gif` files in `GIFROOT=<dir>` `<dir>` and they will be found upon `quickmeme` startup.
* Enter tags in the box below each gif separated by commas and ending with comma.
* Search tags via top text box

## Developer Notes

`quickmeme` relies on an sqlite database called `test.db`. On app start `quickmeme` finds all `*.gif` files from the path specified by the environment variable  `GIFROOT`, creates a database entry for each if an entry doesn't already exist (using the gif path as the lookup).
If the entry does already exist, `quickmeme` will not create a database entry for the gif but will, instead, load the tags associated with the database entry. Tags may then be added or removed using the textbox below the gif. The `,` character triggers the processing of the tags.
Each process trigger clears the associated tags and re-associates them. Tags in the database are unique by name and associated with gif.