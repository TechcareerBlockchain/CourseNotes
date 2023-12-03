"use client"
import redstone from "redstone-api";
import { useEffect, useState } from "react";
const OffChainPage = () => {

    const [price, setPrice] = useState(0);  

    useEffect(() => {
        const getPrice = async () => {
            return await redstone.getPrice("BTC");
        }

        getPrice().then((res) => {
            setPrice(res.value);  
        });
    }, []);

    return(
        <div>
            <h1>OffChainPage</h1>
            <div className="mx-auto"> 
                Price of BTC is {price}
            </div>
        </div>
    )
}

export default OffChainPage;