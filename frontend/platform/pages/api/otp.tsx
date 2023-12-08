import { NextApiRequest, NextApiResponse } from 'next';
import * as PushAPI from "@pushprotocol/restapi";
import * as ethers from "ethers";
import { ENV } from '@pushprotocol/restapi/src/lib/constants';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
    const PK = "";
    console.log("Private key", PK);
    const _signer = new ethers.Wallet(PK);

    function generateRandomNumber(): number {
        const randomNumber = Math.floor(Math.random() * 1000000);
        return randomNumber;
    }

    const sendNotification = async (otp:string) => {
        try {
            const apiResponse = await PushAPI.payloads.sendNotification({
                signer: _signer,
                type: 3, // broadcast
                identityType: 2, // direct payload
                notification: {
                    title: `[Techcareer] OTP Code`,
                    body: `[Techcareer] OTP Code`
                },
                payload: {
                    title: `[Techcareer] OTP Code`,
                    body: otp,
                    cta: '',
                    img: ''
                },
                recipients: '', // recipient address
                channel: '', // your channel address
                env: ENV.STAGING
            });
            
            res.send(otp)
        } catch (err) {
            console.error('Error: ', err);
        }
    }


    try {
        const otp = generateRandomNumber().toString();
        sendNotification(otp);
        
    } catch (err) {
        console.error('Error: ', err);
        res.status(500).json({ error: 'An error occurred while sending notification.' });
    }
}