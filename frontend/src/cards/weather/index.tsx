"use client"

import React from "react"

import styled from "styled-components"

import { useWeather } from "./use-weather"

import CardStyle from "@/src/cards/CardStyle"

export const Component = () => {
    const { weatherData, isLoading, isError } = useWeather()
    if (isLoading || !weatherData) {
        return <>Loading</>
    }
    if (isError) {
        return <>Something happened!</>
    }
    return (
        <WeatherCardContainer>
            <GridContainer>
                <WeatherIconContainer>
                    <Image24x24
                        rel="preload"
                        srcSet={
                            `https://openweathermap.org/img/wn/${weatherData.weatherCurrent.icon}@2x.png`
                        }
                        alt="Weather icon"
                    />
                </WeatherIconContainer>
                <WeatherDataContainer>
                    <TemperatureHumidityGrid>
                        <Temperature>
                            {weatherData.weatherCurrent.temperature} C
                        </Temperature>
                        <Humidity>{weatherData.weatherCurrent.humidity}%</Humidity>
                    </TemperatureHumidityGrid>
                    <Wind>Wind: {weatherData.weatherCurrent.windSpeed}km/h</Wind>
                    <Place>{weatherData.cityName}</Place>
                </WeatherDataContainer>
            </GridContainer>
        </WeatherCardContainer>
    )
}

Component.displayName = "WeatherCard"

const WeatherCardContainer = styled.div`
    ${CardStyle}
`

const Image24x24 = styled.img``

const GridContainer = styled.div``

const WeatherIconContainer = styled.div`
    background-color: white;
`

const WeatherDataContainer = styled.div``

const TemperatureHumidityGrid = styled.div``

const Temperature = styled.p``

const Humidity = styled.p``

const Wind = styled.p``

const Place = styled.p``
