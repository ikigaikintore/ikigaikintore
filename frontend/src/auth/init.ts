import { firebaseConfig } from "./config"
import { initializeApp } from "firebase/app"
import { Auth, getAuth } from "firebase/auth"

const app = initializeApp(firebaseConfig)
const auth: Auth = getAuth(app)

export { app, auth }