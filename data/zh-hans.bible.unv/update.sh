#!/bin/bash

SRC_URL="http://springbible.fhl.net/Bible2/cgic201/Data/hgb.txt"
SRC_ENC="gb18030"
DST_ENC="utf8"
DST_FILE="bible.txt"

curl --silent $SRC_URL | iconv -f $SRC_ENC -t $DST_ENC \
  | sed -e 's/逿/𧼮/g' -e 's/﹖/？/g' -e 's/﹐/，/g' -e 's/﹔/；/g' \
        -e 's/凡于你们有益的/凡与你们有益的/' \
        -e 's/我信神他怎样对我说：/我信神他怎样对我说，/' \
        -e 's/他们对差役说；/他们对差役说：/' \
        -e 's/打杖的时候/打仗的时候/' \
        -e 's/。；/；/' \
        -e 's/，，/，/' \
  > $DST_FILE
