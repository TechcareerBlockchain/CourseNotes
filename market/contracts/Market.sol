// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Product.sol";

contract Market {

    struct ProductOnSale {
        string name;
        uint256 price;
        uint256 decimal;
        address owner;
        address productIdentifier; // Web2 => ID => Primary Key
    }

    ProductOnSale[] products;

    event ProductListed();

    function publishProduct(string calldata name,uint256 price, uint256 decimal) external {
        address deploymentAddress = address(new Product(name,msg.sender,price,decimal));
        products.push(ProductOnSale(name,price,decimal,msg.sender,deploymentAddress));
        emit ProductListed();
    }

    function getProducts() view external returns(ProductOnSale[] memory){
        return products;
    }

}