import * as Domain from '@/domain/me'
import { useSubmitLogIn } from '@/services/login'
import * as React from 'react'

type LoggedMe = (Domain.Me & { authToken: string }) | undefined

export type LoggedMeValue = {
    me: LoggedMe
}

export const useAuthContext = () => {
    const { logIn, isLoading, isError } = useSubmitLogIn()

    const [authMe, setAuthMe] = React.useState<LoggedMe | undefined>(undefined)

    const getLoggedMe = React.useCallback(
        async (username: string, password: string) => {
            try {
                const result = await logIn({
                    email: username,
                    password: password,
                })

                // TODO use it with localStorage
                setAuthMe({
                    authToken: result.token,
                    id: result.id,
                    email: result.email,
                    allowActions: [],
                    isAdmin: result.isAdmin,
                    role: result.role,
                })
            } catch (e) {
                console.error(e)
                setAuthMe(undefined)
            }
        },
        []
    )

    const logOut = React.useCallback(() => {
        // TODO delete it from localStorage
        setAuthMe(undefined)
    }, [])

    return {
        logOut,
        getLoggedMe,
        authMe,
    }
}
