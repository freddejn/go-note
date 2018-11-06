# Small command line application for note taking

# Installation
   * Download repo, unzip to `$GOPATH` run `go install` and put binary in `$PATH`
   * Add export `NOTES_DIR=path_to_where_you_want_your_notes_to_be in your zshrc/bashrc`
   * run using go-note
    
All notes will be put into `$NOTES_PATH`


## TODO

- [ ] Fix error handling
- [x] Add functionaity for search
- [x] Fix so that categories can be created
- [x] Make notes editable with editor of choice

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
* `go-note find / f/ pattern`
* `go-note -h / --help`


## Examples

**To add a note:**

* `go-note add "This is a new note"`

**To list notes:**

* `go-note list`

**To delete a note by id**

* `go-note delete 4`

**To move note to new category**

* `go-note mv 4 -c Work`

**To delete a category**

* `go-note rm -c Work`

**To search for a note** (all searches are boolean)

* `go-note find email boss`

**To search for a note using regex**

* `go-note find "^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"`

