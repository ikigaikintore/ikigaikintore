import { useDashboard } from '@/pages/Home/useDashboard'
import * as Domain from '@/domain'
import React from 'react'
import * as NavBar from '@/components/recipes/NavBar'

type Props = {
    filters: Domain.UserFilter
    setFilter: (filter: Domain.UserFilter) => void
}

export const Component = (props: Props) => {
    const { isError, dashboardState, isLoading } = useDashboard({
        filters: props.filters,
    })

    if (isLoading) {
        return <p>Loading...</p>
    }

    if (isError) {
        return <p>Something went wrong</p>
    }

    return (
        <>
            <NavBar.Component />
            <h1>Home page</h1>
            <ul>
                {dashboardState.users.map((user, i) => (
                    <li key={i}>
                        {user.name} - {user.birthday}
                    </li>
                ))}
            </ul>
        </>
    )
}

Component.displayName = 'Home'
