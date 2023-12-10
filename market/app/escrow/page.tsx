"use client"
import { useState } from "react";

const EscrowPage = () => {
    const [escrowContractAddress,setEscrowContractAddress] = useState("");
    const [price, setPrice] = useState("")
    const [newOwner, setNewOwner] = useState("")

    const buttonOnClickForSeller = () => {
    }
    const buttonOnClickForBuyer = () => {
    }
    const buttonOnClickForChangingOwnership = () => {
    }
    return (
        <div className="bg-white py-6 sm:py-8 lg:py-12">
            <div className="mx-auto max-w-screen-2xl px-4 md:px-8">
                <h2 className="mb-4 text-center text-2xl font-bold text-gray-800 md:mb-8 lg:text-3xl">Escrow</h2>

                <div className="mx-auto max-w-lg rounded-lg border">
                    <div className="flex flex-col gap-4 p-4 md:p-8">
                        <div>
                            <label htmlFor="product_name" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">Escrow Contract Address or Product Address</label>
                            <input value={escrowContractAddress} onChange={(event)=>setEscrowContractAddress(event.target.value)} type="text" name="product_name" className="w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring" />
                        </div>
                        <div>
                            <label htmlFor="price" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">Price</label>
                            <input value={price} onChange={(event)=>setPrice(event.target.value)} type="text" name="price" className="w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring" />
                        </div>
                        <div>
                            <label htmlFor="new_owner" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">New Owner</label>
                            <input value={newOwner} onChange={(event)=>setNewOwner(event.target.value)} type="text" name="new_owner" className="w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring" />
                        </div>

                        {/*  If seller want to change the ownership of the product, this page must be used  */}

                        <button onClick={buttonOnClickForChangingOwnership} className="block rounded-lg bg-gray-800 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-gray-300 transition duration-100 hover:bg-gray-700 focus-visible:ring active:bg-gray-600 md:text-base">Change Ownership</button>
                        <button onClick={buttonOnClickForSeller} className="block rounded-lg bg-gray-800 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-gray-300 transition duration-100 hover:bg-gray-700 focus-visible:ring active:bg-gray-600 md:text-base">Sign By Buyer</button>
                        <button onClick={buttonOnClickForBuyer} className="block rounded-lg bg-gray-800 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-gray-300 transition duration-100 hover:bg-gray-700 focus-visible:ring active:bg-gray-600 md:text-base">Sign By Seller</button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default EscrowPage;