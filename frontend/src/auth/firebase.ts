import React from "react"

import { Auth, connectAuthEmulator, GoogleAuthProvider, signInWithPopup, signOut, } from "firebase/auth"

import { auth } from "./init"

import { Config } from "@/src/config/use-config"

export const useFirebase = () => {
    const { isLocal, emulatorHost, emails } = Config()

    const setupEmulators = React.useCallback(async (auth: Auth) => {
        await fetch(emulatorHost())
        connectAuthEmulator(auth, emulatorHost(), { disableWarnings: true })
    }, [emulatorHost])

    React.useEffect(() => {
        if (isLocal()) {
            setupEmulators(auth).then(() => console.log("loaded")).catch(e => console.error(e))
        }
    }, [emulatorHost, isLocal, setupEmulators])

    const signInUser = React.useCallback(async () => {
        try {
            const resp = await signInWithPopup(auth, new GoogleAuthProvider())
            if (!resp || !resp.user) {
                console.error("No response from Google")
                return
            }
            if (resp.user.isAnonymous || !resp.user.email) {
                console.error("Not authorized")
                return
            }
            if (!emails.includes(resp.user.email)) {
                console.error("Email not authorized")
                return
            }
        } catch (err) {
            console.error(`Error signing ${err}`)
        }
    }, [emails])

    const signOutUser = React.useCallback(async () => {
        try {
            await signOut(auth)
        } catch (err) {
            console.error(`Error signing out ${err}`)
        }
    }, [])

    return {
        auth,
        signInUser,
        signOutUser,
    }
}
