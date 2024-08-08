source ${SCRIPT_DIR}/lib/util.sh

function run() {
    dir="${1}"

    print_header "Running unit tests for ${dir}"

    # Install dependencies if required.
    pre_test "${dir}"

    pushd "${dir}" >/dev/null || exit

    if [[ "x${ORB_NODOWNLOAD}" == "x" ]]; then
        print_msg "Downloading go modules"
        go mod download &>/dev/null 2>&1
        print_msg "done"
    fi

    # Run tests.
    echo $(go env GOPATH)/bin/go test ./...  -v -benchmem -bench=.
    $(go env GOPATH)/bin/richgo test ./... -v -benchmem -bench=.

    rc=$?

    popd >/dev/null || exit

    # Kill all depdency processes.
    post_test "${dir}"

    return ${rc}
}

if [[ "${1}" == "xargs" ]]; then
    out=$(run $2 "${ORB_ROOT}" 2>&1)
    rc=$?
    printf "%s" "${out}"

    exit ${rc}
else
    run $2 "${ORB_ROOT}"
    exit $?
fi