"use client"

import abi from "@/abi.json"
import { useState } from "react"
import { useContractRead, useContractWrite, usePrepareContractWrite, useWaitForTransaction } from 'wagmi'

export function Page() {
  const contractAddress = '0x06183cef76927aefBA7672644A3f0aF35a8A1B07'
  const [textNumber1, setTextNumber1] = useState("")
  const [textNumber2, setTextNumber2] = useState("")
  const [res, setRes] = useState("")
  const [appError, setAppError] = useState("")

  /*
    These hooks are used for writing data or modifying the state of a smart contract. 
    This includes any function that changes the data on the blockchain.
    usePrepareContractWrite prepares the configuration for a contract write operation.
    useContractWrite uses this configuration to execute the write operation. 
    It's used for functions that will typically be marked as nonpayable or payable in Solidity.
    
    const { config } = usePrepareContractWrite({
      address: contractAddress,
      abi: abi,
      functionName: 'addition',
      args: [parseInt(textNumber1), parseInt(textNumber2)]
    });
    const { write } = useContractWrite(config);
  */

    /*
      useContractRead is used to read data from a smart contract. 
      It's typically used for calling functions that do not modify the contract's state. 
      These are often marked as view or pure in Solidity.
      This hook is used to execute functions that don't require a gas fee because they don't make any changes to the blockchain.
    */
    const { refetch } = useContractRead({
      address: contractAddress,
      abi: abi,
      onSuccess: (data:any) => {
        setRes(data.toString())
      },
      functionName: 'addition',
      args: [parseInt(textNumber1), parseInt(textNumber2)],
      enabled: false,
    });
    const buttonOnClick = async () => {
      try {
        await refetch();
      } catch (error:any) {
        setAppError(error.toString());
      }
    };
    
  return (
    <>
      <h1 className="text-center mt-10 font-bold text-xl">Deployed Contract Address Is {contractAddress}</h1>
      <div className="max-w-sm mx-auto">
        <div className="mb-5">
          <label htmlFor="number1" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Number 1</label>
          <input value={textNumber1} onChange={(event) => setTextNumber1(event.target.value)} type="text" id="number1" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="Number 1" required />
        </div>
        <div className="mb-5">
          <label htmlFor="number1" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Number 2</label>
          <input value={textNumber2} onChange={(event) => setTextNumber2(event.target.value)} type="text" id="number2" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="Number 2" required />
        </div>
        <div className="mb-5">
          <button onClick={() => buttonOnClick()} type="submit" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Submit</button>
        </div>
        <div className="mb-5">
          <span>Number 1 is {textNumber1} and Number 2 is {textNumber2}</span>
        </div>
        <div className="mb-5">
          <span>Result is { res }</span>
        </div>
        <div className="mb-5">
          <span>Error is - {appError}</span>
        </div>
      </div>
    </>
  )
}

export default Page
