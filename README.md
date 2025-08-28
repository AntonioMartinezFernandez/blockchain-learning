# Blockchain Learning

This repository contains educational resources and projects related to blockchain technology. It is designed to help beginners and enthusiasts learn about the fundamentals of blockchain, smart contracts, and decentralized applications (DApps).

## Resources

- [Golang ethereum project [video]](https://www.youtube.com/playlist?list=PLay9kDOVd_x7hbhssw4pTKZHzzc6OG0e_)

## Useful Links

- [Etherscan](https://etherscan.io/)
- [Infura](https://www.infura.io/)
- [Metamask](https://metamask.io/en-GB)
- [Sepolia [testnet] ETHER mining faucet](https://sepolia-faucet.pk910.de/)
- [Ganache Install](https://archive.trufflesuite.com/ganache/)
- [Truffle Install](https://archive.trufflesuite.com/docs/truffle/how-to/install/)
- [Hardhat](https://hardhat.org/getting-started/)
- [Remix IDE](https://remix.ethereum.org/)
- [Solidity Docs](https://docs.soliditylang.org/en/v0.8.30/)
- [Solidity by Example](https://solidity-by-example.org/)
- [OpenZeppelin library Docs](https://docs.openzeppelin.com/contracts/5.x/)
- [Solidity Compiler (solc)](https://docs.soliditylang.org/en/latest/installing-solidity.html#docker)
- [Abigen (Go bindings for Ethereum contracts)](https://geth.ethereum.org/docs/tools/abigen)

## Useful Commands

```bash
# Initialize a new Hardhat project
npx hardhat --init

# Start a local blockchain
npx hardhat node

# Run tests
npx hardhat test

# Deploy to local blockchain
npx hardhat ignition deploy ignition/modules/Counter.ts --network localhost

# Compile contracts
npx hardhat compile

# Compile contracts using solc in Docker
docker run \
    --volume "/tmp/some/local/path/:/sources/" \
    ethereum/solc:stable \
        /sources/Contract.sol \
        --abi \
        --bin \
        --output-dir /sources/output/

# Install abigen
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Add Go bin to PATH (add this line to your .bashrc or .zshrc)
export PATH=$PATH:$(go env GOPATH)/bin

# Generate Go bindings for a contract
abigen --bin=./path/to/contract.bin --abi=./path/to/contract.abi --pkg=packageName --out=./path/to/output.go
```
