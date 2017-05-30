#!/bin/sh


DIR="$(cd "$(dirname $0)" && pwd )"  # absolutized and normalized
TTEMPLATE=${TTEMPLATE:-${DIR}/../../bin/t-template}
checksum=md5

exit_code=0
check_if_failed(){
    file1="$1"
    file2="$2"
    if diff "$file1" "$file2" > /dev/null 2>&1;
    then
        echo "[passed]"
    else
        echo "[failed]"
        exit_code=1
    fi
}
# Create a temp dir
mytmpdir=$(mktemp -d 2>/dev/null || mktemp -d -t 'mytmpdir')


test_path="${DIR}/../examples/simple/"
output_file="${mytmpdir}/simple_gen"
touch "${output_file}"
printf "Simple test: "
${TTEMPLATE} "${test_path}/simple.ctmp" \
                    -j "${test_path}/users.json" \
                    -y "${test_path}/users.yml" \
                    -o "${output_file}"

check_if_failed "${DIR}/simple_output" "${output_file}"


test_path="${DIR}/../examples/usernames/"
output_file="${mytmpdir}/userReport_gen"
touch "${output_file}"
printf "user report test: "
${TTEMPLATE} "${test_path}/userReport.ctmp" \
                    -y "${test_path}/cj.yml" \
                    -y "${test_path}/jp.yml" \
                    -y "${test_path}/yo.yml" \
                    -j "${test_path}/ht.json" \
                    -j "${test_path}/pk.json" \
                    -j "${test_path}/ra.json" > "${output_file}"
check_if_failed "${DIR}/userReport_output" "${output_file}"

test_path="${DIR}/../examples/usernames/"
output_file="${mytmpdir}/github_gen"
touch "${output_file}"
printf "github test: "
${TTEMPLATE} "${test_path}/github.ctmp" \
                    -y "${test_path}/cj.yml" \
                    -y "${test_path}/jp.yml" \
                    -y "${test_path}/yo.yml" \
                    -j "${test_path}/ht.json" \
                    -j "${test_path}/pk.json" \
                    -j "${test_path}/ra.json" > "${output_file}"
check_if_failed "${DIR}/github_output" "${output_file}"

test_path="${DIR}/../examples/usernames/"
output_file="${mytmpdir}/delim_gen"
touch "${output_file}"
printf "delim test: "
test_path="${DIR}/../examples/delim/"
cat "${test_path}/test4.yml" | ${TTEMPLATE} "${test_path}/delim.ctmp" -Y \
                    -l "<%" -r "%>" \
                    -j "${test_path}/test1.json" \
                    -j "${test_path}/test2.json" \
                    -j "${test_path}/test3.json" > "${output_file}"
check_if_failed "${DIR}/delim_output" "${output_file}"

rm -rf "${mytmpdir}"
exit ${exit_code}
