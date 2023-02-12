import { Action } from '../role'

export type Me = {
    id: string
    email: string
    isAdmin: boolean
    role: string
    allowActions: Action[]
}
