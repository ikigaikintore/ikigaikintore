import { User, UserFilter } from '@/api/v1/'
import * as Domain from '@/domain'

export const toFilterUser = (filter: UserFilter): Domain.UserFilter => {
    return {
        birthday: filter.birthday,
        name: filter.name,
    }
}

export const toFilteredUser = (user: User): Domain.FilteredUser => {
    return {
        birthday: user.birthday,
        name: user.name,
        id: user.id,
    }
}
