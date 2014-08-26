#!/bin/bash


bin/keywordsplannerApp -locale=it_IT -themes=finance
bin/setphrasesforprod -locale=it_IT -themes=finance
#rm csvfiles/it_IT_finance/*.csv


bin/keywordsplannerApp -locale=fi_FI -themes=finance
bin/setphrasesforprod -locale=fi_FI -themes=finance

bin/keywordsplannerApp -locale=fi_FI -themes=fortune
bin/setphrasesforprod -locale=fi_FI -themes=fortune

bin/keywordsplannerApp -locale=en_US -themes=finance
bin/setphrasesforprod -locale=en_US -themes=finance


find csvfiles -type f -delete


#GOPATH=$GOPATH:/home/juno/git/GKeywordsPlanner/GKeywordsPlanner go test -v




