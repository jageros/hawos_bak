#!/bin/bash
command="protoc --gofast_out=protos/pb protos/pbdef/*.proto --proto_path=protos/pbdef"
echo $command
`$command`

python tools/gen_meta.py