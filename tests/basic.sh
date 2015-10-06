#!/bin/bash
echo "START shebl test"
mkdir tmp
cd tmp
touch testfile
ls -l
rm testfile
cd ..
rmdir tmp
echo "END shebl test"

