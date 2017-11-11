#!/bin/bash

rm -f ebdeploy.zip
zip -r ebdeploy.zip . -x '*.git*' 'node_modules/*'
eb deploy
