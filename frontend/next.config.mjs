export const compiler = {
    styledComponents: true,
};
export const output = "export";
export async const headers = () => {
    return [
        {
            source: "/",
            headers: [
                {
                    key: "Access-Control-Allow-Origin",
                    value: "*",
                },
                {
                    key: "Access-Control-Allow-Methods",
                    value: "GET, POST, PUT, DELETE, OPTIONS",
                },
                {
                    key: "Access-Control-Allow-Headers",
                    value: "Content-Type, Authorization",
                },
            ],
        },
    ];
}