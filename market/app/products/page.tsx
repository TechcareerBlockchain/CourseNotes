"use client"
import { useContractRead } from "wagmi";
import abi from "../../contracts/abi/marketAbi.json";
import { useState } from "react";
const Products = () => {
    const contractAddress = '0x22A30026F44b3F4CD39899653320E77B24B909C4'
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
            <h1>Products</h1>
            <div className="relative overflow-x-auto mx-10">
                <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                    <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
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
                        </tr>
                    </thead>
                    <tbody>

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