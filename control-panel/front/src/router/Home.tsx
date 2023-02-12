import { UserFilter } from '@/domain'
import * as React from 'react'
import * as Router from 'react-router-dom'
import * as Url from '@/modules'
import * as Home from '@/pages/Home'

const filterKey = 'search'

export const Component = () => {
    const [searchParams, setSearchParams] = Router.useSearchParams()

    const setFilter = React.useCallback(
        (filter: UserFilter) => {
            const value = Url.encodeJsonToBase64(filter)
            const params = new URLSearchParams(searchParams)
            if (!filter) {
                params.delete(filterKey)
            } else {
                params.set(filterKey, value)
            }
            setSearchParams(params, { replace: true })
        },
        [searchParams, setSearchParams]
    )

    const filter = React.useMemo(() => {
        const filter = searchParams.get(filterKey)
        if (!filter) {
            return []
        }
        return Url.decodeBase64ToJson(filter)
    }, [searchParams])

    return <Home.Component filters={filter} setFilter={setFilter} />
}
