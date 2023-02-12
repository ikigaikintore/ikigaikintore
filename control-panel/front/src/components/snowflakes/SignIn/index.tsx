import { useValidation } from '@/components/snowflakes/SignIn/useValidation'
import { yupResolver } from '@hookform/resolvers/yup'
import * as React from 'react'
import { useForm } from 'react-hook-form'
import styled from 'styled-components'

type Props = {
    onClickSubmitSignIn: () => void
    onClickCloseSignInWindow: () => void
}

type FormValues = {
    username: string
    password: string
    email: string
}

export const Component = (props: Props) => {
    const { schema } = useValidation()
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<FormValues>({ resolver: yupResolver(schema) })

    return (
        <SigninWrapper>
            <Title>Sign In</Title>
            <Content>
                <>
                    {errors.username && (
                        <ErrorMessageWrap>
                            {errors.username?.message}
                        </ErrorMessageWrap>
                    )}
                    <FormLabel>
                        Username:
                        <InputText
                            className={errors.username ? 'invalid' : ''}
                            {...register('username')}
                            name="username"
                            type="text"
                            placeholder="email or username"
                        />
                    </FormLabel>
                    {errors.password && (
                        <ErrorMessageWrap>
                            {errors.password?.message}
                        </ErrorMessageWrap>
                    )}
                    <FormLabel>
                        Password:
                        <InputText
                            className={errors.password ? 'invalid' : ''}
                            {...register('password')}
                            name="password"
                            type="password"
                            placeholder="password"
                        />
                    </FormLabel>
                    {errors.email && (
                        <ErrorMessageWrap>
                            {errors.email?.message}
                        </ErrorMessageWrap>
                    )}
                    <FormLabel>
                        Email:
                        <InputText
                            className={errors.email ? 'invalid' : ''}
                            {...register('email')}
                            name="email"
                            type="email"
                            placeholder="email address"
                        />
                    </FormLabel>
                </>
                <ButtonsAlignWrap>
                    <button onClick={props.onClickCloseSignInWindow}>
                        Close
                    </button>
                    <button onClick={handleSubmit(props.onClickSubmitSignIn)}>
                        Submit
                    </button>
                </ButtonsAlignWrap>
            </Content>
        </SigninWrapper>
    )
}

Component.displayName = 'SignInForm'

const ErrorMessageWrap = styled.span`
    padding: 10px;
    font-size: 12px;
    font-weight: bold;
`

const SigninWrapper = styled.div`
    position: relative;
    float: right;
    right: 10px;
    top: 0;
    z-index: 100;
    border: 1px solid grey;
    box-shadow: darkgray;
    border-radius: 10px;
    width: 250px;
    background-color: cadetblue;
`

const Title = styled.div`
    align-items: center;
    text-align: center;
    font-weight: bold;
    font-size: 16px;
`

const Content = styled.div``

const ButtonsAlignWrap = styled.div`
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    gap: 10px;
    padding: 4px;
`

const FormLabel = styled.label`
    padding: 6px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin: 0 10px;
`

const InputText = styled.input`
    box-sizing: border-box;
    color: blue;
    background-color: papayawhip;
    border-radius: 2px;
    border: none;
    width: 120px;

    &.invalid {
        border-color: red;
        background-color: brown;
    }
`
