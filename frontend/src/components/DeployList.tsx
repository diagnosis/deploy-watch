import { useDeploys } from "../hooks/useDeploys.ts";
import { DeployItem } from "./DeployItem.tsx";
import { useState } from "react";
import { useSSE } from "../hooks/useSSE.ts";
import { useAuth } from "../hooks/useAuth.ts";
import type { DeployEvent } from "../lib/api.ts";

export function DeployList() {
    const { deploys, isLoading, isError } = useDeploys();
    const { user } = useAuth();
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
        if (!user) {
            return (
                <div className="bg-blue-50 border border-blue-200 rounded-lg p-6">
                    <h3 className="text-lg font-semibold text-blue-900 mb-3">Setup Instructions</h3>
                    <p className="text-blue-800 mb-4">To start tracking your deploys, add a webhook to your GitHub repository:</p>
                    <ol className="list-decimal list-inside space-y-2 text-blue-900 mb-4">
                        <li>Go to your repository settings</li>
                        <li>Click on <strong>Webhooks</strong></li>
                        <li>Click <strong>Add webhook</strong></li>
                        <li>Add the following URL:
                            <div className="bg-white border border-blue-300 rounded px-3 py-2 mt-2 font-mono text-sm break-all">
                                https://backend-cold-bush-2228.fly.dev/webhook
                            </div>
                        </li>
                        <li>Set Content type to <strong>application/json</strong></li>
                        <li>Click <strong>Add webhook</strong></li>
                    </ol>
                    <p className="text-sm text-blue-700">Once configured, all deployment events will appear here in real-time.</p>
                </div>
            );
        }
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