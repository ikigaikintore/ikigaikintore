'use client'

import React from 'react'
import styled from 'styled-components'

import {useWeather} from './use-weather'

type WeatherProps = {}

export const Component = (props: WeatherProps) => {
    const { weatherData, isLoading, isError } = useWeather('Tokyo')
    if (isLoading) {
        return <>Loading</>
    }
    if (isError) {
        return <>Something happened!</>
    }
    console.log(isLoading, isError, weatherData)
    return (
        <WeatherCardContainer>
            <GridContainer>
                <WeatherIconContainer>
                    <img src="path/to/weather-icon.png" alt="Weather icon" />
                </WeatherIconContainer>
                <WeatherDataContainer>
                    <TemperatureHumidityGrid>
                        <Temperature>{weatherData.weatherCurrent.temperature} C</Temperature>
                        <Humidity>{weatherData.weatherCurrent.humidity}%</Humidity>
                    </TemperatureHumidityGrid>
                    <Wind>Wind: {weatherData.weatherCurrent.windSpeed}km/h</Wind>
                    <Place>Tokyo, Japan</Place>
                </WeatherDataContainer>
            </GridContainer>
        </WeatherCardContainer>
    )
}

Component.displayName = 'WeatherCard'

const WeatherCardContainer = styled.div`
  background-color: white;
  border-radius: 10px;
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.1);
  padding: 16px;
`;

const GridContainer = styled.div`
  display: grid;
  grid-template-columns: 1fr;
  @media (min-width: 768px) {
    grid-template-columns: 1fr 1fr;
  }
  column-gap: 16px;
`;

const WeatherIconContainer = styled.div`
  grid-column: span 2;
  @media (min-width: 768px) {
    grid-column: span 1;
  }
`;

const WeatherDataContainer = styled.div`
  grid-column: span 2;
  @media (min-width: 768px) {
    grid-column: span 1;
  }
`;

const TemperatureHumidityGrid = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  column-gap: 16px;
`;

const Temperature = styled.p`
  font-weight: bold;
  font-size: 24px;
`;

const Humidity = styled.p`
  color: #6b7280;
`;

const Wind = styled.p`
  font-size: 14px;
`;

const Place = styled.p`
  font-weight: 600;
`;