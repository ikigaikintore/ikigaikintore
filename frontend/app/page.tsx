"use client"

import * as WeatherCard from '@/src/cards/weather'
import styled from "styled-components";

export default function Page() {
    return (
        <AppPage>
            <MainMenu>
                My menu here
            </MainMenu>
            <MainBody>
                <LinkSection>
                    Link elements
                </LinkSection>
                <CardComponents>
                    <WeatherCard.Component/>
                </CardComponents>
            </MainBody>
        </AppPage>
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