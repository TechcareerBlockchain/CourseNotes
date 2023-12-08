"use client";
import {  SupportedLocale, SwapWidget } from "@uniswap/widgets";
import { useCallback, useRef, useState } from "react";

const UniswapPage = () => {
    const TOKEN_LIST = 'https://gateway.ipfs.io/ipns/tokens.uniswap.org'
    const UNI = '0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984'
    const connectors = useRef<HTMLDivElement>(null)
    const focusConnectors = useCallback(() => connectors.current?.focus(), [])
    const JSON_RPC_URL = 'https://cloudflare-eth.com'
    const [locale, setLocale] = useState<SupportedLocale>('en-US')
    return(
        <SwapWidget
              jsonRpcEndpoint={JSON_RPC_URL}
              tokenList={TOKEN_LIST}
              locale={locale}
              onConnectWallet={focusConnectors}
              defaultInputTokenAddress="NATIVE"
              defaultInputAmount="1"
              defaultOutputTokenAddress={UNI}
            />
    )
}

export default UniswapPage;