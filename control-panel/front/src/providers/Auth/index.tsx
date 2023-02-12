import { LoggedMeValue, useAuthContext } from './useAuth'
import * as React from 'react'

type Props = {
    children: React.ReactNode
}

const LoggedMeContext = React.createContext<LoggedMeValue>({
    me: {
        isAdmin: false,
        allowActions: [],
        role: '',
        id: '',
        email: '',
        authToken: '',
    },
})

export const useAuth = () => React.useContext(LoggedMeContext)

export const Provider: React.FC<Props> = ({ children }) => {
    const { authMe } = useAuthContext()

    return (
        <>
            <LoggedMeContext.Provider value={{ me: authMe }}>
                {children}
            </LoggedMeContext.Provider>
        </>
    )
}

LoggedMeContext.displayName = 'AuthContext'
