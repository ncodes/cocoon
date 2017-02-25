# Launch Script

# start docker daemon
bash dockerd-entrypoint.sh dockerd &
echo "Started docker daemon"
sleep 5

# pull launch-go image
docker pull ncodes/launch-go:latest

# pull cocoon source
git clone --depth=1 https://github.com/ncodes/cocoon

# build the binary
cd cocoon
glide install
echo "Building cocoon source"
go build -v -o /bin/cocoon core/main.go 

# start connector 
repoHash=$(git rev-parse HEAD)
echo "Cocoon Version: $repoHash"
cocoon connector