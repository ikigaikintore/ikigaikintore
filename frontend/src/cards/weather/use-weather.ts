import { useGeolocation } from "@uidotdev/usehooks"
import useSWR from "swr"

import { GetWeather } from "@/src/api/endpoints.v1.pb"
import { token } from "@/src/auth/init"

const fetchWeatherData = async (
    latitude: number,
    longitude: number,
    token: string,
    city?: string,
) => {
    if (!token || token === "") {
        return Promise.reject("No token")
    }
    return GetWeather(
        {
            weatherFilter: {
                latitude: latitude,
                longitude: longitude,
                locationCityName: city,
            },
        },
        {
            baseURL: process.env.NEXT_PUBLIC_BASE_ENDPOINT,
            prefix: "/v1/weather",
            headers: { Authorization: `Bearer ${token}` },
        },
    ).then((res) => {
        return res
    })
}

const useGeoLocation = () => {
    const data = useGeolocation({ timeout: 2, enableHighAccuracy: true })
    return {
        latitude: data.latitude,
        longitude: data.longitude,
        geoError: data.error,
        geoLoading: data.loading,
    }
}

const fetcher = (url) => fetch(url).then(res => res.json())

export const useWeather = () => {
    const fetchWeather = async (latitude: number, longitude: number) => {
        const userToken = await token()
        if (!userToken) {
            throw new Error("No token")
        }
        return fetchWeatherData(latitude, longitude, userToken)
    }

    const geoData = useGeoLocation()
    const latitude = geoData.latitude ?? 36.2360741
    const longitude = geoData.longitude ?? 139.1867545

    const cityData =  useSWR(
        // eslint-disable-next-line max-len
        `https://nominatim.openstreetmap.org/reverse?format=json&lat=${latitude}&lon=${longitude}&zoom=18&addressdetails=1`,
        fetcher,
    )

    const { data, isLoading, error } = useSWR(
        ["GetWeather", latitude, longitude],
        ([, latitude, longitude]) => fetchWeather(latitude, longitude),
    )
    const weatherCurrent = data ? data.weatherCurrent : {
        icon: "nope",
        temperature: "-",
        humidity: "-",
        windSpeed: "-"
    }

    return {
        weatherData: {
            cityName: cityData?.data?.address?.city || "Tokyo",
            weatherCurrent: weatherCurrent,
        },
        isLoading: isLoading || geoData.geoLoading || cityData.isLoading,
        isError: error || geoData.geoError || cityData.isLoading,
    }
}
