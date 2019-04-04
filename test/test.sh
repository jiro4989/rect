#!/bin/bash

set -eu

make build
./bin/rect paste -X 10 <(echo -e '123\n456') <(echo -e 'abc')
echo ---
./bin/rect paste -X 1 -Y 2 <(echo -e 'あ3\n456\n789') <(echo -e 'abcd\nefgh\nabc')
echo ---
./bin/rect paste <(echo -e 'あ3\n456\n789') <(echo -e 'あ3d\nefgh\nabc')
echo ---
./bin/rect paste <(echo -e 'あいうえお') <(echo -e '123456')
