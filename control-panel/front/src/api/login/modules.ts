import { LoginApiFactory, SigninApiFactory } from '@/api/v1'
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

export const LoginApi = LoginApiFactory(undefined, undefined, createAxios())
export const signInApi = SigninApiFactory(undefined, undefined, createAxios())
