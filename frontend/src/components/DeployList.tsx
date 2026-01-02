import {useDeploys} from "../hooks/useDeploys.ts";
import {DeployItem} from "./DeployItem.tsx";

export function DeployList() {
    const {
        deploys,
        isLoading,
        isError,
    } = useDeploys();

    if (isLoading) {
        return <div className="text-gray-500">Loading deploys...</div>;
    }

    if (isError) {
        return <div className="text-red-600">Failed to load deploys</div>;
    }

    if (deploys.length === 0) {
        return <div className="text-gray-400">No deploys yet</div>;
    }

    return(
        <div className="grid grid-cols-1 gap-4">
            {deploys.map((deploy) => (
                <DeployItem key={deploy.id} deploy={deploy} />
            ))}
        </div>
    )
}