#!/bin/sh
set -e -u -x

# Necessary variables:
# $PACKAGE
# $HOOK
# $CHAT
# $SEGMENT

FILEDATE=`date "+%Y%m"`

mkdir csv || true
rm -rf csv/* || true
# download
gsutil -m cp -r "gs://${SEGMENT}/reviews/reviews_${PACKAGE}_${FILEDATE}.csv" csv

mkdir utftmp || true
rm -rf utftmp/* || true

pushd csv
  for f in *.csv; 
  do 
    iconv -f UTF-16LE -t UTF-8 $f > ../utftmp/"$f"; 
  done
popd
rm -rf csv || true
mv utftmp csv

FILENAME="csv/reviews_${PACKAGE}_${FILEDATE}.csv"

go run main.go  -hook=$HOOK -file=$FILENAME -chat=$CHAT

