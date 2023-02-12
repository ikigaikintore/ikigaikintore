import { Login, SignIn } from '@/api/v1/'
import { User } from '@/domain'

export const toSignInUserDto = (signIn: SignIn): User => {
    return {
        email: signIn.email,
        password: signIn.password,
    }
}

export const toLoginUserDto = (logIn: User): Login => {
    return {
        password: logIn.password,
        username: logIn.email,
    }
}
