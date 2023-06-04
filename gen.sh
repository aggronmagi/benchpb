

mkdir -p ./gopb/
protoc --gopb_out=./gopb/ fields.proto

mkdir -p githubgo/
protoc --githubgo_out=./githubgo/ fields.proto

mkdir -p googlego/
protoc --googlego_out=./googlego/ fields.proto


mkdir -p gogo/
protoc --gogo_out=./gogo/ fields.proto


mkdir -p gogofast/
protoc --gogofast_out=./gogofast/ fields.proto


mkdir -p gogofaster/
protoc --gogofaster_out=./gogofaster/ fields.proto
