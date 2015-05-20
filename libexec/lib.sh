# echo cmd and run cmd
run_cmd() {
  cmd="$1"; shift
  echo -e ">> $cmd $@"
  $cmd $@
}

pd() {
  echo -e "$@"
}

say() {
  echo -e "$@"
}

say2() {
  echo -e "$@" >&2
}

error() {
  echo -e "$@" >&2
}

error_exit() {
  echo -e "$@" >&2
  exit 1
}

debug() {
  if [[ $DEBUG != "" ]]; then
    echo -e "$@" >&2
  fi
}

absolutename() {
	readlink -m "$1"
}

filename() {
	base=$(basename "$1")
	echo ${base%.*}
}

extname() {
	base=$(basename "$1")
	[[ "$base" =~ \. ]] && echo "${base##*.}" || echo ""
}
