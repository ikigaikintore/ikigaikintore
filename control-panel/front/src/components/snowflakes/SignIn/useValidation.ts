import * as React from 'react'
import * as Yup from 'yup'

export const useValidation = () => {
    const schema = React.useMemo(() => {
        return Yup.object().shape({
            username: Yup.string()
                .required('username is required')
                .min(5)
                .test('isEmail', 'Invalid email format', (value) => {
                    if (value)
                        return value.includes('@')
                            ? Yup.string().email().isValidSync(value)
                            : true
                    return true
                }),
            email: Yup.string().required('email is required').email(),
            password: Yup.string()
                .min(8)
                .max(32)
                .matches(/[a-zA-Z0-9_\-|!@#$%&*()+=]/),
        })
    }, [])

    return {
        schema,
    }
}
