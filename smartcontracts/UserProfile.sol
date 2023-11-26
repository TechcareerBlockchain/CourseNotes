// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract UserProfile {

    string name;
    uint128 age;

    constructor(string memory _name,uint128 _age){
        name = _name;
        age = _age;
    }

    function getUserInfo() external view returns(string memory,uint128){
        return (name,age);
    }
}