import { initializeApp } from "firebase/app"
import { Auth, getAuth } from "firebase/auth"
import { connectFunctionsEmulator, getFunctions } from "firebase/functions"

import { firebaseConfig } from "./config"

import { Config } from "@/src/config/use-config"

const app = initializeApp(firebaseConfig)
const auth: Auth = getAuth(app)
const token = async () => {
    try {
        const userToken = await auth.currentUser?.getIdToken()
        return userToken || ""
    } catch (err) {
        console.error(`Error getting token: ${err}`)
        return ""
    }
}

const functions = getFunctions(app)
if (Config().isLocal()) {
    const { host, port } = Config().functionsEmulatorHost()
    connectFunctionsEmulator(functions, host, port)
}

export { app, auth, token }