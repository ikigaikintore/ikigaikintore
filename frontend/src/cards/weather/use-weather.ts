import useSWR from "swr";
import {GetWeather} from "@/src/api/endpoints.v1.pb";

export const useWeather = (city: string) => {
    const fetchWeather = async () => {
        return GetWeather({weatherFilter: {location: city}}, {baseURL: process.env.BASE_ENDPOINT, prefix: "/v1/weather"})
            .then(res => {
                return res
            })
    }

    const {data, isLoading, error} = useSWR(`GetWeather/${city}`, fetchWeather)

    return {
        weatherData: data,
        isLoading,
        isError: error,
    }
}
