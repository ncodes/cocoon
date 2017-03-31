#!/bin/bash
# Run the connector 
set -e


# Set up go environment
export GOROOT=/go
PATH=$PATH:$GOROOT/bin
export GOPATH=/gocode
mkdir -p $GOPATH

# Pull cocoon source
branch="master"
repoOwner=github.com/ncodes
repoOwnerDir=$GOPATH/src/$repoOwner
mkdir -p $repoOwnerDir
cd $repoOwnerDir
printf "> Fetching cocoon source. [branch=$branch] [dest=$repoOwnerDir]\n"
git clone --depth=1 -b $branch https://$repoOwner/cocoon

# build the binary
printf "> Building cocoon"
cd cocoon
rm -rf .glide/ && rm -rf vendor
glide --debug install
/go/bin/go build -v -o /bin/cocoon core/main.go
which cocoon
sleep 60

# start connector 
# printf "Running Cocoon Connector"
# cocoon connector