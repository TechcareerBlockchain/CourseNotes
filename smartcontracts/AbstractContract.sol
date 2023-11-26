// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


contract Hayvan{
    function yurumek() external virtual returns(string memory){
        return "yurudu";
    }
}

interface SesVerme {
    function ses() external pure returns(string memory);
} 

contract Kedi is Hayvan,SesVerme {
    function ses() external pure returns(string memory){
        return "miyav";
    }
    function yurumek() override external pure returns(string memory){
        return "kedi yurudu";
    }
}

contract Kopek is Hayvan,SesVerme {
    function ses() external pure returns(string memory){
        return "hav";
    }
}