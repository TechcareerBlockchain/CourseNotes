"use client"
import { useContractRead } from "wagmi";
import abi from "../../contracts/abi/marketAbi.json";
import { useState } from "react";
import Link from "next/link";
import { contractAddress } from "@/data/constants";
const Products = () => {
    const [res, setRes] = useState<Products[]>([])
    useContractRead({
        address: contractAddress,
        abi: abi,
        onSuccess: (data: Products[]) => {
            setRes(data)
        },
        functionName: 'getProducts',
        args: [],
        enabled: true,
    });
    return (
        <div>
            <div className="bg-white py-6 sm:py-8 lg:py-12">
                <div className="mx-auto max-w-screen-2xl px-4 md:px-8">
                    <div className="flex flex-col items-center justify-between gap-4 rounded-lg bg-gray-100 p-4 sm:flex-row md:p-8">
                        <div>
                            <h2 className="text-xl font-bold text-indigo-500 md:text-2xl">Products</h2>
                            <p className="text-gray-600">Onchain Products</p>
                        </div>

                        <Link href="/products/new" className="inline-block rounded-lg bg-indigo-500 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-indigo-300 transition duration-100 hover:bg-indigo-600 focus-visible:ring active:bg-indigo-700 md:text-base">Add Product</Link>
                    </div>
                </div>
            </div>
            <div className="relative overflow-x-auto mx-10">
                <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                    <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr className="text-white">
                            <th scope="col" className="px-6 py-3">
                                Identifier
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Product name
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Price
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Decimal
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Owner
                            </th>
                            <th scope="col" className="px-6 py-3"/>
                        </tr>
                    </thead>
                    <tbody className="text-white">

                        {res.map((product: Products) => {
                            return (
                                <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                                    <th scope="row" className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                        {product.productIdentifier}
                                    </th>
                                    <td className="px-6 py-4">
                                        {product.name}
                                    </td>
                                    <td className="px-6 py-4">
                                        {product.price.toString()}
                                    </td>
                                    <td className="px-6 py-4">
                                        {product.decimal.toString()}
                                    </td>
                                    <td className="px-6 py-4">
                                        {product.owner}
                                    </td>
                                    <td className="px-6 py-4">
                                        <Link href="/products/buy" className="inline-block rounded-lg bg-indigo-500 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-indigo-300 transition duration-100 hover:bg-indigo-600 focus-visible:ring active:bg-indigo-700 md:text-base">Buy</Link>
                                    </td>
                                </tr>
                            )
                        })
                        }
                    </tbody>
                </table>
            </div>


        </div>
    )
}
export default Products;