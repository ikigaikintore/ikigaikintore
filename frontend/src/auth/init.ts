import { initializeApp } from "firebase/app"
import { Auth, getAuth } from "firebase/auth"

import { firebaseConfig } from "./config"

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

export { app, auth, token }