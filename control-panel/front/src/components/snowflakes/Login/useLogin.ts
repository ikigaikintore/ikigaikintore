import * as React from 'react'

type Result = {
    isLoginWindowOpen: boolean

    onClickLogin: () => void
    onClickSubmitLogin: () => void
    onClickCloseLoginWindow: () => void
}

export const useLogin = (): Result => {
    const [isOpen, setOpen] = React.useState<boolean>(false)

    const onClickLogin = React.useCallback(() => {
        setOpen(true)
    }, [])

    const onClickSubmitLogin = React.useCallback(() => {
        return
    }, [])

    const onClickCloseLoginWindow = React.useCallback(() => {
        setOpen(false)
    }, [])

    return {
        isLoginWindowOpen: isOpen,

        onClickLogin,
        onClickSubmitLogin,
        onClickCloseLoginWindow
    }
}