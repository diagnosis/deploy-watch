import type { DeployEvent } from "../lib/api";

interface DeployItemProps {
    deploy: DeployEvent;
    highlight?: boolean;
}

export function DeployItem({ deploy, highlight }: DeployItemProps) {
    return (
        <div
            className={`border rounded-md p-4 text-sm space-y-1 transition-all duration-500
        ${highlight ? "bg-yellow-100 animate-pulse" : "bg-white"}`}
        >
            <div className="font-medium text-gray-900">{deploy.repo_name}</div>
            <div className="text-gray-500">{deploy.branch}</div>
            <div className="text-gray-700">{deploy.commit_message}</div>
            <div className="text-gray-500">
                {deploy.author} · <span className="font-mono">{deploy.commit_sha.slice(0, 7)}</span>
            </div>
            <div className="text-xs text-gray-400">
                {new Date(deploy.created_at).toLocaleString()}
            </div>
        </div>
    );
}