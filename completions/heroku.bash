/Users/jdickey/src/github.com/heroku/heroku-cli/bin/run version
_heroku() {
  local cur cword words opts heroku
  _get_comp_words_by_ref -n : cword words cur
  #heroku="heroku"
  heroku="/Users/jdickey/src/github.com/heroku/heroku-cli/heroku-cli"

  opts=`$heroku completions --cword $cword -- ${words[@]}`

  COMPREPLY=($(compgen -W "${opts}" -- ${cur}))
  __ltrim_colon_completions "$cur"
  return 0
} &&
complete -F _heroku heroku
