"use client"
import { useEffect, useState } from "react";
import axios from "axios";
import { useAccount } from "wagmi";

const OnChainPage = () => {
    const [priceFeedAddress, setPriceFeedAddress] = useState("")
    const [tokenAddress, setTokenAddress] = useState("")
    const [res, setRes] = useState("")
    const [appError, setAppError] = useState("")
    const account = useAccount();

    const buttonOnClick = async () => {
        try {
            const res = await axios.post("http://localhost:3000/chainlink/price",{
                "price_feed_address": priceFeedAddress,
                "token_address": tokenAddress 
            },{
                headers: {
                    'walletAddress': account.address,
                }
            });
            setRes(res.data.data);
            console.log(res);
        } catch (error: any) {
            setAppError(error.toString());
        }
    };


    return (
        <div>
            <h1>OnChainPage</h1>
            <div className="max-w-sm mx-auto">
                <div className="mb-5">
                    <label htmlFor="PriceFeedAddress" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Price Feed Address</label>
                    <input value={priceFeedAddress} onChange={(event) => setPriceFeedAddress(event.target.value)} type="text" id="PriceFeedAddress" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="Price Feed Address" required />
                </div>
                <div className="mb-5">
                    <label htmlFor="TokenAddress" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Token Address</label>
                    <input value={tokenAddress} onChange={(event) => setTokenAddress(event.target.value)} type="text" id="TokenAddress" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="Token Address" required />
                </div>
                <div className="mb-5">
                    <button onClick={() => buttonOnClick()} type="submit" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Submit</button>
                </div>
                <div className="mb-5">
                    <span>Result is {res}</span>
                </div>
                <div className="mb-5">
                    <span>Error is - {appError}</span>
                </div>
            </div>
        </div>
    );
}

export default OnChainPage;