import { toSignInUserDto } from '@/api/login/dto'
import { signInApi } from '@/api/login/modules'
import * as Domain from '@/domain'
import * as React from 'react'
import { useMutation } from 'react-query'

type Results = {
    signIn: (args: Domain.User) => Promise<void>

    isLoading: boolean

    isError: boolean
}

type UseSignIn = () => Results

export const useSubmitSignIn: UseSignIn = () => {
    const mutation = useMutation((args: Domain.User) =>
        signInApi.signIn(toSignInUserDto(args))
    )

    const signIn = React.useCallback(
        async (args: Domain.User) => {
            await mutation.mutateAsync(args)
        },
        [mutation]
    )

    return {
        signIn,

        isLoading: mutation.isLoading,

        isError: mutation.isError,
    }
}
