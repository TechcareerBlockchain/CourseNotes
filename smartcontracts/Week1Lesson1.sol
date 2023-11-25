// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Weel1Lesson1 {
    uint256 public testVariableForInt = 1;
    string public testVariableForString = "deneme";
    bool public testVariableForBoolean = false;
    address testVariableForAddress = 0x676094EE70163AFEEAC9A1b244A1c315b61d79eD;

    event VariableChangedForInt(uint256 newVar);
    event VariableChangedForBool(bool indexed);
    event Counter(uint8 indexed);
    event CounterError();

    constructor() {
        testVariableForAddress = msg.sender;
    }

    function getTestVariableForAddress() external view returns (address) {
        return testVariableForAddress;
    }

    /// @notice changes the testVariableForInt by the parameter
    /// @param newVar new value for testVariableForInt
    function changeTestVariableIntBy(uint256 newVar) external {
        testVariableForInt = newVar;
        emit VariableChangedForInt(newVar);
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
