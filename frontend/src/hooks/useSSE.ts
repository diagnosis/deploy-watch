import {useEffect, useRef} from "react";

const API_URL = import.meta.env.PROD
    ? 'https://backend-cold-bush-2228.fly.dev'
    : '';

export function useSSE(onMessage: (event: MessageEvent) => void) {
    const onMessageRef = useRef(onMessage);
    const reconnectTimeoutRef = useRef<number | null>(null);

    useEffect(() => {
        onMessageRef.current = onMessage;
    }, [onMessage]);

    useEffect(() => {
        let eventSource: EventSource | null = null;
        let isMounted = true;

        const connect = () => {
            if (!isMounted) return;

            const url = `${API_URL}/api/events`;
            eventSource = new EventSource(url, {
                withCredentials: true
            });

            eventSource.onopen = () => {
                console.log("SSE connected");
            };

            eventSource.onmessage = (event) => {
                if (isMounted) {
                    onMessageRef.current(event);
                }
            };

            eventSource.onerror = (err) => {
                console.error("SSE connection error", err);
                eventSource?.close();

                if (isMounted) {
                    reconnectTimeoutRef.current = window.setTimeout(() => {
                        console.log("Reconnecting SSE...");
                        connect();
                    }, 3000);
                }
            };
        };

        connect();

        return () => {
            isMounted = false;
            if (reconnectTimeoutRef.current) {
                clearTimeout(reconnectTimeoutRef.current);
            }
            eventSource?.close();
        };
    }, [])
}