// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./UserProfile.sol";

contract MathContract {

    function addition(uint256 v1,uint256 v2) pure public returns(uint256){
        return v1+v2;
    }

}


interface IAggregator {
    function latestRoundData()
        external
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        );
}

contract Weel1Lesson1 is MathContract {

    struct Person {
        string name;
        uint256 age;
    }

    uint256 public testVariableForInt = 1;
    string public testVariableForString = "deneme";
    bool public testVariableForBoolean = false;
    address testVariableForAddress = 0x676094EE70163AFEEAC9A1b244A1c315b61d79eD;
    uint256[] public testArr; 
    Person[] public users; 
    mapping(address=>string) public userMap;

    event VariableChangedForInt(uint256 newVar);
    event VariableChangedForBool(bool indexed);
    event Counter(uint8 indexed);
    event CounterError();

    constructor() {
        testVariableForAddress = msg.sender;
        testArr.push(addition(1,2));
        testArr.push(addition(4,5));
        users.push(Person("Dogukan",12));
    }

    function createUser(string memory _name,uint128 _age) external returns(address){
        UserProfile user = new UserProfile(_name,_age);
        address userContractAddress = address(user);
        userMap[userContractAddress] = _name;
        return userContractAddress;
    }

    function getTestVariableForAddress() external view returns (address) {
        return testVariableForAddress;
    }

    // BTC/USDC 0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43
    function getPriceOfBTC() external view returns(int256){
        IAggregator aggregator = IAggregator(0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43);
        (, int256 answer,,,) = aggregator.latestRoundData();
        return answer;
    }

    function bytesToUint(bytes memory b) internal pure returns (uint256){
        uint256 number;
        for(uint i=0;i<b.length;i++){
            number = number + uint(uint8(b[i]))*(2**(8*(b.length-(i+1))));
        }
        return number;
    }

    function callMathContract(address mathContractAddress) external returns(uint256){
        bytes memory payload = abi.encodeWithSignature("addition(uint256,uint256)",1,1);
        (bool success,bytes memory data) = mathContractAddress.call(payload);
        require(success,"Feed is not available");
        return bytesToUint(data);
    }

    /// @notice changes the testVariableForInt by the parameter
    /// @param newVar new value for testVariableForInt
    function changeTestVariableIntBy(uint256 newVar) external {
        testVariableForInt = newVar;
        emit VariableChangedForInt(newVar);
    }

    function getLastElement() view external returns(uint256){
        return testArr[testArr.length-1];
    }

    function popLastElement() external{
        testArr.pop();
    }

    function changeTestVariableInt() public payable {
        testVariableForInt = msg.value;
        emit VariableChangedForInt(msg.value);
    }

    function changeTestVariableBool() public {
        testVariableForBoolean = true;
        emit VariableChangedForBool(testVariableForBoolean);
    }

    function countToTenWithEmit() external {
        for (uint8 counter = 0; counter < 10; counter++) {
            emit Counter(counter);
        }
    }

    /// @notice counts to ten end logs only even numbers
    function countToTenOnlyForEvenWithEmit() external {
        for (uint8 counter = 0; counter < 10; counter++) {
            if (counter == 0) {
                emit CounterError();
            } else if (counter % 2 == 0) {
                emit Counter(counter);
            } else {
                emit CounterError();
            }

            /*            
                if (counter == 0){
                    emit CounterError();
                } 
                if (counter % 2 == 0){
                    emit CounterError();
                } else {
                    emit CounterError();
                }
            */
        }
    }
}
