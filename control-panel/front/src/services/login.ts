import { toLoginUserDto } from '@/api/login/dto'
import { LoginApi } from '@/api/login/modules'
import { Logged } from '@/api/v1'
import * as Domain from '@/domain'
import * as React from 'react'
import { useMutation } from 'react-query'

type Results = {
    logIn: (args: Domain.User) => Promise<Logged>

    isLoading: boolean

    isError: boolean
}

type UseLoginIn = () => Results

export const useSubmitLogIn: UseLoginIn = () => {
    const mutation = useMutation((args: Domain.User) =>
        LoginApi.login(toLoginUserDto(args))
    )

    const logIn = React.useCallback(
        async (args: Domain.User) => {
            const { data } = await mutation.mutateAsync(args)
            return data
        },
        [mutation]
    )

    return {
        logIn,

        isLoading: mutation.isLoading,

        isError: mutation.isError,
    }
}
