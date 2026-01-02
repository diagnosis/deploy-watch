import { useDeploys } from "../hooks/useDeploys.ts";
import { DeployItem } from "./DeployItem.tsx";
import { useState } from "react";
import { useSSE } from "../hooks/useSSE.ts";
import type { DeployEvent } from "../lib/api.ts";

export function DeployList() {
    const { deploys, isLoading, isError } = useDeploys();
    const [events, setEvents] = useState<DeployEvent[]>([]);
    const [highlightId, setHighlightId] = useState<string | null>(null);

    // Listen to SSE
    useSSE((event) => {
        const deploy: DeployEvent = JSON.parse(event.data);
        setEvents((prev) => [deploy, ...prev]);
        setHighlightId(deploy.id);

        // Remove highlight after 10s
        setTimeout(() => setHighlightId(null), 10000);
    });

    if (isLoading) {
        return <div className="text-gray-500">Loading deploys...</div>;
    }

    if (isError) {
        return <div className="text-red-600">Failed to load deploys</div>;
    }

    const allDeploys = [...events, ...deploys];

    if (allDeploys.length === 0) {
        return <div className="text-gray-400">No deploys yet</div>;
    }

    return (
        <div className="grid grid-cols-1 gap-4">
            {allDeploys.map((deploy) => (
                <DeployItem
                    key={deploy.id}
                    deploy={deploy}
                    highlight={deploy.id === highlightId}
                />
            ))}
        </div>
    );
}