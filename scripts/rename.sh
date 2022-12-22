#!/usr/bin/env bash

cd "$(dirname "$0")/.." || exit

if [ "$#" != 2 ]; then
    echo "Usage: $0 <package-name> <binary-name>"
    exit 1
fi

old_package_name="github.com/leonardoce/gostructure"
old_binary_name="myapp"

new_package_name="$1"
new_binary_name="$2"

find . -name \*.go -exec sed -e "s#${old_package_name}#${new_package_name}#" -e "s#${old_binary_name}#${new_binary_name}#" -i '{}' \;
find . -name \*.sh -not -name \*rename.sh -exec sed -e "s#${old_package_name}#${new_package_name}#" -e "s#${old_binary_name}#${new_binary_name}#" -i '{}' \;
find . -name go.\* -exec sed -e "s#${old_package_name}#${new_package_name}#" -e "s#${old_binary_name}#${new_binary_name}#" -i '{}' \;
find . -name \*.yml -exec sed -e "s#${old_package_name}#${new_package_name}#" -e "s#${old_binary_name}#${new_binary_name}#" -i '{}' \;
find . -name \*.qtpl -exec sed -e "s#${old_package_name}#${new_package_name}#" -e "s#${old_binary_name}#${new_binary_name}#" -i '{}' \;

find . -type d -name "*${old_binary_name}*" | rename "s#${old_binary_name}#${new_binary_name}#"
find . -type f -name "*${old_binary_name}*" | rename "s#${old_binary_name}#${new_binary_name}#"