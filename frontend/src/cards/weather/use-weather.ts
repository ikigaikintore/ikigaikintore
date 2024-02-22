import useSWR from "swr"

import { GetWeather } from "@/src/api/endpoints.v1.pb"
import { token } from "@/src/auth/init"

const fetchWeatherData = async (city: string, token: string) => {
  if (!token || token === "") {
    return Promise.reject("No token")
  }
  return GetWeather(
    { weatherFilter: { location: city } },
    {
      baseURL: process.env.NEXT_PUBLIC_BASE_ENDPOINT,
      prefix: "/v1/weather",
      headers: { Authorization: `Bearer ${token}` },
    }
  ).then((res) => {
    return res
  })
}

export const useWeather = (city: string) => {
  const fetchWeather = async () => {
    const userToken = await token()
    if (!userToken) {
      throw new Error("No token")
    }
    return fetchWeatherData(city, userToken)
  }

  const { data, isLoading, error } = useSWR(`GetWeather/${city}`, fetchWeather)

  return {
    weatherData: data,
    isLoading,
    isError: error,
  }
}
