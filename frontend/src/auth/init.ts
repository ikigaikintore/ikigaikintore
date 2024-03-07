import { initializeApp } from "firebase/app"
import { Auth, getAuth } from "firebase/auth"
import { connectFunctionsEmulator, getFunctions } from "firebase/functions"

import { firebaseConfig } from "./config"

import { useConfig } from "@/src/config/use-config"

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
if (useConfig().isLocal()) {
  const { host, port } = useConfig().functionsEmulatorHost()
  connectFunctionsEmulator(functions, host, port)
}

export { app, auth, token }