RUN_NAME="autoclock"

mkdir -p output/bin output/conf
cp script/bootstrap.sh output
chmod +x output/bootstrap.sh

go build -i -v -o output/bin/${RUN_NAME}
