import { FilteredUser, UserFilter } from '@/domain'
import { useFilterUsers } from '@/services'
import * as React from 'react'

type Args = {
    filters: UserFilter
}

type Results = {
    dashboardState: DashboardState

    isError: boolean

    isLoading: boolean
}

type DashboardState = {
    users: FilteredUser[]
    sort: {
        type: 'name' | 'birthday'
        order: 'asc' | 'desc'
    }
}

export const useDashboard = (args: Args): Results => {
    const {
        filterUsers,
        isError: isErrorFilterUsers,
        isLoading: isLoadingFilterUsers,
    } = useFilterUsers()

    const [dashboardState, setDashboardState] = React.useState<DashboardState>({
        users: [],
        sort: {
            type: 'name',
            order: 'asc',
        },
    })

    React.useEffect(() => {
        ;(async () => {
            try {
                const users = await filterUsers(args.filters)
                if (!users) {
                    return
                }
                setDashboardState((old) => ({ ...old, users: users }))
            } catch (err) {
                console.error(err)
            }
        })()
    }, [])

    return {
        dashboardState,

        isError: isErrorFilterUsers,

        isLoading: isLoadingFilterUsers,
    }
}
