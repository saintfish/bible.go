#!/bin/bash

SRC_URL="http://springbible.fhl.net/Bible2/cgic201/Data/bible.txt"
SRC_ENC="big5"
DST_ENC="utf8"
DST_FILE="bible.txt"

curl --silent $SRC_URL | iconv -f $SRC_ENC -t $DST_ENC \
  | sed -e 's/&#x27F2E;/𧼮/g' -e 's/&#x233B4;/𣎴/g' -e 's/榖/殼/g' \
  > $DST_FILE
