#!/bin/sh

DIR="$(cd "$(dirname $0)" && pwd )"  # absolutized and normalized
TTEMPLATE=${DIR}/../../bin/t-template
exit_code=0
# Create a temp dir
mytmpdir=$(mktemp -d 2>/dev/null || mktemp -d -t 'mytmpdir')

check_if_failed(){
    file1="$1"
    file2="$2"
    if diff "$file1" "$file2" > /dev/null 2>&1;
    then
        printf "%-20s \t\t[passed]\n" "${3}"
    else
        printf  "%-20s \t\t[failed]\n" "${3}"
        exit_code=1
    fi
}

# TEST 1 simple
test_path="${DIR}/../examples/simple/"
output_file="${mytmpdir}/simple_gen"
touch "${output_file}"
PATH="${DIR}/../../bin/:${PATH}"
"${test_path}/simple.ctmp" \
                    -j "${test_path}/users.json" \
                    -y "${test_path}/users.yml" \
                    -o "${output_file}"

check_if_failed "${DIR}/simple_output" "${output_file}" "Simple test"

# TEST 2 userreport
test_path="${DIR}/../examples/usernames/"
output_file="${mytmpdir}/userReport_gen"
${TTEMPLATE} "${test_path}/userReport.ctmp" \
                    -y "${test_path}/cj.yml" \
                    -y "${test_path}/jp.yml" \
                    -y "${test_path}/yo.yml" \
                    -j "${test_path}/ht.json" \
                    -j "${test_path}/pk.json" \
                    -j "${test_path}/ra.json" > "${output_file}"
check_if_failed "${DIR}/userReport_output" "${output_file}" "User report test"

# TEST 3 username
test_path="${DIR}/../examples/usernames/"
output_file="${mytmpdir}/github_gen"
"${test_path}/github.ctmp" \
                    -y "${test_path}/cj.yml" \
                    -y "${test_path}/jp.yml" \
                    -y "${test_path}/yo.yml" \
                    -j "${test_path}/ht.json" \
                    -j "${test_path}/pk.json" \
                    -j "${test_path}/ra.json" > "${output_file}"
check_if_failed "${DIR}/github_output" "${output_file}" "Github test"

# TEST 4 delim
output_file="${mytmpdir}/delim_gen"
test_path="${DIR}/../examples/delim/"
cat "${test_path}/test4.yml" | ${TTEMPLATE} "${test_path}/delim.ctmp" -Y \
                    -l "<%" -r "%>" \
                    -j "${test_path}/test1.json" \
                    -j "${test_path}/test2.json" \
                    -j "${test_path}/test3.json" > "${output_file}"
check_if_failed "${DIR}/delim_output" "${output_file}" "Delim test"

# TEST 5 no input
output_file="${mytmpdir}/noinput_gen"
test_path="${DIR}/../examples/noinput/"
export NAME_X="MAX"
export LAST_X="POWER"
${TTEMPLATE} "${test_path}/noinput.ctmp" > "${output_file}"
check_if_failed "${DIR}/noinput_output" "${output_file}" "No input"

rm -rf "${mytmpdir}"
exit ${exit_code}
