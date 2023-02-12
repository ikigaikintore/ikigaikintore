import * as Service from '@/services'
import * as React from 'react'

type Result = {
    isSiginWindowOpen: boolean

    isError: boolean
    isLoading: boolean

    onClickSignIn: () => void
    onClickSubmitSignIn: () => void
    onClickCloseSiginWindow: () => void
}

export const useSignIn = (): Result => {
    const {
        signIn,
        isLoading: signInIsLoading,
        isError: signInIsError,
    } = Service.useSubmitSignIn()

    const [isOpen, setOpen] = React.useState<boolean>(false)

    const onClickSignIn = React.useCallback(() => {
        setOpen(true)
    }, [])

    const onClickSubmitSignIn = React.useCallback(async () => {
        try {
            await signIn({ username: '', password: '', email: '' })
        } catch (err) {
            console.error(err)
        }
    }, [signIn])

    const onClickCloseSiginWindow = React.useCallback(() => {
        setOpen(false)
    }, [])

    return {
        isSiginWindowOpen: isOpen,

        isError: signInIsError,
        isLoading: signInIsLoading,

        onClickSignIn,
        onClickSubmitSignIn,
        onClickCloseSiginWindow,
    }
}
