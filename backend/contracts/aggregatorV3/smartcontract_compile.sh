solc --abi $1.sol -o build
solc --bin $1.sol -o build
abigen --bin=./build/$0.bin --abi=./build/$0.abi --pkg=$2 --out=$3.go
