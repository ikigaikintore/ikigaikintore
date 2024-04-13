import React from "react"

import { User } from "firebase/auth"

import * as fb from "./firebase"

export type AuthContextType = {
  user: User | null
  loading: boolean
  signInUser: () => Promise<void>
  signOutUser: () => Promise<void>
}

export const useAuth = () => {
    const [user, setUser] = React.useState<User | null>(null)
    const [loading, setLoading] = React.useState(true)
    const { auth, signInUser, signOutUser } = fb.useFirebase()

    React.useEffect(() => {
        const unsubscribe = auth.onAuthStateChanged((user) => {
            setUser(user)
            setLoading(false)
        })

        return () => unsubscribe()
    }, [auth, user])

    const signIn = async () => {
        setLoading(true)
        try {
            await signInUser()
        } catch (err) {
            console.error(`Error signing in: ${err}`)
        }
    }

    const signOut = async () => {
        setLoading(true)
        try {
            await signOutUser()
        } catch (err) {
            console.error(`Error signing out: ${err}`)
        }
    }

    return {
        user,
        loading,
        signInUser: signIn,
        signOutUser: signOut,
    }
}
