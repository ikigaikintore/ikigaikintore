import React from "react"
import {
  signInWithPopup,
  GoogleAuthProvider,
  signOut,
  connectAuthEmulator,
} from "firebase/auth"
import { useConfig } from "@/src/config/use-config"
import { auth } from "./init"

export const useFirebase = () => {
  const { isLocal, emulatorHost, emails } = useConfig()

  React.useEffect(() => {
    if (isLocal()) {
      connectAuthEmulator(auth, emulatorHost())
    }
  }, [auth])

  const signInUser = React.useCallback(async () => {
    try {
      const resp = await signInWithPopup(auth, new GoogleAuthProvider())
      if (!resp || !resp.user) {
        throw new Error("No response from Google")
      }
      if (resp.user.isAnonymous || !resp.user.email) {
        throw new Error("Not authorized")
      }
      if (!emails.includes(resp.user.email)) {
        throw new Error("Email not authorized")
      }
    } catch (err) {
      console.error(`Error signing ${err}`)
    }
  }, [auth, emails])

  const signOutUser = React.useCallback(async () => {
    try {
      await signOut(auth)
    } catch (err) {
      console.error(`Error signing out ${err}`)
    }
  }, [signOut, auth])

  return {
    auth,
    signInUser,
    signOutUser,
  }
}
