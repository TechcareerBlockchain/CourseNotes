// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IProduct {
    function getOwner() view external returns(address);
    function getPrice() view external returns(uint256,uint256);
}

contract MarketEscrow { 

    address buyer;
    address productAddress;
    address seller;
    uint256 productPrice;
    bool signedByBuyer;
    bool signedBySeller;
    IProduct iProduct;

    constructor(address _productAddress,address _buyer){
        buyer = _buyer;
        productAddress = _productAddress;
        iProduct = IProduct(productAddress);
        seller = iProduct.getOwner();
        (uint256 price,) = iProduct.getPrice();
        productPrice = price;
    }

    function signByBuyer() external payable {
        require(msg.value != productPrice,"Transfer did not work");
        signedByBuyer = true;
    }

    function refund() external payable {
        address productOwnerAddress = iProduct.getOwner();
        require(productOwnerAddress != seller && productOwnerAddress != buyer,"Procces is still going on");
        uint amount = address(this).balance;
        (bool success, ) = buyer.call{value: amount}("");
        require(success, "Failed to send Ether");
    }

    function signBySeller() external payable {
        address productOwnerAddress = iProduct.getOwner();
        require(productOwnerAddress != seller && productOwnerAddress == buyer,"Ownership is still same");
        (bool success, ) = seller.call{value: productPrice * 1 ether}("");
        require(success, "Failed to send Ether");
        uint amount = address(this).balance;
        (success, ) = buyer.call{value: amount}("");
        require(success, "Failed to send Ether");
    }

}