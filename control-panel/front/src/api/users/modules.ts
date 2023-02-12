import { UsersApiFactory } from '@/api/v1'
import axios, { AxiosError } from 'axios'

type Options = {
    withCredentials: boolean
}

const createAxios = (opts?: Options) => {
    const instance = axios.create({
        withCredentials: opts?.withCredentials || false,
    })

    instance.interceptors.response.use(
        (response) => response,
        (err: AxiosError) => {
            return Promise.reject(err.response)
        }
    )

    return instance
}

export const UsersApiInstance = UsersApiFactory(
    undefined,
    undefined,
    createAxios()
)
export const usersApi = UsersApiFactory(undefined, undefined, createAxios())
