'use client'

import React from 'react'
import styled from 'styled-components'

import {useWeather} from './use-weather'

type WeatherProps = {}

export const Component = (props: WeatherProps) => {
    const { weatherData, isLoading, isError } = useWeather('Tokyo')
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
                    <Image24x24 rel="preload" srcSet={`https://openweathermap.org/img/wn/${weatherData.weatherCurrent.icon}@2x.png`} alt="Weather icon" />
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
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  padding: 8px;
`;

const Image24x24 = styled.img`
    width: 128px;
    height: 128px;

    border-radius: 50%;
    //object-fit: cover;
`;

const GridContainer = styled.div`
  display: grid;
  grid-template-columns: 2fr 2fr;
  //column-gap: 16px;
`;

const WeatherIconContainer = styled.div`
  grid-column: span 4;
`;

const WeatherDataContainer = styled.div`
  grid-column: span 4;
`;

const TemperatureHumidityGrid = styled.div`
    display: grid;
    grid-auto-flow: column;
    align-items: center;
`;

const Temperature = styled.p`
  font-weight: bold;
  font-size: 16px;
`;

const Humidity = styled.p`
  color: darkblue;
`;

const Wind = styled.p`
  font-size: 14px;
`;

const Place = styled.p`
  font-weight: bold;
`;