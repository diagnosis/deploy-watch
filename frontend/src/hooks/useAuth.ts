import {useQuery} from "@tanstack/react-query";
import {appClient} from "../lib/api.ts";

export function useAuth() {
    const {
        data,
        isLoading,
        isError,
    } = useQuery({
        queryKey: ['auth'],
        queryFn: () => appClient.getCurrentUser(),
        retry: false,
    });

    return { user: data ?? null, isLoading, isError }

}