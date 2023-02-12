import { toFilteredUser, toFilterUser } from '@/api/users/dto'
import { usersApi } from '@/api/users/modules'
import * as Domain from '@/domain'
import * as React from 'react'
import { useMutation } from 'react-query'

type Results = {
    filterUsers: (args: Domain.UserFilter) => Promise<Domain.FilteredUser[]>

    isLoading: boolean

    isError: boolean
}

type UseFilterUsers = () => Results

export const useFilterUsers: UseFilterUsers = () => {
    const mutation = useMutation((args: Domain.UserFilter) =>
        usersApi.filterUsers(toFilterUser(args))
    )

    const filterUsers = React.useCallback(
        async (args: Domain.UserFilter) => {
            const response = await mutation.mutateAsync(args)
            return response.data.map((u) => toFilteredUser(u))
        },
        [mutation]
    )

    return {
        filterUsers,

        isLoading: mutation.isLoading,

        isError: mutation.isError,
    }
}
