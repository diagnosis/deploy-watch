import {useEffect} from "react";

const API_URL = import.meta.env.PROD
    ? 'https://backend-cold-bush-2228.fly.dev'
    : '';

export function useSSE(onMessage: (event: MessageEvent) => void) {
    useEffect(() => {
        const url = `${API_URL}/api/events`;
        const eventSource = new EventSource(url, {
            withCredentials: true
        });

        eventSource.onmessage = (event) => onMessage(event)

        eventSource.onerror = (err) => {
            console.log("SSE error", err)
            eventSource.close()
        }
        return  () =>{
            eventSource.close()
        }
    }, [onMessage])
}