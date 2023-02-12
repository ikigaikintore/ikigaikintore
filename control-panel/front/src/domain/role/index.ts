export type Action = {
    name: string
    permissions: Permission[]
}

export type Permission = {
    name: string
    type: string
}
