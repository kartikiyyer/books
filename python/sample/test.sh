#!/usr/bin/env bash

#OLD_VERSION=$(echo "17.2R2-22" | cut -d'-' -f1)
#NEW_VERSION=$(echo "19.1R1-20" | cut -d'-' -f1)
#SCRIPTS_LIST_FILE=/tmp/upgrade_pe_scripts$$
#
#awk '{out=""; for(i=5;i<=NF;i++){out=out" "$i}; print $1 " " $3 out}' /Users/kiyyer/Documents/Git/sdsn/dev_19_1_r1/tr-platform/upgrades/pe_upgrade/files/install_pip_packages_map | grep "^$OLD_VERSION $NEW_VERSION" | cut -d' ' -f3- > $SCRIPTS_LIST_FILE
#
#(for script in $(cat ${SCRIPTS_LIST_FILE})
# do
#    echo "Executing ${script}"
#    (while IFS= read line
#     do
#        echo ${line}
#        if [ $? -ne 0 ]
#            then exit 1
#        fi
#     done < "/Users/kiyyer/Documents/Git/sdsn/dev_19_1_r1/tr-platform/upgrades/pe_upgrade/files/${script}")
#
# done)

#
#if [ $# -ne 2 ]
#    then echo "Usage: $0 <oldVersion> <newVersion>"
#         exit 1
#fi
#
#
#echo "Hi"
#

proceed=false

read -r -p "It is highly recommended to take snapshot of Policy Enforcer VM before upgrade.
Are you sure you want to proceed with upgrade? [yes/no] " response
case "$response" in
    [yY][eE][sS])
        proceed=true
        ;;
    *)
        proceed=false
        ;;
esac


if [ $proceed = "false" ]
    then echo "Upgrade aborted!"
         exit 1
fi