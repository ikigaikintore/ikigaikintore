export const encodeJsonToBase64 = (value: any): string => {
    const str = JSON.stringify(value)
    return Buffer.from(str, 'base64').toString()
}

export const decodeBase64ToJson = (value: string): any => {
    return JSON.parse(value.toString())
}
