import {useEffect} from "react";

export function useSSE(onMessage: (event: MessageEvent) => void) {
    useEffect(() => {
        const eventSource = new EventSource('/api/events');
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