import {GetWeather} from "@/src/api/endpoints.v1.pb";
import useSWR from "swr";

export const useWeather = (city: string) => {
    const fetchWeather = async () => {
        return GetWeather({weatherFilter: {location: city}}, {baseURL: "http://localhost:8999", prefix: "/v1/weather"})
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
