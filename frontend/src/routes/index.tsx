// routes/index.tsx
import { createFileRoute } from '@tanstack/react-router'
import {DeployList} from "../components/DeployList.tsx";

function Home() {
    return (
        <div>
            <div className="bg-blue-50 border border-blue-200 rounded-lg p-6 mb-6">
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
                <p className="text-sm text-blue-700">Once configured, all deployment events will appear below in real-time.</p>
            </div>

            <h2 className="text-2xl font-semibold mb-4">Recent Deploys</h2>
            <DeployList />
        </div>
    )
}

export const Route = createFileRoute('/')({
    component: Home,
})