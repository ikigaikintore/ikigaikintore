"use client"

import styled from "styled-components"

import { AuthProvider } from "@/src/auth/index"
import { useAuth } from "@/src/auth/use-auth"
import * as WeatherCard from "@/src/cards/weather"

export default function Page() {
  const { user, signInUser, signOutUser } = useAuth()

  return (
    <AuthProvider>
      <AppPage>
        <MainMenu>
          <AuthButtons>
            {user ? (
              <button onClick={signOutUser}>Sign out</button>
            ) : (
              <button onClick={signInUser}>Sign in</button>
            )}
          </AuthButtons>
        </MainMenu>
        {user ? (
          <MainBody>
            <LinkSection>Link elements</LinkSection>
            <CardComponents>
              <WeatherCard.Component />
            </CardComponents>
          </MainBody>
        ) : (<div></div>)}
      </AppPage>
    </AuthProvider>
  )
}

const AppPage = styled.section`
  display: grid;
  grid-template-rows: auto 1fr;
  height: 100vh;
`

const MainMenu = styled.header`
  background-color: #333; /* Example background color */
  color: white; /* Example text color */
  padding: 10px;
`

const LinkSection = styled.aside`
  background-color: #ddd; /* Example background color */
  padding: 10px;
`

const MainBody = styled.main`
  display: grid;
  grid-template-columns: 10% 1fr; /* 10% for LinkSection and 90% for CardComponents */
  height: 100%;
`

const CardComponents = styled.div`
  display: grid;
  padding: 10px;
  gap: 16px;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
`

const AuthButtons = styled.div`
  display: flex;
  gap: 8px;
  padding: 10px;
  align-items: center;

  button {
    background-color: #f0f0f0; /* Light grey background */
    color: #333; /* Dark text color for contrast */
    border: none;
    padding: 8px 16px; /* Top and Bottom, Left and Right padding */
    border-radius: 4px; /* Rounded corners */
    cursor: pointer; /* Cursor changes to pointer to indicate clickable */
    transition:
      background-color 0.3s,
      color 0.3s; /* Smooth transition for hover effect */

    &:hover {
      background-color: #e0e0e0; /* Slightly darker on hover */
      color: #000; /* Darker text color on hover */
    }

    &:focus {
      outline: none; /* Remove default focus outline */
      box-shadow: 0 0 0 2px #aaa; /* Custom focus style to improve accessibility */
    }
  }
`
