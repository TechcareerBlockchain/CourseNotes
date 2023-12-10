"use client"
import { contractAddress } from "@/data/constants";
import { useContractRead, useContractWrite, usePrepareContractWrite } from "wagmi";
import abi from "../../../contracts/abi/marketAbi.json";
import { useEffect, useState } from "react";
import { error } from "console";
import { useRouter } from "next/navigation";

const AddProductPage = () => {
    const [res, setRes] = useState("")
    const [productName, setProductName] = useState("")
    const [price, setPrice] = useState("")
    const [decimal, setDecimal] = useState("")
    const router = useRouter()

    const { config } = usePrepareContractWrite({
        address: contractAddress,
        abi: abi,
        functionName: 'publishProduct',
        args: [productName,Number(price),Number(decimal)],
    });
    const { data, isLoading, isSuccess, write, error  } = useContractWrite(config);

    useEffect(()=>{
        if(data?.hash && !isLoading){
            router.push("/products")
        }
    },[data,isLoading])

    const buttonOnClick = () => {
        if(write){
            write()
        }
    }

    return (
        <div className="bg-white py-6 sm:py-8 lg:py-12">
            <div className="mx-auto max-w-screen-2xl px-4 md:px-8">
                <h2 className="mb-4 text-center text-2xl font-bold text-gray-800 md:mb-8 lg:text-3xl">Add New Product</h2>

                <div className="mx-auto max-w-lg rounded-lg border">
                    <div className="flex flex-col gap-4 p-4 md:p-8">
                        <div>
                            <label htmlFor="product_name" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">Product Name</label>
                            <input value={productName} onChange={(event)=>setProductName(event.target.value)} type="text" name="product_name" className="w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring" />
                        </div>
                        <div>
                            <label htmlFor="price" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">Price</label>
                            <input value={price} onChange={(event)=>setPrice(event.target.value)} type="text" name="price" className="w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring" />
                        </div>
                        <div>
                            <label htmlFor="decimal" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">Decimal</label>
                            <input value={decimal} onChange={(event)=>setDecimal(event.target.value)} type="text" name="decimal" className="w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring" />
                        </div>
                        <div>
                            {isSuccess ? "Saved succesfully" : error?.message}
                        </div>

                        <button onClick={buttonOnClick} className="block rounded-lg bg-gray-800 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-gray-300 transition duration-100 hover:bg-gray-700 focus-visible:ring active:bg-gray-600 md:text-base">Save</button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default AddProductPage;