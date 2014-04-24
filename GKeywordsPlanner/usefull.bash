#!/bin/bash


bin/keywordsplannerApp -locale=it_IT -themes=finance
bin/setphrasesforprod -locale=it_IT -themes=finance
#rm csvfiles/it_IT_finance/*.csv


bin/keywordsplannerApp -locale=fi_FI -themes=finance
bin/setphrasesforprod -locale=fi_FI -themes=finance

bin/keywordsplannerApp -locale=fi_FI -themes=fortune
bin/setphrasesforprod -locale=fi_FI -themes=fortune

find csvfiles -type f -delete




