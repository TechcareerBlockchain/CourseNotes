"use client";
import { CredentialType, IDKitWidget, ISuccessResult } from "@worldcoin/idkit";

const Login = () => {

    const onSuccess = (response: ISuccessResult) => {
        console.log(response);
    }
    return (
        <div>
            <IDKitWidget
                app_id={process.env.NEXT_PUBLIC_WORLDCOIN} // obtained from the Developer Portal
                action="login" // this is your action name from the Developer Portal
                onSuccess={onSuccess}
                credential_types={[CredentialType.Device, CredentialType.Orb]} // the credentials you want to accept
            >
                {({ open }) => <button onClick={open}>Verify with World ID</button>}
            </IDKitWidget>
        </div>

    )
}

export default Login;