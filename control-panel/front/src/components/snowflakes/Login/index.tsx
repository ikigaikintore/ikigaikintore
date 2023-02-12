import * as React from 'react'
import styled from 'styled-components'

type Props = {
    onClickSubmitLogin: () => void
    onClickCloseLoginWindow: () => void
}

export const Component = (props: Props) => {
    return (
        <LoginWrapper>
            <Title>Log In</Title>
            <Content>
                <FormContentWrap>
                    <FormLabel>Username:
                        <InputText type="text" placeholder="email or username"/>
                    </FormLabel>
                    <FormLabel>Password:
                        <InputText type="password" placeholder="password"/>
                    </FormLabel>
                </FormContentWrap>
                <ButtonsAlignWrap>
                    <button onClick={props.onClickCloseLoginWindow}>Close</button>
                    <button onClick={props.onClickSubmitLogin}>Submit</button>
                </ButtonsAlignWrap>
            </Content>
        </LoginWrapper>
    )
}

Component.displayName = 'LoginForm'

const LoginWrapper = styled.div`
  position: relative;
  float: right;
  right: 50px;
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

const FormContentWrap = styled.div`
  justify-content: space-between;
`

const FormLabel = styled.label`
  padding: 6px;
  display: flex;
  justify-content: space-evenly;
  align-items: center;
`

const InputText = styled.input`
  box-sizing: border-box;
  color: blue;
  background-color: papayawhip;
  border-radius: 2px;
  border: none;
  width: 120px;
  margin-left: 10px;
`
