import {useEffect, useRef} from "react";

const API_URL = import.meta.env.PROD
    ? 'https://backend-cold-bush-2228.fly.dev'
    : '';

export function useSSE(onMessage: (event: MessageEvent) => void) {
    const onMessageRef = useRef(onMessage);

    useEffect(() => {
        onMessageRef.current = onMessage;
    }, [onMessage]);

    useEffect(() => {
        const url = `${API_URL}/api/events`;
        const eventSource = new EventSource(url, {
            withCredentials: true
        });

        eventSource.onmessage = (event) => onMessageRef.current(event)

        eventSource.onerror = (err) => {
            console.error("SSE connection error", err)
        }

        return () => {
            eventSource.close()
        }
    }, [])
}