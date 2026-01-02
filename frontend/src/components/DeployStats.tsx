import { useDeploys } from "../hooks/useDeploys.ts";
import {StatCard} from "./StatCard.tsx";

export function DeployStats() {
    const { deploys, isLoading, isError } = useDeploys();

    if (isLoading) return <div>Loading stats...</div>;
    if (isError) return <div>Error loading stats</div>;

    const now = new Date();

    const isSameDay = (a: Date, b: Date) =>
        a.getFullYear() === b.getFullYear() &&
        a.getMonth() === b.getMonth() &&
        a.getDate() === b.getDate();

    const todayDeploys = deploys.filter((d) =>
        isSameDay(new Date(d.created_at), now)
    );

    const startOfWeek = new Date(now);
    startOfWeek.setDate(now.getDate() - now.getDay());

    const thisWeekDeploys = deploys.filter((d) => {
        const date = new Date(d.created_at);
        return date >= startOfWeek && date <= now;
    });

    const startOfMonth = new Date(now.getFullYear(), now.getMonth(), 1);

    const thisMonthDeploys = deploys.filter((d) => {
        const date = new Date(d.created_at);
        return date >= startOfMonth && date <= now;
    });

    const repoCount = deploys.reduce<Record<string, number>>((acc, d) => {
        acc[d.repo_name] = (acc[d.repo_name] || 0) + 1;
        return acc;
    }, {});

    const mostActiveRepo = Object.entries(repoCount).sort(
        (a, b) => b[1] - a[1]
    )[0];
    const repoStats = Object.entries(
        deploys.reduce<Record<string, number>>((acc, d) => {
            acc[d.repo_name] = (acc[d.repo_name] || 0) + 1;
            return acc;
        }, {})
    ).sort((a, b) => b[1] - a[1]);

    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
            <StatCard label="Total Deploys" value={deploys.length} />
            <StatCard label="Today" value={todayDeploys.length} />
            <StatCard label="This Week" value={thisWeekDeploys.length} />
            <StatCard label="This Month" value={thisMonthDeploys.length} />

            {mostActiveRepo && (
                <div className="md:col-span-2 lg:col-span-4 border rounded-md p-4 text-sm">
                    <div className="text-gray-500">Most Active Repo</div>
                    <div className="font-medium text-gray-900">
                        {mostActiveRepo[0]} ({mostActiveRepo[1]} deploys)
                    </div>
                </div>
            )}
            <div className="md:col-span-2 lg:col-span-4 border rounded-md p-4">
                <div className="text-sm text-gray-500 mb-2">Deploys per Repository</div>

                <div className="space-y-1">
                    {repoStats.map(([repo, count]) => (
                        <div
                            key={repo}
                            className="flex justify-between text-sm text-gray-800"
                        >
                            <span className="font-medium">{repo}</span>
                            <span className="text-gray-500">{count}</span>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}
