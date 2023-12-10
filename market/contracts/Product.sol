// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Product {

    string name;
    address owner;
    uint256 price;
    uint256 decimal;

    event ProductSold();

    constructor(string memory _name, address _owner, uint256 _price, uint256 _decimal){
        name = _name;
        owner = _owner;
        price = _price;
        decimal = _decimal;
    }

    modifier onlyOwner {
        require(owner == msg.sender,"You are not the owner");
        _;
    }

    function getPrice() view external returns(uint256,uint256){
        return (price,decimal);
    }

    function getOwner() view external returns(address){
        return owner;
    }

    function sellProduct(address newOwner) external onlyOwner{
        owner = newOwner;
        emit ProductSold();
    }

}