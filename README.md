# Small command line application to store notes


# Installation

    * Download run `go install` and put binary in `$PATH`
    * Add export `NOTES_DIR=path_to_where_you_want_your_notes_to_be in your zshrc/bashrc`
    * run using go-note

All notes and will be put into `$NOTES_PATH`

## TODO

- [] Make so that notes can be added just by typing go-note "note"

```
notes
|-- Category1
|   |-- 181105140821.5310_Send_email_..._.note
|-- Category2
|   |-- 181105092759.7175_Apply_for_t..._.note
|-- Category3
    |-- 181104091716.8531_Remember_to..._.note
    |-- 181104210414.9005_Need_to_che..._.note
```

# Interface

* `go-note a / add "note" [-c/--category "category"]`
* `go-note l / list [-c/--category "category"]`
* `go-note rm / remove/ id [-c/--category "category"]`
* `go-note mv / move/ id category`
* `go-note -h / --help```
