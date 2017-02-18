# set coccon firewall and groups
bash ${NOMAD_META_SCRIPTS_DIR}/${COCOON_GRPS_SCRIPT_NAME}

# pull cocoon source
git clone -b master https://github.com/ncodes/cocoon
git config --global http.https://gopkg.in.followRedirects true

# compilte data directory to binary
cd cocoon/core/data
go get -u -v github.com/jteeuwen/go-bindata/...
go-bindata --pkg data ./...

# launcher launcher
cd ../launcher
glide install
go build -o /bin/launcher launcher.go

# start launcher
launcher start