import { error } from "@sveltejs/kit";
import type { PageLoad } from "./$types";
import { Configuration, DefaultApi } from "$lib/models/api";

export const load: PageLoad = async ({ url, params }) => {
    const apiConfiguration = new Configuration({
        basePath: `${url.protocol}//${url.host}/api`,
    });
    const api = new DefaultApi(apiConfiguration);

    try {
        const response = await api.getGreeting({
            requestBodyOfTheGetGreetingMethod: {
                jsonrpc: "2.0",
                id: "something",
                method: "get_greeting",
                params: {
                    thing: params.thing,
                },
            },
        });
        return { greeting: response.result?.greeting || "" };
    } catch (e) {
        console.log(e);
        error(500, "Internal server error");
    }
};
